//go:build integration && solver

package solver_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/solver"
	"github.com/lilypad-tech/lilypad/pkg/solver/matcher"
	"go.opentelemetry.io/otel"
)

func TestMatchingSortingLogic(t *testing.T) {
	testCases := []struct {
		name           string
		jobOffer       data.JobOfferContainer
		resourceOffers []data.ResourceOfferContainer
		expectedWinner string
	}{
		{
			name:     "market price - selects cheapest and oldest offer",
			jobOffer: getJobOffer("job-market", data.MarketPrice, 0),
			resourceOffers: []data.ResourceOfferContainer{
				getResourceOffer("cheap-old", 1, 100),
				getResourceOffer("expensive-new", 2, 200),
				getResourceOffer("cheap-newer", 3, 100),
			},
			expectedWinner: "cheap-old",
		},
		{
			name:     "market price - prefers cheaper over older",
			jobOffer: getJobOffer("job-market", data.MarketPrice, 0),
			resourceOffers: []data.ResourceOfferContainer{
				getResourceOffer("expensive-old", 1, 200),
				getResourceOffer("cheap-new", 2, 100),
			},
			expectedWinner: "cheap-new",
		},
		{
			name:     "fixed price - selects oldest affordable offer",
			jobOffer: getJobOffer("job-fixed", data.FixedPrice, 150),
			resourceOffers: []data.ResourceOfferContainer{
				getResourceOffer("cheap-old", 1, 100),
				getResourceOffer("expensive-new", 2, 200),
				getResourceOffer("cheap-newer", 3, 100),
			},
			expectedWinner: "cheap-old",
		},
		{
			name:     "fixed price - exact price match prefers oldest",
			jobOffer: getJobOffer("job-fixed", data.FixedPrice, 100),
			resourceOffers: []data.ResourceOfferContainer{
				getResourceOffer("exact-old", 1, 100),
				getResourceOffer("exact-newer", 2, 100),
				getResourceOffer("exact-newest", 3, 100),
			},
			expectedWinner: "exact-old",
		},
	}

	storeConfigs := solver.SetupTestStores(t)
	for _, config := range storeConfigs {
		getStore, clearStore := config.Init()

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s/%s", config.Name, tc.name), func(t *testing.T) {
				store := getStore()
				defer clearStore()

				// Add job offer
				_, err := store.AddJobOffer(tc.jobOffer)
				if err != nil {
					t.Fatalf("Failed to add job offer: %v", err)
				}

				// The database store implementation sorts by database timestamp,
				// so we insert with a small delay to create an ordering. The memory implementation
				// sorts by resource offer created time, so we assign one in our test cases.
				for i, offer := range tc.resourceOffers {
					_, err := store.AddResourceOffer(offer)
					if err != nil {
						t.Fatalf("Failed to add resource offer: %v", err)
					}

					if i < len(tc.resourceOffers)-1 {
						time.Sleep(50 * time.Millisecond)
					}
				}

				// Run matcher
				deals, err := matcher.GetMatchingDeals(
					context.Background(),
					store,
					func(string, string, uint8) (*data.JobOfferContainer, error) { return nil, nil },
					otel.Tracer("test"),
					otel.Meter("test"),
				)
				if err != nil {
					t.Fatalf("GetMatchingDeals failed: %v", err)
				}

				if len(deals) != 1 {
					t.Fatalf("Expected 1 deal, got %d", len(deals))
				}
				if deals[0].ResourceOffer.ID != tc.expectedWinner {
					t.Errorf("Expected resource offer %s to win, got %s",
						tc.expectedWinner, deals[0].ResourceOffer.ID)
				}
			})
		}
	}
}

func getJobOffer(id string, mode data.PricingMode, maxPrice uint64) data.JobOfferContainer {
	return data.JobOfferContainer{
		ID: id,
		JobOffer: data.JobOffer{
			ID:   id,
			Mode: mode,
			Pricing: data.DealPricing{
				InstructionPrice: maxPrice,
			},
			Services: data.ServiceConfig{
				Solver:   "default-solver",
				Mediator: []string{"mediator1"},
			},
		},
	}
}

func getResourceOffer(id string, createdAt int, price uint64) data.ResourceOfferContainer {
	return data.ResourceOfferContainer{
		ID: id,
		ResourceOffer: data.ResourceOffer{
			ID:        id,
			CreatedAt: createdAt,
			DefaultPricing: data.DealPricing{
				InstructionPrice: price,
			},
			Services: data.ServiceConfig{
				Solver:   "default-solver",
				Mediator: []string{"mediator1"},
			},
		},
	}
}
