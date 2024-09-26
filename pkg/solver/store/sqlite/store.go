package store

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/sqlite"
	"fmt"
	"strings"
	"sync"

	"github.com/lilypad-tech/lilypad/pkg/data"
	// "github.com/lilypad-tech/lilypad/pkg/jsonl"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
)

type SolverStoreSqlite struct {
	DB					*gorm.DB
	mutex				sync.RWMutex
}

func getMatchID(resourceOffer string, jobOffer string) string {
	return fmt.Sprintf("%s-%s", resourceOffer, jobOffer)
}

func NewSolverStoreSqlite(dbPath string) (*SolverStoreSqlite, error) {
	var database *gorm.DB
    var err error

    database, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
        Logger: logger.Default.LogMode(logger.Silent),
    })

    if err != nil {
        return nil, err
    }

	// create tables if they haven't been created yet
    // err = database.AutoMigrate()
	if err != nil {
		fmt.Println("Failed to auto migrate:", err)
        return nil, err
	}

	return &SolverStoreSqlite{DB: database}, err
}

func (s *SolverStoreSqlite) AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return &jobOffer, nil
}

func (s *SolverStoreSqlite) AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return &resourceOffer, nil
}

func (s *SolverStoreSqlite) AddDeal(deal data.DealContainer) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return &deal, nil
}

func (s *SolverStoreSqlite) AddResult(result data.Result) (*data.Result, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return &result, nil
}

func (s *SolverStoreSqlite) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return decision, nil
}

func (s *SolverStoreSqlite) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return jobOffers, nil
}

func (s *SolverStoreSqlite) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return resourceOffers, nil
}

func (s *SolverStoreSqlite) GetDeals(query store.GetDealsQuery) ([]data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return deals, nil
}

func (s *SolverStoreSqlite) GetJobOffer(id string) (*data.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return jobOffer, nil
}

func (s *SolverStoreSqlite) GetResourceOffer(id string) (*data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return resourceOffer, nil
}

func (s *SolverStoreSqlite) GetResourceOfferByAddress(address string) (*data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return nil, nil
}

func (s *SolverStoreSqlite) GetDeal(id string) (*data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return deal, nil
}

func (s *SolverStoreSqlite) GetResult(id string) (*data.Result, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return result, nil
}

func (s *SolverStoreSqlite) GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return decision, nil
}

func (s *SolverStoreSqlite) UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return jobOffer, nil
}

func (s *SolverStoreSqlite) UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return resourceOffer, nil
}

func (s *SolverStoreSqlite) UpdateDealState(id string, state uint8) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return deal, nil
}

func (s *SolverStoreSqlite) UpdateDealMediator(id string, mediator string) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return deal, nil
}

func (s *SolverStoreSqlite) UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return deal, nil
}

func (s *SolverStoreSqlite) UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return deal, nil
}

func (s *SolverStoreSqlite) UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return deal, nil
}

func (s *SolverStoreSqlite) RemoveJobOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return nil
}

func (s *SolverStoreSqlite) RemoveResourceOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return nil
}

// Compile-time interface check:
var _ store.SolverStore = (*SolverStoreSqlite)(nil)
