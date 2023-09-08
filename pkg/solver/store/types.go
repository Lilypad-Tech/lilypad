package store

import "github.com/bacalhau-project/lilypad/pkg/data"

type SolverStore interface {
	AddJobOffer(jobOffer data.JobOffer) error
	AddResourceOffer(jobOffer data.ResourceOffer) error
	GetJobOffers() ([]data.JobOffer, error)
	GetResourceOffers() ([]data.ResourceOffer, error)
	GetJobOffer(id string) (data.JobOffer, error)
	GetResourceOffer(id string) (data.ResourceOffer, error)
	RemoveJobOffer(id string) error
	RemoveResourceOffer(id string) error
}
