package mediator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/executor"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/module"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
)

type MediatorController struct {
	solverClient *solver.SolverClient
	options      MediatorOptions
	web3SDK      *web3.Web3SDK
	web3Events   *web3.EventChannels
	loop         *system.ControlLoop
	log          *system.ServiceLogger
	executor     executor.Executor
	// keep track of which jobs are running
	// this is because no remote state will change
	// whilst we are actually running a job
	runningJobsMutex sync.RWMutex
	runningJobs      map[string]bool
}

// the background "even if we have not heard of an event" loop
// i.e. things will not wait 10 seconds - the control loop
// reacts to events in the system - this 10 second background
// loop is just for in case we miss any events
const CONTROL_LOOP_INTERVAL = 10 * time.Second

func NewMediatorController(
	options MediatorOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
) (*MediatorController, error) {
	log.Debug().Msgf("begin NewMediatorController")
	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Services.Solver)
	if err != nil {
		log.Error().Msgf("error GetSolverUrl")
		return nil, err
	}

	solverClient, err := solver.NewSolverClient(http.ClientOptions{
		URL:        solverUrl,
		PrivateKey: options.Web3.PrivateKey,
	})
	if err != nil {
		log.Error().Msgf("error NewSolverClient")
		return nil, err
	}

	controller := &MediatorController{
		solverClient: solverClient,
		options:      options,
		web3SDK:      web3SDK,
		web3Events:   web3.NewEventChannels(),
		log:          system.NewServiceLogger(system.MediatorService),
		executor:     executor,
		runningJobs:  map[string]bool{},
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

// we subscribe to the solver because all this subscription will do is trigger
// a control loop which will load deals we've been assigned to from the solver
// we could of course subscribe directly to the web3 event from the mediation contract
// (and there is nothing preventing a client from doing this)
// however - we want to avoid a race condition between the mediator hearing the web3
// event and the then triggering a control loop against the solver
// before it's heard about the same event
//
// the worst that happens is a 10 second
// delay before the mediator control loop kicks in the background
// but the idea of event driven control loop triggering is for the system
// to run fast and so we will wait for the solver to trigger the event
// before we trigger our control loop (rambling comment of the day award goes to....)
func (controller *MediatorController) subscribeToSolver() error {
	controller.solverClient.SubscribeEvents(func(ev solver.SolverEvent) {
		// we need to agree to the deal now we've heard about it
		if ev.EventType == solver.DealMediatorUpdated {
			if ev.Deal == nil {
				controller.log.Error("solver event", fmt.Errorf("RP received nil deal"))
				return
			}

			// check if this deal is for us
			if ev.Deal.Mediator != controller.web3SDK.GetAddress().String() {
				return
			}

			solver.ServiceLogSolverEvent(system.MediatorService, ev)

			// trigger the solver
			controller.loop.Trigger()
		}
	})
	return nil
}

func (controller *MediatorController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error)
	// get the local subscriptions setup
	err := controller.subscribeToSolver()
	if err != nil {
		errorChan <- err
		return errorChan
	}

	err = controller.solverClient.Start(ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}
	// activate the web3 event listeners
	err = controller.web3Events.Start(controller.web3SDK, ctx, cm)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	controller.loop = system.NewControlLoop(
		system.ResourceProviderService,
		ctx,
		CONTROL_LOOP_INTERVAL,
		func() error {
			err := controller.solve()
			if err != nil {
				errorChan <- err
			}
			return err
		},
	)

	err = controller.loop.Start(true)
	if err != nil {
		errorChan <- err
		return errorChan
	}

	return errorChan
}

/*
*
*
*

# Solve

*
*
*
*/

func (controller *MediatorController) solve() error {
	controller.log.Debug("solving", "")

	// if there are jobs that have had both sides agree then we should run the job
	err := controller.runJobs()
	if err != nil {
		return err
	}
	return nil
}

/*
 *
 *
 *

 Run jobs

 *
 *
 *
*/

func (controller *MediatorController) runJobs() error {
	checkedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			Mediator: controller.web3SDK.GetAddress().String(),
			State:    "ResultsChecked",
		},
		func(dealContainer data.DealContainer) bool {
			controller.runningJobsMutex.RLock()
			defer controller.runningJobsMutex.RUnlock()
			_, ok := controller.runningJobs[dealContainer.ID]
			return !ok
		},
	)
	if err != nil {
		return err
	}
	if len(checkedDeals) <= 0 {
		return nil
	}
	// TODO - we need some kind of queue here
	// the reource provider is capacity managing by the rate
	// at which it's posting offers to the solver
	// but here we have no such rate limiting
	for _, dealContainer := range checkedDeals {
		func() {
			controller.runningJobsMutex.Lock()
			defer controller.runningJobsMutex.Unlock()
			controller.runningJobs[dealContainer.ID] = true
		}()

		go controller.runJob(dealContainer)
	}

	return err
}

func (controller *MediatorController) runJob(deal data.DealContainer) {
	controller.log.Info("mediator run job", deal)
	mediatorResult := data.Result{
		DealID: deal.ID,
		Error:  "",
	}
	err := func() error {
		module, err := module.LoadModule(deal.Deal.JobOffer.Module, deal.Deal.JobOffer.Inputs)
		if err != nil {
			return fmt.Errorf("error loading module: %s", err.Error())
		}
		executorResult, err := controller.executor.RunJob(deal, *module)
		if err != nil {
			return fmt.Errorf("error running job: %s", err.Error())
		}
		mediatorResult.InstructionCount = uint64(executorResult.InstructionCount)
		mediatorResult.DataID = executorResult.ResultsCID

		return nil
	}()

	if err != nil {
		mediatorResult.Error = err.Error()
	}

	// we should have the same result as the resource provider posted to the solver
	// so before we make a decision - let's load the result that the RP posted
	rpResult, err := controller.solverClient.GetResult(deal.ID)
	if err != nil {
		controller.log.Error("error loading existing result for deal", err)
		return
	}

	isResultCorrect := true

	if rpResult.DataID != mediatorResult.DataID {
		controller.log.Info("mediation data results different", fmt.Sprintf("deal %s, mediator: %s, rp: %s", deal.ID, mediatorResult.DataID, rpResult.DataID))
		isResultCorrect = false
	}

	if rpResult.InstructionCount != mediatorResult.InstructionCount {
		controller.log.Info("mediation instruction count different", fmt.Sprintf("deal %s, mediator: %d, rp: %d", deal.ID, mediatorResult.InstructionCount, rpResult.InstructionCount))
		isResultCorrect = false
	}

	if isResultCorrect {
		txHash, err := controller.web3SDK.MediationAcceptResult(
			deal.Deal.ID,
		)
		if err != nil {
			controller.log.Error("error calling mediation accept result tx for job", err)
			return
		}

		_, err = controller.solverClient.UpdateTransactionsMediator(deal.ID, data.DealTransactionsMediator{
			MediationAcceptResult: txHash,
		})
		if err != nil {
			controller.log.Error("error adding mediation accept result tx hash for deal", err)
			return
		}
	} else {
		txHash, err := controller.web3SDK.MediationRejectResult(
			deal.Deal.ID,
		)
		if err != nil {
			controller.log.Error("error calling mediation reject result tx for job", err)
			return
		}

		_, err = controller.solverClient.UpdateTransactionsMediator(deal.ID, data.DealTransactionsMediator{
			MediationRejectResult: txHash,
		})
		if err != nil {
			controller.log.Error("error adding mediation reject result tx hash for deal", err)
			return
		}
	}
}
