package executor

import (
	"github.com/lilypad-tech/lilypad/pkg/data"
)

type ExecutorResults struct {
	ResultsDir       string
	ResultsCID       string
	InstructionCount int
}

type Executor interface {
	Id() (string, error)
	IsAvailable() (bool, error)
	GetMachineSpecs() ([]data.MachineSpec, error)
	// run the given job and return a local folder
	// that contains the results
	RunJob(
		deal data.DealContainer,
		module data.Module,
	) (*ExecutorResults, error)
}
