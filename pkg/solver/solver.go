package solver

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type SolverOptions struct {
	Web3   web3.Web3Options
	Server server.ServerOptions
}

type Solver struct {
	Web3 *web3.ContractSDK
}

func NewSolver(
	options SolverOptions,
) (*Solver, error) {
	web3, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return nil, err
	}
	solver := &Solver{
		Web3: web3,
	}
	return solver, nil
}

func (solver *Solver) Start(ctx context.Context) error {
	return nil
}
