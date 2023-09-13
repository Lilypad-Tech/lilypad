package solver

import (
	"testing"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

func TestDoOffersMatch(t *testing.T) {
	trustedParties := data.TrustedParties{
		Mediator:  []string{"apples"},
		Directory: []string{"oranges"},
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
		Mode:           data.FixedPrice,
		TrustedParties: trustedParties,
	}

	basicJobOffer := data.JobOffer{
		Spec: data.MachineSpec{
			CPU: 1000,
			GPU: 1000,
			RAM: 1024,
		},
		Mode:           data.MarketPrice,
		TrustedParties: trustedParties,
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
				offer.TrustedParties.Mediator = []string{}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.TrustedParties.Mediator = []string{}
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Mis-matched mediators",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.TrustedParties.Mediator = []string{"apples2"}
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
				offer.TrustedParties.Mediator = []string{"apples2", "apples"}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "Different but matching directories",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.TrustedParties.Directory = []string{"oranges2", "oranges"}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: true,
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
