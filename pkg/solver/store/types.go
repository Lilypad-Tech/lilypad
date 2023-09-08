package store

import "github.com/bacalhau-project/lilypad/pkg/data"

type GetJobOffersQuery struct {
	JobCreator string `json:"job_creator"`
}

type GetResourceOffersQuery struct {
	ResourceProvider string `json:"resource_provider"`
}

type SolverStore interface {
	AddJobOffer(jobOffer data.JobOffer) (*data.JobOffer, error)
	AddResourceOffer(jobOffer data.ResourceOffer) (*data.ResourceOffer, error)
	GetJobOffers(query GetJobOffersQuery) ([]data.JobOffer, error)
	GetResourceOffers(query GetResourceOffersQuery) ([]data.ResourceOffer, error)
	GetJobOffer(id string) (*data.JobOffer, error)
	GetResourceOffer(id string) (*data.ResourceOffer, error)
	RemoveJobOffer(id string) error
	RemoveResourceOffer(id string) error
}
