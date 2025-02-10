package stats

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/module/shortcuts"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
)

type StatsOptions struct {
	Enabled bool
	URL     string
}

type Stats interface {
	PostJobRun(store store.SolverStore, deal *data.DealContainer) error
	PostReputation(address string, reputation Reputation) error
}

func NewStats(service system.Service, options StatsOptions, web3Options web3.Web3Options, web3SDK *web3.Web3SDK) (Stats, error) {
	if !options.Enabled {
		return &NoopStats{}, nil
	}

	return &HTTPStats{
		ClientOptions: http.ClientOptions{
			URL:           options.URL,
			PrivateKey:    web3Options.PrivateKey,
			Type:          string(service),
			PublicAddress: web3SDK.GetAddress().String(),
		}}, nil
}

// Stats API implementation

type HTTPStats struct {
	ClientOptions http.ClientOptions
}

func (stat *HTTPStats) PostJobRun(store store.SolverStore, deal *data.DealContainer) error {
	// ⚠️ The memory store reports the time when the job offer was created by the job creator.
	// The database store uses database timestamps to avoid forged timestamps from the job creator.
	// Avoid using the memory store when you need reliable job run time measurements.
	createdAt, err := store.GetJobOfferCreatedAt(deal.Deal.JobOffer.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get job offer created time")
		return err
	}
	jobRunTime := time.Now().UnixMilli() - createdAt

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
		Tflops:                    nil,
		TotalDurationMilliseconds: float64(jobRunTime),
		ExtraData:                 extraData,
		JobOfferCID:               deal.JobOffer,
		ResourceOfferCID:          deal.ResourceOffer,
	}

	_, err = http.PostRequest[JobRun, JobRun](stat.ClientOptions, "/job-run", jobRun)
	if err != nil {
		return fmt.Errorf("failed to post job run: %s", err)
	}

	return nil
}

func (stat *HTTPStats) PostReputation(address string, reputation Reputation) error {
	path := fmt.Sprintf("/resource-provider/%s/reputation", address)
	_, err := http.PostRequest[Reputation, Reputation](stat.ClientOptions, path, reputation)
	if err != nil {
		return fmt.Errorf("failed to post reputation: %s", err)
	}

	return nil
}

// Noop implementation

type NoopStats struct{}

func (stat *NoopStats) PostJobRun(store store.SolverStore, deal *data.DealContainer) error {
	return nil
}

func (stat *NoopStats) PostReputation(address string, reputation Reputation) error {
	return nil
}
