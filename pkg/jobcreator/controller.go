package jobcreator

import (
	"context"
	"fmt"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
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

/*
*
*
*

	Setup

*
*
*
*/
func (controller *JobCreatorController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(ev solver.SolverEvent) {
		solver.ServiceLogSolverEvent(system.JobCreatorService, ev)
		// we need to agree to the deal now we've heard about it
		if ev.EventType == solver.DealAdded {
			if ev.Deal == nil {
				system.Error(system.JobCreatorService, "solver event", fmt.Errorf("RP received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.JobCreator != controller.web3SDK.GetAddress().String() {
				return
			}
		}
	})
	return nil
}

func (controller *JobCreatorController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		system.Info(system.JobCreatorService, "StorageDealStateChange", ev)
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

	go func() {
		for {
			err := controller.solve()
			if err != nil {
				log.Error().Err(err).Msgf("error solving")
				errorChan <- err
				return
			}
			select {
			case <-time.After(1 * time.Second):
			case <-ctx.Done():
				return
			}
		}
	}()
	return errorChan
}

/*
 *
 *
 *

 Solve

 *
 *
 *
*/

func (controller *JobCreatorController) solve() error {
	system.Debug(system.JobCreatorService, "solving", "")
	return nil
}

/*
 *
 *
 *

 Data

 *
 *
 *
*/

func (controller *JobCreatorController) AddJobOffer(offer data.JobOffer) (data.JobOfferContainer, error) {
	return controller.solverClient.AddJobOffer(offer)
}
