package web3

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	lilypadproxy "github.com/lilypad-tech/lilypad/v2/pkg/web3/bindingsV2/LilypadProxy"
)

func (sdk *Web3SDKV2) SaveResult(result lilypadproxy.SharedStructsResult) (bool, error) {
	transaction, err := sdk.Contracts.LilypadProxy.SetResult(sdk.TransactOpts, result)
	if err != nil {
		sdk.Log.Error().Err(err).
			Str("resultId", result.ResultId).
			Str("dealId", result.DealId).
			Str("resultCID", result.ResultCID).
			Msg("failed to save result")
		return false, err
	}

	_, err = sdk.WaitTx(context.Background(), transaction)
	if err != nil {
		sdk.Log.Error().Err(err).
			Str("transactionHash", transaction.Hash().Hex()).
			Str("resultId", result.ResultId).
			Str("dealId", result.DealId).
			Str("resultCID", result.ResultCID).
			Msg("failed while executing transaction to save result")
		return false, err
	}

	return true, nil
}

func (sdk *Web3SDKV2) GetResult(resultId string) (lilypadproxy.SharedStructsResult, error) {
	result, err := sdk.Contracts.LilypadProxy.GetResult(sdk.CallOpts, resultId)
	if err != nil {
		sdk.Log.Error().Err(err).Str("resultId", resultId).Msg("failed to get result")
		return lilypadproxy.SharedStructsResult{}, err
	}

	return result, nil
}

func (sdk *Web3SDKV2) SaveDeal(deal lilypadproxy.SharedStructsDeal) (bool, error) {
	transaction, err := sdk.Contracts.LilypadProxy.SetDeal(sdk.TransactOpts, deal)
	if err != nil {
		sdk.Log.Error().Err(err).
			Str("dealId", deal.DealId).
			Str("jobCreator", deal.JobCreator.String()).
			Str("resourceProvider", deal.ResourceProvider.String()).
			Str("jobOfferCID", deal.JobOfferCID).
			Str("resourceOfferCID", deal.ResourceOfferCID).
			Msg("failed to save deal")
		return false, err
	}

	_, err = sdk.WaitTx(context.Background(), transaction)
	if err != nil {
		sdk.Log.Error().Err(err).
			Str("transactionHash", transaction.Hash().Hex()).
			Str("jobCreator", deal.JobCreator.String()).
			Str("resourceProvider", deal.ResourceProvider.String()).
			Str("jobOfferCID", deal.JobOfferCID).
			Str("resourceOfferCID", deal.ResourceOfferCID).
			Msg("failed while executing transaction to save deal")
		return false, err
	}

	return true, nil
}

func (sdk *Web3SDKV2) GetDeal(dealId string) (lilypadproxy.SharedStructsDeal, error) {
	deal, err := sdk.Contracts.LilypadProxy.GetDeal(sdk.CallOpts, dealId)
	if err != nil {
		sdk.Log.Error().Err(err).Str("dealId", dealId).Msg("failed to get deal")
		return lilypadproxy.SharedStructsDeal{}, err
	}

	return deal, nil
}

func (sdk *Web3SDKV2) AcceptJobPayment(amount *big.Int) (bool, error) {
	transaction, err := sdk.Contracts.LilypadProxy.AcceptJobPayment(sdk.TransactOpts, amount)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("failed to accept job creator collateral")
		return false, err
	}

	_, err = sdk.WaitTx(context.Background(), transaction)
	if err != nil {
		sdk.Log.Error().Err(err).Str("transactionHash", transaction.Hash().String()).Msg("failed to accept job creator collateral after submitting transaction")
		return false, err
	}

	return true, nil
}

func (sdk *Web3SDKV2) AcceptResourceProviderCollateral(amount *big.Int) (bool, error) {
	transaction, err := sdk.Contracts.LilypadProxy.AcceptResourceProviderCollateral(sdk.TransactOpts, amount)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("failed to accept resource provider collateral")
		return false, err
	}

	_, err = sdk.WaitTx(context.Background(), transaction)
	if err != nil {
		sdk.Log.Error().Err(err).Str("transactionHash", transaction.Hash().String()).Msg("failed to accept resource provider collateral after submitting transaction")
		return false, err
	}

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

func (sdk *Web3SDKV2) GetMinimumResourceProviderCollateralAmount() (*big.Int, error) {
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
