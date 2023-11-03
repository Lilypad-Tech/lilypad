package bacalhau

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/data/bacalhau"
	executorlib "github.com/bacalhau-project/lilypad/pkg/executor"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
}

type BacalhauExecutor struct {
	Options     BacalhauExecutorOptions
	bacalhauEnv []string
}

func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {
	bacalhauEnv := []string{
		fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)
	return &BacalhauExecutor{
		Options:     options,
		bacalhauEnv: bacalhauEnv,
	}, nil
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	id, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	resultsDir, err := executor.copyJobResults(deal.ID, id)
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

	if jobState.State.State != bacalhau.JobStateCompleted {
		return nil, fmt.Errorf("job %s did not complete successfully: %s", id, jobState.State.State.String())
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

	runOutput, err := runCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutput)
	}

	id := strings.TrimSpace(string(runOutput))
	fmt.Printf("Got bacalhau job ID: %s\n", id)

	return id, nil
}

func (executor *BacalhauExecutor) copyJobResults(dealID string, jobID string) (string, error) {
	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, dealID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder of results %s -> %s", dealID, err.Error())
	}

	copyResultsCmd := exec.Command(
		"bacalhau",
		"get",
		jobID,
		"--output-dir", resultsDir,
	)
	copyResultsCmd.Env = executor.bacalhauEnv

	_, err = copyResultsCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error copying results %s -> %s", dealID, err.Error())
	}

	return resultsDir, nil
}

func (executor *BacalhauExecutor) getJobState(dealID string, jobID string) (*bacalhau.JobWithInfo, error) {
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

	var job bacalhau.JobWithInfo
	err = json.Unmarshal(output, &job)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling job JSON %s -> %s", dealID, err.Error())
	}

	return &job, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)
