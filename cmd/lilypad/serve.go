package lilypad

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/jobcreator"
	optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"

	// optionsfactory "github.com/lilypad-tech/lilypad/pkg/options"
	"github.com/lilypad-tech/lilypad/pkg/solver"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// var options jobcreator.JobCreatorOptions
var network string
var gcmd *cobra.Command

func newServeCmd() *cobra.Command {
	// options := optionsfactory.NewJobCreatorOptions()
	runCmd := &cobra.Command{
		Use:     "serve",
		Short:   "Run a job on the Lilypad network.",
		Long:    "Run a job on the Lilypad network.",
		Example: "run cowsay:v0.0.1 -i Message=moo",
		RunE: func(cmd *cobra.Command, args []string) error {
			gcmd = cmd
			//options := optionsfactory.NewJobCreatorOptions()
			// network, _ := cmd.Flags().GetString("network")
			// options, err := optionsfactory.ProcessJobCreatorOptions(options, args, network)
			// if err != nil {
			// 	return err
			// }
			// return runServeJob(cmd, options)
			// options = optionsfactory.NewJobCreatorOptions()
			network, _ = cmd.Flags().GetString("network")
			ctx := context.Background()
			if err := run(ctx, os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			return nil
		},
	}

	// optionsfactory.AddJobCreatorCliFlags(runCmd, &options)

	return runCmd
}
func run(ctx context.Context, w io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// logger := log.New(w, "", log.Ldate|log.Ltime)
	// config := config.New()
	// db := db.New(config, logger)
	// lw := logsWriter.New(config, logger)
	// geo := geo.New(logger)
	srv := New() //, db, lw, geo)

	httpServer := &http.Server{
		Addr:         net.JoinHostPort("127.0.0.1", "9876"),
		Handler:      srv,
		WriteTimeout: 120 * time.Second,
		ReadTimeout:  120 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "Error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()

	return nil
}
func runServeJob(cmd *cobra.Command, options jobcreator.JobCreatorOptions) (cid string, dataid string, err error) {

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
		return "", "", fmt.Errorf("failed to start spinner: %w", err)
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
		return "", "", fmt.Errorf("failed to stop spinner: %w", err)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	commandCtx := system.NewCommandContext(cmd)
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
		case "JobOfferCancelled":
			desc = "Job cancelled..."
			emoji = "😭"
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
		return "", "", err
	}
	spinner.Stop()
	fmt.Printf("\n🍂 Lilypad job completed, try 👇\n    open %s\n    cat %s/stdout\n    cat %s/stderr\n    https://ipfs.io/ipfs/%s\n",
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		result.Result.DataID,
	)
	return result.Result.DataID, result.JobOffer.DealID, err
}

// func createSpinner(message string, emoji string) (*yacspin.Spinner, error) {
// 	// build the configuration, each field is documented
// 	cfg := yacspin.Config{
// 		Frequency:         100 * time.Millisecond,
// 		CharSet:           yacspin.CharSets[69],
// 		Suffix:            " ", // puts a least one space between the animating spinner and the Message
// 		Message:           message,
// 		SuffixAutoColon:   true,
// 		ColorAll:          false,
// 		Colors:            []string{"fgMagenta"},
// 		StopCharacter:     emoji,
// 		StopColors:        []string{"fgGreen"},
// 		StopMessage:       message,
// 		StopFailCharacter: "✗",
// 		StopFailColors:    []string{"fgRed"},
// 		StopFailMessage:   "failed",
// 	}

// 	s, err := yacspin.New(cfg)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to make spinner from struct: %w", err)
// 	}

// 	stopOnSignal(s)
// 	return s, nil
// }

// func stopOnSignal(spinner *yacspin.Spinner) {
// 	// ensure we stop the spinner before exiting, otherwise cursor will remain
// 	// hidden and terminal will require a `reset`
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		<-sigCh

// 		spinner.StopFailMessage("interrupted")

// 		// ignoring error intentionally
// 		_ = spinner.StopFail()

//			os.Exit(0)
//		}()
//	}
func New(
// logger *log.Logger,
// db *gorm.DB,
// lw *logsWriter.Writer,
// geo *geo.Geo,
) http.Handler {
	router := mux.NewRouter()
	// middleware
	router.Use(cors)
	// router.Use(loggerMiddleware(logger))
	//
	router.HandleFunc("/ping", handlePing).Methods("GET")
	// routes
	handler := addRoutes(
		router,
	)

	return handler
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func loggerMiddleware(logger *log.Logger) func(n http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}
func addRoutes(
	router *mux.Router,
) http.Handler {
	// metricsDashboardRouter := router.PathPrefix("/metrics-dashboard").Subrouter()

	powLogsRouter := router.PathPrefix("/run").Subrouter()

	//todo need other api?
	powLogsRouter.Handle("/job", handlePostCardInfos()).Methods("POST")
	// powLogsRouter.Handle("/cardinfos/{address}/{signature}", handleGetCardInfoByAddress(logger, db)).Methods("GET")

	// // Temporary until self host RPC is deprecated
	// loggerRouter := router.PathPrefix("/log").Subrouter()
	// loggerRouter.Handle("/error", HandleErrorLog(logger, db)).Methods("POST")
	return router
}

type JobModule struct {
	gorm.Model
	//todo card info field
	Module string `json:"module"`
	Input  string `json:"input"`
}

func handlePostCardInfos() http.Handler {

	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
			var entry JobModule
			err := json.NewDecoder(r.Body).Decode(&entry)
			if err != nil {
				// logger.Printf("error parsing card info request body %s", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println(entry)
			options := optionsfactory.NewJobCreatorOptions()
			options.Offer.Module.Name = entry.Module
			inputArgs := strings.Split(entry.Module, " ")
			options, err = optionsfactory.ProcessJobCreatorOptions(options, inputArgs, network)
			if err != nil {
				fmt.Println(err)
				// return err
			}
			cid, dataid, err := runServeJob(gcmd, options)
			// // Extract the signature from the entry
			// signatureBytes, err := hex.DecodeString(entry.Signature[2:]) // Remove "0x" prefix
			// if err != nil {
			// 	logger.Printf("error decoding signature %s", err)
			// 	http.Error(w, "invalid signature format", http.StatusBadRequest)
			// 	return
			// }

			// // Compute the hash of the message
			// message := []byte(entry.Address + entry.Timestamp + entry.PowID + entry.Uuid + entry.GpuName + entry.Executable + entry.NvidiaSmi + entry.Challenge + entry.Difficulty) // Adjust the message as needed
			// hash := crypto.Keccak256Hash(message)

			// // Verify the signature
			// sigPublicKey, err := crypto.SigToPub(hash.Bytes(), signatureBytes)
			// if err != nil {
			// 	logger.Printf("error verifying signature %s", err)
			// 	http.Error(w, "invalid signature", http.StatusBadRequest)
			// 	return
			// }

			// // Derive the address from the recovered public key
			// recoveredAddress := crypto.PubkeyToAddress(*sigPublicKey).Hex()

			// // Compare the derived address with the provided wallet address
			// if recoveredAddress != entry.Address {
			// 	logger.Printf("signature does not match wallet address")
			// 	http.Error(w, "invalid signature", http.StatusBadRequest)
			// 	return
			// }

			// if err = db.Save(&entry).Error; err != nil {
			// 	logger.Printf("error saving card info %s", err)
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }

			// w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf("{\"cid\":\"%s\",\"dataid\":\"%s\"}", cid, dataid))) //`{"result:ipfs://" + x))
			// json.NewEncoder(w).Encode([]byte(fmt.Sprintf("{\"result\":\"%s\"}", x)))
		},
	)
}
