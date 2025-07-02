package anura

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

type ModuleDiscoveryResponse struct {
	Data    ResponseData `json:"data"`
	Message string       `json:"message"`
	Status  int          `json:"status"`
}

type ResponseData struct {
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
	if len(anuraOptions.ApiKey) == 0 {
		err := errors.New("Anura API key not set")
		log.Error().Err(err).Msg("Anura API Key is required. You can learn how to obtain one by reading the Anura docs: https://docs.lilypad.tech/lilypad/developer-resources/inference-api")
		return nil, err
	}

	if len(anuraOptions.BaseURL) == 0 {
		err := errors.New("Anura base URL not set")
		log.Error().Err(err).Msg("Anura Base URL is required. You can learn what the Base URL is by reading the Anura docs: https://docs.lilypad.tech/lilypad/developer-resources/inference-api")
		return nil, err
	}

	return &anuraClient{
		clientOptions: anuraOptions,
		client:        http.Client{Timeout: time.Duration(30) * time.Second}}, nil
}

func (a *anuraClient) GetAnuraModules() (ModuleDiscoveryResponse, error) {
	getModulesUrl := a.clientOptions.BaseURL + "/modules"
	fmt.Println(getModulesUrl)
	req, err := http.NewRequest("GET", getModulesUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to build anura get modules request")
		return ModuleDiscoveryResponse{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.clientOptions.ApiKey))
	resp, err := a.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed anura get modules request")
		return ModuleDiscoveryResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading get modules response")
		return ModuleDiscoveryResponse{}, err
	}

	var response ModuleDiscoveryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("Failed unmarshelling get modules response")
		return ModuleDiscoveryResponse{}, err
	}

	return response, nil
}
