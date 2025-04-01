package metricsDashboard

import (
	"encoding/json"
	"time"

	"github.com/lilypad-tech/lilypad/v2/pkg/data"
	"github.com/lilypad-tech/lilypad/v2/pkg/http"
)

const jobsEndpoint = "jobs"
const nodeInfoEndpoint = "nodes"
const nodeConnectionEndpoint = "uptimes"
const dealsEndpoint = "deals"
const namespace = "metrics-dashboard"

var host string

type NodeConnectionParams struct {
	Event       string
	ID          string
	CountryCode string
	IP          string
}

type DealPayload struct {
	ID               string
	JobCreator       string
	ResourceProvider string
	JobID            string
}

func Init(h string) {
	host = h
}

func TrackJobOfferUpdate(evOffer data.JobOfferContainer) {
	if host == "" {
		return
	}
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

	url := host + namespace + "/" + jobsEndpoint
	http.GenericJSONPostClient(url, payload)
}

func TrackNodeInfo(resourceOffer data.ResourceOffer) {
	if host == "" {
		return
	}
	data := map[string]interface{}{
		"ID":      resourceOffer.ResourceProvider,
		"GPU":     resourceOffer.Spec.GPU,
		"CPU":     resourceOffer.Spec.CPU,
		"RAM":     resourceOffer.Spec.RAM,
		"Modules": resourceOffer.Modules,
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	url := host + namespace + "/" + nodeInfoEndpoint
	http.GenericJSONPostClient(url, payload)
}

func TrackNodeConnectionEvent(params NodeConnectionParams) {
	if host == "" {
		return
	}
	data := map[string]interface{}{
		"ID":          params.ID,
		"Event":       params.Event,
		"CountryCode": params.CountryCode,
		"IP":          params.IP,
		"Time":        time.Now().UnixMilli(),
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	url := host + namespace + "/" + nodeConnectionEndpoint
	http.GenericJSONPostClient(url, payload)
}

func TrackDeal(params DealPayload) {
	if host == "" {
		return
	}
	byts, _ := json.Marshal(params)
	payload := string(byts)

	url := host + namespace + "/" + dealsEndpoint
	http.GenericJSONPostClient(url, payload)
}
