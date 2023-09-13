package solver

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

// add an enum for various types of event
type SolverEventType int

const (
	JobOfferAdded SolverEventType = iota
	ResourceOfferAdded
	DealAdded
	JobOfferStateUpdated
	ResourceOfferStateUpdated
	DealStateUpdated
)

type SolverEvent struct {
	EventType     SolverEventType              `json:"event_type"`
	JobOffer      *data.JobOfferContainer      `json:"job_offer"`
	ResourceOffer *data.ResourceOfferContainer `json:"resource_offer"`
	Deal          *data.DealContainer          `json:"deal"`
}

type SolverController struct {
	web3SDK         *web3.Web3SDK
	web3Events      *web3.EventChannels
	store           store.SolverStore
	solverEventSubs []func(SolverEvent)
	options         SolverOptions
}

func NewSolverController(
	web3SDK *web3.Web3SDK,
	store store.SolverStore,
	options SolverOptions,
) (*SolverController, error) {
	controller := &SolverController{
		web3SDK:    web3SDK,
		web3Events: web3.NewEventChannels(),
		store:      store,
		options:    options,
	}
	return controller, nil
}

func (controller *SolverController) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error)
	// get the local subscriptions setup
	err := controller.subscribeToWeb3()
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

	// make sure we are registered as a solver
	// so that users can lookup our URL
	err = controller.registerAsSolver()
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

func (controller *SolverController) solve() error {

	// find out which deals we can make from matching the offers
	deals, err := getDeals(controller.store)
	if err != nil {
		return err
	}

	// loop over each of the deals add add them to the store and emit events
	for _, deal := range deals {
		_, err := controller.addDeal(deal)
		if err != nil {
			return err
		}
	}
	return nil
}

// * get the deal id
// * see if we have the deal locally
// * update the deal state locally
func (controller *SolverController) subscribeToWeb3() error {
	controller.web3Events.Storage.SubscribeDealStateChange(func(ev storage.StorageDealStateChange) {
		system.Info(system.SolverService, "StorageDealStateChange", ev)
	})
	return nil
}

// return a new event channel that will hear about events
// coming out of this controller
func (controller *SolverController) subscribeEvents(handler func(SolverEvent)) {
	controller.solverEventSubs = append(controller.solverEventSubs, handler)
}

// write the given event to all generated event channels
func (controller *SolverController) writeEvent(ev SolverEvent) {
	for _, handler := range controller.solverEventSubs {
		handler(ev)
	}
}

func (controller *SolverController) registerAsSolver() error {
	selfAddress := controller.web3SDK.GetAddress()
	solverType, err := data.GetServiceType("Solver")
	if err != nil {
		return err
	}

	selfUser, err := controller.web3SDK.GetUser(selfAddress)
	if err != nil {
		return err
	}

	// TODO: check the other props and call update if they have changed
	if selfUser.Url != controller.options.Server.URL {
		system.Info(system.SolverService, "url change", fmt.Sprintf("solver will be updated because URL has changed: %s %s != %s", selfAddress.String(), selfUser.Url, controller.options.Server.URL))
		err = controller.web3SDK.UpdateUser(
			big.NewInt(0),
			controller.options.Server.URL,
			[]uint8{solverType},
			[]common.Address{},
			[]common.Address{},
		)
		if err != nil {
			return err
		}
	} else {
		system.Info(system.SolverService, "url same", fmt.Sprintf("solver url already correct: %s %s", selfAddress.String(), controller.options.Server.URL))
	}

	existingSolvers, err := controller.web3SDK.GetSolverAddresses()
	if err != nil {
		return err
	}
	foundSolver := false
	for _, existingSolver := range existingSolvers {
		if existingSolver.String() == selfAddress.String() {
			system.Info(system.SolverService, "solver exists", selfAddress.String())
			foundSolver = true
			break
		}
	}
	if !foundSolver {
		system.Info(system.SolverService, "solver registering", "")
		// add the solver to the storage contract
		err = controller.web3SDK.AddUserToList(
			solverType,
		)
		if err != nil {
			return err
		}
		system.Info(system.SolverService, "solver registered", selfAddress.String())
	}
	return nil
}

func (controller *SolverController) addJobOffer(jobOffer data.JobOffer) (*data.JobOfferContainer, error) {
	id, err := data.GetJobOfferID(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOffer.ID = id

	system.Info(system.SolverService, "add job offer", jobOffer)

	ret, err := controller.store.AddJobOffer(getJobOfferContainer(jobOffer))
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: JobOfferAdded,
		JobOffer:  ret,
	})
	return ret, nil
}

func (controller *SolverController) addResourceOffer(resourceOffer data.ResourceOffer) (*data.ResourceOfferContainer, error) {
	id, err := data.GetResourceOfferID(resourceOffer)
	if err != nil {
		return nil, err
	}
	resourceOffer.ID = id

	system.Info(system.SolverService, "add resource offer", resourceOffer)

	ret, err := controller.store.AddResourceOffer(getResourceOfferContainer(resourceOffer))
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType:     ResourceOfferAdded,
		ResourceOffer: ret,
	})
	return ret, nil
}

func (controller *SolverController) addDeal(deal data.Deal) (*data.DealContainer, error) {
	id, err := data.GetDealID(deal)
	if err != nil {
		return nil, err
	}
	deal.ID = id

	system.Info(system.SolverService, "add deal", deal)

	ret, err := controller.store.AddDeal(getDealContainer(deal))
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: DealAdded,
		Deal:      ret,
	})
	_, err = controller.updateJobOfferState(ret.JobOffer, ret.ID, ret.State)
	if err != nil {
		return nil, err
	}
	_, err = controller.updateResourceOfferState(ret.ResourceOffer, ret.ID, ret.State)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (controller *SolverController) updateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	system.Info(system.SolverService, "update job offer", fmt.Sprintf("%s %s", id, data.GetAgreementStateString(state)))

	ret, err := controller.store.UpdateJobOfferState(id, dealID, state)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: JobOfferStateUpdated,
		JobOffer:  ret,
	})
	return ret, nil
}

func (controller *SolverController) updateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	system.Info(system.SolverService, "update resource offer", fmt.Sprintf("%s %s", id, data.GetAgreementStateString(state)))

	ret, err := controller.store.UpdateResourceOfferState(id, dealID, state)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType:     ResourceOfferStateUpdated,
		ResourceOffer: ret,
	})
	return ret, nil
}

// this will also update the job and resource offer states
func (controller *SolverController) updateDealState(id string, state uint8) (*data.DealContainer, error) {
	system.Info(system.SolverService, "update deal", fmt.Sprintf("%s %s", id, data.GetAgreementStateString(state)))

	ret, err := controller.store.UpdateDealState(id, state)
	if err != nil {
		return nil, err
	}
	controller.writeEvent(SolverEvent{
		EventType: DealStateUpdated,
		Deal:      ret,
	})
	_, err = controller.updateJobOfferState(ret.JobOffer, ret.ID, ret.State)
	if err != nil {
		return nil, err
	}
	_, err = controller.updateResourceOfferState(ret.ResourceOffer, ret.ID, ret.State)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//log.Info().Msgf("solver solving")

// // THIS IS JUST FOR TESTING
// log.Info().Msgf("sending tx")
// tx, err := controller.web3SDK.Contracts.Token.Transfer(
// 	controller.web3SDK.TransactOpts,
// 	common.HexToAddress("0x2546BcD3c84621e976D8185a91A922aE77ECEc30"),
// 	big.NewInt(1),
// )
// if err != nil {
// 	log.Info().Msgf("error sending tx: %s\n", err.Error())

// } else {
// 	log.Info().Msgf("tx sent: %s\n", tx.Hash())
// }
