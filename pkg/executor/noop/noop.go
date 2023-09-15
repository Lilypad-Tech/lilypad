package noop

import (
	"fmt"
	"path/filepath"

	"github.com/bacalhau-project/lilypad/pkg/data"
	executorlib "github.com/bacalhau-project/lilypad/pkg/executor"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

const RESULTS_DIR = "noop-results"

type NoopExecutorOptions struct {
	BadActor bool
}

type NoopExecutor struct {
	Options NoopExecutorOptions
}

func NewNoopExecutor(options NoopExecutorOptions) (*NoopExecutor, error) {
	return &NoopExecutor{
		Options: options,
	}, nil
}

func (executor *NoopExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, deal.ID))
	if err != nil {
		return nil, fmt.Errorf("error creating a local folder of results %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stdout"), []byte("Hello World!"))
	if err != nil {
		return nil, fmt.Errorf("error creating stdout file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stderr"), []byte(""))
	if err != nil {
		return nil, fmt.Errorf("error creating stderr file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "exitCode"), []byte("0"))
	if err != nil {
		return nil, fmt.Errorf("error creating exitCode file %s -> %s", deal.ID, err.Error())
	}
	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       "123",
		InstructionCount: 1,
	}
	return results, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*NoopExecutor)(nil)
