package web3

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/system"
)

type Web3Options struct {
	RpcURL     string `json:"rpc_url"`
	PrivateKey string `json:"private_key"`
	ChainID    int    `json:"chain_id"`

	// the deplpoyed contract addresses
	ControllerAddress string `json:"controller_address"`
	PaymentsAddress   string `json:"payments_address"`
	StorageAddress    string `json:"storage_address"`
	TokenAddress      string `json:"token_address"`

	// the addresses of services we want to connect to
	SolverAddress    string `json:"solver_address"`
	DirectoryAddress string `json:"directory_address"`
}

type EventChannelCollection interface {
	Start(
		sdk *Web3SDK,
		ctx context.Context,
		cm *system.CleanupManager,
	) error
}
