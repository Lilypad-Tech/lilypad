//go:build bench && solver

package matcher

import "testing"

func BenchmarkMatchOffers(b *testing.B) {
	testCases := getMatchTestCases()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			result := matchOffers(tc.resourceOffer, tc.jobOffer)
			if result.matched() != tc.shouldMatch {
				b.Fatalf("expected match to be %v for case %s", tc.shouldMatch, tc.name)
			}
		}
	}
}
