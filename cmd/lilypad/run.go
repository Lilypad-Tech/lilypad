package lilypad

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/jobcreator"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/solver"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/spf13/cobra"

	"github.com/theckman/yacspin"
)

func newRunCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()
	runCmd := &cobra.Command{
		Use:     "run",
		Short:   "Run a job on the Lilypad network.",
		Long:    "Run a job on the Lilypad network.",
		Example: "run cowsay:v0.0.1 -i Message=moo",
		RunE: func(cmd *cobra.Command, args []string) error {
			optionsfactory.CheckDeprecation(options.Offer.Services, options.Web3)

			network, _ := cmd.Flags().GetString("network")
			options, err := optionsfactory.ProcessJobCreatorOptions(options, args, network)
			if err != nil {
				return err
			}
			return runJob(cmd, options)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(runCmd, &options)

	return runCmd
}

func runJob(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	c := color.New(color.FgCyan).Add(color.Bold)
	header := `
⠀⠀⠀⠀⠀⠀⣀⣤⣤⢠⣤⣀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢴⣿⣿⣿⣿⢸⣿⡟⠀⠀⠀⠀⠀    ██╗     ██╗██╗  ██╗   ██╗██████╗  █████╗ ██████╗
⠀⠀⣰⣿⣦⡙⢿⣿⣿⢸⡿⠀⠀⠀⠀⢀⠀    ██║     ██║██║  ╚██╗ ██╔╝██╔══██╗██╔══██╗██╔══██╗
⠀⢰⣿⣿⣿⣿⣦⡙⣿⢸⠁⢀⣠⣴⣾⣿⡆    ██║     ██║██║   ╚████╔╝ ██████╔╝███████║██║  ██║
⠀⣛⣛⣛⣛⣛⣛⣛⠈⠀⣚⣛⣛⣛⣛⣛⣛    ██║     ██║██║    ╚██╔╝  ██╔═══╝ ██╔══██║██║  ██║
⠀⢹⣿⣿⣿⣿⠟⣡⣿⢸⣮⡻⣿⣿⣿⣿⡏    ███████╗██║███████╗██║   ██║     ██║  ██║██████╔╝
⠀⠀⢻⣿⡟⣩⣾⣿⣿⢸⣿⣿⣌⠻⣿⡟⠀    ╚══════╝╚═╝╚══════╝╚═╝   ╚═╝     ╚═╝  ╚═╝╚═════╝ v2
⠀⠀⠀⠉⢾⣿⣿⣿⣿⢸⣿⣿⣿⡷⠈⠀⠀
⠀⠀⠀⠀⠀⠈⠙⠛⠛⠘⠛⠋⠁⠀ ⠀⠀⠀   Decentralized Compute Network  https://lilypad.tech

`
	if system.Version != "" {
		header = strings.Replace(header, "v2", system.Version, 1)
	}
	c.Print(header)

	spinner, err := createSpinner("Lilypad submitting job", "🌟")
	if err != nil {
		fmt.Printf("failed to make spinner from config struct: %v\n", err)
		os.Exit(1)
	}

	// start the spinner animation
	if err := spinner.Start(); err != nil {
		return fmt.Errorf("failed to start spinner: %w", err)
	}

	// update message
	// spinner.Message("uploading files")

	// let spinner render some more
	// time.Sleep(1 * time.Second)

	// if you wanted to print a failure message...
	//
	// if err := spinner.StopFail(); err != nil {
	// 	return fmt.Errorf("failed to stop spinner: %w", err)
	// }

	if err := spinner.Stop(); err != nil {
		return fmt.Errorf("failed to stop spinner: %w", err)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tc := system.TelemetryConfig{Service: system.JobCreatorService, CollectorURL: options.Offer.Services.TelemetryURL, Enabled: false}
	commandCtx := system.NewCommandContext(cmd, tc)
	defer commandCtx.Cleanup()
	result, err := jobcreator.RunJob(commandCtx, options, func(evOffer data.JobOfferContainer) {
		spinner.Stop()
		st := data.GetAgreementStateString(evOffer.State)
		var desc string
		var emoji string
		switch st {
		case "DealNegotiating":
			desc = "Job submitted. Negotiating deal..."
			emoji = "🤝"
		case "DealAgreed":
			desc = "Deal agreed. Running job..."
			emoji = "💌"
		case "ResultsSubmitted":
			desc = "Results submitted. Awaiting verification..."
			emoji = "🤔"
		case "ResultsAccepted":
			desc = "Results accepted. Downloading result..."
			emoji = "✅"
		case "ResultsRejected":
			desc = "Results rejected! Getting refund..."
			emoji = "🙀"
		default:
			desc = st
			emoji = "🌟"
		}
		spinner, err = createSpinner(desc, emoji)
		if err != nil {
			fmt.Printf("failed to make spinner from config struct: %v\n", err)
			os.Exit(1)
		}

		// start the spinner animation
		if err := spinner.Start(); err != nil {
			fmt.Printf("failed to start spinner: %s", err)
			os.Exit(1)
		}

		// UPDATE FUNCTION
		// fmt.Printf("evOffer: %s --------------------------------------\n")
		// spew.Dump(evOffer)

	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	spinner.Stop()
	fmt.Printf("\n🍂 Lilypad job completed, try 👇\n    open %s\n    cat %s/stdout\n    cat %s/stderr\n    https://ipfs.io/ipfs/%s\n",
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		result.Result.DataID,
	)
	return err
}

func createSpinner(message string, emoji string) (*yacspin.Spinner, error) {
	// build the configuration, each field is documented
	cfg := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[69],
		Suffix:            " ", // puts a least one space between the animating spinner and the Message
		Message:           message,
		SuffixAutoColon:   true,
		ColorAll:          false,
		Colors:            []string{"fgMagenta"},
		StopCharacter:     emoji,
		StopColors:        []string{"fgGreen"},
		StopMessage:       message,
		StopFailCharacter: "✗",
		StopFailColors:    []string{"fgRed"},
		StopFailMessage:   "failed",
	}

	s, err := yacspin.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to make spinner from struct: %w", err)
	}

	stopOnSignal(s)
	return s, nil
}

func stopOnSignal(spinner *yacspin.Spinner) {
	// ensure we stop the spinner before exiting, otherwise cursor will remain
	// hidden and terminal will require a `reset`
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh

		spinner.StopFailMessage("interrupted")

		// ignoring error intentionally
		_ = spinner.StopFail()

		os.Exit(0)
	}()
}
