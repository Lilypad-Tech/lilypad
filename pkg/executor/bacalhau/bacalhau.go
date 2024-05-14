package bacalhau

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/data/bacalhau"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/lilymetrics"
	"github.com/lilypad-tech/lilypad/pkg/system"
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
	span := lilymetrics.Trace(context.Background())
	defer span.End()
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
	time.Sleep(5 * time.Second)
	span := lilymetrics.Trace(context.Background())
	defer span.End()
	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, dealID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder of results %s -> %s", dealID, err.Error())
	}

	//bacalhau list --id-filter 2f5b845a-6ad3-41f1-969c-18a4209082e7 --output json
	info := exec.Command("bacalhau", "list", "--id-filter", jobID, "--output", "json")
	info.Env = executor.bacalhauEnv
	output, err := info.Output()
	fmt.Println("bacalhau list output ", string(output))

	var jobData []JobData
	err = json.Unmarshal([]byte(output), &jobData)
	if err != nil {
		fmt.Println("Error:", err)

	}
	fmt.Println("cid ", jobData[0].State.Executions[0].PublishedResults.CID)
	exec.Command("ipfs", "get", jobData[0].State.Executions[0].PublishedResults.CID, "-o", resultsDir).Run()
	return resultsDir, nil
	copyCmdText := fmt.Sprintf("bacalhau get %s --output-dir %s --ipfs-swarm-addrs==/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWAW3hUQmCnzRdja86pAy9S5dBFXMRvsVF2vyWvAD3vnj1 ", jobID, resultsDir)
	log.Debug().Msgf("executor.bacalhauEnv: %s", executor.bacalhauEnv)
	log.Debug().Msgf("Executing command: %s", copyCmdText) // Log the command before execution for debugging
	copyResultsCmd := exec.Command("bacalhau", "get", jobID, "--output-dir", resultsDir)
	copyResultsCmd.Env = executor.bacalhauEnv
	fmt.Println("solved job")
	stdout, err := copyResultsCmd.StdoutPipe()
	// if err != nil {
	// 	// log.Fatal(err)
	// }

	if err := copyResultsCmd.Start(); err != nil {
		//log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		r := scanner.Text()
		fmt.Println("bacalhau get result ", r)
		// if strings.Contains(r, "Job result:") {
		// 	words := strings.Fields(r)
		// 	lastWord := words[len(words)-1]
		// 	fmt.Println("JOB " + lastWord)
		// 	cmdstd := exec.Command("cat", "/tmp/lilypad/data/downloaded-files/"+lastWord+"/stdout")

		// 	output, err := cmdstd.Output()
		// 	fmt.Println(string(output))
		// 	server.BroadcastToNamespace("/", "reply", output, "data")
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		return true
		// 	}

		// 	break
		// } else {
		// server.BroadcastToNamespace("/", "reply", r, "data")
		// }

	}

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = copyResultsCmd.CombinedOutput()
	// if err != nil {
	// 	return "", fmt.Errorf("error copying results %s -> %s, command executed: %s", dealID, err.Error(), copyCmdText)
	// }
	// wait for the results to be copied
	return resultsDir, nil
}

func (executor *BacalhauExecutor) getJobState(dealID string, jobID string) (*bacalhau.JobWithInfo, error) {
	span := lilymetrics.Trace(context.Background())
	defer span.End()
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

type JobMetadata struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	ClientID  string    `json:"ClientID"`
	Requester struct {
		RequesterNodeID    string `json:"RequesterNodeID"`
		RequesterPublicKey string `json:"RequesterPublicKey"`
	} `json:"Requester"`
}

type JobSpec struct {
	Engine        string `json:"Engine"`
	Verifier      string `json:"Verifier"`
	Publisher     string `json:"Publisher"`
	PublisherSpec struct {
		Type string `json:"Type"`
	} `json:"PublisherSpec"`
	Docker struct {
		Image      string   `json:"Image"`
		Entrypoint []string `json:"Entrypoint"`
	} `json:"Docker"`
}

type JobResources struct {
	GPU string `json:"GPU"`
}

type JobNetwork struct {
	Type string `json:"Type"`
}

type JobDeal struct {
	Concurrency int `json:"Concurrency"`
}

type Job struct {
	APIVersion string       `json:"APIVersion"`
	Metadata   JobMetadata  `json:"Metadata"`
	Spec       JobSpec      `json:"Spec"`
	Resources  JobResources `json:"Resources"`
	Network    JobNetwork   `json:"Network"`
	Timeout    int          `json:"Timeout"`
	Deal       JobDeal      `json:"Deal"`
}

type JobExecution struct {
	JobID              string `json:"JobID"`
	NodeId             string `json:"NodeId"`
	ComputeReference   string `json:"ComputeReference"`
	State              string `json:"State"`
	AcceptedAskForBid  bool   `json:"AcceptedAskForBid"`
	VerificationResult struct {
		Complete bool `json:"Complete"`
		Result   bool `json:"Result"`
	} `json:"VerificationResult"`
	PublishedResults struct {
		StorageSource string `json:"StorageSource"`
		Name          string `json:"Name"`
		CID           string `json:"CID"`
	} `json:"PublishedResults"`
	RunOutput struct {
		Stdout          string `json:"stdout"`
		Stdouttruncated bool   `json:"stdouttruncated"`
		Stderr          string `json:"stderr"`
		Stderrtruncated bool   `json:"stderrtruncated"`
		ExitCode        int    `json:"exitCode"`
		RunnerError     string `json:"runnerError"`
	} `json:"RunOutput"`
	Version    int       `json:"Version"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

type JobState struct {
	JobID      string         `json:"JobID"`
	Executions []JobExecution `json:"Executions"`
	State      string         `json:"State"`
	Version    int            `json:"Version"`
	CreateTime time.Time      `json:"CreateTime"`
	UpdateTime time.Time      `json:"UpdateTime"`
	TimeoutAt  time.Time      `json:"TimeoutAt"`
}

type JobData struct {
	Job   Job      `json:"Job"`
	State JobState `json:"State"`
}
