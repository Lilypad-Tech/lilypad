package store

import (
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
)

type SolverStoreDatabase struct {
	// TODO Remove placeholder which we needed to implement the interface with
	// package name shadowed
	placeholderDeal *data.DealContainer
}

func (store *SolverStoreDatabase) AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error) {
	return &jobOffer, nil
}

func (store *SolverStoreDatabase) AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error) {
	return &resourceOffer, nil
}

func (store *SolverStoreDatabase) AddDeal(deal data.DealContainer) (*data.DealContainer, error) {
	return &deal, nil
}

func (store *SolverStoreDatabase) AddResult(result data.Result) (*data.Result, error) {
	return &result, nil
}

func (store *SolverStoreDatabase) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error) {
	decision := &data.MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Deal:          deal,
		Result:        result,
	}
	return decision, nil
}

func (store *SolverStoreDatabase) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	jobOffers := []data.JobOfferContainer{}
	return jobOffers, nil
}

func (store *SolverStoreDatabase) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	resourceOffers := []data.ResourceOfferContainer{}
	return resourceOffers, nil
}

func (store *SolverStoreDatabase) GetDeals(query store.GetDealsQuery) ([]data.DealContainer, error) {
	deals := []data.DealContainer{}
	return deals, nil
}

func (store *SolverStoreDatabase) GetDealsAll() ([]data.DealContainer, error) {
	deals := []data.DealContainer{}
	return deals, nil
}

func (store *SolverStoreDatabase) GetJobOffer(id string) (*data.JobOfferContainer, error) {
	jobOffer := &data.JobOfferContainer{}
	return jobOffer, nil
}

func (store *SolverStoreDatabase) GetResourceOffer(id string) (*data.ResourceOfferContainer, error) {
	resourceOffer := &data.ResourceOfferContainer{}
	return resourceOffer, nil
}

func (store *SolverStoreDatabase) GetResourceOfferByAddress(address string) (*data.ResourceOfferContainer, error) {
	resourceOffer := &data.ResourceOfferContainer{}
	return resourceOffer, nil
}

func (store *SolverStoreDatabase) GetDeal(id string) (*data.DealContainer, error) {
	deal := &data.DealContainer{}
	return deal, nil
}

func (store *SolverStoreDatabase) GetResult(id string) (*data.Result, error) {
	result := &data.Result{}
	return result, nil
}

func (store *SolverStoreDatabase) GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error) {
	decision := &data.MatchDecision{}
	return decision, nil
}

func (store *SolverStoreDatabase) UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	jobOffer := &data.JobOfferContainer{}
	return jobOffer, nil
}

func (store *SolverStoreDatabase) UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	resourceOffer := &data.ResourceOfferContainer{}
	return resourceOffer, nil
}

func (store *SolverStoreDatabase) UpdateDealState(id string, state uint8) (*data.DealContainer, error) {
	deal := &data.DealContainer{}
	return deal, nil
}

func (store *SolverStoreDatabase) UpdateDealMediator(id string, mediator string) (*data.DealContainer, error) {
	deal := &data.DealContainer{}
	return deal, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	return store.placeholderDeal, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	return store.placeholderDeal, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error) {
	return store.placeholderDeal, nil
}

func (store *SolverStoreDatabase) RemoveJobOffer(id string) error {
	return nil
}

func (store *SolverStoreDatabase) RemoveResourceOffer(id string) error {
	return nil
}
