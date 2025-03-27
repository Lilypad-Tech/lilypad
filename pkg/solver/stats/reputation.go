package stats

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/http"
)

type Reputation struct {
	JobMatched               *bool  `json:"job_matched"`
	JobCompletedNoValidation *bool  `json:"job_completed_no_validation"`
	ValidationLost           *bool  `json:"validation_lost"`
	RuntimeMillis            *int64 `json:"runtime_millis"`
	ModuleID                 string `json:"module_id"`
	JobFailed                *bool  `json:"job_failed"`
}

// The reputation builder constructs reputation from
// reputation events. For now, it does not expose runtime
// millis or module ID.
type ReputationBuilder struct {
	reputation Reputation
}

func NewReputationBuilder() *ReputationBuilder {
	return &ReputationBuilder{}
}

func (b *ReputationBuilder) WithJobMatched(val bool) *ReputationBuilder {
	b.reputation.JobMatched = &val
	return b
}

func (b *ReputationBuilder) WithJobCompletedNoValidation(val bool) *ReputationBuilder {
	b.reputation.JobCompletedNoValidation = &val
	return b
}

func (b *ReputationBuilder) WithValidationLost(val bool) *ReputationBuilder {
	b.reputation.ValidationLost = &val
	return b
}

func (b *ReputationBuilder) WithJobFailed(val bool) *ReputationBuilder {
	b.reputation.JobFailed = &val
	return b
}

func (b *ReputationBuilder) Build() Reputation {
	return b.reputation
}

// Stats API implementation

func (stat *HTTPStats) PostReputation(address string, reputation Reputation) error {
	path := fmt.Sprintf("/resource-provider/%s/reputation", address)
	_, err := http.PostRequest[Reputation, Reputation](stat.ClientOptions, path, reputation)
	if err != nil {
		log.Error().Err(err).Msg("failed to post reputation")
		return fmt.Errorf("failed to post reputation: %s", err)
	}

	return nil
}

// Noop implementation

func (stat *NoopStats) PostReputation(address string, reputation Reputation) error {
	return nil
}
