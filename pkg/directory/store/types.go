package store

import "github.com/bacalhau-project/lilypad/pkg/data"

type GetDealsQuery struct {
	JobCreator       string `json:"job_creator"`
	ResourceProvider string `json:"resource_provider"`
}

type DirectoryStore interface {
	AddDeal(deal data.Deal) (*data.Deal, error)
	GetDeals(query GetDealsQuery) ([]data.Deal, error)
	GetDeal(id string) (*data.Deal, error)
}
