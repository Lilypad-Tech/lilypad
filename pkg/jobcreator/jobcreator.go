package jobcreator

import (
	"context"
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type JobCreatorOptions struct {
	Web3 web3.Web3Options
}

type JobCreator struct {
	web3SDK    *web3.Web3SDK
	controller *JobCreatorController
}

func NewJobCreator(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
) (*JobCreator, error) {
	if options.Web3.SolverAddress == "" {
		return nil, fmt.Errorf("--web3-solver-address or WEB3_SOLVER_ADDRESS is empty")
	}
	controller, err := NewJobCreatorController(web3SDK)
	if err != nil {
		return nil, err
	}
	solver := &JobCreator{
		controller: controller,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (JobCreator *JobCreator) Start(ctx context.Context, cm *system.CleanupManager) error {
	return JobCreator.controller.Start(ctx, cm)
}
