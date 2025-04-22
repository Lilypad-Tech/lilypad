package web3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/pow"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/users"
	"github.com/ethereum/go-ethereum/common"
)

func (sdk *Web3SDK) GetServiceAddresses(serviceType string) ([]common.Address, error) {
	solverType, err := data.GetServiceType(serviceType)
	if err != nil {
		return nil, err
	}
	return sdk.Contracts.Users.ShowUsersInList(
		sdk.CallOpts,
		solverType,
	)
}

func (sdk *Web3SDK) GetSolverAddresses() ([]common.Address, error) {
	return sdk.GetServiceAddresses("Solver")
}

func (sdk *Web3SDK) GetUser(
	address common.Address,
) (users.SharedStructsUser, error) {
	return sdk.Contracts.Users.GetUser(
		sdk.CallOpts,
		address,
	)
}

func (sdk *Web3SDK) UpdateUser(
	metadataCID string,
	url string,
	roles []uint8,
) error {
	tx, err := sdk.Contracts.Users.UpdateUser(
		sdk.TransactOpts,
		metadataCID,
		url,
		roles,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting Users.UpdateUser")
		return err
	} else {
		sdk.Log.Info().Str("transactionHash", tx.Hash().String()).Msg("Submitted users.UpdateUser")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}

func (sdk *Web3SDK) AddUserToList(
	serviceType uint8,
) error {
	tx, err := sdk.Contracts.Users.AddUserToList(
		sdk.TransactOpts,
		serviceType,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("error submitting Users.AddUserToList")
		return err
	} else {
		sdk.Log.Info().Str("transactionHash", tx.Hash().String()).Msg("Submitted users.AddUserToList")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}

func (sdk *Web3SDK) GetSolverUrl(address string) (string, error) {
	sdk.Log.Debug().Str("address", address).Msg("Begin GetSolverUrl from contract address")
	solver, err := sdk.Contracts.Users.GetUser(
		sdk.CallOpts,
		common.HexToAddress(address),
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Failed to discover solver URL")
		return "", err
	}

	if solver.UserAddress == common.HexToAddress("0x0") {
		return "", fmt.Errorf("no solver found for address: %s", address)
	}
	return solver.Url, nil
}

func (sdk *Web3SDK) Agree(
	deal data.Deal,
) (string, error) {
	mediators := []common.Address{}
	for _, mediator := range deal.Members.Mediators {
		mediators = append(mediators, common.HexToAddress(mediator))
	}
	tx, err := sdk.Contracts.Controller.Agree(
		sdk.TransactOpts,
		deal.ID,
		data.ConvertDealMembers(deal.Members),
		data.ConvertDealTimeouts(deal.Timeouts),
		data.ConvertDealPricing(deal.Pricing),
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting controller.Agree() tx")
		return "", err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted controller.Agree() tx")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (sdk *Web3SDK) AddResult(
	dealId string,
	resultsId string,
	dataId string,
	instructionCount uint64,
) (string, error) {
	tx, err := sdk.Contracts.Controller.AddResult(
		sdk.TransactOpts,
		dealId,
		resultsId,
		dataId,
		big.NewInt(int64(instructionCount)),
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting controller.AddResult() tx")
		return "", err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted controller.AddResult() tx")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (sdk *Web3SDK) AcceptResult(
	dealId string,
) (string, error) {
	tx, err := sdk.Contracts.Controller.AcceptResult(
		sdk.TransactOpts,
		dealId,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting controller.AcceptResult() tx")
		return "", err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted controller.AcceptResult() tx")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (sdk *Web3SDK) CheckResult(
	dealId string,
) (string, error) {
	tx, err := sdk.Contracts.Controller.CheckResult(
		sdk.TransactOpts,
		dealId,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting controller.CheckResult() tx")
		return "", err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted controller.CheckResult() tx")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (sdk *Web3SDK) MediationAcceptResult(
	dealId string,
) (string, error) {
	tx, err := sdk.Contracts.Controller.MediationAcceptResult(
		sdk.TransactOpts,
		dealId,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting controller.MediationAcceptResult() tx")
		return "", err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted controller.MediationAcceptResult() tx")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (sdk *Web3SDK) MediationRejectResult(
	dealId string,
) (string, error) {
	tx, err := sdk.Contracts.Controller.MediationRejectResult(
		sdk.TransactOpts,
		dealId,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting controller.MediationRejectResult() tx")
		return "", err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted controller.MediationRejectResult() tx")
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (sdk *Web3SDK) GetGenerateChallenge(
	ctx context.Context,
	nodeId string,
) (string, *pow.PowGenerateChallenge, error) {
	tx, err := sdk.Contracts.Pow.GenerateChallenge(
		sdk.TransactOpts,
		nodeId,
	)
	if err != nil {
		sdk.Log.Error().Err(err).Msg("Error submitting pow.GenerateChallenge() tx")
		return "", nil, err
	} else {
		sdk.Log.Debug().Str("transactionHash", tx.Hash().String()).Msg("Submitted pow.GenerateChallenge() tx")
		system.DumpObjectDebug(tx)
	}
	receipt, err := sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", nil, err
	}

	if receipt.Status == 0 {
		return tx.Hash().String(), nil, fmt.Errorf("execute challenge fail")
	}

	challenge, err := sdk.Contracts.Pow.ParseGenerateChallenge(*receipt.Logs[0])
	if err != nil {
		return "", nil, err
	}
	return tx.Hash().String(), challenge, nil
}

func (sdk *Web3SDK) SubmitWork(
	ctx context.Context,
	nonce *big.Int,
	nodeId string,
) (common.Hash, error) {
	tx, err := sdk.Contracts.Pow.SubmitWork(sdk.TransactOpts, nonce, nodeId)
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := sdk.WaitTx(ctx, tx)
	if err != nil {
		return common.Hash{}, err
	}

	if receipt.Status == 0 {
		return tx.Hash(), fmt.Errorf("excute transaction fail")
	}

	return tx.Hash(), nil
}

type PowValidPOWSubmission struct {
	WalletAddress    common.Address
	NodeId           string
	Nonce            *big.Int
	StartTimestap    *big.Int
	CompleteTimestap *big.Int
	Challenge        [32]byte
	Difficulty       *big.Int
}

func (sdk *Web3SDK) SendPowSignal(ctx context.Context) (*pow.PowNewPowRound, error) {
	tx, err := sdk.Contracts.Pow.TriggerNewPowRound(sdk.TransactOpts)
	if err != nil {
		return nil, err
	}

	receipt, err := sdk.WaitTx(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("wait new pow siganl tx(%s) %w", tx.Hash(), err)
	}

	if receipt.Status == 0 {
		return nil, fmt.Errorf("send new pow signal successfully but execute fail status(%d) tx(%s)", receipt.Status, tx.Hash())
	}

	newPowRoundEvent, err := sdk.Contracts.Pow.ParseNewPowRound(*receipt.Logs[0])
	if err != nil {
		return nil, fmt.Errorf("parse new pow round event fail tx(%s) %w", tx.Hash(), err)
	}
	return newPowRoundEvent, nil
}
