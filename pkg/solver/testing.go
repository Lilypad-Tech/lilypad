package solver

import (
	"testing"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
	databasestore "github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store/db"
	memorystore "github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store/memory"
)

const DB_CONN_STR = "postgres://postgres:postgres@localhost:5432/solver-db?sslmode=disable"

type TestStoreConfig struct {
	Name string
	Init func() (getStore func() store.SolverStore, clearStore func())
}

func SetupTestStores(t *testing.T) []TestStoreConfig {
	initMemory := func() (func() store.SolverStore, func()) {
		// Get store function creates a new memory store
		// which effectively clears data between runs
		getStore := func() store.SolverStore {
			s, err := memorystore.NewSolverStoreMemory()
			if err != nil {
				t.Fatalf("Failed to create memory store: %v", err)
			}
			return s
		}
		clearStore := func() {}

		return getStore, clearStore
	}

	initDatabase := func() (func() store.SolverStore, func()) {
		db, err := databasestore.NewSolverStoreDatabase(DB_CONN_STR, "silent")
		if err != nil {
			t.Fatalf("Failed to create database store: %v", err)
		}

		// Clear data at initialization
		clearStoreDatabase(t, db)

		// Get store functions clear data and returns
		// the same store instance
		getStore := func() store.SolverStore {
			clearStoreDatabase(t, db)
			return db
		}
		clearStore := func() { clearStoreDatabase(t, db) }

		return getStore, clearStore
	}

	return []TestStoreConfig{
		{Name: "memory", Init: initMemory},
		{Name: "database", Init: initDatabase},
	}
}

func clearStoreDatabase(t *testing.T, s store.SolverStore) {
	// Delete job offers
	jobOffers, err := s.GetJobOffers(store.GetJobOffersQuery{})
	if err != nil {
		t.Fatalf("Failed to get existing job offers: %v", err)
	}

	for _, result := range jobOffers {
		err := s.RemoveJobOffer(result.ID)
		if err != nil {
			t.Fatalf("Failed to remove existing job offer: %v", err)
		}
	}

	// Delete resource offers
	resourceOffers, err := s.GetResourceOffers(store.GetResourceOffersQuery{})
	if err != nil {
		t.Fatalf("Failed to get existing resource offers: %v", err)
	}

	for _, result := range resourceOffers {
		err := s.RemoveResourceOffer(result.ID)
		if err != nil {
			t.Fatalf("Failed to remove existing resource offer: %v", err)
		}
	}

	// Delete deals
	deals, err := s.GetDealsAll()
	if err != nil {
		t.Fatalf("Failed to get existing deals: %v", err)
	}

	for _, deal := range deals {
		err := s.RemoveDeal(deal.ID)
		if err != nil {
			t.Fatalf("Failed to remove existing deal: %v", err)
		}
	}

	// Delete results
	results, err := s.GetResults()
	if err != nil {
		t.Fatalf("Failed to get existing results: %v", err)
	}

	for _, result := range results {
		err := s.RemoveResult(result.DealID)
		if err != nil {
			t.Fatalf("Failed to remove existing result: %v", err)
		}
	}

	// Delete match decisions
	decisions, err := s.GetMatchDecisions()
	if err != nil {
		t.Fatalf("Failed to get existing match decisions: %v", err)
	}

	for _, decision := range decisions {
		err := s.RemoveMatchDecision(decision.ResourceOffer, decision.JobOffer)
		if err != nil {
			t.Fatalf("Failed to remove existing match decision: %v", err)
		}
	}

	// Delete allowed resource providers
	providers, err := s.GetAllowedResourceProviders()
	if err != nil {
		t.Fatalf("Failed to get existing allowed resource providers: %v", err)
	}

	for _, provider := range providers {
		err := s.RemoveAllowedResourceProvider(provider)
		if err != nil {
			t.Fatalf("Failed to remove existing allowed resource provider: %v", err)
		}
	}
}
