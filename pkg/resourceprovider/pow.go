package resourceprovider

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/powtoken"
	"github.com/rs/zerolog/log"
)

var (
	bigOne = big.NewInt(1)
)

func Mine(ctx context.Context, challenge [32]byte, minerAddress common.Address, difficulty *big.Int) (*big.Int, error) {
	nonce := big.NewInt(0)
	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("cancel by context")
		default:
			// Combine miner address, nonce, and challenge string
			// must compatable with contrack
			//
			data, err := formatMinerArgs(challenge, minerAddress, nonce)
			if err != nil {
				return nil, err
			}

			// Calculate Keccak-256 hash
			hashResult := crypto.Keccak256(data)

			// Check if the hash is below the target difficulty
			hashNumber := new(big.Int).SetBytes(hashResult)
			fmt.Println("number ", hashNumber.Text(2))
			fmt.Println("difficulty", difficulty.Text(2))
			if hashNumber.Cmp(difficulty) == -1 {
				log.Info().Msgf("Success! Nonce: %d, Elapsed Time: %d, Hash: %v", time.Since(startTime), nonce, hashResult)
				return nonce, nil
			}

			nonce.Add(nonce, bigOne)
		}
	}
}

func SubmitWork(ctx context.Context, sdk web3.Web3SDK, nonce *big.Int, nodeId string) (common.Hash, *powtoken.PowtokenValidPOWSubmitted, error) {
	return sdk.SubmitWork(ctx, nonce, nodeId)
}

func formatMinerArgs(challenge [32]byte, minerAddress common.Address, nonce *big.Int) ([]byte, error) {
	//todo use nonce in replace instead of building from scratch for better performance
	// keccak256(abi.encodePacked(lastChallenge, msg.sender, nodeId));
	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
	addressTy, _ := abi.NewType("address", "", nil)
	uint256Ty, _ := abi.NewType("uint256", "", nil)

	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: addressTy,
		},
		{
			Type: uint256Ty,
		},
	}

	bytes, err := arguments.Pack(
		challenge,
		minerAddress,
		nonce,
	)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
