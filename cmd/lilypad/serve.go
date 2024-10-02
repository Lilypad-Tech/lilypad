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
	return runCmd
}
func run(ctx context.Context, w io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	srv := New()

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

func New() http.Handler {
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
	router.PathPrefix("/files").Handler(http.StripPrefix("/files", http.FileServer(http.Dir("/tmp/lilypad/data/downloaded-files/"))))

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
	powLogsRouter := router.PathPrefix("/run").Subrouter()
	powLogsRouter.Handle("/job", handleJob()).Methods("POST")
	return router
}

type JobModule struct {
	gorm.Model
	Module string                 `json:"module"`
	Input  map[string]interface{} `json:"input"`
}

func handleJob() http.Handler {

	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
			var entry JobModule
			err := json.NewDecoder(r.Body).Decode(&entry)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println(entry)
			options := optionsfactory.NewJobCreatorOptions()
			options.Offer.Module.Name = entry.Module
			options.Offer.Inputs = map[string]string{entry.Input["name"].(string): entry.Input["value"].(string)}
			inputArgs := []string{entry.Module} //strings.Split(entry.Module, " ") //[]string{entry.Input} //
			options, err = optionsfactory.ProcessJobCreatorOptions(options, inputArgs, network)
			if err != nil {
				fmt.Println(err)
				// return err
			}
			cid, dataid, err := runServeJob(gcmd, options)
			if err != nil {
				fmt.Println(err)

			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf("{\"cid\":\"%s\",\"dataid\":\"%s\"}", cid, dataid)))
		},
	)
}
