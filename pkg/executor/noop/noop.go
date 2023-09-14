package noop

import (
	"fmt"
	"path/filepath"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/executor"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

const RESULTS_DIR = "noop-results"

type NoopExecutor struct {
}

func NewNoopExecutor() (*NoopExecutor, error) {
	return &NoopExecutor{}, nil
}

func (executor *NoopExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (string, int, error) {
	resultsDir, err := system.DataDir(filepath.Join(RESULTS_DIR, deal.ID))
	if err != nil {
		return "", 0, fmt.Errorf("error creating a local folder of results %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stdout"), []byte("Hello World!"))
	if err != nil {
		return "", 0, fmt.Errorf("error creating stdout file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stderr"), []byte(""))
	if err != nil {
		return "", 0, fmt.Errorf("error creating stderr file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "exitCode"), []byte("0"))
	if err != nil {
		return "", 0, fmt.Errorf("error creating exitCode file %s -> %s", deal.ID, err.Error())
	}
	return resultsDir, 1, nil
}

// Compile-time interface check:
var _ executor.Executor = (*NoopExecutor)(nil)
