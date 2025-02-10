package stats

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/module/shortcuts"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/rs/zerolog/log"
)

type JobRun struct {
	DealID                    string          `json:"dealId"`
	ModuleID                  string          `json:"moduleId"`
	ResourceProvider          string          `json:"resourceProvider"`
	JobCreator                string          `json:"jobCreator"`
	TotalDurationMilliseconds float64         `json:"totalDuration_ms"`
	ExtraData                 json.RawMessage `json:"extraData"`
	JobOffer                  string          `json:"jobOffer"`
	ResourceOffer             string          `json:"resourceOffer"`
}

// Stats API implementation

func (stat *HTTPStats) PostJobRun(store store.SolverStore, deal *data.DealContainer) error {
	jobOffer, err := store.GetJobOffer(deal.Deal.JobOffer.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get job offer")
		return err
	}
	jobRunTime := time.Now().UnixMilli() - int64(jobOffer.JobOffer.CreatedAt)

	extraData, err := json.Marshal(map[string]string{
		"state": data.GetAgreementStateString(deal.State),
	})
	if err != nil {
		log.Error().Msgf("unable to marshal extra data: %s", err)
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

func (stat *NoopStats) PostJobRun(store store.SolverStore, deal *data.DealContainer) error {
	return nil
}
