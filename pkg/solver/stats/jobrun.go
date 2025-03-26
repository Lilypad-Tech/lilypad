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
	jobRunTime := time.Now().UnixMilli() - int64(deal.Deal.JobOffer.CreatedAt)
	extraData, err := json.Marshal(map[string]string{
		"state": data.GetAgreementStateString(deal.State),
	})
	if err != nil {
		stat.log.Error().Err(err).Msg("unable to marshal extra data")
		return err
	}

	jobRun := JobRun{
		DealID:                    deal.ID,
		ModuleID:                  shortcuts.GetShortcut(deal.Deal.JobOffer.Module.Repo, deal.Deal.JobOffer.Module.Hash),
		ResourceProvider:          deal.ResourceProvider,
		JobCreator:                deal.JobCreator,
		TotalDurationMilliseconds: float64(jobRunTime),
		ExtraData:                 extraData,
		JobOffer:                  deal.JobOffer,
		ResourceOffer:             deal.ResourceOffer,
	}

	_, err = http.PostRequest[JobRun, JobRun](stat.ClientOptions, "/job-run", jobRun)
	if err != nil {
		return fmt.Errorf("failed to post job run: %s", err)
	}

	return nil
}

// Noop implementation

func (stat *NoopStats) PostJobRun(deal *data.DealContainer) error {
	return nil
}
