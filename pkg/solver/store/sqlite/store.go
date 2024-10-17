package store

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/sqlite"
	"fmt"
	// "strings"
	"sync"

	"github.com/lilypad-tech/lilypad/pkg/data"
	// "github.com/lilypad-tech/lilypad/pkg/jsonl"
	// "github.com/lilypad-tech/lilypad/pkg/solver/store"
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
    err = database.AutoMigrate()
	if err != nil {
		fmt.Println("Failed to auto migrate:", err)
        return nil, err
	}

	return &SolverStoreSqlite{DB: database}, err
}

func (s *SolverStoreSqlite) AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	jobOfferData := ToJobOfferData(jobOffer)
	result := s.DB.Create(&jobOfferData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &jobOffer, nil
}

// func (s *SolverStoreSqlite) AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error) {
// 	s.mutex.Lock()
// 	defer s.mutex.Unlock()
//
// 	dealData := ToRe(deal)
// 	result := s.DB.Create(&dealData)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
//
// 	return &resourceOffer, nil
// }

func (s *SolverStoreSqlite) AddDeal(deal data.DealContainer) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	dealData := ToDealData(deal)
	result := s.DB.Create(&dealData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &deal, nil
}

func (s *SolverStoreSqlite) AddResult(result data.Result) (*data.Result, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return &result, nil
}

// Removed everything else needs redo
