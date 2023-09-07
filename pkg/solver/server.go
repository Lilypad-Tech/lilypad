package solver

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

type solverServer struct {
	options server.ServerOptions
}

func NewSolverServer(
	options server.ServerOptions,
) (*solverServer, error) {
	server := &solverServer{
		options: options,
	}
	return server, nil
}

func (server *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {

	return nil
}
