//go:build unit

package matcher

import "testing"


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
