package stats

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/http"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3"
	"github.com/rs/zerolog"
)

type StatsOptions struct {
	Enabled bool
	URL     string
}

type Stats interface {
	PostJobRun(deal *data.DealContainer) error
	PostReputation(address string, reputation Reputation) error
}

var log = system.GetLogger(system.SolverService)

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
		},
		log: system.GetLogger(system.SolverService),
	}, nil
}

// Stats API implementation

type HTTPStats struct {
	ClientOptions http.ClientOptions
	log           *zerolog.Logger
}

// Noop implementation

type NoopStats struct{}
