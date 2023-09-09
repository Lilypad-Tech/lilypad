package directory

import (
	"context"

	"github.com/bacalhau-project/lilypad/pkg/directory/store"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
)

type DirectoryOptions struct {
	Web3   web3.Web3Options
	Server http.ServerOptions
}

type Directory struct {
	web3SDK    *web3.ContractSDK
	server     *directoryServer
	controller *DirectoryController
	store      store.DirectoryStore
}

func NewDirectory(
	options DirectoryOptions,
	store store.DirectoryStore,
	web3SDK *web3.ContractSDK,
) (*Directory, error) {
	controller, err := NewDirectoryController(web3SDK, store)
	if err != nil {
		return nil, err
	}
	server, err := NewSolverServer(options.Server, controller, store)
	if err != nil {
		return nil, err
	}
	solver := &Directory{
		controller: controller,
		store:      store,
		server:     server,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (directory *Directory) Start(ctx context.Context, cm *system.CleanupManager) error {
	err := directory.controller.Start(ctx, cm)
	if err != nil {
		return err
	}

	return directory.server.ListenAndServe(ctx, cm)
}

func (directory *Directory) GetEventChannel() DirectoryEventChannel {
	return directory.controller.getEventChannel()
}
