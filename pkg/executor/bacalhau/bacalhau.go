package bacalhau

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"path/filepath"

	"github.com/lilypad-tech/lilypad/pkg/data"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/ipfs"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/bacalhau/pkg/models"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
	ApiPort string
}

type BacalhauExecutor struct {
	Options     BacalhauExecutorOptions
	bacalhauEnv []string

	bacalhauClient BacalhauClient
	ipfsClient     *ipfs.Client
}

func NewBacalhauExecutor(options BacalhauExecutorOptions, ipfsClient *ipfs.Client) (*BacalhauExecutor, error) {

	apiHost := ""
	if options.ApiHost != "DO_NOT_SET" {
		apiHost = fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost)
	}
	bacalhauEnv := []string{
		apiHost,
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)

	client, err := NewBacalhauClient(options.ApiHost)
	if err != nil {
		return nil, err
	}

	return &BacalhauExecutor{
		Options:        options,
		bacalhauEnv:    bacalhauEnv,
		bacalhauClient: *client,
		ipfsClient:     ipfsClient,
	}, nil
}

func (executor *BacalhauExecutor) Id() (string, error) {
	id, err := executor.bacalhauClient.GetID()
	if err != nil {
		return "", fmt.Errorf("error getting bacalhau ID %s", err.Error())
	}
	return id, nil
}

// Checks that Bacalhau is installed, correctly configured, and available
func (executor *BacalhauExecutor) IsAvailable() (bool, error) {
	alive, err := executor.bacalhauClient.Alive()
	if !alive || err != nil {
		return false, fmt.Errorf("Bacalhau is not currently available. Please ensure that Bacalhau is running, then try again. %w", err)
	}

	// Check that we have the right version of Bacalhau
	version, err := executor.bacalhauClient.GetVersion()
	if err != nil {
		return false, fmt.Errorf("error getting bacalhau version %s", err.Error())
	}
	// TODO: we may want to relax this
	if version != "v1.5.1" {
		return false, errors.New("Bacalhau version must be v1.5.1")
	}

	return true, nil
}

func (executor *BacalhauExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	return executor.bacalhauClient.GetMachineSpecs()
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	id, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	jobInfo, err := executor.bacalhauClient.GetJob(id)
	if err != nil {
		return nil, err
	}

	if len(jobInfo.Executions.Items) <= 0 {
		return nil, fmt.Errorf("no executions found for job %s", id)
	}

	// Check that the job completed successfully
	if jobInfo.Job.State.StateType != models.JobStateTypeCompleted {
		return nil, fmt.Errorf("job %s did not complete successfully: %s", id, jobInfo.Job.State.StateType.String())
	}

	// TODO! Directly get the results to solver?
	// system.EnsureDataDir(RESULTS_DIR)
	// resultsDir := system.GetDataDir(filepath.Join(RESULTS_DIR, deal.ID))
	// cidString := jobInfo.Executions.Items[0].PublishedResult
	// err = executor.ipfsClient.Get(context.Background(), cidString, resultsDir)
	// if err != nil {
	// 	return nil, fmt.Errorf("error getting results from IPFS %s -> %s", deal.ID, err)
	// }

	// // TODO: we should think about WASM and instruction count here
	// results := &executorlib.ExecutorResults{
	// 	ResultsDir:       resultsDir,
	// 	ResultsCID:       jobInfo.Executions.Items[0].PublishedResult,
	// 	InstructionCount: 1,
	// }

	return nil, nil
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

	putJobResponse, err := executor.bacalhauClient.PostJob(module.Job)
	if err != nil {
		return "", fmt.Errorf("error creating job %s -> %s", deal.ID, err.Error())
	}

	id := putJobResponse.JobID

	fmt.Printf("Got bacalhau job ID: %s\n", id)

	return id, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)
