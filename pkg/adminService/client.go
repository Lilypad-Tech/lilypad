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

type ResourceProviderInAllowListResponse struct {
	InAllowList bool `json:"inAllowList"`
}

type ResourceProviderInTestListResponse struct {
	InTestList bool `json:"inTestList"`
}

type AdminServiceClient interface {
	GetAllowList() ([]ResourceProviderListItem, error)
	GetTestList() ([]ResourceProviderListItem, error)
	IsRPonAllowList(resourceProvider string) (ResourceProviderInAllowListResponse, error)
	IsRPonTestList(resourceProvider string) (ResourceProviderInTestListResponse, error)
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

func (a *adminServiceClient) IsRPonAllowList(resourceProvider string) (ResourceProviderInAllowListResponse, error) {
	isRpOnAllowListUrl := a.clientOptions.BaseURL + "/api/v1/allow-list/contains/" + resourceProvider

	req, err := http.NewRequest("GET", isRpOnAllowListUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to build admin service rp contained on allow list request")
		return ResourceProviderInAllowListResponse{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed admin service call to check if RP is on allow list request")
		return ResourceProviderInAllowListResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading get rp contained on allow list response")
		return ResourceProviderInAllowListResponse{}, err
	}

	var response ResourceProviderInListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("Failed unmarshelling get rp contained on allow list response")
		return ResourceProviderInAllowListResponse{}, err
	}

	return response, nil
}

func (a *adminServiceClient) IsRPonTestList(resourceProvider string) (ResourceProviderInTestListResponse, error) {
	isRpOnAllowListUrl := a.clientOptions.BaseURL + "/api/v1/test-list/contains/" + resourceProvider

	req, err := http.NewRequest("GET", isRpOnAllowListUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to build admin service rp contained on test list request")
		return ResourceProviderInTestListResponse{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed admin service call to check if RP is on test list request")
		return ResourceProviderInTestListResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading get rp contained on test list response")
		return ResourceProviderInTestListResponse{}, err
	}

	var response ResourceProviderInTestListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("Failed unmarshelling get rp contained on test list response")
		return ResourceProviderInTestListResponse{}, err
	}

	return response, nil
}
