package web3

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (sdk *Web3SDKV2) SaveResult(result interface{}) (bool, error) {
	//TODO: Implement method
	return true, nil
}

func (sdk *Web3SDKV2) GetResult(resultId string) (interface{}, error) {
	//TODO: Implement method
	return nil, nil
}

func (sdk *Web3SDKV2) SaveDeal(deal interface{}) (bool, error) {
	//TODO: implement method
	return true, nil
}

func (sdk *Web3SDKV2) GetDeal(dealId string) (interface{}, error) {
	//TODO: implement method
	return nil, nil
}

func (sdk *Web3SDKV2) AcceptJobPayment(amount *big.Int) (bool, error) {
	//TODO: implement method
	return true, nil
}

func (sdk *Web3SDKV2) acceptResourceProviderCollateral(amount *big.Int) (bool, error) {
	//TODO: implement method
	return true, nil
}

func (sdk *Web3SDKV2) ApproveTokenTransfer(amount *big.Int, spender common.Address) (string, error) {
	transaction, err := sdk.Contracts.LilypadToken.Approve(sdk.TransactOpts, spender, amount)
	if err != nil {
		sdk.Log.Error().Err(err).Str("spender", spender.String()).Str("amount", amount.String()).Msg("Failed to approve token transfer")
		return "", err
	}

	_, err = sdk.WaitTx(context.Background(), transaction)
	if err != nil {
		sdk.Log.Error().Err(err).Str("spender", spender.String()).Str("amount", amount.String()).Str("TransactionHash", transaction.Hash().String()).Msg("Transaction failed attempting to approve token transfer")
		return "", err
	}

	return transaction.Hash().String(), nil
}

func (sdk *Web3SDKV2) GetEscrowBalance(address common.Address) (*big.Int, error) {
	balance, err := sdk.Contracts.LilypadProxy.GetEscrowBalance(sdk.CallOpts, address)
	if err != nil {
		sdk.Log.Error().Err(err).Str("address", address.String()).Msg("Failed to get escrow balance")
		return nil, err
	}
	return balance, nil
}

func (sdk *Web3SDKV2) getMinimumResourceProviderCollateralAmount() (*big.Int, error) {
	minimumAmount, err := sdk.Contracts.LilypadProxy.GetMinimumResourceProviderCollateralAmount(sdk.CallOpts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get minimum resource provider collateral amount")
		return nil, err
	}
	return minimumAmount, nil
}

func (sdk *Web3SDKV2) GetLilypadPaymentEngineAddress() (common.Address, error) {
	paymentEngineAddress, err := sdk.Contracts.LilypadProxy.GetPaymentEngineAddress(sdk.CallOpts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get payment engine contract address")
		return common.Address{}, err
	}
	return paymentEngineAddress, nil
}

func (sdk *Web3SDKV2) GetLilypadUserAddress() (common.Address, error) {
	userAddress, err := sdk.Contracts.LilypadProxy.GetUserAddress(sdk.CallOpts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get user contract address")
		return common.Address{}, err
	}
	return userAddress, nil
}

func (sdk *Web3SDKV2) GetLilypadStorageAddress() (common.Address, error) {
	storageAddress, err := sdk.Contracts.LilypadProxy.GetStorageAddress(sdk.CallOpts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get storage contract address")
		return common.Address{}, err
	}
	return storageAddress, nil
}

func (sdk *Web3SDKV2) GetLilypadL2TokenAddress() (common.Address, error) {
	l2TokenAddress, err := sdk.Contracts.LilypadProxy.Getl2LilypadTokenAddress(sdk.CallOpts)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to get l2LilypadToken contract address")
		return common.Address{}, err
	}
	return l2TokenAddress, nil
}
