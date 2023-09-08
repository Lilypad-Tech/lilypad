package store

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/directory/store"
)

type DealStoreMemory struct {
	deals   []data.Deal
	dealMap map[string]data.Deal
}

func NewDealStoreMemory() (*DealStoreMemory, error) {
	return &DealStoreMemory{
		deals:   []data.Deal{},
		dealMap: map[string]data.Deal{},
	}, nil
}

func (s *DealStoreMemory) AddDeal(deal data.Deal) (*data.Deal, error) {
	s.deals = append(s.deals, deal)
	s.dealMap[deal.ID] = deal
	return &deal, nil
}

func (s *DealStoreMemory) GetDeals(query store.GetDealsQuery) ([]data.Deal, error) {
	if query.JobCreator != "" {
		deals := []data.Deal{}
		for _, deal := range s.deals {
			if deal.JobCreator == query.JobCreator {
				deals = append(deals, deal)
			}
		}
		return deals, nil
	} else if query.ResourceProvider != "" {
		deals := []data.Deal{}
		for _, deal := range s.deals {
			if deal.ResourceProvider == query.ResourceProvider {
				deals = append(deals, deal)
			}
		}
		return deals, nil
	}
	return s.deals, nil
}

func (s *DealStoreMemory) GetDeal(id string) (*data.Deal, error) {
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return &deal, nil
}
