package web3

import (
	"fmt"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/controller"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/storage"
	"github.com/ethereum/go-ethereum/common"
)

func (sdk *Web3SDK) GetServiceAddresses(serviceType string) ([]common.Address, error) {
	solverType, err := data.GetServiceType(serviceType)
	if err != nil {
		return nil, err
	}
	return sdk.Contracts.Storage.ShowUsersInList(
		sdk.CallOpts,
		solverType,
	)
}

func (sdk *Web3SDK) GetSolverAddresses() ([]common.Address, error) {
	return sdk.GetServiceAddresses("Solver")
}

func (sdk *Web3SDK) GetUser(
	address common.Address,
) (storage.SharedStructsUser, error) {
	return sdk.Contracts.Storage.GetUser(
		sdk.CallOpts,
		address,
	)
}

func (sdk *Web3SDK) UpdateUser(
	metadataCID string,
	url string,
	roles []uint8,
) error {
	updateUserTx, err := sdk.Contracts.Storage.UpdateUser(
		sdk.TransactOpts,
		metadataCID,
		url,
		roles,
	)
	if err != nil {
		return err
	}
	_, err = sdk.waitTx(updateUserTx)
	if err != nil {
		return err
	}
	return nil
}

func (sdk *Web3SDK) AddUserToList(
	serviceType uint8,
) error {
	addToListTx, err := sdk.Contracts.Storage.AddUserToList(
		sdk.TransactOpts,
		serviceType,
	)
	if err != nil {
		return err
	}
	_, err = sdk.waitTx(addToListTx)
	if err != nil {
		return err
	}
	return nil
}

func (sdk *Web3SDK) GetSolverUrl(address string) (string, error) {
	solver, err := sdk.Contracts.Storage.GetUser(
		sdk.CallOpts,
		common.HexToAddress(address),
	)
	if err != nil {
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
	if len(deal.ID) != 32 {
		return "", fmt.Errorf("The deal ID must be exactly 32 bytes long: %s", deal.ID)
	}
	var dealID big.Int
	dealID.SetBytes([]byte(deal.ID))
	tx, err := sdk.Contracts.Controller.Agree(
		sdk.TransactOpts,
		&dealID,
		controller.SharedStructsDealMembers{},
		controller.SharedStructsDealTimeouts{},
		controller.SharedStructsDealPricing{},
	)
	if err != nil {
		return "", err
	}
	_, err = sdk.waitTx(tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
