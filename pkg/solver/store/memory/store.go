package store

import "github.com/bacalhau-project/lilypad/pkg/data"

type SolverStoreMemory struct {
	jobOffers        []data.JobOffer
	resourceOffers   []data.ResourceOffer
	jobOfferMap      map[string]data.JobOffer
	resourceOfferMap map[string]data.ResourceOffer
}

func NewSolverStoreMemory() *SolverStoreMemory {
	return &SolverStoreMemory{}
}

func (s *SolverStoreMemory) AddJobOffer(jobOffer data.JobOffer) error {
	id, err := data.CalculateCID(jobOffer)
	if err != nil {
		return err
	}
	s.jobOffers = append(s.jobOffers, jobOffer)
	s.jobOfferMap[id] = jobOffer
	return nil
}

func (s *SolverStoreMemory) AddResourceOffer(resourceOffer data.ResourceOffer) error {
	id, err := data.CalculateCID(resourceOffer)
	if err != nil {
		return err
	}
	s.resourceOffers = append(s.resourceOffers, resourceOffer)
	s.resourceOfferMap[id] = resourceOffer
	return nil
}

func (s *SolverStoreMemory) GetJobOffers() ([]data.JobOffer, error) {
	return s.jobOffers, nil
}

func (s *SolverStoreMemory) GetResourceOffers() ([]data.ResourceOffer, error) {
	return s.resourceOffers, nil
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

func (s *SolverStoreMemory) RemoveJobOffer(id string) error {
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
