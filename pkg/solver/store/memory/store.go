package store

import (
	"fmt"
	"sync"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
)

type SolverStoreMemory struct {
	jobOffers        []data.JobOfferContainer
	jobOfferMap      map[string]data.JobOfferContainer
	resourceOffers   []data.ResourceOfferContainer
	resourceOfferMap map[string]data.ResourceOfferContainer
	deals            []data.DealContainer
	dealMap          map[string]data.DealContainer
	matchDecisionMap map[string]data.MatchDecision
	mutex            sync.RWMutex
}

func getMatchID(resourceOffer string, jobOffer string) string {
	return fmt.Sprintf("%s-%s", resourceOffer, jobOffer)
}

func NewSolverStoreMemory() (*SolverStoreMemory, error) {
	return &SolverStoreMemory{
		jobOffers:        []data.JobOfferContainer{},
		jobOfferMap:      map[string]data.JobOfferContainer{},
		resourceOffers:   []data.ResourceOfferContainer{},
		resourceOfferMap: map[string]data.ResourceOfferContainer{},
		deals:            []data.DealContainer{},
		dealMap:          map[string]data.DealContainer{},
		matchDecisionMap: map[string]data.MatchDecision{},
	}, nil
}

func (s *SolverStoreMemory) AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.jobOffers = append(s.jobOffers, jobOffer)
	s.jobOfferMap[jobOffer.ID] = jobOffer
	return &jobOffer, nil
}

func (s *SolverStoreMemory) AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resourceOffers = append(s.resourceOffers, resourceOffer)
	s.resourceOfferMap[resourceOffer.ID] = resourceOffer
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) AddDeal(deal data.DealContainer) (*data.DealContainer, error) {
	s.deals = append(s.deals, deal)
	s.dealMap[deal.ID] = deal
	return &deal, nil
}

func (s *SolverStoreMemory) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := getMatchID(resourceOffer, jobOffer)
	_, ok := s.matchDecisionMap[id]
	if ok {
		return nil, fmt.Errorf("that match already exists")
	}
	decision := data.MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Deal:          deal,
		Result:        result,
	}
	s.matchDecisionMap[id] = decision
	return &decision, nil
}

func (s *SolverStoreMemory) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	jobOffers := []data.JobOfferContainer{}
	for _, jobOffer := range s.jobOffers {
		matching := true
		if query.JobCreator != "" && jobOffer.JobCreator != query.JobCreator {
			matching = false
		}
		if query.NotMatched {
			if jobOffer.DealID != "" {
				matching = false
			}
		}
		if matching {
			jobOffers = append(jobOffers, jobOffer)
		}
	}
	return jobOffers, nil
}

func (s *SolverStoreMemory) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	resourceOffers := []data.ResourceOfferContainer{}
	for _, resourceOffer := range s.resourceOffers {
		matching := true
		if query.ResourceProvider != "" && resourceOffer.ResourceProvider != query.ResourceProvider {
			matching = false
		}
		if query.Active && !data.IsActiveAgreementState(resourceOffer.State) {
			matching = false
		}
		if query.NotMatched {
			if resourceOffer.DealID != "" {
				matching = false
			}
		}
		if matching {
			resourceOffers = append(resourceOffers, resourceOffer)
		}
	}
	return resourceOffers, nil
}

func (s *SolverStoreMemory) GetDeals(query store.GetDealsQuery) ([]data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	ret := s.deals
	if query.JobCreator != "" {
		filteredDeals := []data.DealContainer{}
		for _, deal := range ret {
			if deal.JobCreator == query.JobCreator {
				filteredDeals = append(filteredDeals, deal)
			}
		}
		ret = filteredDeals
	}
	if query.ResourceProvider != "" {
		filteredDeals := []data.DealContainer{}
		for _, deal := range ret {
			if deal.ResourceProvider == query.ResourceProvider {
				filteredDeals = append(filteredDeals, deal)
			}
		}
		ret = filteredDeals
	}
	if query.State != "" {
		state, err := data.GetAgreementState(query.State)
		if err != nil {
			return nil, err
		}
		filteredDeals := []data.DealContainer{}
		for _, deal := range ret {
			if deal.State == state {
				filteredDeals = append(filteredDeals, deal)
			}
		}
		ret = filteredDeals
	}
	return ret, nil
}

func (s *SolverStoreMemory) GetJobOffer(id string) (*data.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	return &jobOffer, nil
}

func (s *SolverStoreMemory) GetResourceOffer(id string) (*data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) GetDeal(id string) (*data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return &deal, nil
}

func (s *SolverStoreMemory) GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	id := getMatchID(resourceOffer, jobOffer)
	decision, ok := s.matchDecisionMap[id]
	if !ok {
		return nil, nil
	}
	return &decision, nil
}

func (s *SolverStoreMemory) UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	jobOffer.DealID = dealID
	jobOffer.State = state
	return &jobOffer, nil
}

func (s *SolverStoreMemory) UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	resourceOffer.DealID = dealID
	resourceOffer.State = state
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) UpdateDealState(id string, state uint8) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	deal.State = state
	return &deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	deal.TransactionsJobCreator = data
	return &deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	deal.TransactionsResourceProvider = data
	return &deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	deal.TransactionsMediator = data
	return &deal, nil
}

func (s *SolverStoreMemory) RemoveJobOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	newJobOffers := []data.JobOfferContainer{}
	for _, jobOffer := range s.jobOffers {
		jobOfferId, err := data.CalculateCID(jobOffer)
		if err != nil {
			return err
		}
		if jobOfferId != id {
			newJobOffers = append(newJobOffers, jobOffer)
		} else {
			delete(s.jobOfferMap, id)
		}
	}
	s.jobOffers = newJobOffers
	return nil
}

func (s *SolverStoreMemory) RemoveResourceOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	newResourceOffers := []data.ResourceOfferContainer{}
	for _, resourceOffer := range s.resourceOffers {
		resourceOfferId, err := data.CalculateCID(resourceOffer)
		if err != nil {
			return err
		}
		if resourceOfferId != id {
			newResourceOffers = append(newResourceOffers, resourceOffer)
		} else {
			delete(s.resourceOfferMap, id)
		}
	}
	s.resourceOffers = newResourceOffers
	return nil
}

// Compile-time interface check:
var _ store.SolverStore = (*SolverStoreMemory)(nil)
