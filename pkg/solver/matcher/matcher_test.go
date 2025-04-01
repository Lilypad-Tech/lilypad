//go:build unit

package matcher

import (
	"testing"

	"github.com/lilypad-tech/lilypad/v2/pkg/data"
)

func TestMatchOffers(t *testing.T) {
	testCases := getMatchTestCases()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := matchOffers(tc.resourceOffer, tc.jobOffer)
			if result.matched() != tc.shouldMatch {
				t.Errorf("Expected match to be %v, but got %+v", tc.shouldMatch, result)
			}
		})
	}
}

func TestIsCheaperOrOlder(t *testing.T) {
	testCases := []struct {
		name     string
		offerA   data.ResourceOffer
		offerB   data.ResourceOffer
		expected bool
	}{
		{
			name: "cheaper wins",
			offerA: data.ResourceOffer{
				CreatedAt: 2,
				DefaultPricing: data.DealPricing{
					InstructionPrice: 100,
				},
			},
			offerB: data.ResourceOffer{
				CreatedAt: 1,
				DefaultPricing: data.DealPricing{
					InstructionPrice: 200,
				},
			},
			expected: true,
		},
		{
			name: "older wins when same price",
			offerA: data.ResourceOffer{
				CreatedAt: 1,
				DefaultPricing: data.DealPricing{
					InstructionPrice: 100,
				},
			},
			offerB: data.ResourceOffer{
				CreatedAt: 2,
				DefaultPricing: data.DealPricing{
					InstructionPrice: 100,
				},
			},
			expected: true,
		},
		{
			name: "more expensive loses regardless of age",
			offerA: data.ResourceOffer{
				CreatedAt: 1,
				DefaultPricing: data.DealPricing{
					InstructionPrice: 200,
				},
			},
			offerB: data.ResourceOffer{
				CreatedAt: 2,
				DefaultPricing: data.DealPricing{
					InstructionPrice: 100,
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isCheaperOrOlder(tc.offerA, tc.offerB)
			if result != tc.expected {
				t.Errorf("isCheaperOrOlder(%+v, %+v) = %v, expected %v",
					tc.offerA, tc.offerB, result, tc.expected)
			}
		})
	}
}
