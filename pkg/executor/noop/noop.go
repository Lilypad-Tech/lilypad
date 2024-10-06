package noop

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/lilypad-tech/lilypad/pkg/data"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/system"
)

const RESULTS_DIR = "noop-results"

type NoopExecutorOptions struct {
	Id               string
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

func NewNoopExecutorOptions() NoopExecutorOptions {
	return NoopExecutorOptions{
		Id:               uuid.NewString(),
		BadActor:         false,
		ResultsCID:       "123",
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

func (executor *NoopExecutor) Id() (string, error) {
	return executor.Options.Id, nil
}

func (executor *NoopExecutor) IsAvailable() (bool, error) {
	return true, nil
}

func (executor *NoopExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	return []data.MachineSpec{}, nil
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
