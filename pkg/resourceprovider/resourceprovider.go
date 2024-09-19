package resourceprovider

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/holiman/uint256"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/executor/bacalhau"
	"github.com/lilypad-tech/lilypad/pkg/powLogs"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/rs/zerolog/log"
)

// this configures the resource offers we will keep track of
type ResourceProviderOfferOptions struct {
	// if we are configuring a single machine then
	// these values are populated by the flags
	OfferSpec data.MachineSpec
	// we can dupliate the single spec to create a list of specs
	OfferCount int
	// this represents how many machines we will keep
	// offering to the network
	// we can configure this with a config file
	// to start with we will just add --cpu --gpu and --ram flags
	// to the resource provider CLI which constrains them to a single machine
	Specs []data.MachineSpec
	// the list of modules we are willing to run
	// an empty list means anything
	Modules []string

	// this will normally be FixedPrice for RP's
	Mode data.PricingMode

	// the default pricing for this resource provider
	// for all modules that don't have a specific price
	DefaultPricing  data.DealPricing
	DefaultTimeouts data.DealTimeouts

	// allow different pricing for different modules
	ModulePricing  map[string]data.DealPricing
	ModuleTimeouts map[string]data.DealTimeouts

	// which mediators and directories this RP will trust
	Services data.ServiceConfig
}

// this configures the pow we will keep track of
type ResourceProviderPowOptions struct {
	DisablePow bool
	NumWorkers int

	CudaGridSize       int
	CudaBlockSize      int
	CudaHashsPerThread int
}

type ResourceProviderOptions struct {
	Bacalhau   bacalhau.BacalhauExecutorOptions
	Offers     ResourceProviderOfferOptions
	Web3       web3.Web3Options
	Pow        ResourceProviderPowOptions
	Telemetry  system.TelemetryOptions
	Standalone bool
}

type ResourceProvider struct {
	web3SDK    *web3.Web3SDK
	options    ResourceProviderOptions
	controller *ResourceProviderController
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
	telemetry system.Telemetry,
) (*ResourceProvider, error) {

	controller, err := NewResourceProviderController(options, web3SDK, executor, telemetry)
	if err != nil {
		return nil, err
	}
	solver := &ResourceProvider{
		controller: controller,
		options:    options,
		web3SDK:    web3SDK,
	}
	powLogs.Init(options.Offers.Services.APIHost)
	return solver, nil
}

func (resourceProvider *ResourceProvider) Start(ctx context.Context, cm *system.CleanupManager) chan error {

	if !resourceProvider.options.Pow.DisablePow {
		if errCh := resourceProvider.StartMineLoop(ctx); errCh != nil {
			return errCh
		}
	}
	return resourceProvider.controller.Start(ctx, cm)
}

func (resourceProvider *ResourceProvider) StartMineLoop(ctx context.Context) chan error {
	errorChan := make(chan error, 1)
	walletAddress := resourceProvider.web3SDK.GetAddress()
	nodeId, err := resourceProvider.controller.executor.Id()
	if err != nil {
		errorChan <- err
		return errorChan
	}
	log.Info().Msgf("Wallet %s node id %s is ready for mine", walletAddress, nodeId)

	taskCh := make(chan Task)
	resourceProvider.controller.web3Events.Pow.SubscribenewPowRound(func(newPowRound pow.PowNewPowRound) {

		_, challenge, err := resourceProvider.web3SDK.GetGenerateChallenge(ctx, nodeId)

		if err != nil {
			log.Err(err).Msgf("Unable to fetch challenge")
			return
		}

		log.Info().Msgf("Receive a new pow signal challenge hex %s, difficulty %s", "0x"+hex.EncodeToString(challenge.Challenge[:]), challenge.Difficulty.String())

		difficulty, _ := uint256.FromBig(challenge.Difficulty)
		uuid := uuid.New()

		err = PostCard(uuid.String(), "0x"+hex.EncodeToString(challenge.Challenge[:]), challenge.Difficulty.String())
		if err != nil {
			log.Err(err).Msgf("Unable to post card")
		}

		taskCh <- Task{
			Id:         uuid,
			Challenge:  challenge.Challenge,
			Difficulty: difficulty,
		}

	})

	submitWork := func(nonce *big.Int, hashrate float64) {
		finishTime := time.Now().Unix()
		id := walletAddress.String() + strconv.FormatInt(finishTime, 10)
		powLogs.TrackHashrate(data.MinerHashRate{
			ID:       id,
			Address:  walletAddress.String(),
			Date:     finishTime,
			Hashrate: hashrate,
		})
		txId, err := resourceProvider.web3SDK.SubmitWork(ctx, nonce, nodeId)
		if err != nil {
			log.Err(err).Msgf("Submit work fail")
			return
		}
		log.Info().Str("address", walletAddress.Hex()).
			Str("nodeid", nodeId).
			Str("Nonce", nonce.String()).
			Str("txid", txId.String()).
			Msgf("Mine and submit successfully")
	}

	log.Info().Msgf("Listen to new pow round signal")
	miner := NewMinerController(nodeId, resourceProvider.options.Pow, taskCh, submitWork)
	go miner.Start(ctx)
	return nil
}

func TriggerNewPowRound(ctx context.Context, web3SDK *web3.Web3SDK) (common.Hash, error) {
	tx, err := web3SDK.Contracts.Pow.TriggerNewPowRound(web3SDK.TransactOpts)
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := web3SDK.WaitTx(ctx, tx)
	if err != nil {
		return common.Hash{}, err
	}

	if receipt.Status != 1 {
		return tx.Hash(), fmt.Errorf("trigger new pow round")
	}
	return tx.Hash(), nil
}
