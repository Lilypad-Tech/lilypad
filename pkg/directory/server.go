package directory

import (
	"context"
	"encoding/json"
	"fmt"
	corehttp "net/http"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/directory/store"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type directoryServer struct {
	options    http.ServerOptions
	controller *DirectoryController
	store      store.DirectoryStore
}

func NewSolverServer(
	options http.ServerOptions,
	controller *DirectoryController,
	store store.DirectoryStore,
) (*directoryServer, error) {
	server := &directoryServer{
		options:    options,
		controller: controller,
		store:      store,
	}
	return server, nil
}

func (directoryServer *directoryServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(http.API_SUB_PATH).Subrouter()

	subrouter.Use(http.CorsMiddleware)

	subrouter.HandleFunc("/deals", http.Wrapper(directoryServer.getDeals)).Methods("GET")
	subrouter.HandleFunc("/deals", http.Wrapper(directoryServer.addDeal)).Methods("POST")

	writeEventChannel := make(chan []byte)

	// listen to events coming out of the controller and write them to all
	// connected websocket connections
	go func() {
		select {
		case ev := <-directoryServer.controller.getEventChannel():
			// we write this event to all connected web socket connections
			fmt.Printf("got controller event: %+v\n", ev)
			evBytes, err := json.Marshal(ev)
			if err != nil {
				log.Error().Msgf("Error marshalling event: %s", err.Error())
			}
			writeEventChannel <- evBytes
			return
		case <-ctx.Done():
			return
		}
	}()

	http.StartWebSocketServer(
		subrouter,
		"/ws",
		writeEventChannel,
		ctx,
	)

	srv := &corehttp.Server{
		Addr:              fmt.Sprintf("%s:%d", directoryServer.options.Host, directoryServer.options.Port),
		WriteTimeout:      time.Minute * 15,
		ReadTimeout:       time.Minute * 15,
		ReadHeaderTimeout: time.Minute * 15,
		IdleTimeout:       time.Minute * 60,
		Handler:           router,
	}

	// Create a channel to receive errors from ListenAndServe
	serverErrors := make(chan error, 1)

	// Run ListenAndServe in a goroutine because it blocks
	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return err
	case <-ctx.Done():
		// Create a context with a timeout for the server to close
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt to gracefully shut down the server
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to stop server: %w", err)
		}
	}

	return nil
}

func (directoryServer *directoryServer) getDeals(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.Deal, error) {
	query := store.GetDealsQuery{}
	// if there is a job_creator query param then assign it
	if jobCreator := req.URL.Query().Get("job_creator"); jobCreator != "" {
		query.JobCreator = jobCreator
	}
	if resourceProvider := req.URL.Query().Get("resource_provider"); resourceProvider != "" {
		query.ResourceProvider = resourceProvider
	}
	return directoryServer.store.GetDeals(query)
}

func (directoryServer *directoryServer) addDeal(res corehttp.ResponseWriter, req *corehttp.Request) (*data.Deal, error) {
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		return nil, err
	}
	deal, err := http.ReadBody[data.Deal](req)
	if err != nil {
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != deal.JobCreator && signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("job creator or resource provider address does not match signer address")
	}
	return directoryServer.controller.addDeal(deal)
}
