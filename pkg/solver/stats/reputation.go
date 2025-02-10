package stats

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
