package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/executor/noop"
	"github.com/bacalhau-project/lilypad/pkg/jobcreator"
	"github.com/bacalhau-project/lilypad/pkg/mediator"
	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/resourceprovider"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	solvermemorystore "github.com/bacalhau-project/lilypad/pkg/solver/store/memory"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type testOptions struct {
	moderationChance int
	executor         noop.NoopExecutorOptions
}

func getSolver(t *testing.T, options testOptions) (*solver.Solver, error) {
	solverOptions := optionsfactory.NewSolverOptions()
	solverOptions.Web3.PrivateKey = os.Getenv("SOLVER_PRIVATE_KEY")
	solverOptions.Server.Port = 8080
	solverOptions.Server.URL = "http://localhost:8080"

	// test that the solver private key is defined
	if solverOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("SOLVER_PRIVATE_KEY is not defined")
	}

	web3SDK, err := web3.NewContractSDK(solverOptions.Web3)
	if err != nil {
		return nil, err
	}

	solverStore, err := solvermemorystore.NewSolverStoreMemory()
	if err != nil {
		return nil, err
	}

	return solver.NewSolver(solverOptions, solverStore, web3SDK)
}

func getResourceProvider(
	t *testing.T,
	systemContext *system.CommandContext,
	options testOptions,
) (*resourceprovider.ResourceProvider, error) {
	resourceProviderOptions := optionsfactory.NewResourceProviderOptions()
	resourceProviderOptions.Web3.PrivateKey = os.Getenv("RESOURCE_PROVIDER_PRIVATE_KEY")
	if resourceProviderOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("RESOURCE_PROVIDER_PRIVATE_KEY is not defined")
	}
	resourceProviderOptions, err := optionsfactory.ProcessResourceProviderOptions(resourceProviderOptions)
	if err != nil {
		return nil, err
	}

	web3SDK, err := web3.NewContractSDK(resourceProviderOptions.Web3)
	if err != nil {
		return nil, err
	}

	executor, err := noop.NewNoopExecutor(options.executor)
	if err != nil {
		return nil, err
	}

	return resourceprovider.NewResourceProvider(resourceProviderOptions, web3SDK, executor)
}

func getMediator(
	t *testing.T,
	systemContext *system.CommandContext,
	options testOptions,
) (*mediator.Mediator, error) {
	mediatorOptions := optionsfactory.NewMediatorOptions()
	mediatorOptions.Web3.PrivateKey = os.Getenv("MEDIATOR_PRIVATE_KEY")
	if mediatorOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("MEDIATOR_PRIVATE_KEY is not defined")
	}
	mediatorOptions, err := optionsfactory.ProcessMediatorOptions(mediatorOptions)
	if err != nil {
		return nil, err
	}

	web3SDK, err := web3.NewContractSDK(mediatorOptions.Web3)
	if err != nil {
		return nil, err
	}

	options.executor.BadActor = false

	executor, err := noop.NewNoopExecutor(options.executor)
	if err != nil {
		return nil, err
	}

	return mediator.NewMediator(mediatorOptions, web3SDK, executor)
}

func getJobCreatorOptions(options testOptions) (jobcreator.JobCreatorOptions, error) {
	jobCreatorOptions := optionsfactory.NewJobCreatorOptions()
	jobCreatorOptions.Web3.PrivateKey = os.Getenv("JOB_CREATOR_PRIVATE_KEY")
	if jobCreatorOptions.Web3.PrivateKey == "" {
		return jobCreatorOptions, fmt.Errorf("JOB_CREATOR_PRIVATE_KEY is not defined")
	}
	ret, err := optionsfactory.ProcessJobCreatorOptions(jobCreatorOptions, []string{
		// this should point to the shortcut
		"cowsay:v0.0.2",
	})

	if err != nil {
		return ret, err
	}

	jobCreatorOptions.Mediation.CheckResultsPercentage = options.moderationChance
	return ret, nil
}

func testStackWithOptions(
	t *testing.T,
	commandCtx *system.CommandContext,
	options testOptions,
) (*jobcreator.RunJobResults, error) {

	solver, err := getSolver(t, options)
	if err != nil {
		return nil, err
	}

	solver.Start(commandCtx.Ctx, commandCtx.Cm)

	// give the solver server a chance to boot before we get all the websockets
	// up and trying to connect to it
	time.Sleep(100 * time.Millisecond)

	resourceProvider, err := getResourceProvider(t, commandCtx, options)
	if err != nil {
		return nil, err
	}

	resourceProvider.Start(commandCtx.Ctx, commandCtx.Cm)

	mediator, err := getMediator(t, commandCtx, options)
	if err != nil {
		return nil, err
	}

	mediator.Start(commandCtx.Ctx, commandCtx.Cm)

	jobCreatorOptions, err := getJobCreatorOptions(options)
	if err != nil {
		return nil, err
	}

	result, err := jobcreator.RunJob(commandCtx, jobCreatorOptions, func(evOffer data.JobOfferContainer) {

	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func TestNoModeration(t *testing.T) {
	commandCtx := system.NewTestingContext()
	defer commandCtx.Cleanup()

	message := "hello apples this is a message"

	executorOptions := noop.NewNoopExecutorOptions()
	executorOptions.Stdout = message

	result, err := testStackWithOptions(t, commandCtx, testOptions{
		moderationChance: 0,
		executor:         executorOptions,
	})

	assert.NoError(t, err, "there was an error running the job")
	assert.Equal(t, noop.NOOP_RESULTS_CID, result.Result.DataID, "the data ID was correct")

	localPath := solver.GetDownloadsFilePath(result.Result.DealID)

	fmt.Printf("result --------------------------------------\n")
	spew.Dump(localPath)
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal().Str("err", err.Error()).Msgf(".env not found")
	}

}
