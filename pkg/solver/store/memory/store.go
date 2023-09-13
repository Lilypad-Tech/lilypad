package store

import (
	"sync"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
)

type SolverStoreMemory struct {
	jobOffers        []data.JobOffer
	jobOfferMap      map[string]data.JobOffer
	resourceOffers   []data.ResourceOffer
	resourceOfferMap map[string]data.ResourceOffer
	deals            []data.Deal
	dealMap          map[string]data.Deal
	mutex            sync.Mutex
}

func NewSolverStoreMemory() (*SolverStoreMemory, error) {
	return &SolverStoreMemory{
		jobOffers:        []data.JobOffer{},
		jobOfferMap:      map[string]data.JobOffer{},
		resourceOffers:   []data.ResourceOffer{},
		resourceOfferMap: map[string]data.ResourceOffer{},
		deals:            []data.Deal{},
		dealMap:          map[string]data.Deal{},
	}, nil
}

func (s *SolverStoreMemory) AddJobOffer(jobOffer data.JobOffer) (*data.JobOffer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.jobOffers = append(s.jobOffers, jobOffer)
	s.jobOfferMap[jobOffer.ID] = jobOffer
	return &jobOffer, nil
}

func (s *SolverStoreMemory) AddResourceOffer(resourceOffer data.ResourceOffer) (*data.ResourceOffer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resourceOffers = append(s.resourceOffers, resourceOffer)
	s.resourceOfferMap[resourceOffer.ID] = resourceOffer
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) AddDeal(deal data.Deal) (*data.Deal, error) {
	s.deals = append(s.deals, deal)
	s.dealMap[deal.ID] = deal
	return &deal, nil
}

func (s *SolverStoreMemory) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOffer, error) {
	if query.JobCreator != "" {
		jobOffers := []data.JobOffer{}
		for _, jobOffer := range s.jobOffers {
			if jobOffer.JobCreator == query.JobCreator {
				jobOffers = append(jobOffers, jobOffer)
			}
		}
		return jobOffers, nil
	}
	return s.jobOffers, nil
}

func (s *SolverStoreMemory) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOffer, error) {
	if query.ResourceProvider != "" {
		resourceOffers := []data.ResourceOffer{}
		for _, resourceOffer := range s.resourceOffers {
			if resourceOffer.ResourceProvider == query.ResourceProvider {
				resourceOffers = append(resourceOffers, resourceOffer)
			}
		}
		return resourceOffers, nil
	}
	return s.resourceOffers, nil
}

func (s *SolverStoreMemory) GetDeals(query store.GetDealsQuery) ([]data.Deal, error) {
	if query.JobCreator != "" {
		deals := []data.Deal{}
		for _, deal := range s.deals {
			if deal.Members.JobCreator == query.JobCreator {
				deals = append(deals, deal)
			}
		}
		return deals, nil
	} else if query.ResourceProvider != "" {
		deals := []data.Deal{}
		for _, deal := range s.deals {
			if deal.Members.ResourceProvider == query.ResourceProvider {
				deals = append(deals, deal)
			}
		}
		return deals, nil
	}
	return s.deals, nil
}

func (s *SolverStoreMemory) GetJobOffer(id string) (*data.JobOffer, error) {
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	return &jobOffer, nil
}

func (s *SolverStoreMemory) GetResourceOffer(id string) (*data.ResourceOffer, error) {
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) GetDeal(id string) (*data.Deal, error) {
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return &deal, nil
}

func (s *SolverStoreMemory) RemoveJobOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	newJobOffers := []data.JobOffer{}
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
	newResourceOffers := []data.ResourceOffer{}
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
