package server

import (
	"net/http"

	"github.com/bacalhau-project/lilypad/pkg/types"
)

type statusResponse struct {
	User *types.User `json:"user"`
}

func (apiServer *lilypadAPIServer) status(res http.ResponseWriter, req *http.Request) (statusResponse, error) {
	user := getRequestUser(req.Context())
	return statusResponse{
		User: user,
	}, nil
}
