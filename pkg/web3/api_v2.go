package web3

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (sdk *Web3SDKV2) SaveResult(result interface{}, opts *bind.CallOpts) (bool, error) {
	//TODO: Implement method
	return true, nil
}

func (sdk *Web3SDKV2) GetResult(resultId string, opts *bind.CallOpts) (interface{}, error) {
	//TODO: Implement method
	return nil, nil
}

func (sdk *Web3SDKV2) SaveDeal(deal interface{}, opts *bind.CallOpts) (bool, error) {
	//TODO: implement method
	return true, nil
}

func (sdk *Web3SDKV2) GetDeal(dealId string, opts *bind.CallOpts) (interface{}, error) {
	//TODO: implement method
	return nil, nil
}

func (sdk *Web3SDKV2) AcceptJobPayment(amount *big.Int, opts *bind.CallOpts) (bool, error) {
	//TODO: implement method
	return true, nil
}

func (sdk *Web3SDKV2) acceptResourceProviderCollateral(amount *big.Int, opts *bind.CallOpts) (bool, error) {
	//TODO: implement method
	return true, nil
}

func (sdk *Web3SDKV2) ApproveTokenTransfer(amount *big.Int, spender common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	transaction, err := sdk.Contracts.LilypadToken.Approve(opts, spender, amount)
	if err != nil {
		sdk.Log.Error().Err(err).Str("spender", spender.String()).Str("amount", amount.String()).Msg("Failed to approve token transfer")
		return common.Hash{}, err
	}
	return transaction.Hash(), nil
}

func (sdk *Web3SDKV2) GetEscrowBalance(address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	balance, err := sdk.Contracts.LilypadProxy.GetEscrowBalance(opts, address)
	if err != nil {
		sdk.Log.Error().Err(err).Str("address", address.String()).Msg("Failed to get escrow balance")
		return nil, err
	}
	return balance, nil
}

func (sdk *Web3SDKV2) getMinimumResourceProviderCollateralAmount(opts *bind.CallOpts) (*big.Int, error) {
	minimumAmount, err := sdk.Contracts.LilypadProxy.GetMinimumResourceProviderCollateralAmount(opts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get minimum resource provider collateral amount")
		return nil, err
	}
	return minimumAmount, nil
}

func (sdk *Web3SDKV2) GetLilypadPaymentEngineAddress(opts *bind.CallOpts) (common.Address, error) {
	paymentEngineAddress, err := sdk.Contracts.LilypadProxy.GetPaymentEngineAddress(opts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get payment engine contract address")
		return common.Address{}, err
	}
	return paymentEngineAddress, nil
}

func (sdk *Web3SDKV2) GetLilypadUserAddress(opts *bind.CallOpts) (common.Address, error) {
	userAddress, err := sdk.Contracts.LilypadProxy.GetUserAddress(opts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get user contract address")
		return common.Address{}, err
	}
	return userAddress, nil
}

func (sdk *Web3SDKV2) GetLilypadStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	storageAddress, err := sdk.Contracts.LilypadProxy.GetStorageAddress(opts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get storage contract address")
		return common.Address{}, err
	}
	return storageAddress, nil
}

func (sdk *Web3SDKV2) GetLilypadL2TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	l2TokenAddress, err := sdk.Contracts.LilypadProxy.Getl2LilypadTokenAddress(opts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get l2LilypadToken contract address")
		return common.Address{}, err
	}
	return l2TokenAddress, nil
}
