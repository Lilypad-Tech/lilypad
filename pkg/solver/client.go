package solver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

type SolverClient struct {
	options         http.ClientOptions
	solverEventSubs []func(SolverEvent)
}

func NewSolverClient(
	options http.ClientOptions,
) (*SolverClient, error) {
	client := &SolverClient{
		options:         options,
		solverEventSubs: []func(SolverEvent){},
	}
	return client, nil
}

// connect the websocket to the solver server
func (client *SolverClient) Start(ctx context.Context, cm *system.CleanupManager) error {
	websocketEventChannel := make(chan []byte)
	go func() {
		for {
			select {
			case evBytes := <-websocketEventChannel:
				// parse the ev into a new SolverEvent
				var ev SolverEvent
				if err := json.Unmarshal(evBytes, &ev); err != nil {
					log.Error().Msgf("Error unmarshalling event: %s", err.Error())
					continue
				}
				fmt.Printf("got websocket event: %+v\n", ev)
				// loop over each event channel and write the event to it
				for _, handler := range client.solverEventSubs {
					go handler(ev)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	http.ConnectWebSocket(
		http.WebsocketURL(client.options, http.WEBSOCKET_SUB_PATH),
		websocketEventChannel,
		ctx,
	)
	return nil
}

func (client *SolverClient) SubscribeEvents(handler func(SolverEvent)) {
	client.solverEventSubs = append(client.solverEventSubs, handler)
}

func (client *SolverClient) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	return http.GetRequest[[]data.JobOfferContainer](client.options, "/job_offers")
}

func (client *SolverClient) AddJobOffer(jobOffer data.JobOffer) (data.JobOfferContainer, error) {
	return http.PostRequest[data.JobOffer, data.JobOfferContainer](client.options, "/job_offers", jobOffer)
}

func (client *SolverClient) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	return http.GetRequest[[]data.ResourceOfferContainer](client.options, "/resource_offers")
}

func (client *SolverClient) AddResourceOffer(resourceOffer data.ResourceOffer) (data.ResourceOfferContainer, error) {
	return http.PostRequest[data.ResourceOffer, data.ResourceOfferContainer](client.options, "/resource_offers", resourceOffer)
}
