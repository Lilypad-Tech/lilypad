package store

import (
	"sync"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
)

type SolverStoreMemory struct {
	jobOffers        []data.JobOffer
	resourceOffers   []data.ResourceOffer
	jobOfferMap      map[string]data.JobOffer
	resourceOfferMap map[string]data.ResourceOffer
	matches          []data.Match
	matchMap         map[string]data.Match
	mutex            sync.Mutex
}

func NewSolverStoreMemory() (*SolverStoreMemory, error) {
	return &SolverStoreMemory{
		jobOffers:        []data.JobOffer{},
		resourceOffers:   []data.ResourceOffer{},
		matches:          []data.Match{},
		jobOfferMap:      map[string]data.JobOffer{},
		resourceOfferMap: map[string]data.ResourceOffer{},
		matchMap:         map[string]data.Match{},
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

func (s *SolverStoreMemory) AddMatch(match data.Match) (*data.Match, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.matches = append(s.matches, match)
	s.matchMap[match.ID] = match
	return &match, nil
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

func (s *SolverStoreMemory) GetMatches(query store.GetMatchesQuery) ([]data.Match, error) {
	if query.JobCreator != "" && query.ResourceProvider != "" {
		matches := []data.Match{}
		for _, match := range s.matches {
			if match.JobCreator == query.JobCreator && match.ResourceProvider == query.ResourceProvider {
				matches = append(matches, match)
			}
		}
		return matches, nil
	}
	return s.matches, nil
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

func (s *SolverStoreMemory) GetMatch(id string) (*data.Match, error) {
	match, ok := s.matchMap[id]
	if !ok {
		return nil, nil
	}
	return &match, nil
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

func (s *SolverStoreMemory) RemoveMatch(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	newMatches := []data.Match{}
	for _, match := range s.matches {
		matchId, err := data.CalculateCID(match)
		if err != nil {
			return err
		}
		if matchId != id {
			newMatches = append(newMatches, match)
		} else {
			delete(s.matchMap, id)
		}
	}
	s.matches = newMatches
	return nil
}
