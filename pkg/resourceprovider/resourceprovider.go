package resourceprovider

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/holiman/uint256"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/executor/bacalhau"
	"github.com/lilypad-tech/lilypad/pkg/ipfs"
	"github.com/lilypad-tech/lilypad/pkg/powLogs"
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider/preflight"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
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
	DisablePow         bool
	NumWorkers         int
	CudaGridSize       int
	CudaBlockSize      int
	CudaHashsPerThread int
}

type ResourceProviderOptions struct {
	Bacalhau  bacalhau.BacalhauExecutorOptions
	Offers    ResourceProviderOfferOptions
	Web3      web3.Web3Options
	Pow       ResourceProviderPowOptions
	IPFS      ipfs.IPFSOptions
	Telemetry system.TelemetryOptions
	Preflight preflight.PreflightConfig
}

type ResourceProvider struct {
	web3SDK    *web3.Web3SDK
	options    ResourceProviderOptions
	controller *ResourceProviderController
	gpuInfo    []preflight.GPUInfo
}

func NewResourceProvider(
	options ResourceProviderOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
	tracer trace.Tracer,
) (*ResourceProvider, error) {
	if err := runPreflightChecks(context.Background(), options.Preflight); err != nil {
		return nil, fmt.Errorf("preflight checks failed: %w", err)
	}

	controller, err := NewResourceProviderController(options, web3SDK, executor, tracer)
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

func runPreflightChecks(ctx context.Context, config preflight.PreflightConfig) error {
	log.Info().Msg("Starting preflight checks...")
	checker := preflight.NewPreflightChecker()

	// Logging GPU requirements
	log.Info().
		Int64("min_memory_gb", config.GPU.MinMemoryGB).
		Msg("GPU requirements")

	err := checker.RunAllChecks(ctx, config)
	if err != nil {
		log.Error().Err(err).Msg("Preflight checks failed")
		return err
	}

	log.Info().Msg("All preflight checks passed successfully")
	return nil
}

func (resourceProvider *ResourceProvider) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error, 1)

	checker := preflight.NewPreflightChecker()
	gpuInfo, err := checker.GetGPUInfo(ctx)
	if err != nil {
		// Instead of returning error, just log warning and continue
		log.Warn().Err(err).Msg("⚠️  GPU capabilities will not be available - continuing in CPU-only mode")
		// Set empty GPU info
		resourceProvider.gpuInfo = []preflight.GPUInfo{}
	} else {
		resourceProvider.gpuInfo = gpuInfo
		log.Info().Msgf("🎮 Successfully initialized with %d GPUs", len(gpuInfo))
	}

	if !resourceProvider.options.Pow.DisablePow {
		if errCh := resourceProvider.StartMineLoop(ctx); errCh != nil {
			// Forward any mining errors to the main error channel
			go func() {
				if err := <-errCh; err != nil {
					errorChan <- fmt.Errorf("mining error: %w", err)
				}
			}()
		}
	}

	// Start the controller and forward any errors
	controllerErrChan := resourceProvider.controller.Start(ctx, cm)
	go func() {
		if err := <-controllerErrChan; err != nil {
			errorChan <- err
		}
	}()

	return errorChan
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

		os.Setenv("WEB3_PRIVATE_KEY", resourceProvider.options.Web3.PrivateKey) //to ensure this env exit because binary need it but if user specific this with flag, no environment variable exit.
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
