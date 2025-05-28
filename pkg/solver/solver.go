package solver

import (
	"context"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/http"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/stats"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/metric"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type SolverTimeoutOptions struct {
	MatchSeconds     uint64 `json:"match_seconds"`
	ExecutionSeconds uint64 `json:"execution_seconds"`
	DownloadSeconds  uint64 `json:"download_seconds"`
	TotalSeconds     uint64 `json:"total_seconds"`
}

type SolverOptions struct {
	Server    http.ServerOptions
	Store     store.StoreOptions
	Web3      web3.Web3Options
	Services  data.ServiceConfig
	Stats     stats.StatsOptions
	Telemetry system.TelemetryOptions
	Metrics   system.MetricsOptions
	Logs      system.LogsOptions
	Timeouts  SolverTimeoutOptions
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
	versionConfig *system.VersionConfig,
) (*Solver, error) {
	controller, err := NewSolverController(web3SDK, store, stats, options, tracer, meter)
	if err != nil {
		return nil, err
	}
	server, err := NewSolverServer(options.Server, controller, store, stats, versionConfig, options.Services, meter)
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
