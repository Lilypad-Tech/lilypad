package lilypad

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/holiman/uint256"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider"
	"github.com/spf13/cobra"
)

func newPowMockCmd() *cobra.Command {
	options := optionsfactory.GetDefaultResourceProviderPowOptions()

	var difficulty string
	resourceProviderCmd := &cobra.Command{
		Use:     "mock-pow",
		Short:   "Mock pow.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			waitCh := make(chan struct{})
			taskCh := make(chan resourceprovider.Task)
			minerCtl := resourceprovider.NewMinerController("mock node id", resourceprovider.ResourceProviderPowOptions{
				DisablePow:         true,
				NumWorkers:         options.NumWorkers,
				CudaGridSize:       options.CudaGridSize,
				CudaBlockSize:      options.CudaBlockSize,
				CudaHashsPerThread: options.CudaHashsPerThread,
			}, taskCh, func(nonce *big.Int, hashrate float64) {
				fmt.Println("receive solution %s, %f", nonce, hashrate)
				waitCh <- struct{}{}
			})

			minerCtl.Start(cmd.Context())

			challenge := [32]byte{}
			rand.Read(challenge[:])

			difficulty, _ := new(big.Int).SetString(difficulty, 10)
			fmt.Printf("challenge %s difficulty %s\n", hex.EncodeToString(challenge[:]), difficulty.String())
			taskCh <- resourceprovider.Task{
				Id:         uuid.New(),
				Challenge:  challenge,
				Difficulty: uint256.MustFromBig(difficulty),
			}
			<-waitCh
			return nil
		},
	}

	optionsfactory.AddResourceProviderPowCliFlags(resourceProviderCmd, &options)
	resourceProviderCmd.PersistentFlags().StringVar(
		&difficulty, "difficulty", "555460709263765739036470010701196062214039696708679004195670928130048",
		"difficulty  555460709263765739036470010701196062214039696708679004195670928130048(8M) 92576780592126171815437600338300430792573009392238517278497593884672(40M)",
	)
	return resourceProviderCmd
}
