package solver

import (
	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type SolverOptions struct {
	Web3   web3.Web3Options
	Server server.ServerOptions
}

type Solver struct {
	// Contract contract.Contract
	// Store    store.Store
}

func NewSolver(
	options SolverOptions,
) (*Solver, error) {
	solver := &Solver{
		// Contract: options.Contract,
		// Store:    options.Store,
	}
	return solver, nil
}
