package stats

import (
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/system"

	"github.com/lilypad-tech/lilypad/pkg/web3"
)

type StatsOptions struct {
	Enabled bool
	URL     string
}

type Stats interface{}

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

// Noop implementation

type NoopStats struct{}
