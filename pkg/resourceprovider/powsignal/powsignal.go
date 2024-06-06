package powsignal

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
)

type PowSignalOptions struct {
	Web3 web3.Web3Options
}

type ResourceProvider struct {
	web3SDK *web3.Web3SDK
	options PowSignalOptions
}

func SendPowSignal(ctx context.Context, web3SDK *web3.Web3SDK, options PowSignalOptions) (*pow.PowNewPowRound, error) {
	tx, err := web3SDK.Contracts.Pow.TriggerNewPowRound(web3SDK.TransactOpts)
	if err != nil {
		return nil, err
	}

	receipt, err := web3SDK.WaitTx(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("wait new pow siganl tx(%s) %w", tx.Hash(), err)
	}

	if receipt.Status == 0 {
		return nil, fmt.Errorf("send new pow signal successfully but execute fail status(%d) tx(%s)", receipt.Status, tx.Hash())
	}

	newPowRoundEvent, err := web3SDK.Contracts.Pow.ParseNewPowRound(*receipt.Logs[0])
	if err != nil {
		return nil, fmt.Errorf("parse new pow round event fail tx(%s) %w", tx.Hash(), err)
	}
	return newPowRoundEvent, nil
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

func GetPowSubmission(ctx context.Context, web3SDK *web3.Web3SDK) (map[common.Address][]PowValidPOWSubmission, error) {
	miners, err := web3SDK.Contracts.Pow.GetMiners(web3SDK.CallOpts)
	if err != nil {
		return nil, err
	}

	results := make(map[common.Address][]PowValidPOWSubmission)
	for _, minerAddr := range miners {
		validProofCount, err := web3SDK.Contracts.Pow.MinerSubmissionCount(web3SDK.CallOpts, minerAddr)
		if err != nil {
			return nil, err
		}

		var powSubmissions []PowValidPOWSubmission
		for i := uint64(0); i < validProofCount.Uint64(); i++ { //enough large
			submission, err := web3SDK.Contracts.Pow.PowSubmissions(web3SDK.CallOpts, minerAddr, new(big.Int).SetUint64(i))
			if err != nil {
				return nil, err
			}
			powSubmissions = append(powSubmissions, PowValidPOWSubmission{
				WalletAddress:    submission.WalletAddress,
				NodeId:           submission.NodeId,
				Nonce:            submission.Nonce,
				StartTimestap:    submission.StartTimestap,
				CompleteTimestap: submission.CompleteTimestap,
				Challenge:        submission.Challenge,
				Difficulty:       submission.Difficulty,
			})
		}
		results[minerAddr] = powSubmissions
	}
	return results, nil
}
