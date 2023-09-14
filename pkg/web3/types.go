package web3

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
)

type Web3Options struct {

	// core settings
	RpcURL     string `json:"rpc_url"`
	PrivateKey string `json:"private_key"`
	ChainID    int    `json:"chain_id"`

	// contract addresses
	ControllerAddress string `json:"controller_address"`
	PaymentsAddress   string `json:"payments_address"`
	StorageAddress    string `json:"storage_address"`
	UsersAddress      string `json:"users_address"`
	MediationAddress  string `json:"mediation_address"`
	TokenAddress      string `json:"token_address"`

	// service addresses
	SolverAddress string `json:"solver_address"`
}

type EventChannelCollection interface {
	Start(
		sdk *Web3SDK,
		ctx context.Context,
		cm *system.CleanupManager,
	) error
}
