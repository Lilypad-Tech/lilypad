package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/controller"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/gorilla/mux"
)

type lilypadAPIServer struct {
	Options    ServerOptions
	Controller *controller.Controller
}

func NewServer(
	options ServerOptions,
	controller *controller.Controller,
) (*lilypadAPIServer, error) {
	if options.URL == "" {
		return nil, fmt.Errorf("server url is required")
	}
	if options.Host == "" {
		return nil, fmt.Errorf("server host is required")
	}
	if options.Port == 0 {
		return nil, fmt.Errorf("server port is required")
	}

	return &lilypadAPIServer{
		Options:    options,
		Controller: controller,
	}, nil
}

func (apiServer *lilypadAPIServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.Use(authMiddleware)
	subrouter.Use(corsMiddleware)

	subrouter.HandleFunc("/status", wrapper(apiServer.status)).Methods("GET")

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", apiServer.Options.Host, apiServer.Options.Port),
		WriteTimeout:      time.Minute * 15,
		ReadTimeout:       time.Minute * 15,
		ReadHeaderTimeout: time.Minute * 15,
		IdleTimeout:       time.Minute * 60,
		Handler:           router,
	}
	return srv.ListenAndServe()
}
