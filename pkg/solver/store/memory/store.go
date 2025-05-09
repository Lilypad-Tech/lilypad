package store

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
)

type SolverStoreMemory struct {
	jobOfferMap                map[string]*data.JobOfferContainer
	resourceOfferMap           map[string]*data.ResourceOfferContainer
	dealMap                    map[string]*data.DealContainer
	resultMap                  map[string]*data.Result
	matchDecisionMap           map[string]*data.MatchDecision
	allowedResourceProviderMap map[string]string
	mutex                      sync.RWMutex
}

func NewSolverStoreMemory() (*SolverStoreMemory, error) {
	return &SolverStoreMemory{
		jobOfferMap:                map[string]*data.JobOfferContainer{},
		resourceOfferMap:           map[string]*data.ResourceOfferContainer{},
		dealMap:                    map[string]*data.DealContainer{},
		resultMap:                  map[string]*data.Result{},
		matchDecisionMap:           map[string]*data.MatchDecision{},
		allowedResourceProviderMap: map[string]string{},
	}, nil
}

func (s *SolverStoreMemory) AddJobOffer(jobOffer data.JobOfferContainer) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.jobOfferMap[jobOffer.ID] = &jobOffer

	return &jobOffer, nil
}

func (s *SolverStoreMemory) AddResourceOffer(resourceOffer data.ResourceOfferContainer) (*data.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resourceOfferMap[resourceOffer.ID] = &resourceOffer

	return &resourceOffer, nil
}

func (s *SolverStoreMemory) AddDeal(deal data.DealContainer) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.dealMap[deal.ID] = &deal

	return &deal, nil
}

func (s *SolverStoreMemory) AddResult(result data.Result) (*data.Result, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resultMap[result.DealID] = &result

	return &result, nil
}

func (s *SolverStoreMemory) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*data.MatchDecision, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := store.GetMatchID(resourceOffer, jobOffer)
	_, ok := s.matchDecisionMap[id]
	if ok {
		return nil, fmt.Errorf("that match already exists")
	}
	decision := &data.MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Deal:          deal,
		Result:        result,
	}
	s.matchDecisionMap[id] = decision

	return decision, nil
}

func (store *SolverStoreMemory) AddAllowedResourceProvider(resourceProvider string) (string, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.allowedResourceProviderMap[resourceProvider] = resourceProvider

	return resourceProvider, nil
}

func (s *SolverStoreMemory) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	jobOffers := []data.JobOfferContainer{}
	for _, jobOffer := range s.jobOfferMap {
		matching := true
		if query.JobCreator != "" && jobOffer.JobCreator != query.JobCreator {
			matching = false
		}
		if query.NotMatched {
			if jobOffer.DealID != "" {
				matching = false
			}
		}
		if query.Active &&
			jobOffer.State != data.GetAgreementStateIndex("DealNegotiating") &&
			jobOffer.State != data.GetAgreementStateIndex("DealAgreed") &&
			jobOffer.State != data.GetAgreementStateIndex("ResultsSubmitted") {
			matching = false
		}
		if query.Cancelled != nil {
			isCancelled := jobOffer.State == data.GetAgreementStateIndex("JobOfferCancelled") ||
				jobOffer.State == data.GetAgreementStateIndex("JobTimedOut")

			wantCancelled := *query.Cancelled
			if wantCancelled && !isCancelled {
				matching = false
			} else if !wantCancelled && isCancelled {
				matching = false
			}
		}
		if matching {
			jobOffers = append(jobOffers, *jobOffer)
		}
	}

	// Sort oldest first
	if query.OrderOldestFirst {
		sort.Slice(jobOffers, func(i, j int) bool {
			return jobOffers[i].JobOffer.CreatedAt < jobOffers[j].JobOffer.CreatedAt
		})
	}

	return jobOffers, nil
}

func (s *SolverStoreMemory) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	resourceOffers := []data.ResourceOfferContainer{}
	for _, resourceOffer := range s.resourceOfferMap {
		matching := true
		if query.ResourceProvider != "" && resourceOffer.ResourceProvider != query.ResourceProvider {
			matching = false
		}
		if query.Active && !data.IsActiveAgreementState(resourceOffer.State) {
			matching = false
		}
		if query.NotMatched {
			if resourceOffer.DealID != "" {
				matching = false
			}
		}
		if matching {
			resourceOffers = append(resourceOffers, *resourceOffer)
		}
	}

	// Sort oldest first
	if query.OrderOldestFirst {
		sort.Slice(resourceOffers, func(i, j int) bool {
			return resourceOffers[i].ResourceOffer.CreatedAt < resourceOffers[j].ResourceOffer.CreatedAt
		})
	}

	return resourceOffers, nil
}

func (s *SolverStoreMemory) GetDeals(query store.GetDealsQuery) ([]data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	deals := []data.DealContainer{}
	queryState := uint8(0)
	if query.State != "" {
		parsedState, err := data.GetAgreementState(query.State)
		if err != nil {
			return nil, err
		}
		queryState = parsedState
	}
	for _, deal := range s.dealMap {
		matching := true
		if query.JobCreator != "" && deal.JobCreator != query.JobCreator {
			matching = false
		}
		if query.ResourceProvider != "" && deal.ResourceProvider != query.ResourceProvider {
			matching = false
		}
		if query.Mediator != "" && deal.Mediator != query.Mediator {
			matching = false
		}
		if query.State != "" && deal.State != queryState {
			matching = false
		}
		if matching {
			deals = append(deals, *deal)
		}
	}
	return deals, nil
}

func (s *SolverStoreMemory) GetDealsAll() ([]data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	deals := []data.DealContainer{}
	for _, deal := range s.dealMap {
		deals = append(deals, *deal)
	}
	return deals, nil
}

func (s *SolverStoreMemory) GetResults() ([]data.Result, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	results := []data.Result{}
	for _, result := range s.resultMap {
		results = append(results, *result)
	}
	return results, nil
}

func (s *SolverStoreMemory) GetMatchDecisions() ([]data.MatchDecision, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	results := []data.MatchDecision{}
	for _, decision := range s.matchDecisionMap {
		results = append(results, *decision)
	}
	return results, nil
}

func (store *SolverStoreMemory) GetAllowedResourceProviders() ([]string, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	providers := []string{}
	for provider := range store.allowedResourceProviderMap {
		providers = append(providers, provider)
	}

	return providers, nil
}

func (s *SolverStoreMemory) GetJobOffer(id string) (*data.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	return jobOffer, nil
}

func (s *SolverStoreMemory) GetResourceOffer(id string) (*data.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	return resourceOffer, nil
}

func (s *SolverStoreMemory) GetDeal(id string) (*data.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return deal, nil
}

func (s *SolverStoreMemory) GetResult(id string) (*data.Result, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result, ok := s.resultMap[id]
	if !ok {
		return nil, nil
	}
	return result, nil
}

func (s *SolverStoreMemory) GetMatchDecision(resourceOffer string, jobOffer string) (*data.MatchDecision, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	id := store.GetMatchID(resourceOffer, jobOffer)
	decision, ok := s.matchDecisionMap[id]
	if !ok {
		return nil, nil
	}
	return decision, nil
}

func (s *SolverStoreMemory) UpdateJobOfferState(id string, dealID string, state uint8) (*data.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, fmt.Errorf("job offer not found: %s", id)
	}
	jobOffer.DealID = dealID
	jobOffer.State = state
	s.jobOfferMap[id] = jobOffer
	return jobOffer, nil
}

func (s *SolverStoreMemory) UpdateResourceOfferState(id string, dealID string, state uint8) (*data.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, fmt.Errorf("resource offer not found: %s", id)
	}
	resourceOffer.DealID = dealID
	resourceOffer.State = state
	s.resourceOfferMap[id] = resourceOffer
	return resourceOffer, nil
}

func (s *SolverStoreMemory) UpdateDealState(id string, state uint8) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	deal.State = state
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealUploadTime(id string, timestamp int) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	deal.UploadAt = timestamp
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealDownloadTime(id string, timestamp int) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	deal.DownloadAt = timestamp
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealMediator(id string, mediator string) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	deal.Mediator = mediator
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsResourceProvider(id string, data data.DealTransactionsResourceProvider) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	txs := &deal.Transactions.ResourceProvider
	if data.Agree != "" {
		txs.Agree = data.Agree
	}
	if data.AddResult != "" {
		txs.AddResult = data.AddResult
	}
	if data.TimeoutAgree != "" {
		txs.TimeoutAgree = data.TimeoutAgree
	}
	if data.TimeoutJudgeResult != "" {
		txs.TimeoutJudgeResult = data.TimeoutJudgeResult
	}
	if data.TimeoutMediateResult != "" {
		txs.TimeoutMediateResult = data.TimeoutMediateResult
	}
	return deal, nil
}
func (s *SolverStoreMemory) UpdateDealTransactionsJobCreator(id string, data data.DealTransactionsJobCreator) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	txs := &deal.Transactions.JobCreator
	if data.Agree != "" {
		txs.Agree = data.Agree
	}
	if data.AcceptResult != "" {
		txs.AcceptResult = data.AcceptResult
	}
	if data.CheckResult != "" {
		txs.CheckResult = data.CheckResult
	}
	if data.TimeoutAgree != "" {
		txs.TimeoutAgree = data.TimeoutAgree
	}
	if data.TimeoutSubmitResult != "" {
		txs.TimeoutSubmitResult = data.TimeoutSubmitResult
	}
	if data.TimeoutMediateResult != "" {
		txs.TimeoutMediateResult = data.TimeoutMediateResult
	}
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsMediator(id string, data data.DealTransactionsMediator) (*data.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	txs := &deal.Transactions.Mediator
	if data.MediationAcceptResult != "" {
		txs.MediationAcceptResult = data.MediationAcceptResult
	}
	if data.MediationRejectResult != "" {
		txs.MediationRejectResult = data.MediationRejectResult
	}
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) RemoveJobOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.jobOfferMap, id)
	return nil
}

func (s *SolverStoreMemory) RemoveResourceOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.resourceOfferMap, id)
	return nil
}

func (s *SolverStoreMemory) RemoveDeal(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.dealMap, id)
	return nil
}

func (s *SolverStoreMemory) RemoveResult(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.resultMap, id)
	return nil
}

func (s *SolverStoreMemory) RemoveMatchDecision(resourceOffer string, jobOffer string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for k := range s.matchDecisionMap {
		if strings.Contains(k, jobOffer) || strings.Contains(k, resourceOffer) {
			delete(s.matchDecisionMap, k)
		}
	}
	return nil
}

func (store *SolverStoreMemory) RemoveAllowedResourceProvider(resourceProvider string) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	delete(store.allowedResourceProviderMap, resourceProvider)
	return nil
}

// Strictly speaking, the compiler will check the interface
// implementation without this check. But some code editors
// report errors more effectively when we have it.
var _ store.SolverStore = (*SolverStoreMemory)(nil)
