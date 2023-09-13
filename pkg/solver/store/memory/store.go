package store

import (
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
	mutex            sync.Mutex
}

func NewSolverStoreMemory() (*SolverStoreMemory, error) {
	return &SolverStoreMemory{
		jobOffers:        []data.JobOfferContainer{},
		jobOfferMap:      map[string]data.JobOfferContainer{},
		resourceOffers:   []data.ResourceOfferContainer{},
		resourceOfferMap: map[string]data.ResourceOfferContainer{},
		deals:            []data.DealContainer{},
		dealMap:          map[string]data.DealContainer{},
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

func (s *SolverStoreMemory) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
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
	if query.JobCreator != "" {
		deals := []data.DealContainer{}
		for _, deal := range s.deals {
			if deal.JobCreator == query.JobCreator {
				deals = append(deals, deal)
			}
		}
		return deals, nil
	} else if query.ResourceProvider != "" {
		deals := []data.DealContainer{}
		for _, deal := range s.deals {
			if deal.ResourceProvider == query.ResourceProvider {
				deals = append(deals, deal)
			}
		}
		return deals, nil
	}
	return s.deals, nil
}

func (s *SolverStoreMemory) GetJobOffer(id string) (*data.JobOfferContainer, error) {
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	return &jobOffer, nil
}

func (s *SolverStoreMemory) GetResourceOffer(id string) (*data.ResourceOfferContainer, error) {
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) GetDeal(id string) (*data.DealContainer, error) {
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return &deal, nil
}

func (s *SolverStoreMemory) UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	jobOffer.DealID = dealID
	jobOffer.State = state
	return &jobOffer, nil
}

func (s *SolverStoreMemory) UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	resourceOffer.DealID = dealID
	resourceOffer.State = state
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) UpdateDealState(id string, state uint8) (*data.DealContainer, error) {
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	deal.State = state
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
