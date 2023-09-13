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
