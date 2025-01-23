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

type ResourceProviderOfferOptions struct {
	// ... [keep existing fields]
	OfferSpec       data.MachineSpec
	OfferCount      int
	Specs           []data.MachineSpec
	Modules         []string
	Mode            data.PricingMode
	DefaultPricing  data.DealPricing
	DefaultTimeouts data.DealTimeouts
	ModulePricing   map[string]data.DealPricing
	ModuleTimeouts  map[string]data.DealTimeouts
	Services        data.ServiceConfig
}

type ResourceProviderPowOptions struct {
	// ... [keep existing fields]
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
	Preflight preflight.PreflightConfig // Add preflight config
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
	checker := preflight.NewPreflightChecker()
	return checker.RunAllChecks(ctx, config)
}

func (resourceProvider *ResourceProvider) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	errorChan := make(chan error, 1)

	// Update GPU info if enabled
	if resourceProvider.options.Preflight.GPU.Enabled {
		checker := preflight.NewPreflightChecker()
		gpuInfo, err := checker.GetGPUInfo(ctx)
		if err != nil {
			errorChan <- fmt.Errorf("failed to get GPU information: %w", err)
			return errorChan
		}
		resourceProvider.gpuInfo = gpuInfo
	}

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
