package stats

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/module/shortcuts"
)

type JobRun struct {
	DealID                        string          `json:"dealId"`
	ModuleID                      string          `json:"moduleId"`
	JobCreator                    string          `json:"jobCreator"`
	ResourceProvider              string          `json:"resourceProvider"`
	JobOffer                      string          `json:"jobOffer"`
	ResourceOffer                 string          `json:"resourceOffer"`
	JobState                      string          `json:"jobState"`
	MatchDurationMilliseconds     float64         `json:"matchDuration_ms"`
	ExecutionDurationMilliseconds float64         `json:"executionDuration_ms"`
	RetrievalDurationMilliseconds float64         `json:"retrievalDuration_ms"`
	TotalDurationMilliseconds     float64         `json:"totalDuration_ms"`
	ExtraData                     json.RawMessage `json:"extraData"`
}

// Stats API implementation

func (stat *HTTPStats) PostJobRun(deal *data.DealContainer) error {
	// The time from job offer creation until a match is made
	matchDuration := int64(deal.CreatedAt) - int64(deal.Deal.JobOffer.CreatedAt)

	// The time from when a match is made until job outputs are uploaded
	var executionDuration int64
	if deal.UploadAt > 0 {
		executionDuration = int64(deal.UploadAt) - int64(deal.CreatedAt)
	}

	// The time from when resource provider upload to job creator download
	var retrievalDuration int64
	// The total job run duration until download or failure
	var jobRunDuration int64
	if deal.DownloadAt > 0 {
		retrievalDuration = int64(deal.DownloadAt) - int64(deal.UploadAt)
		jobRunDuration = int64(deal.DownloadAt) - int64(deal.Deal.JobOffer.CreatedAt)
	} else {
		// The job failed, so we record the job run duration until now
		jobRunDuration = time.Now().UnixMilli() - int64(deal.Deal.JobOffer.CreatedAt)
	}

	jobRun := JobRun{
		DealID:                        deal.ID,
		ModuleID:                      shortcuts.GetShortcut(deal.Deal.JobOffer.Module.Repo, deal.Deal.JobOffer.Module.Hash),
		ResourceProvider:              deal.ResourceProvider,
		JobCreator:                    deal.JobCreator,
		ResourceOffer:                 deal.ResourceOffer,
		JobOffer:                      deal.JobOffer,
		JobState:                      data.GetAgreementStateString(deal.State),
		MatchDurationMilliseconds:     float64(matchDuration),
		ExecutionDurationMilliseconds: float64(executionDuration),
		RetrievalDurationMilliseconds: float64(retrievalDuration),
		TotalDurationMilliseconds:     float64(jobRunDuration),
	}

	_, err := http.PostRequest[JobRun, JobRun](stat.ClientOptions, "/job-run", jobRun)
	if err != nil {
		return fmt.Errorf("failed to post job run: %s", err)
	}

	return nil
}

// Noop implementation

func (stat *NoopStats) PostJobRun(deal *data.DealContainer) error {
	return nil
}
