package solver

import (
	"context"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type SolverOptions struct {
	Web3      web3.Web3Options
	Server    http.ServerOptions
	Services  data.ServiceConfig
	Telemetry system.TelemetryOptions
	Metrics   system.MetricsOptions
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
	tracer trace.Tracer,
) (*Solver, error) {
	controller, err := NewSolverController(web3SDK, store, options, tracer)
	if err != nil {
		return nil, err
	}
	server, err := NewSolverServer(options.Server, controller, store, options.Services)
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

func (solver *Solver) Start(ctx context.Context, cm *system.CleanupManager, tracerProvider *sdkTrace.TracerProvider) chan error {
	errorChan := solver.controller.Start(ctx, cm)
	log.Debug().Msgf("solver.server.ListenAndServe")
	go func() {
		err := solver.server.ListenAndServe(ctx, cm, tracerProvider)
		if err != nil {
			errorChan <- err
		}
	}()
	return errorChan
}
