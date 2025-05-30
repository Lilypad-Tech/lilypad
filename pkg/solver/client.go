package solver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/http"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/rs/zerolog/log"
)

type SolverClient struct {
	options         http.ClientOptions
	solverEventSubs []func(SolverEvent)
}

func NewSolverClient(
	options http.ClientOptions,
) (*SolverClient, error) {
	client := &SolverClient{
		options:         options,
		solverEventSubs: []func(SolverEvent){},
	}
	return client, nil
}

// connect the websocket to the solver server
func (client *SolverClient) Start(ctx context.Context, cm *system.CleanupManager) error {

	websocketURL := fmt.Sprintf("%s%s%s%s%s", http.WEBSOCKET_SUB_PATH, "?&Type=", client.options.Type, "&ID=", client.options.PublicAddress)
	websocketEventChannel := http.ConnectWebSocket(http.WebsocketURL(client.options, websocketURL), ctx)
	go func() {
		for {
			select {
			case evBytes := <-websocketEventChannel:
				// parse the ev into a new SolverEvent
				var ev SolverEvent
				if err := json.Unmarshal(evBytes, &ev); err != nil {
					log.Error().Msgf("Error unmarshalling event: %s", err.Error())
					continue
				}
				// loop over each event channel and write the event to it
				for _, handler := range client.solverEventSubs {
					go handler(ev)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (client *SolverClient) SubscribeEvents(handler func(SolverEvent)) {
	client.solverEventSubs = append(client.solverEventSubs, handler)
}

func (client *SolverClient) GetJobOffers(query store.GetJobOffersQuery) ([]data.JobOfferContainer, error) {
	queryParams := map[string]string{}
	if query.JobCreator != "" {
		queryParams["job_creator"] = query.JobCreator
	}
	if query.NotMatched {
		queryParams["not_matched"] = "true"
	}
	return http.GetRequest[[]data.JobOfferContainer](client.options, "/job_offers", queryParams)
}

func (client *SolverClient) GetResourceOffers(query store.GetResourceOffersQuery) ([]data.ResourceOfferContainer, error) {
	queryParams := map[string]string{}
	if query.ResourceProvider != "" {
		queryParams["resource_provider"] = query.ResourceProvider
	}
	if query.Active {
		queryParams["active"] = "true"
	}
	if query.NotMatched {
		queryParams["not_matched"] = "true"
	}
	return http.GetRequest[[]data.ResourceOfferContainer](client.options, "/resource_offers", queryParams)
}

func (client *SolverClient) GetDeals(query store.GetDealsQuery) ([]data.DealContainer, error) {
	queryParams := map[string]string{}
	if query.JobCreator != "" {
		queryParams["job_creator"] = query.JobCreator
	}
	if query.ResourceProvider != "" {
		queryParams["resource_provider"] = query.ResourceProvider
	}
	if query.State != "" {
		queryParams["state"] = query.State
	}
	return http.GetRequest[[]data.DealContainer](client.options, "/deals", queryParams)
}

func (client *SolverClient) GetDeal(id string) (data.DealContainer, error) {
	return http.GetRequest[data.DealContainer](client.options, fmt.Sprintf("/deals/%s", id), map[string]string{})
}

func (client *SolverClient) GetResult(id string) (data.Result, error) {
	return http.GetRequest[data.Result](client.options, fmt.Sprintf("/deals/%s/result", id), map[string]string{})
}

func (client *SolverClient) GetDealsWithFilter(query store.GetDealsQuery, filter func(data.DealContainer) bool) ([]data.DealContainer, error) {
	deals, err := client.GetDeals(query)
	if err != nil {
		return nil, err
	}
	ret := []data.DealContainer{}
	for _, deal := range deals {
		if filter(deal) {
			ret = append(ret, deal)
		}
	}
	return ret, nil
}

func (client *SolverClient) AddJobOffer(jobOffer data.JobOffer) (data.JobOfferContainer, error) {
	return http.PostRequest[data.JobOffer, data.JobOfferContainer](client.options, "/job_offers", jobOffer)
}

func (client *SolverClient) AddJobOfferWithFiles(jobOffer data.JobOffer, inputsPath string) (data.JobOfferContainer, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add job offer part
	jobOfferPart, err := writer.CreateFormFile("job_offer.json", "job_offer.json")
	if err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error creating job offer part: %w", err)
	}
	jobOfferJson, err := json.Marshal(jobOffer)
	if err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error marshaling job offer: %w", err)
	}
	if _, err := jobOfferPart.Write([]byte(jobOfferJson)); err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error writing job offer part: %w", err)
	}

	// Add input files part
	inputBuf, err := system.GetTarBuffer(inputsPath)
	if err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error creating inputs file tar buffer: %w", err)
	}
	inputPart, err := writer.CreateFormFile("inputs.tar", "inputs.tar")
	if err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error creating inputs file part: %w", err)
	}
	if _, err := io.Copy(inputPart, inputBuf); err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error writing inputs file part: %w", err)
	}
	if err := writer.Close(); err != nil {
		return data.JobOfferContainer{}, fmt.Errorf("error closing multipart writer: %w", err)
	}

	// Add content type header
	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}

	return http.PostRequestBufferWithHeaders[data.JobOfferContainer](
		client.options,
		"/job_offers/with_files",
		headers,
		&requestBody,
	)
}

func (client *SolverClient) AddResourceOffer(resourceOffer data.ResourceOffer) (data.ResourceOfferContainer, error) {
	return http.PostRequest[data.ResourceOffer, data.ResourceOfferContainer](client.options, "/resource_offers", resourceOffer)
}

func (client *SolverClient) AddResult(result data.Result) (data.Result, error) {
	return http.PostRequest[data.Result, data.Result](client.options, fmt.Sprintf("/deals/%s/result", result.DealID), result)
}

func (client *SolverClient) UpdateTransactionsResourceProvider(id string, payload data.DealTransactionsResourceProvider) (data.DealContainer, error) {
	return http.PostRequest[data.DealTransactionsResourceProvider, data.DealContainer](client.options, fmt.Sprintf("/deals/%s/txs/resource_provider", id), payload)
}

func (client *SolverClient) UpdateTransactionsJobCreator(id string, payload data.DealTransactionsJobCreator) (data.DealContainer, error) {
	return http.PostRequest[data.DealTransactionsJobCreator, data.DealContainer](client.options, fmt.Sprintf("/deals/%s/txs/job_creator", id), payload)
}

func (client *SolverClient) UpdateTransactionsMediator(id string, payload data.DealTransactionsMediator) (data.DealContainer, error) {
	return http.PostRequest[data.DealTransactionsMediator, data.DealContainer](client.options, fmt.Sprintf("/deals/%s/txs/mediator", id), payload)
}

func (client *SolverClient) UploadResultFiles(id string, localPath string) (data.Result, error) {
	buf, err := system.GetTarBuffer(localPath)
	if err != nil {
		return data.Result{}, err
	}
	return http.PostRequestBuffer[data.Result](client.options, fmt.Sprintf("/deals/%s/files", id), buf)
}

func (client *SolverClient) DownloadInputFiles(id string, localPath string) error {
	buf, err := http.GetRequestBuffer(client.options, fmt.Sprintf("/deals/%s/input_files", id), map[string]string{})
	if err != nil {
		return err
	}
	return system.ExpandTarBuffer(buf, localPath)
}

func (client *SolverClient) DownloadResultFiles(id string, localPath string) error {
	buf, err := http.GetRequestBuffer(client.options, fmt.Sprintf("/deals/%s/files", id), map[string]string{})
	if err != nil {
		return err
	}
	return system.ExpandTarBuffer(buf, localPath)
}

// Validation service

func (client *SolverClient) GetValidationToken() (http.ValidationToken, error) {
	return http.GetRequest[http.ValidationToken](client.options, fmt.Sprintf("/validation_token"), map[string]string{})
}
