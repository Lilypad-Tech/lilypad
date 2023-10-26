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
	BadActor         bool
	ResultsCID       string
	Stdout           string
	Stderr           string
	ExitCode         string
	InstructionCount int
}

type NoopExecutor struct {
	Options NoopExecutorOptions
}

const NOOP_RESULTS_CID = "123"

func NewNoopExecutorOptions() NoopExecutorOptions {
	return NoopExecutorOptions{
		BadActor:         false,
		ResultsCID:       NOOP_RESULTS_CID,
		Stdout:           "Hello World!",
		Stderr:           "",
		ExitCode:         "0",
		InstructionCount: 1,
	}
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
	err = system.WriteFile(filepath.Join(resultsDir, "stdout"), []byte(executor.Options.Stdout))
	if err != nil {
		return nil, fmt.Errorf("error creating stdout file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stderr"), []byte(executor.Options.Stdout))
	if err != nil {
		return nil, fmt.Errorf("error creating stderr file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "exitCode"), []byte(executor.Options.ExitCode))
	if err != nil {
		return nil, fmt.Errorf("error creating exitCode file %s -> %s", deal.ID, err.Error())
	}
	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       executor.Options.ResultsCID,
		InstructionCount: executor.Options.InstructionCount,
	}
	return results, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*NoopExecutor)(nil)
