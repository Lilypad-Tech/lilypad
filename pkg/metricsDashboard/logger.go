package metricsDashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
)

const jobsEndpoint = "jobs"
const nodeInfoEndpoint = "nodes"
const nodeConnectionEndpoint = "uptimes"
const dealsEndpoint = "deals"

var host = os.Getenv("API_HOST") + "metrics-dashboard/"

func TrackEvent(url string, json string) {
	data := []byte(json)

	client := &http.Client{Timeout: time.Second * 1}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
}

func TrackJobOfferUpdate(evOffer data.JobOfferContainer) {
	var url = host + jobsEndpoint
	var module = evOffer.JobOffer.Module.Name
	if module == "" {
		module = evOffer.JobOffer.Module.Repo + ":" + evOffer.JobOffer.Module.Hash
	}

	data := map[string]interface{}{
		"ID":                evOffer.ID,
		"JobOfferCreatedAt": evOffer.JobOffer.CreatedAt,
		"JobCreator":        evOffer.JobCreator,
		"DealID":            evOffer.DealID,
		"State":             evOffer.State,
		"Module":            module,
		"Event":             "JobOfferUpdate",
		"EventUpdatedAt":    time.Now().UnixMilli(),
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	TrackEvent(url, payload)
}

func TrackNodeInfo(resourceOffer data.ResourceOffer) {
	var url = host + nodeInfoEndpoint

	data := map[string]interface{}{
		"ID":      resourceOffer.ResourceProvider,
		"GPU":     resourceOffer.Spec.GPU,
		"CPU":     resourceOffer.Spec.CPU,
		"RAM":     resourceOffer.Spec.RAM,
		"Modules": resourceOffer.Modules,
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	TrackEvent(url, payload)
}

type NodeConnectionParams struct {
	Event       string
	ID          string
	CountryCode string
	IP          string
}

func TrackNodeConnectionEvent(params NodeConnectionParams) {
	var url = host + nodeConnectionEndpoint
	data := map[string]interface{}{
		"ID":          params.ID,
		"Event":       params.Event,
		"CountryCode": params.CountryCode,
		"IP":          params.IP,
		"Time":        time.Now().UnixMilli(),
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	TrackEvent(url, payload)
}

type DealPayload struct {
	ID               string
	JobCreator       string
	ResourceProvider string
	JobID            string
}

func TrackDeal(params DealPayload) {
	var url = host + dealsEndpoint
	byts, _ := json.Marshal(params)
	payload := string(byts)

	TrackEvent(url, payload)
}
