package jobcreator

import (
	"context"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/rs/zerolog/log"
)

type JobCreatorController struct {
	solverClient *solver.SolverClient
	options      JobCreatorOptions
	web3SDK      *web3.Web3SDK
	web3Events   *web3.EventChannels
}

func NewJobCreatorController(
	options JobCreatorOptions,
	web3SDK *web3.Web3SDK,
) (*JobCreatorController, error) {
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Web3.SolverAddress)
	if err != nil {
		return nil, err
	}

	solverClient, err := solver.NewSolverClient(http.ClientOptions{
		URL:        solverUrl,
		PrivateKey: options.Web3.PrivateKey,
	})
	if err != nil {
		return nil, err
	}

	controller := &JobCreatorController{
		solverClient: solverClient,
		options:      options,
		web3SDK:      web3SDK,
		web3Events:   web3.NewEventChannels(),
	}
	return controller, nil
}

func (controller *JobCreatorController) solve() error {
	log.Info().Msgf("solving")
	return nil
}

func (controller *JobCreatorController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(event solver.SolverEvent) {
		log.Info().Msgf("New solver event %+v", event)
	})
	return nil
}

func (controller *JobCreatorController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("New SubscribeTransfer. From: %s, Value: %d", event.From.Hex(), event.Value)
	})
	return nil
}

func (controller *JobCreatorController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error)
	err := controller.subscribeToSolver()
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.subscribeToWeb3()
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.solverClient.Start(ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := controller.solve()
				if err != nil {
					log.Error().Err(err).Msgf("error solving")
					errorChan <- err
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return errorChan
}
