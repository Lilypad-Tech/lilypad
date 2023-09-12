package store

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/directory/store"
)

type DirectoryStoreMemory struct {
	deals   []data.Deal
	dealMap map[string]data.Deal
}

func NewDirectoryStoreMemory() (*DirectoryStoreMemory, error) {
	return &DirectoryStoreMemory{
		deals:   []data.Deal{},
		dealMap: map[string]data.Deal{},
	}, nil
}

func (s *DirectoryStoreMemory) AddDeal(deal data.Deal) (*data.Deal, error) {
	s.deals = append(s.deals, deal)
	s.dealMap[deal.ID] = deal
	return &deal, nil
}

func (s *DirectoryStoreMemory) GetDeals(query store.GetDealsQuery) ([]data.Deal, error) {
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

func (s *DirectoryStoreMemory) GetDeal(id string) (*data.Deal, error) {
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return &deal, nil
}
