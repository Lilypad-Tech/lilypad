package metricsDashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lilypad-tech/lilypad/pkg/data"
)

const jobsEndpoint = "jobs"
const nodeInfoEndpoint = "nodes"
const nodeConnectionEndpoint = "uptimes"
const dealsEndpoint = "deals"

var host = os.Getenv("API_HOST")

func trackEvent(path string, json string) {
	fmt.Println("===========<<")
	fmt.Println()
	fmt.Println(host, json)
	fmt.Println()
	fmt.Println("===========<<")
	if host == "" {
		return
	}

	fmt.Println("===========>>")
	fmt.Println()
	var url = host + "metrics-dashboard/" + path
	fmt.Println()
	fmt.Println("===========>>")
	fmt.Println(url)

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
	fmt.Println("===========")
	fmt.Println()
	fmt.Println("TrackJobOfferUpdate")
	fmt.Println(evOffer)
	spew.Dump(evOffer)
	fmt.Println()
	fmt.Println("===========")
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

	trackEvent(jobsEndpoint, payload)
}

func TrackNodeInfo(resourceOffer data.ResourceOffer) {
	data := map[string]interface{}{
		"ID":      resourceOffer.ResourceProvider,
		"GPU":     resourceOffer.Spec.GPU,
		"CPU":     resourceOffer.Spec.CPU,
		"RAM":     resourceOffer.Spec.RAM,
		"Modules": resourceOffer.Modules,
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	trackEvent(nodeInfoEndpoint, payload)
}

type NodeConnectionParams struct {
	Event       string
	ID          string
	CountryCode string
	IP          string
}

func TrackNodeConnectionEvent(params NodeConnectionParams) {
	data := map[string]interface{}{
		"ID":          params.ID,
		"Event":       params.Event,
		"CountryCode": params.CountryCode,
		"IP":          params.IP,
		"Time":        time.Now().UnixMilli(),
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	trackEvent(nodeConnectionEndpoint, payload)
}

type DealPayload struct {
	ID               string
	JobCreator       string
	ResourceProvider string
	JobID            string
}

func TrackDeal(params DealPayload) {
	byts, _ := json.Marshal(params)
	payload := string(byts)

	trackEvent(dealsEndpoint, payload)
}
