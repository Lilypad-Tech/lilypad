package store

import "github.com/bacalhau-project/lilypad/pkg/data"

type GetJobOffersQuery struct {
	JobCreator string `json:"job_creator"`
	// this means job offers that have not been matched at all yet
	// the solver will use this to load only non matched resource offers

	// we use the DealID property of the jobOfferContainer to tell if it's been matched
	NotMatched bool `json:"not_matched"`
}

type GetResourceOffersQuery struct {
	ResourceProvider string `json:"resource_provider"`

	// this means "currently occupied" - any free floating resource offers count
	// as active (because they could be matched any moment)
	// any resource offers of the following states are considered active:
	// - DealNegotiating
	// - DealAgreed
	// if we hit results submitted (or anything after that point)
	// then the resource offer is no longer considered active
	// (because the compute side is done and now we are onto payment & mediation)
	// this flag is used by the resource provider to ask "give me all my active resource offers"
	// so that it knows when to post more reosurce offers to the solver
	Active bool `json:"active"`

	// this means resource offers that have not been matched at all yet
	// the solver will use this to load only non matched resource offers

	// we use the DealID property of the resourceOfferContainer to tell if it's been matched
	NotMatched bool `json:"not_matched"`
}

type GetDealsQuery struct {
	JobCreator       string `json:"job_creator"`
	ResourceProvider string `json:"resource_provider"`
	Mediator         string `json:"mediator"`

	// only deals that are in this state will be returned
	State string `json:"state"`
}

type SolverStore interface {
	AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error)
	AddResourceOffer(jobOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error)
	AddDeal(deal data.DealContainer) (*data.DealContainer, error)
	AddResult(result data.Result) (*data.Result, error)
	AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error)
	GetJobOffers(query GetJobOffersQuery) ([]data.JobOfferContainer, error)
	GetResourceOffers(query GetResourceOffersQuery) ([]data.ResourceOfferContainer, error)
	GetDeals(query GetDealsQuery) ([]data.DealContainer, error)
	GetJobOffer(id string) (*data.JobOfferContainer, error)
	GetResourceOffer(id string) (*data.ResourceOfferContainer, error)
	GetDeal(id string) (*data.DealContainer, error)
	GetResult(id string) (*data.Result, error)
	GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error)
	UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error)
	UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error)
	UpdateDealState(id string, state uint8) (*data.DealContainer, error)
	UpdateDealMediator(id string, mediator string) (*data.DealContainer, error)
	UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error)
	UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error)
	UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error)
	RemoveJobOffer(id string) error
	RemoveResourceOffer(id string) error
}
