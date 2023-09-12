package options

import "github.com/bacalhau-project/lilypad/pkg/solver"

func NewSolverOptions() solver.SolverOptions {
	return solver.SolverOptions{
		Server: GetDefaultServerOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
}
