package solver

import (
	"testing"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

func TestDoOffersMatch(t *testing.T) {
	services := data.ServiceConfig{
		Solver:   "oranges",
		Mediator: []string{"apples"},
	}

	basicResourceOffer := data.ResourceOffer{
		Spec: data.MachineSpec{
			CPU: 1000,
			GPU: 1000,
			RAM: 1024,
		},
		DefaultPricing: data.DealPricing{
			InstructionPrice: 10,
		},
		Mode:     data.FixedPrice,
		Services: services,
	}

	basicJobOffer := data.JobOffer{
		Spec: data.MachineSpec{
			CPU: 1000,
			GPU: 1000,
			RAM: 1024,
		},
		Mode:     data.MarketPrice,
		Services: services,
	}

	testCases := []struct {
		name          string
		resourceOffer func(offer data.ResourceOffer) data.ResourceOffer
		jobOffer      func(offer data.JobOffer) data.JobOffer
		shouldMatch   bool
	}{
		{
			name: "Basic match",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "CPU mis-match",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Spec.CPU = 2000
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Empty mediators",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Mediator = []string{}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Services.Mediator = []string{}
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Mis-matched mediators",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Mediator = []string{"apples2"}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Different but matching mediators",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Mediator = []string{"apples2", "apples"}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "Different solver",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Solver = "pears"
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Fixed price - too expensive",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Mode = data.FixedPrice
				offer.Pricing.InstructionPrice = 9
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Fixed price - can afford",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Mode = data.FixedPrice
				offer.Pricing.InstructionPrice = 11
				return offer
			},
			shouldMatch: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := doOffersMatch(tc.resourceOffer(basicResourceOffer), tc.jobOffer(basicJobOffer))
			if result != tc.shouldMatch {
				t.Errorf("Expected match to be %v, but got %v", tc.shouldMatch, result)
			}
		})
	}
}
