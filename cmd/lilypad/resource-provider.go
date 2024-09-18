package lilypad

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/executor/bacalhau"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/spf13/cobra"
)

func newResourceProviderCmd() *cobra.Command {
	options := optionsfactory.NewResourceProviderOptions()

	resourceProviderCmd := &cobra.Command{
		Use:     "resource-provider",
		Short:   "Start the lilypad resource-provider service.",
		Long:    "Start the lilypad resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {

			network, _ := cmd.Flags().GetString("network")
			options, err := optionsfactory.ProcessResourceProviderOptions(options, network)
			if err != nil {
				return err
			}
			return runResourceProvider(cmd, options, network)
		},
	}

	optionsfactory.AddResourceProviderCliFlags(resourceProviderCmd, &options)

	return resourceProviderCmd
}

func runResourceProvider(cmd *cobra.Command, options resourceprovider.ResourceProviderOptions, network string) error {
	fmt.Println("Starting Bacalhau")
	go resourceprovider.StartIpfs()
	wait("http://127.0.0.1:5001/webui")
	//time.Sleep(10 * time.Second)

	fmt.Println("Starting Bacalhau")
	go resourceprovider.StartBacalhau()
	wait("http://localhost:1234/api/v1/agent/alive")
	// url := ""
	// time.Sleep(10 * time.Second)

	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	executor, err := bacalhau.NewBacalhauExecutor(options.Bacalhau)
	if err != nil {
		return err
	}

	// Ensure that our executor is available
	status, err := executor.IsAvailable()
	if !status || err != nil {
		return err
	}

	tc := system.TelemetryConfig{
		TelemetryURL:   options.Telemetry.URL,
		TelemetryToken: options.Telemetry.Token,
		Enabled:        !options.Telemetry.Disable,
		Service:        system.ResourceProviderService,
		Network:        network,
		Address:        web3SDK.GetAddress().String(),
		GPU:            system.GetGPUInfo(),
	}
	telemetry, err := system.SetupOTelSDK(commandCtx.Ctx, tc)
	if err != nil {
		fmt.Printf("failed to setup opentelemetry: %s", err)
	}
	commandCtx.Cm.RegisterCallbackWithContext(telemetry.Shutdown)

	resourceProviderService, err := resourceprovider.NewResourceProvider(options, web3SDK, executor, telemetry)
	if err != nil {
		return err
	}

	resourecProviderErrors := resourceProviderService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-resourecProviderErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}

func wait(url string) {
	for {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()

			break
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(1 * time.Second)
	}
}
