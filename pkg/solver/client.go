package solver

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
)

type SolverClient struct {
	options http.ClientOptions
}

func NewSolverClient(
	options http.ClientOptions,
) (*SolverClient, error) {
	client := &SolverClient{
		options: options,
	}
	return client, nil
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
