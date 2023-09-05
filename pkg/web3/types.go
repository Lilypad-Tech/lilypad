package web3

type Web3Options struct {
	RpcURL     string `json:"rpc_url"`
	PrivateKey string `json:"private_key"`
	ChainID    int    `json:"chain_id"`

	// the deplpoyed contract addresses
	ControllerAddress string `json:"controller_address"`
	PaymentsAddress   string `json:"payments_address"`
	StorageAddress    string `json:"storage_address"`
	TokenAddress      string `json:"token_address"`
}
