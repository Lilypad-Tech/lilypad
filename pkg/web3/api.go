package web3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/users"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/http"
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

func (sdk *Web3SDK) GetGenerateChallenge(
	ctx context.Context,
	nodeId string,
) (string, *pow.PowGenerateChallenge, error) {
	tx, err := sdk.Contracts.Pow.GenerateChallenge(
		sdk.TransactOpts,
		nodeId,
	)
	if err != nil {
		system.Error(sdk.Options.Service, "error submitting pow.GenerateChallenge", err)
		return "", nil, err
	} else {
		system.Debug(sdk.Options.Service, "submitted pow.GenerateChallenge", tx.Hash().String())
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

func (sdk *Web3SDK) SubmitWorkForBatching(ctx context.Context, nonce *big.Int, nodeId string) (common.Hash, error) {
	optsCopy := sdk.TransactOpts
	optsCopy.NoSend = true
	tx, err := sdk.Contracts.Pow.SubmitWork(optsCopy, nonce, nodeId)
	if err != nil {
		return common.Hash{}, err
	}
	//call the endpoint here
	req, err := createRequest(sdk.Options.PowBatchWsUrl, tx.Data())
	if err != nil {
		fmt.Errorf("failed to create request: %v", err)
	}

	err = executeRequest(req)
	if err != nil {
		fmt.Errorf("failed to execute request: %v", err)
	}

	fmt.Println("Transaction submitted successfully")
	return tx.Hash(), nil
}

func createRequest(url string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func executeRequest(req *http.Request) error {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
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
