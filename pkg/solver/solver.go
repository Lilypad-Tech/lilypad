package solver

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
)

type SolverOptions struct {
	Web3   web3.Web3Options
	Server http.ServerOptions
}

type Solver struct {
	web3SDK    *web3.Web3SDK
	server     *solverServer
	controller *SolverController
	store      store.SolverStore
	options    SolverOptions
}

func NewSolver(
	options SolverOptions,
	store store.SolverStore,
	web3SDK *web3.Web3SDK,
) (*Solver, error) {
	controller, err := NewSolverController(web3SDK, store, options)
	if err != nil {
		return nil, err
	}
	server, err := NewSolverServer(options.Server, controller, store)
	if err != nil {
		return nil, err
	}
	solver := &Solver{
		controller: controller,
		store:      store,
		server:     server,
		web3SDK:    web3SDK,
		options:    options,
	}
	return solver, nil
}

func (solver *Solver) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := solver.controller.Start(ctx, cm)
	log.Debug().Msgf("solver.server.ListenAndServe")
	go func() {
		err := solver.server.ListenAndServe(ctx, cm)
		if err != nil {
			errorChan <- err
		}
	}()
	return errorChan
}
