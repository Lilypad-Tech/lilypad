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

var host = os.Getenv("API_HOST")
var endpoint = "metrics-dashboard/logs"
var url = host + endpoint

func TrackEvent(json string) {
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
	var module = evOffer.JobOffer.Module.Name
	if module == "" {
		module = evOffer.JobOffer.Module.Repo + ":" + evOffer.JobOffer.Module.Hash
	}

	data := map[string]interface{}{
		"ID":         evOffer.ID,
		"JobOfferID": evOffer.JobOffer.ID,
		"CreatedAt":  evOffer.JobOffer.CreatedAt,
		"JobCreator": evOffer.JobCreator,
		"DealID":     evOffer.DealID,
		"State":      evOffer.State,
		"Module":     module,
		"Timestamp":  time.Now().UnixMilli(),
		"Event":      "JobOfferUpdate",
	}
	byts, _ := json.Marshal(data)
	payload := string(byts)

	TrackEvent(payload)
}
