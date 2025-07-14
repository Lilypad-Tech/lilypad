package adminService

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

type AllowListResponse struct {
	Data AllowListResponseData `json:"data"`
}

type TestListResponse struct {
	Data TestListResponseData `json:"data"`
}

type AllowListResponseData struct {
	AllowList []ListItem `json:"allowlist"`
}

type TestListResponseData struct {
	TestList []ListItem `json:"testlist"`
}

type ListItem struct {
	ID               string `json:"id"`
	ResourceProvider int64  `json:"resource_provider"`
}

type AdminServiceClient interface {
	GetAllowList() (AllowListResponse, error)
	GetTestList() (TestListResponse, error)
}

type AdminServiceClientOptions struct {
	BaseURL string
	ApiKey  string
}

type adminServiceClient struct {
	clientOptions AdminServiceClientOptions
	client        http.Client
}

func NewAdminServiceClient(adminServiceOptions AdminServiceClientOptions) (AdminServiceClient, error) {
	if len(adminServiceOptions.ApiKey) == 0 {
		err := errors.New("Admin Service API key not set")
		log.Error().Err(err).Msg("Admin Service API Key is required")
		return nil, err
	}

	if len(adminServiceOptions.BaseURL) == 0 {
		err := errors.New("Admin Service base URL not set")
		log.Error().Err(err).Msg("Admin Service Base URL is required")
		return nil, err
	}

	return &adminServiceClient{
		clientOptions: adminServiceOptions,
		client:        http.Client{Timeout: time.Duration(30) * time.Second}}, nil
}

func (a *adminServiceClient) GetTestList() (TestListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adminServiceClient) GetAllowList() (AllowListResponse, error) {
	getAllowListUrl := a.clientOptions.BaseURL + "/allow-list"

	req, err := http.NewRequest("GET", getAllowListUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to build anura get modules request")
		return AllowListResponse{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed anura get modules request")
		return AllowListResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading get modules response")
		return AllowListResponse{}, err
	}

	var response AllowListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("Failed unmarshelling get modules response")
		return AllowListResponse{}, err
	}

	return response, nil
}
