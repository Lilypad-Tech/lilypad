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
	db *gorm.DB
}

func NewSolverStoreDatabase(connStr string, gormLogLevel string) (*SolverStoreDatabase, error) {
	config := &gorm.Config{}
	switch gormLogLevel {
	case "silent":
		config.Logger = logger.Default.LogMode(logger.Silent)
	case "info":
		config.Logger = logger.Default.LogMode(logger.Info)
	case "error":
		config.Logger = logger.Default.LogMode(logger.Error)
	case "warn":
		config.Logger = logger.Default.LogMode(logger.Warn)
	default:
		return nil, fmt.Errorf("expected solver store gorm log level silent, info, error, or warn, but received: %s", gormLogLevel)
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

	return &SolverStoreDatabase{db}, nil
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
	record := Deal{
		CID:              deal.ID,
		JobCreator:       deal.JobCreator,
		ResourceProvider: deal.ResourceProvider,
		Mediator:         deal.Mediator,
		State:            deal.State,
		Attributes:       datatypes.NewJSONType(deal),
	}

	result := store.db.Create(&record)
	if result.Error != nil {
		return nil, result.Error
	}

	return &deal, nil
}

func (store *SolverStoreDatabase) AddResult(result data.Result) (*data.Result, error) {
	record := Result{
		DealID:     result.DealID,
		CID:        result.ID,
		Attributes: datatypes.NewJSONType(result),
	}

	res := store.db.Create(&record)
	if res.Error != nil {
		return nil, res.Error
	}

	return &result, nil
}

func (store *SolverStoreDatabase) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error) {
	decision := &data.MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Deal:          deal,
		Result:        result,
	}
	record := MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Attributes:    datatypes.NewJSONType(*decision),
	}

	res := store.db.Create(&record)
	if res.Error != nil {
		return nil, res.Error
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
	q := store.db.Where([]Deal{})

	// Apply filters
	if query.JobCreator != "" {
		q = q.Where("job_creator = ?", query.JobCreator)
	}
	if query.ResourceProvider != "" {
		q = q.Where("resource_provider = ?", query.ResourceProvider)
	}
	if query.Mediator != "" {
		q = q.Where("mediator = ?", query.Mediator)
	}
	if query.State != "" {
		parsedState, err := data.GetAgreementState(query.State)
		if err != nil {
			return nil, err
		}
		q = q.Where("state = ?", parsedState)
	}

	var records []Deal
	if err := q.Find(&records).Error; err != nil {
		return nil, err
	}

	deals := make([]data.DealContainer, len(records))
	for i, record := range records {
		deals[i] = record.Attributes.Data()
	}

	return deals, nil
}

func (store *SolverStoreDatabase) GetDealsAll() ([]data.DealContainer, error) {
	var records []Deal
	if err := store.db.Find(&records).Error; err != nil {
		return nil, err
	}

	deals := make([]data.DealContainer, len(records))
	for i, record := range records {
		deals[i] = record.Attributes.Data()
	}

	return deals, nil
}

func (store *SolverStoreDatabase) GetResults() ([]data.Result, error) {
	var records []Result
	if err := store.db.Find(&records).Error; err != nil {
		return nil, err
	}

	results := make([]data.Result, len(records))
	for i, record := range records {
		results[i] = record.Attributes.Data()
	}

	return results, nil
}

func (store *SolverStoreDatabase) GetMatchDecisions() ([]data.MatchDecision, error) {
	var records []MatchDecision
	if err := store.db.Find(&records).Error; err != nil {
		return nil, err
	}

	decisions := make([]data.MatchDecision, len(records))
	for i, record := range records {
		decisions[i] = record.Attributes.Data()
	}

	return decisions, nil
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
	// Deals are unique by CID, so we can query first
	var record Deal
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	deal := record.Attributes.Data()
	return &deal, nil
}

func (store *SolverStoreDatabase) GetResult(id string) (*data.Result, error) {
	// Results are queried by deal ID for now
	// Deal IDs are unique, so we can query first
	var record Result
	res := store.db.Where("deal_id = ?", id).First(&record)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}

	result := record.Attributes.Data()
	return &result, nil
}

func (store *SolverStoreDatabase) GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error) {
	// The resource offer and job offer are unique
	// CIDs, so we can query first
	var record MatchDecision
	result := store.db.Where("resource_offer = ? AND job_offer = ?", resourceOffer, jobOffer).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	decision := record.Attributes.Data()
	return &decision, nil
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
	var record Deal
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deal not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.State = state

	if err := store.db.Model(&record).
		Select("State", "Attributes").
		Updates(Deal{
			State:      state,
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
}

func (store *SolverStoreDatabase) UpdateDealMediator(id string, mediator string) (*data.DealContainer, error) {
	var record Deal
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deal not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.Mediator = mediator

	if err := store.db.Model(&record).
		Select("Mediator", "Attributes").
		Updates(Deal{
			Mediator:   mediator,
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	var record Deal
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deal not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.Transactions.JobCreator = data

	if err := store.db.Model(&record).
		Select("Attributes").
		Updates(Deal{
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	var record Deal
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deal not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.Transactions.ResourceProvider = data

	if err := store.db.Model(&record).
		Select("Attributes").
		Updates(Deal{
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
}

func (store *SolverStoreDatabase) UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error) {
	var record Deal
	result := store.db.Where("c_id = ?", id).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deal not found: %s", id)
		}
		return nil, result.Error
	}

	// Update the jsonb data
	inner := record.Attributes.Data()
	inner.Transactions.Mediator = data

	if err := store.db.Model(&record).
		Select("Attributes").
		Updates(Deal{
			Attributes: datatypes.NewJSONType(inner),
		}).Error; err != nil {
		return nil, err
	}

	return &inner, nil
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

func (store *SolverStoreDatabase) RemoveDeal(id string) error {
	var record Deal
	result := store.db.Where("c_id = ?", id).Delete(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *SolverStoreDatabase) RemoveResult(id string) error {
	var record Result
	result := store.db.Where("deal_id = ?", id).Delete(&record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *SolverStoreDatabase) RemoveMatchDecision(resourceOffer string, jobOffer string) error {
	if resourceOffer == "" && jobOffer == "" {
		return fmt.Errorf("resource offer or job offer must be set")
	}

	// Build the query based on which parameters are provided
	query := store.db.Where([]MatchDecision{})

	if resourceOffer != "" {
		query = query.Where("resource_offer = ?", resourceOffer)
	}
	if jobOffer != "" {
		query = query.Where("job_offer = ?", jobOffer)
	}

	// Execute the delete
	result := query.Delete(&MatchDecision{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Strictly speaking, the compiler will check the interface
// implementation without this check. But some code editors
// report errors more effectively when we have it.
var _ store.SolverStore = (*SolverStoreDatabase)(nil)
