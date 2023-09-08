package solver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/gorilla/mux"
)

type solverServer struct {
	options    server.ServerOptions
	controller *SolverController
	store      store.SolverStore
}

func NewSolverServer(
	options server.ServerOptions,
	controller *SolverController,
	store store.SolverStore,
) (*solverServer, error) {
	server := &solverServer{
		options:    options,
		controller: controller,
		store:      store,
	}
	return server, nil
}

func (solverServer *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.Use(server.CorsMiddleware)
	subrouter.Use(server.AuthMiddleware)

	subrouter.HandleFunc("/job_offers", server.Wrapper(solverServer.getJobOffers)).Methods("GET")
	subrouter.HandleFunc("/job_offers", server.Wrapper(solverServer.addJobOffer)).Methods("POST")

	subrouter.HandleFunc("/resource_offers", server.Wrapper(solverServer.getResourceOffers)).Methods("GET")
	subrouter.HandleFunc("/resource_offers", server.Wrapper(solverServer.addResourceOffer)).Methods("POST")

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", solverServer.options.Host, solverServer.options.Port),
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

func (solverServer *solverServer) getJobOffers(res http.ResponseWriter, req *http.Request) ([]string, error) {
	return []string{"job1", "job2"}, nil
}

func (solverServer *solverServer) getResourceOffers(res http.ResponseWriter, req *http.Request) ([]string, error) {
	return []string{"job1", "job2"}, nil
}

func (solverServer *solverServer) addJobOffer(res http.ResponseWriter, req *http.Request) (*data.JobOffer, error) {
	signerAddress := server.GetContextAddress(req.Context())
	jobOffer, err := server.ReadBody[data.JobOffer](req)
	if err != nil {
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != jobOffer.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}
	return solverServer.controller.addJobOffer(jobOffer)
}

func (solverServer *solverServer) addResourceOffer(res http.ResponseWriter, req *http.Request) (*data.ResourceOffer, error) {
	signerAddress := server.GetContextAddress(req.Context())
	resourceOffer, err := server.ReadBody[data.ResourceOffer](req)
	if err != nil {
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != resourceOffer.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	return solverServer.controller.addResourceOffer(resourceOffer)
}
