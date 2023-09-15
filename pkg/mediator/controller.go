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

	// we know the address of the solver but what is it's url?
	solverUrl, err := web3SDK.GetSolverUrl(options.Services.Solver)
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

// list the deals we have been assigned to that we have not yet posted and agree tx to the contract for
func (controller *MediatorController) runJobs() error {

	// load the deals that are in DealNegotiating
	// and do not have a TransactionsResourceProvider.Agree tx
	agreedDeals, err := controller.solverClient.GetDealsWithFilter(
		store.GetDealsQuery{
			ResourceProvider: controller.web3SDK.GetAddress().String(),
			State:            "DealAgreed",
		},
		// this is where the solver has found us a match and we need to agree to it
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
	if len(agreedDeals) <= 0 {
		return nil
	}

	// TODO - we are relying on the rate at which we post resource offers
	// as our capacity manager right now
	// this will work because we only post resource offers as fast we handle jobs
	// but it would be worth putting some kind of queue here that is also aware
	// of the underlying capacity of the machine

	// map over the deals and run them
	for _, dealContainer := range agreedDeals {
		func() {
			controller.runningJobsMutex.Lock()
			defer controller.runningJobsMutex.Unlock()
			controller.runningJobs[dealContainer.ID] = true
		}()

		go controller.runJob(dealContainer)
	}

	return err
}

// this is run in it's own go-routine
// we've already updated controller.runningJobs so we know this will only
// run once
func (controller *MediatorController) runJob(deal data.DealContainer) {
	controller.log.Info("run job", deal)
	result := data.Result{
		DealID: deal.ID,
		Error:  nil,
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
		result.InstructionCount = uint64(executorResult.InstructionCount)
		result.DataID = executorResult.ResultsCID

		controller.log.Info(fmt.Sprintf("uploading results: %s %s %s", deal.ID, executorResult.ResultsDir, executorResult.ResultsCID), executorResult.ResultsDir)

		// upload the tarball to the solver service
		// TODO: we need some kind of on-chain attestation that the solver has the results
		uploadedResult, err := controller.solverClient.UploadResultFiles(deal.ID, executorResult.ResultsDir)
		if err != nil {
			return fmt.Errorf("error uploading results: %s", err.Error())
		}

		result.DataID = uploadedResult.DataID

		return nil
	}()

	// if this error is defined then it is probably the fault of the job not us
	// and we expect a mediator to get the same error
	if err != nil {
		result.Error = err
	}

	// the tarball of the results has been uploaded
	// now let's post the result data itself to the solver
	// then we will post the results on-chain
	createdResult, err := controller.solverClient.AddResult(result)
	if err != nil {
		// TODO: what should we do here?
		// the current path would be the post results times out
		// and the JC can claim a refund
		// but it's not really the fault of the RP that the solver refused to upload the results
		controller.log.Error("error posting result", err)
		return
	}

	txHash, err := controller.web3SDK.AddResult(
		deal.Deal.ID,
		createdResult.ID,
		result.InstructionCount,
	)
	if err != nil {
		controller.log.Error("error calling add result tx for job", err)
		return
	}

	_, err = controller.solverClient.UpdateTransactionsResourceProvider(deal.ID, data.DealTransactionsResourceProvider{
		AddResult: txHash,
	})
	if err != nil {
		// TODO: we need a way of deciding based on certain classes of error what happens
		// some will be retryable - otherwise will be fatal
		// we need a way to exit a job loop as a baseline
		controller.log.Error("error adding add result tx hash for deal", err)
		return
	}
}
