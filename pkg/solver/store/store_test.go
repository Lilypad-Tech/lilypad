//go:build integration && solver

package store_test

import (
	"fmt"
	"slices"
	"sort"
	"sync"
	"testing"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	databasestore "github.com/lilypad-tech/lilypad/pkg/solver/store/db"
	memorystore "github.com/lilypad-tech/lilypad/pkg/solver/store/memory"
	"golang.org/x/exp/rand"
)

const DB_CONN_STR = "postgres://postgres:postgres@localhost:5432/solver-db?sslmode=disable"

type storeConfig struct {
	name string
	init func() (getStore func() store.SolverStore)
}

func setupStores(t *testing.T) []storeConfig {
	initMemory := func() func() store.SolverStore {
		// Get store function creates a new memory store
		// which effectively clears data between runs
		return func() store.SolverStore {
			s, err := memorystore.NewSolverStoreMemory()
			if err != nil {
				t.Fatalf("Failed to create memory store: %v", err)
			}
			return s
		}
	}

	initDatabase := func() func() store.SolverStore {
		db, err := databasestore.NewSolverStoreDatabase(DB_CONN_STR, true)
		if err != nil {
			t.Fatalf("Failed to create database store: %v", err)
		}

		// Clear data at initialization
		clearStoreDatabase(t, db)

		// Get store functions clear data and returns
		// the same store instance
		return func() store.SolverStore {
			clearStoreDatabase(t, db)
			return db
		}
	}

	return []storeConfig{
		{name: "memory", init: initMemory},
		{name: "database", init: initDatabase},
	}
}

func clearStoreDatabase(t *testing.T, s store.SolverStore) {
	jobOffers, err := s.GetJobOffers(store.GetJobOffersQuery{
		IncludeCancelled: true,
	})
	if err != nil {
		t.Fatalf("Failed to get existing job offers: %v", err)
	}

	for _, result := range jobOffers {
		err := s.RemoveJobOffer(result.ID)
		if err != nil {
			t.Fatalf("Failed to remove existing job offer: %v", err)
		}
	}
}

// Job offers

func TestJobOfferOps(t *testing.T) {
	storeConfigs := setupStores(t)
	for _, config := range storeConfigs {
		// Test multiple job offers in a single test run
		t.Run(config.name, func(t *testing.T) {
			getStore := config.init()
			store := getStore()

			// Generate multiple job offers
			jobOffers := generateJobOffers(5, 50)

			for _, jobOffer := range jobOffers {
				// Add job offer
				added, err := store.AddJobOffer(jobOffer)
				if err != nil {
					t.Fatalf("Failed to add job offer: %v", err)
				}
				if added.ID != jobOffer.ID {
					t.Errorf("Expected ID %s, got %s", jobOffer.ID, added.ID)
				}

				// Get job offer
				retrieved, err := store.GetJobOffer(jobOffer.ID)
				if err != nil {
					t.Fatalf("Failed to get job offer: %v", err)
				}
				if retrieved == nil {
					t.Fatalf("Expected job offer, got nil")
				}
				if retrieved.ID != jobOffer.ID {
					t.Errorf("Expected ID %s, got %s", jobOffer.ID, retrieved.ID)
				}

				// Update job offer
				newDealID := generateCID()
				newState := generateState()

				updated, err := store.UpdateJobOfferState(jobOffer.ID, newDealID, newState)
				if err != nil {
					t.Fatalf("Failed to update job offer state: %v", err)
				}
				if updated.DealID != newDealID || updated.State != newState {
					t.Errorf("Update failed: expected dealID=%s state=%d, got dealID=%s state=%d",
						newDealID, newState, updated.DealID, updated.State)
				}

				// Remove job offer
				err = store.RemoveJobOffer(jobOffer.ID)
				if err != nil {
					t.Fatalf("Failed to remove job offer: %v", err)
				}

				// Verify removal
				removed, err := store.GetJobOffer(jobOffer.ID)
				if err != nil {
					t.Fatalf("Error checking removed job offer: %v", err)
				}
				if removed != nil {
					t.Error("Job offer still exists after removal")
				}
			}
		})
	}
}

func TestJobOfferQuery(t *testing.T) {
	// Test cases set offer fields relevant to querying.
	// All other fields are left with their zero-values.
	testCases := []struct {
		name     string
		offers   []data.JobOfferContainer
		query    store.GetJobOffersQuery
		expected []string // expected IDs in results
	}{
		{
			name: "filter by job creator",
			offers: []data.JobOfferContainer{
				{
					ID:         "QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      0,
				},
				{
					ID:         "QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky",
					JobCreator: "0xabcdef0123456789abcdef0123456789abcdef01",
					DealID:     "",
					State:      0,
				},
			},
			query: store.GetJobOffersQuery{
				JobCreator: "0x1234567890123456789012345678901234567890",
			},
			expected: []string{"QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx"},
		},
		{
			name: "filter not matched offers",
			offers: []data.JobOfferContainer{
				{
					ID:         "QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "QmZ9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kz",
					State:      0,
				},
				{
					ID:         "QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      0,
				},
			},
			query: store.GetJobOffersQuery{
				NotMatched: true,
			},
			expected: []string{"QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky"},
		},
		{
			name: "filter out cancelled offers",
			offers: []data.JobOfferContainer{
				{
					ID:         "QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      data.GetDefaultAgreementState(),
				},
				{
					ID:         "QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      data.GetAgreementStateIndex("JobOfferCancelled"),
				},
			},
			query: store.GetJobOffersQuery{
				IncludeCancelled: false,
			},
			expected: []string{"QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx"},
		},
		{
			name: "include cancelled offers",
			offers: []data.JobOfferContainer{
				{
					ID:         "QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      data.GetDefaultAgreementState(),
				},
				{
					ID:         "QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      data.GetAgreementStateIndex("JobOfferCancelled"),
				},
			},
			query: store.GetJobOffersQuery{
				IncludeCancelled: true,
			},
			expected: []string{
				"QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx",
				"QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky",
			},
		},
		{
			name: "combined filters",
			offers: []data.JobOfferContainer{
				{
					ID:         "QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      data.GetDefaultAgreementState(),
				},
				{
					ID:         "QmX9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Ky",
					JobCreator: "0xabcdef0123456789abcdef0123456789abcdef01",
					DealID:     "",
					State:      data.GetDefaultAgreementState(),
				},
				{
					ID:         "QmZ9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kz",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "QmW9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kw",
					State:      data.GetDefaultAgreementState(),
				},
				{
					ID:         "QmV9JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kv",
					JobCreator: "0x1234567890123456789012345678901234567890",
					DealID:     "",
					State:      data.GetAgreementStateIndex("JobOfferCancelled"),
				},
			},
			query: store.GetJobOffersQuery{
				JobCreator:       "0x1234567890123456789012345678901234567890",
				NotMatched:       true,
				IncludeCancelled: false,
			},
			expected: []string{"QmY8JwJh3bYDUuAnwfpxwStjUY1nQwyhJJ4SPpdV3bZ9Kx"},
		},
	}

	storeConfigs := setupStores(t)
	for _, config := range storeConfigs {
		getStore := config.init()

		for _, tc := range testCases {
			// Test each case in a separate test run
			t.Run(fmt.Sprintf("%s/%s", config.name, tc.name), func(t *testing.T) {
				store := getStore()

				// Add test job offers
				for _, jo := range tc.offers {
					_, err := store.AddJobOffer(jo)
					if err != nil {
						t.Fatalf("Failed to add job offer: %v", err)
					}
				}

				// Run query
				results, err := store.GetJobOffers(tc.query)
				if err != nil {
					t.Fatalf("GetJobOffers failed: %v", err)
				}

				// Extract result IDs
				resultIDs := make([]string, len(results))
				for i, r := range results {
					resultIDs[i] = r.ID
				}

				// Sort both slices for comparison
				sort.Strings(resultIDs)
				sort.Strings(tc.expected)

				if !slices.Equal(resultIDs, tc.expected) {
					t.Errorf("Expected results %v, got %v", tc.expected, resultIDs)
				}
			})
		}
	}
}

func TestJobOfferConcurrentOps(t *testing.T) {
	offers := generateJobOffers(4, 10)

	storeConfigs := setupStores(t)
	for _, config := range storeConfigs {
		// Test concurrent adds in a single test run
		t.Run(config.name, func(t *testing.T) {
			getStore := config.init()
			store := getStore()

			// Add offers concurrently
			errCh := make(chan error, len(offers))
			var wg sync.WaitGroup

			for _, offer := range offers {
				wg.Add(1)
				go func(o data.JobOfferContainer) {
					defer wg.Done()
					_, err := store.AddJobOffer(o)
					if err != nil {
						errCh <- err
					}
				}(offer)
			}

			wg.Wait()
			close(errCh)

			// Check for any errors during concurrent operations
			for err := range errCh {
				if err != nil {
					t.Errorf("Concurrent operation error: %v", err)
				}
			}

			// Verify all offers were added
			for _, offer := range offers {
				retrieved, err := store.GetJobOffer(offer.ID)
				if err != nil {
					t.Errorf("Failed to get job offer %s: %v", offer.ID, err)
				}
				if retrieved == nil {
					t.Errorf("Job offer %s not found after concurrent add", offer.ID)
				}
				if retrieved != nil && retrieved.ID != offer.ID {
					t.Errorf("Retrieved job offer ID mismatch: expected %s, got %s", offer.ID, retrieved.ID)
				}
			}
		})
	}
}

// Generators

func generateCID() string {
	randBytes := make([]byte, 22)
	rand.Read(randBytes)

	// Mock CIDv0
	return fmt.Sprintf("Qm%44x", randBytes)
}

func generateState() uint8 {
	return uint8(rand.Intn(len(data.AgreementState)))
}

func generateJobOffer() data.JobOfferContainer {
	return data.JobOfferContainer{
		ID: generateCID(),
	}
}

func generateJobOffers(min int, max int) []data.JobOfferContainer {
	count := min + rand.Intn(max-min+1)
	offers := make([]data.JobOfferContainer, count)

	for i := 0; i < count; i++ {
		offers[i] = generateJobOffer()
	}

	return offers
}
