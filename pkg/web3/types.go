package web3

import (
	"context"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
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
	PowAddress        string `json:"pow_address" toml:"pow_address"`

	// V2 Protocol
	LilypadProxyAddress            string `json:"lilypad_proxy_address" toml:"lilypad_proxy_address"`
	LilypadPaymentEngineAddress    string `json:"lilypad_payment_engine_address" toml:"lilypad_payment_engine_address"`
	LilypadStorageAddress          string `json:"lilypad_storage_address" toml:"lilypad_storage_address"`
	LilypadUserAddress             string `json:"lilypad_user_address" toml:"lilypad_user_address"`
	LilypadL2TokenAddress          string `json:"lilypad_l2_token_address" toml:"lilypad_l2_token_address"`
	LilypadL1TokenAddress          string `json:"lilypad_l1_token_address" toml:"lilypad_l1_token_address"`
	LilypadTokenomicsAddress       string `json:"lilypad_tokenomics_address" toml:"lilypad_tokenomics_address"`
	LilypadModuleDirectoryAddress  string `json:"lilypad_module_directory_address" toml:"lilypad_module_directory_address"`
	LilypadContractRegistryAddress string `json:"lilypad_contract_registry_address" toml:"lilypad_contract_registry_address"`

	// this is injected by whatever service we are running
	// it's used for logging tx's
	Service system.Service `json:"-" toml:"-"`
}

type EventChannelCollection interface {
	Start(
		ctx context.Context,
		cm *system.CleanupManager,
		sdk *Web3SDK,
	) error
}
