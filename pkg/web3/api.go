package web3

import (
	"fmt"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/ethereum/go-ethereum/common"
)

func (sdk *Web3SDK) GetServiceAddresses(serviceType string) ([]common.Address, error) {
	solverType, err := data.GetServiceType(serviceType)
	if err != nil {
		return nil, err
	}
	fmt.Printf("List type: %d\n", solverType)
	return sdk.Contracts.Storage.ShowUsersInList(
		sdk.CallOpts,
		solverType,
	)
}

func (sdk *Web3SDK) GetSolverAddresses() ([]common.Address, error) {
	return sdk.GetServiceAddresses("Solver")
}

func (sdk *Web3SDK) GetDirectoryAddresses() ([]common.Address, error) {
	return sdk.GetServiceAddresses("Directory")
}

func (sdk *Web3SDK) AddSolver(
	metadataCID *big.Int,
	url string,
	trustedMediators []common.Address,
	trustedDirectories []common.Address,
) error {
	solverType, err := data.GetServiceType("Solver")
	if err != nil {
		return err
	}
	tx, err := sdk.Contracts.Storage.UpdateUser(
		sdk.TransactOpts,
		metadataCID,
		url,
		[]uint8{solverType},
		trustedMediators,
		trustedDirectories,
	)
	if err != nil {
		return err
	}
	_, err = sdk.waitTx(tx)
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
