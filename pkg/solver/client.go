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
	options             http.ClientOptions
	solverEventChannels []SolverEventChannel
}

func NewSolverClient(
	options http.ClientOptions,
) (*SolverClient, error) {
	client := &SolverClient{
		options: options,
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
				for _, eventChannel := range client.solverEventChannels {
					eventChannel <- ev
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (client *SolverClient) GetEventChannel() SolverEventChannel {
	eventChannel := make(SolverEventChannel)
	client.solverEventChannels = append(client.solverEventChannels, eventChannel)
	return eventChannel
}

func (client *SolverClient) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOffer, error) {
	return http.Get[[]data.JobOffer](client.options, "/job_offers")
}

func (client *SolverClient) AddJobOffer(jobOffer data.JobOffer) (data.JobOffer, error) {
	return http.Post[data.JobOffer, data.JobOffer](client.options, "/job_offers", jobOffer)
}

func (client *SolverClient) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOffer, error) {
	return http.Get[[]data.ResourceOffer](client.options, "/resource_offers")
}

func (client *SolverClient) AddResourceOffer(resourceOffer data.ResourceOffer) (data.ResourceOffer, error) {
	return http.Post[data.ResourceOffer, data.ResourceOffer](client.options, "/resource_offers", resourceOffer)
}
