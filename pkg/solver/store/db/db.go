package store

import (
	"errors"
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SolverStoreDatabase struct {
	// TODO Remove placeholder which we needed to implement the interface with
	// package name shadowed
	placeholderDeal *data.DealContainer
	db              *gorm.DB
	// TODO Log writers?
	// logWriters       map[string]jsonl.Writer
}

func NewSolverStoreDatabase(connStr string, silenceLogs bool) (*SolverStoreDatabase, error) {
	config := &gorm.Config{}
	if silenceLogs {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(connStr), config)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&JobOffer{})
	db.AutoMigrate(&ResourceOffer{})
	db.AutoMigrate(&Deal{})
	db.AutoMigrate(&Result{})
	db.AutoMigrate(&MatchDecision{})

	return &SolverStoreDatabase{&data.DealContainer{}, db}, nil
}

func (store *SolverStoreDatabase) AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error) {
	record := JobOffer{
		CID:        jobOffer.ID,
		JobCreator: jobOffer.JobCreator,
		DealID:     jobOffer.DealID,
		State:      jobOffer.State,
		Attributes: datatypes.NewJSONType(jobOffer),
	}

	result := store.db.Create(&record)
	if result.Error != nil {
		return nil, result.Error
	}

	return &jobOffer, nil
}

func (store *SolverStoreDatabase) AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error) {
	record := ResourceOffer{
		CID:              resourceOffer.ID,
		ResourceProvider: resourceOffer.ResourceProvider,
		DealID:           resourceOffer.DealID,
		State:            resourceOffer.State,
		Attributes:       datatypes.NewJSONType(resourceOffer),
	}

	result := store.db.Create(&record)
	if result.Error != nil {
		return nil, result.Error
	}

	return &resourceOffer, nil
}

func (store *SolverStoreDatabase) AddDeal(deal data.DealContainer) (*data.DealContainer, error) {
	return &deal, nil
}

func (store *SolverStoreDatabase) AddResult(result data.Result) (*data.Result, error) {
	return &result, nil
}

func (store *SolverStoreDatabase) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error) {
	decision := &data.MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Deal:          deal,
		Result:        result,
	}
	return decision, nil
}

func (store *SolverStoreDatabase) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	q := store.db.Where([]JobOffer{})

	// Apply filters
	if query.JobCreator != "" {
		q = q.Where("job_creator = ?", query.JobCreator)
	}
	if query.NotMatched {
		q = q.Where("deal_id = ''")
	}
	if !query.IncludeCancelled {
		q = q.Where("state != ?", data.GetAgreementStateIndex("JobOfferCancelled"))
	}

	var records []JobOffer
	if err := q.Find(&records).Error; err != nil {
		return nil, err
	}

	jobOffers := make([]data.JobOfferContainer, len(records))
	for i, record := range records {
		jobOffers[i] = record.Attributes.Data()
	}

	return jobOffers, nil
}

func (store *SolverStoreDatabase) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	q := store.db.Where([]ResourceOffer{})

	// Apply filters
	if query.ResourceProvider != "" {
		q = q.Where("resource_provider = ?", query.ResourceProvider)
	}
	if query.NotMatched {
		q = q.Where("deal_id = ''")
	}
	if query.Active {
		q = q.Where("state IN (?)", []uint8{
			data.GetAgreementStateIndex("DealNegotiating"),
			data.GetAgreementStateIndex("DealAgreed"),
		})
	}

	var records []ResourceOffer
	if err := q.Find(&records).Error; err != nil {
		return nil, err
	}

	resourceOffers := make([]data.ResourceOfferContainer, len(records))
	for i, record := range records {
		resourceOffers[i] = record.Attributes.Data()
	}

	return resourceOffers, nil
}

func (store *SolverStoreDatabase) GetDeals(query store.GetDealsQuery) ([]data.DealContainer, error) {
	deals := []data.DealContainer{}
	return deals, nil
}

func (store *SolverStoreDatabase) GetDealsAll() ([]data.DealContainer, error) {
	deals := []data.DealContainer{}
	return deals, nil
}

func (store *SolverStoreDatabase) GetJobOffer(id string) (*data.JobOfferContainer, error) {
	// Offers are unique by CID, so we can query first
	var record JobOffer
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	jobOffer := record.Attributes.Data()
	return &jobOffer, nil
}

func (store *SolverStoreDatabase) GetResourceOffer(id string) (*data.ResourceOfferContainer, error) {
	// Offers are unique by CID, so we can query first
	var record ResourceOffer
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	resourceOffer := record.Attributes.Data()
	return &resourceOffer, nil
}

func (store *SolverStoreDatabase) GetResourceOfferByAddress(address string) (*data.ResourceOfferContainer, error) {
	var record ResourceOffer
	result := store.db.Where("resource_provider = ?", address).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	resourceOffer := record.Attributes.Data()
	return &resourceOffer, nil
}

func (store *SolverStoreDatabase) GetDeal(id string) (*data.DealContainer, error) {
	deal := &data.DealContainer{}
	return deal, nil
}

func (store *SolverStoreDatabase) GetResult(id string) (*data.Result, error) {
	result := &data.Result{}
	return result, nil
}

func (store *SolverStoreDatabase) GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error) {
	decision := &data.MatchDecision{}
	return decision, nil
}

func (store *SolverStoreDatabase) UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	var record JobOffer
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("job offer not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.DealID = dealID
	inner.State = state

	if err := store.db.Model(&record).
		Select("DealID", "State", "Attributes").
		Updates(JobOffer{
			DealID:     dealID,
			State:      state,
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
}

func (store *SolverStoreDatabase) UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	var record ResourceOffer
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("resource offer not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.DealID = dealID
	inner.State = state

	if err := store.db.Model(&record).
		Select("DealID", "State", "Attributes").
		Updates(ResourceOffer{
			DealID:     dealID,
			State:      state,
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
}

func (store *SolverStoreDatabase) UpdateDealState(id string, state uint8) (*data.DealContainer, error) {
	deal := &data.DealContainer{}
	return deal, nil
}

func (store *SolverStoreDatabase) UpdateDealMediator(id string, mediator string) (*data.DealContainer, error) {
	deal := &data.DealContainer{}
	return deal, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	return store.placeholderDeal, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	return store.placeholderDeal, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error) {
	return store.placeholderDeal, nil
}

func (store *SolverStoreDatabase) RemoveJobOffer(id string) error {
	var record JobOffer
	result := store.db.Where("c_id = ?", id).Delete(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *SolverStoreDatabase) RemoveResourceOffer(id string) error {
	var record ResourceOffer
	result := store.db.Where("c_id = ?", id).Delete(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
