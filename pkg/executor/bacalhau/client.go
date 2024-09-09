package bacalhau

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bacalhau-project/bacalhau/pkg/publicapi/apimodels"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

const API_VERSION = "v1"

type BacalhauClient struct {
	http    *retryablehttp.Client
	options BacalhauExecutorOptions
}

func newClient(options BacalhauExecutorOptions) (*BacalhauClient, error) {
	http := retryablehttp.NewClient()
	return &BacalhauClient{
		http:    http,
		options: options,
	}, nil
}

func (client *BacalhauClient) alive() (bool, error) {
	result, err := getRequest[apimodels.IsAliveResponse](client, "agent/alive")
	if err != nil {
		return false, err
	}

	return result.IsReady(), nil
}

func (client *BacalhauClient) getVersion() (apimodels.GetVersionResponse, error) {
	return getRequest[apimodels.GetVersionResponse](client, "agent/version")
}

func getRequest[ResultType any](client *BacalhauClient, requestPath string) (ResultType, error) {
	var result ResultType
	url := fmt.Sprintf("%s/%s", client.getBaseUrl(), requestPath)
	request, err := retryablehttp.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.http.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (client *BacalhauClient) getBaseUrl() string {
	return fmt.Sprintf("http://%s:%s/api/%s", client.options.ApiHost, client.options.ApiPort, API_VERSION)
}
