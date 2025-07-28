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

type ResourceProviderListItem struct {
	ID               int    `json:"ID"`
	ResourceProvider string `json:"resource_provider"`
}

type AdminServiceClient interface {
	GetAllowList() ([]ResourceProviderListItem, error)
	GetTestList() ([]ResourceProviderListItem, error)
}

type AdminServiceClientOptions struct {
	BaseURL            string
	ApiKey             string
	EnableAdminService bool
}

type adminServiceClient struct {
	clientOptions AdminServiceClientOptions
	client        http.Client
}

func NewAdminServiceClient(adminServiceOptions AdminServiceClientOptions) (AdminServiceClient, error) {
	if len(adminServiceOptions.ApiKey) == 0 {
		err := errors.New("admin Service API key not set")
		log.Error().Err(err).Msg("Admin Service API Key is required")
		return nil, err
	}

	if len(adminServiceOptions.BaseURL) == 0 {
		err := errors.New("admin Service base URL not set")
		log.Error().Err(err).Msg("Admin Service Base URL is required")
		return nil, err
	}

	return &adminServiceClient{
		clientOptions: adminServiceOptions,
		client:        http.Client{Timeout: time.Duration(30) * time.Second}}, nil
}

func (a *adminServiceClient) GetTestList() ([]ResourceProviderListItem, error) {
	getAllowListUrl := a.clientOptions.BaseURL + "/api/v1/test-list"

	req, err := http.NewRequest("GET", getAllowListUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to build admin service get test list request")
		return []ResourceProviderListItem{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed admin service get test list request")
		return []ResourceProviderListItem{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading get test list response")
		return []ResourceProviderListItem{}, err
	}

	var response []ResourceProviderListItem
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("Failed unmarshelling get test list response")
		return []ResourceProviderListItem{}, err
	}

	return response, nil
}

func (a *adminServiceClient) GetAllowList() ([]ResourceProviderListItem, error) {
	getAllowListUrl := a.clientOptions.BaseURL + "/api/v1/allow-list"

	req, err := http.NewRequest("GET", getAllowListUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to build admin service get allow list request")
		return []ResourceProviderListItem{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed admin service get allow list request")
		return []ResourceProviderListItem{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading get allow list response")
		return []ResourceProviderListItem{}, err
	}

	var response []ResourceProviderListItem
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("Failed unmarshelling get allow list response")
		return []ResourceProviderListItem{}, err
	}

	return response, nil
}
