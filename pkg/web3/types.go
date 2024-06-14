package web3

import (
	"context"

	"github.com/lilypad-tech/lilypad/pkg/system"
)

type Web3Options struct {

	// core settings
	RpcURL     string `json:"rpc_url" toml:"rpc_url"`
	PrivateKey string `json:"private_key" toml:"private_key"`
	ChainID    int    `json:"chain_id" toml:"chain_id"`

	// contract addresses
	ControllerAddress string `json:"controller_address" toml:"controller_address"`
	PaymentsAddress   string `json:"payments_address" toml:"payments_address"`
	StorageAddress    string `json:"storage_address" toml:"storage_address"`
	UsersAddress      string `json:"users_address" toml:"users_address"`
	TokenAddress      string `json:"token_address" toml:"token_address"`
	MediationAddress  string `json:"mediation_address" toml:"mediation_address"`
	JobCreatorAddress string `json:"jobcreator_address" toml:"jobcreator_address"`
	PowAddress        string `json:"pow_address" toml:"pow_address`
	// this is injected by whatever service we are running
	// it's used for logging tx's
	Service system.Service `json:"-" toml:"-"`
}

type EventChannelCollection interface {
	Start(
		sdk *Web3SDK,
		ctx context.Context,
		cm *system.CleanupManager,
	) error
}
