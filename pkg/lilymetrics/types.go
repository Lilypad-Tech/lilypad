package lilymetrics

import "time"

type JobMetadata struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	ClientID  string    `json:"ClientID"`
	Requester struct {
		RequesterNodeID    string `json:"RequesterNodeID"`
		RequesterPublicKey string `json:"RequesterPublicKey"`
	} `json:"Requester"`
}

type JobSpec struct {
	Engine        string `json:"Engine"`
	Verifier      string `json:"Verifier"`
	Publisher     string `json:"Publisher"`
	PublisherSpec struct {
		Type string `json:"Type"`
	} `json:"PublisherSpec"`
	Docker struct {
		Image      string   `json:"Image"`
		Entrypoint []string `json:"Entrypoint"`
	} `json:"Docker"`
}

type JobResources struct {
	GPU string `json:"GPU"`
}

type JobNetwork struct {
	Type string `json:"Type"`
}

type JobDeal struct {
	Concurrency int `json:"Concurrency"`
}

type Job struct {
	APIVersion string       `json:"APIVersion"`
	Metadata   JobMetadata  `json:"Metadata"`
	Spec       JobSpec      `json:"Spec"`
	Resources  JobResources `json:"Resources"`
	Network    JobNetwork   `json:"Network"`
	Timeout    int          `json:"Timeout"`
	Deal       JobDeal      `json:"Deal"`
}

type JobExecution struct {
	JobID              string `json:"JobID"`
	NodeId             string `json:"NodeId"`
	ComputeReference   string `json:"ComputeReference"`
	State              string `json:"State"`
	AcceptedAskForBid  bool   `json:"AcceptedAskForBid"`
	VerificationResult struct {
		Complete bool `json:"Complete"`
		Result   bool `json:"Result"`
	} `json:"VerificationResult"`
	PublishedResults struct {
		StorageSource string `json:"StorageSource"`
		Name          string `json:"Name"`
		CID           string `json:"CID"`
	} `json:"PublishedResults"`
	RunOutput struct {
		Stdout          string `json:"stdout"`
		Stdouttruncated bool   `json:"stdouttruncated"`
		Stderr          string `json:"stderr"`
		Stderrtruncated bool   `json:"stderrtruncated"`
		ExitCode        int    `json:"exitCode"`
		RunnerError     string `json:"runnerError"`
	} `json:"RunOutput"`
	Version    int       `json:"Version"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

type JobState struct {
	JobID      string         `json:"JobID"`
	Executions []JobExecution `json:"Executions"`
	State      string         `json:"State"`
	Version    int            `json:"Version"`
	CreateTime time.Time      `json:"CreateTime"`
	UpdateTime time.Time      `json:"UpdateTime"`
	TimeoutAt  time.Time      `json:"TimeoutAt"`
}

type JobData struct {
	Job   Job      `json:"Job"`
	State JobState `json:"State"`
}
type Log struct {
	id        string
	timestamp string
	details   string
}
type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
type Update struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
type Data struct {
	Dealid   string `json:"dealid"`
	State    string `json:"state"`
	Metadata string `json:"metadata"`
}
type JobUpdate struct {
	ID         string `json:"id"`
	ModuleID   string `json:"module_id"`
	JobID      string `json:"job_id"`
	Status     string `json:"status"`
	TimeUpdate string `json:"time_update"`
	Details    string `json:"details"`
	TimeStart  string `json:"time_start"`
	// Message     string `json:"message"`
}
type DbJob struct {
	id          string
	module_id   string
	job_id      string
	status      string
	time_update string
	details     string
	time_start  string
}