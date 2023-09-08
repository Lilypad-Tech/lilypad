package solver

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
)

type SolverClient struct {
	options server.ClientOptions
}

func NewSolverClient(
	options server.ClientOptions,
) (*SolverClient, error) {
	client := &SolverClient{
		options: options,
	}
	return client, nil
}

func (client *SolverClient) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOffer, error) {
	//url := fmt.Sprintf("%s/api/v1/job_offers", client.options.URL)
	// do a http request and get the response
	// parse the response and return the data
	return []data.JobOffer{}, nil
}
