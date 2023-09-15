package bacalhau

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/executor"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
}

type BacalhauExecutor struct {
	Options BacalhauExecutorOptions
}

func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {
	return &BacalhauExecutor{
		Options: options,
	}, nil
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (string, int, error) {
	bacalhauEnv := append(os.Environ(), fmt.Sprintf("BACALHAU_API_HOST=%s", executor.Options.ApiHost))
	runCmd := exec.Command(
		"bacalhau",
		"create",
		"--id-only",
		"--wait",
		"-",
	)
	runCmd.Env = bacalhauEnv
	stdin, err := runCmd.StdinPipe()
	if err != nil {
		return "", 1, fmt.Errorf("error getting stdin pipe for deal %s -> %s", deal.ID, err.Error())
	}
	stdout, err := runCmd.StdoutPipe()
	if err != nil {
		return "", 1, fmt.Errorf("error getting stdout pipe for deal %s -> %s", deal.ID, err.Error())
	}

	input, err := json.Marshal(module.Job)
	if err != nil {
		return "", 1, fmt.Errorf("error getting job JSON for deal %s -> %s", deal.ID, err.Error())
	}

	_, err = stdin.Write(input)
	if err != nil {
		return "", 1, fmt.Errorf("error writing job JSON %s -> %s", deal.ID, err.Error())
	}
	stdin.Close()

	jobIDOutput, err := io.ReadAll(stdout)
	if err != nil {
		return "", 1, fmt.Errorf("error reading stdout bytes from create job %s -> %s", deal.ID, err.Error())
	}

	err = runCmd.Wait()
	if err != nil {
		return "", 1, fmt.Errorf("error waiting for job to complete %s -> %s", deal.ID, err.Error())
	}

	id := strings.TrimSpace(string(jobIDOutput))

	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, id))
	if err != nil {
		return "", 1, fmt.Errorf("error creating a local folder of results %s -> %s", deal.ID, err.Error())
	}

	copyResultsCmd := exec.Command(
		"bacalhau",
		"get",
		id,
		"--output-dir", resultsDir,
	)
	copyResultsCmd.Env = bacalhauEnv

	_, err = copyResultsCmd.CombinedOutput()
	if err != nil {
		return "", 1, fmt.Errorf("error copying results %s -> %s", deal.ID, err.Error())
	}

	return resultsDir, 1, nil
}

// Compile-time interface check:
var _ executor.Executor = (*BacalhauExecutor)(nil)
