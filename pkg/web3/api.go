package web3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/users"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
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
		system.Error(sdk.Options.Service, "error submitting Users.UpdateUser", err)
		return err
	} else {
		system.Info(sdk.Options.Service, "submitted users.UpdateUser", tx.Hash().String())
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
		system.Error(sdk.Options.Service, "error submitting Users.AddUserToList", err)
		return err
	} else {
		system.Info(sdk.Options.Service, "submitted users.AddUserToList", tx.Hash().String())
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}

func (sdk *Web3SDK) GetSolverUrl(address string) (string, error) {
	log.Debug().Msgf("begin GetSolverUrl from contract at address: %s", address)
	solver, err := sdk.Contracts.Users.GetUser(
		sdk.CallOpts,
		common.HexToAddress(address),
	)
	if err != nil {
		log.Error().Msgf("GetUser error")
		log.Error().Msgf("error: %s", err)
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
		system.Error(sdk.Options.Service, "error submitting controller.Agree() tx", err)
		return "", err
	} else {
		system.Debug(sdk.Options.Service, "submitted controller.Agree() tx", tx.Hash().String())
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
		system.Error(sdk.Options.Service, "error submitting controller.AddResult", err)
		return "", err
	} else {
		system.Debug(sdk.Options.Service, "submitted controller.AddResult", tx.Hash().String())
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
		system.Error(sdk.Options.Service, "error submitting controller.AcceptResult", err)
		return "", err
	} else {
		system.Debug(sdk.Options.Service, "submitted controller.AcceptResult", tx.Hash().String())
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
		system.Error(sdk.Options.Service, "error submitting controller.CheckResult", err)
		return "", err
	} else {
		system.Debug(sdk.Options.Service, "submitted controller.CheckResult", tx.Hash().String())
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
		system.Error(sdk.Options.Service, "error submitting controller.MediationAcceptResult", err)
		return "", err
	} else {
		system.Debug(sdk.Options.Service, "submitted controller.MediationAcceptResult", tx.Hash().String())
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
		system.Error(sdk.Options.Service, "error submitting controller.MediationRejectResult", err)
		return "", err
	} else {
		system.Debug(sdk.Options.Service, "submitted controller.MediationRejectResult", tx.Hash().String())
		system.DumpObjectDebug(tx)
	}
	_, err = sdk.WaitTx(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
