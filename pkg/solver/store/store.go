package store

import (
	"fmt"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
)

type StoreOptions struct {
	Type         string
	ConnStr      string
	GormLogLevel string
}

type GetJobOffersQuery struct {
	JobCreator string `json:"job_creator"`
	// this means job offers that have not been matched at all yet
	// the solver will use this to load only non matched resource offers

	// we use the DealID property of the jobOfferContainer to tell if it's been matched
	NotMatched bool `json:"not_matched"`

	// Active job offers are umatched or in an in-progress deal.
	// This includes the DealNegotiating, DealAgreed, or ResultsSubmitted states.
	Active bool `json:"in_progress"`

	// Cancelled job offers are in a JobOfferCancelled or a timed out state.
	// All job offers are included when Cancelled is nil.
	Cancelled *bool `json:"cancelled"`

	// Sort job offers oldest first
	OrderOldestFirst bool `json:"order_oldest_first"`
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

	// Sort resource offers oldest first
	OrderOldestFirst bool `json:"order_oldest_first"`

	PendingTesting bool `json:"pending_test"`
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
	AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error)
	AddDeal(deal data.DealContainer) (*data.DealContainer, error)
	AddResult(result data.Result) (*data.Result, error)
	AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error)
	AddBulkMatchDecisions(records []data.MatchDecision) error
	AddAllowedResourceProvider(resourceProvider string) (string, error)
	GetJobOffers(query GetJobOffersQuery) ([]data.JobOfferContainer, error)
	GetResourceOffers(query GetResourceOffersQuery) ([]data.ResourceOfferContainer, error)
	GetDeals(query GetDealsQuery) ([]data.DealContainer, error)
	GetDealsAll() ([]data.DealContainer, error)
	GetResults() ([]data.Result, error)
	GetMatchDecisions() ([]data.MatchDecision, error)
	GetAllowedResourceProviders() ([]string, error)
	GetJobOffer(id string) (*data.JobOfferContainer, error)
	GetResourceOffer(id string) (*data.ResourceOfferContainer, error)
	GetDeal(id string) (*data.DealContainer, error)
	GetResult(id string) (*data.Result, error)
	GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error)
	UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error)
	UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error)
	UpdateDealState(id string, state uint8) (*data.DealContainer, error)
	UpdateDealUploadTime(id string, timestamp int) (*data.DealContainer, error)
	UpdateDealDownloadTime(id string, timestamp int) (*data.DealContainer, error)
	UpdateDealMediator(id string, mediator string) (*data.DealContainer, error)
	UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error)
	UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error)
	UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error)
	RemoveJobOffer(id string) error
	RemoveResourceOffer(id string) error
	RemoveDeal(id string) error
	RemoveResult(id string) error
	RemoveMatchDecision(resourceOffer string, jobOffer string) error
	RemoveAllowedResourceProvider(resourceProvider string) error
}

func GetMatchID(resourceOffer string, jobOffer string) string {
	return fmt.Sprintf("%s-%s", resourceOffer, jobOffer)
}
