package store

import "github.com/bacalhau-project/lilypad/pkg/data"

type GetJobOffersQuery struct {
	JobCreator string `json:"job_creator"`
}

type GetResourceOffersQuery struct {
	ResourceProvider string `json:"resource_provider"`
}

type GetMatchesQuery struct {
	JobCreator       string `json:"job_creator"`
	ResourceProvider string `json:"resource_provider"`
}

type GetDealsQuery struct {
	JobCreator       string `json:"job_creator"`
	ResourceProvider string `json:"resource_provider"`
}

type SolverStore interface {
	AddJobOffer(jobOffer data.JobOffer) (*data.JobOffer, error)
	AddResourceOffer(jobOffer data.ResourceOffer) (*data.ResourceOffer, error)
	AddMatch(match data.Match) (*data.Match, error)
	AddDeal(deal data.Deal) (*data.Deal, error)
	GetJobOffers(query GetJobOffersQuery) ([]data.JobOffer, error)
	GetResourceOffers(query GetResourceOffersQuery) ([]data.ResourceOffer, error)
	GetMatches(query GetMatchesQuery) ([]data.Match, error)
	GetDeals(query GetDealsQuery) ([]data.Deal, error)
	GetJobOffer(id string) (*data.JobOffer, error)
	GetResourceOffer(id string) (*data.ResourceOffer, error)
	GetMatch(id string) (*data.Match, error)
	GetDeal(id string) (*data.Deal, error)
	RemoveJobOffer(id string) error
	RemoveResourceOffer(id string) error
	RemoveMatch(id string) error
}
