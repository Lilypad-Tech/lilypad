package solver

import (
	"context"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/solver/stats"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/metric"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type SolverOptions struct {
	Server            http.ServerOptions
	Store             store.StoreOptions
	Web3              web3.Web3Options
	Services          data.ServiceConfig
	Stats             stats.StatsOptions
	Telemetry         system.TelemetryOptions
	Metrics           system.MetricsOptions
	JobTimeoutSeconds int
}

type Solver struct {
	web3SDK    *web3.Web3SDK
	server     *solverServer
	controller *SolverController
	store      store.SolverStore
	stats      stats.Stats
	options    SolverOptions
	log        *zerolog.Logger
}

func NewSolver(
	options SolverOptions,
	store store.SolverStore,
	web3SDK *web3.Web3SDK,
	stats stats.Stats,
	tracer trace.Tracer,
	meter metric.Meter,
) (*Solver, error) {
	controller, err := NewSolverController(web3SDK, store, options, tracer, meter)
	if err != nil {
		return nil, err
	}
	server, err := NewSolverServer(options.Server, controller, store, stats, options.Services)
	if err != nil {
		return nil, err
	}
	solver := &Solver{
		controller: controller,
		store:      store,
		server:     server,
		web3SDK:    web3SDK,
		options:    options,
		log:        system.GetLogger(system.SolverService),
	}
	return solver, nil
}

func (solver *Solver) Start(ctx context.Context, cm *system.CleanupManager, tracerProvider *sdkTrace.TracerProvider) chan error {
	errorChan := solver.controller.Start(ctx, cm)
	solver.log.Debug().Msgf("solver.server.ListenAndServe")
	go func() {
		err := solver.server.ListenAndServe(ctx, cm, tracerProvider)
		if err != nil {
			errorChan <- err
		}
	}()
	return errorChan
}
