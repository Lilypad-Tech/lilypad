package stats

import "encoding/json"

type JobRun struct {
	DealID                    string          `json:"dealId"`
	ModuleID                  string          `json:"moduleId"`
	ResourceProvider          string          `json:"resourceProvider"`
	JobCreator                string          `json:"jobCreator"`
	Tflops                    *int64          `json:"tflops"`
	TotalDurationMilliseconds float64         `json:"totalDuration_ms"`
	ExtraData                 json.RawMessage `json:"extraData"`
	JobOfferCID               string          `json:"jobOfferCid"`
	ResourceOfferCID          string          `json:"resourceOfferCid"`
}
