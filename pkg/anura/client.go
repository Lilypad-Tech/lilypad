package anura

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ModuleDiscoveryResponse struct {
	Modules []Module `json:"modules"`
}

type Module struct {
	Name             string `json:"name"`
	ImageSizeInBytes int64  `json:"image_size_in_bytes"`
	Repo             string `json:"repo"`
	CommitHash       string `json:"commit_hash"`
	JobSpecLocation  string `json:"job_spec_location"`
	Tier             string `json:"tier"`
}

type AnuraClient interface {
	GetAnuraModules() (ModuleDiscoveryResponse, error)
}

type AnuraClientOptions struct {
	BaseURL string
	ApiKey  string
}

type anuraClient struct {
	clientOptions AnuraClientOptions
	client        http.Client
}

func NewAnuraClient(anuraOptions AnuraClientOptions) (AnuraClient, error) {
	if len(anuraOptions.ApiKey) == 0 || len(anuraOptions.BaseURL) == 0 {
		return nil, errors.New("anura api key and baseUrl are required")
	}
	return &anuraClient{
		clientOptions: anuraOptions,
		client:        http.Client{Timeout: time.Duration(30) * time.Second}}, nil
}

func (a *anuraClient) GetAnuraModules() (ModuleDiscoveryResponse, error) {
	getModulesUrl := a.clientOptions.BaseURL + "/modules"
	req, err := http.NewRequest("GET", getModulesUrl, nil)
	if err != nil {
		return ModuleDiscoveryResponse{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		return ModuleDiscoveryResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ModuleDiscoveryResponse{}, err
	}

	var response ModuleDiscoveryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ModuleDiscoveryResponse{}, err
	}

	return response, nil
}
