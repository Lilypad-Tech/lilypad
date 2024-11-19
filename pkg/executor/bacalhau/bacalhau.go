package bacalhau

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/data/bacalhau"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/ipfs"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
	ApiPort string
}

type BacalhauExecutor struct {
	Options        BacalhauExecutorOptions
	bacalhauEnv    []string
	bacalhauClient BacalhauClient
	ipfsClient     *ipfs.Client
}

func NewBacalhauExecutor(options BacalhauExecutorOptions, ipfsClient *ipfs.Client) (*BacalhauExecutor, error) {
	client, err := newClient(options)
	if err != nil {
		return nil, err
	}

	apiHost := ""
	if options.ApiHost != "DO_NOT_SET" {
		apiHost = fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost)
	}
	bacalhauEnv := []string{
		apiHost,
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)

	return &BacalhauExecutor{
		Options:        options,
		bacalhauEnv:    bacalhauEnv,
		bacalhauClient: *client,
		ipfsClient:     ipfsClient,
	}, nil
}

func (executor *BacalhauExecutor) Id() (string, error) {
	nodeIdCmd := exec.Command(
		"bacalhau",
		"id",
	)
	nodeIdCmd.Env = executor.bacalhauEnv

	runOutputRaw, err := nodeIdCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error calling get id results %s, %s", err.Error(), runOutputRaw)
	}

	splitOutputs := strings.Split(strings.Trim(string(runOutputRaw), " \t\n"), "\n")
	runOutput := splitOutputs[len(splitOutputs)-1]

	var idResult struct {
		ID       string
		ClientID string
	}
	err = json.Unmarshal([]byte(runOutput), &idResult)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling job JSON %s %s", err.Error(), runOutputRaw)
	}

	return idResult.ID, nil
}

// Checks that Bacalhau is installed, correctly configured, and available
func (executor *BacalhauExecutor) IsAvailable() (bool, error) {
	isAlive, err := executor.bacalhauClient.alive()
	if !isAlive || err != nil {
		return false, fmt.Errorf("Bacalhau is not currently available. Please ensure that Bacalhau is running, then try again. %w", err)
	}

	// Check that we have the right version of Bacalhau
	version, err := executor.bacalhauClient.getVersion()
	if err != nil {
		return false, err
	}
	// TODO: we may want to relax this
	if version.GitVersion != "v1.3.2" {
		return false, errors.New("Bacalhau version must be v1.3.2")
	}

	return true, nil
}

func (executor *BacalhauExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	var specs []data.MachineSpec
	result, err := executor.bacalhauClient.getNodes()
	if err != nil {
		return specs, err
	}

	for _, node := range result.Nodes {
		spec := data.MachineSpec{
			CPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.CPU) * 1000, // convert float to "milli-CPU"
			RAM:  int(node.Info.ComputeNodeInfo.MaxCapacity.Memory),
			GPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.GPU),
			Disk: int(node.Info.ComputeNodeInfo.MaxCapacity.Disk),
		}
		for _, gpu := range node.Info.ComputeNodeInfo.MaxCapacity.GPUs {
			spec.GPUs = append(spec.GPUs, data.GPUSpec{
				Name:   gpu.Name,
				Vendor: string(gpu.Vendor),
				VRAM:   int(gpu.Memory),
			})
		}
		specs = append(specs, spec)
	}
	return specs, nil
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	id, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	jobState, err := executor.getJobState(deal.ID, id)
	if err != nil {
		return nil, err
	}

	if len(jobState.State.Executions) <= 0 {
		return nil, fmt.Errorf("no executions found for job %s", id)
	}

	// Check that the job completed successfully
	if jobState.State.State != bacalhau.JobStateCompleted {
		return nil, fmt.Errorf("job %s did not complete successfully: %s", id, jobState.State.State.String())
	}

	system.EnsureDataDir(RESULTS_DIR)
	resultsDir := system.GetDataDir(filepath.Join(RESULTS_DIR, deal.ID))
	cidString := jobState.State.Executions[0].PublishedResult.CID
	err = executor.ipfsClient.Get(context.Background(), cidString, resultsDir)
	if err != nil {
		return nil, fmt.Errorf("error getting results from IPFS %s -> %s", deal.ID, err)
	}

	// TODO: we should think about WASM and instruction count here
	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       jobState.State.Executions[0].PublishedResult.CID,
		InstructionCount: 1,
	}

	return results, nil
}

// run the bacalhau job and return the job ID
func (executor *BacalhauExecutor) getJobID(
	deal data.DealContainer,
	module data.Module,
) (string, error) {
	// get a JSON string of the job
	jsonBytes, err := json.Marshal(module.Job)
	if err != nil {
		return "", fmt.Errorf("error getting job JSON for deal %s -> %s", deal.ID, err.Error())
	}
	bacalhauJobSpecDir, err := system.EnsureDataDir(filepath.Join("bacalhau-job-specs", deal.ID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder for job specs %s -> %s", deal.ID, err.Error())
	}
	jobPath := filepath.Join(bacalhauJobSpecDir, "job.json")
	err = system.WriteFile(jobPath, jsonBytes)
	if err != nil {
		return "", fmt.Errorf("error writing job JSON %s -> %s", deal.ID, err.Error())
	}

	runCmd := exec.Command(
		"bacalhau",
		"create",
		"--id-only",
		"--wait",
		jobPath,
	)
	runCmd.Env = executor.bacalhauEnv

	runOutputRaw, err := runCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutputRaw)
	}
	splitOutputs := strings.Split(string(runOutputRaw), "\n")
	runOutput := splitOutputs[0]
	outputError := strings.Join(strings.Fields(strings.Join(splitOutputs[1:], " ")), " ")

	if outputError != "" {
		log.Error().Msgf("error parsing output %s -> %s, %s", deal.ID, outputError, runOutput)
	}

	id := strings.TrimSpace(string(runOutput))
	fmt.Printf("Got bacalhau job ID: %s\n", id)

	return id, nil
}

func (executor *BacalhauExecutor) getJobState(dealID string, jobID string) (*bacalhau.JobWithInfo, error) {
	var job bacalhau.JobWithInfo

	for len(job.State.Executions) == 0 {
		describeCmd := exec.Command(
			"bacalhau",
			"describe",
			"--json",
			jobID,
		)
		describeCmd.Env = executor.bacalhauEnv

		output, err := describeCmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("error calling describe command results %s -> %s", dealID, err.Error())
		}

		err = json.Unmarshal(output, &job)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling job JSON %s -> %s", dealID, err.Error())
		}
	}

	return &job, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)
