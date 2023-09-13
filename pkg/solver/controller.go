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
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

// add an enum for various types of event
type SolverEventType int

const (
	JobOfferAdded SolverEventType = iota
	ResourceOfferAdded
	JobOfferUpdated
	ResourceOfferUpdated
	MatchFound
)

type SolverEvent struct {
	EventType     SolverEventType              `json:"event_type"`
	JobOffer      *data.JobOfferContainer      `json:"job_offer"`
	ResourceOffer *data.ResourceOfferContainer `json:"resource_offer"`
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
	log.Debug().Msgf("Solver solving")
	_, err := getMatches(controller)
	if err != nil {
		return err
	}

	return nil
}

func (controller *SolverController) subscribeToWeb3() error {
	controller.web3Events.Token.SubscribeTransfer(func(event token.TokenTransfer) {
		log.Info().Msgf("solver Transfer. From: %s, Value: %d", event.From.Hex(), event.Value)
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
		go handler(ev)
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
		log.Info().Msgf("solver will be updated because URL has changed: %s %s != %s", selfAddress.String(), selfUser.Url, controller.options.Server.URL)
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
		log.Info().Msgf("solver url already correct: %s %s", selfAddress.String(), controller.options.Server.URL)
	}

	existingSolvers, err := controller.web3SDK.GetSolverAddresses()
	if err != nil {
		return err
	}
	foundSolver := false
	for _, existingSolver := range existingSolvers {
		if existingSolver.String() == selfAddress.String() {
			log.Info().Msgf("solver has already been listed: %s", selfAddress.String())
			foundSolver = true
			break
		}
	}
	if !foundSolver {
		log.Info().Msgf("solver has not been listed: %s", selfAddress.String())
		// add the solver to the storage contract
		err = controller.web3SDK.AddUserToList(
			solverType,
		)
		if err != nil {
			return err
		}
		log.Info().Msgf("solver now registered: %s", selfAddress.String())
	}
	return nil
}

func (controller *SolverController) addJobOffer(jobOffer data.JobOffer) (*data.JobOfferContainer, error) {
	jobOffer.ID = ""
	id, err := data.CalculateCID(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOffer.ID = id

	log.Info().
		Str("solver add job offer", fmt.Sprintf("%+v", jobOffer)).
		Msgf("")

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
	resourceOffer.ID = ""
	id, err := data.CalculateCID(resourceOffer)
	if err != nil {
		return nil, err
	}
	resourceOffer.ID = id

	log.Info().
		Str("solver add resource offer", fmt.Sprintf("%+v", resourceOffer)).
		Msgf("")

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
