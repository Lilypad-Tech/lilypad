package bacalhau

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"time"

	"k8s.io/apimachinery/pkg/selection"
)

type EngineSpec struct {
	Type   string
	Params map[string]interface{}
}

type PublisherSpec struct {
	Type   Publisher              `json:"Type,omitempty"`
	Params map[string]interface{} `json:"Params,omitempty"`
}

// for VM style executors
type JobSpecDocker struct {
	// this should be pullable by docker
	Image string `json:"Image,omitempty"`
	// optionally override the default entrypoint
	Entrypoint []string `json:"Entrypoint,omitempty"`
	// Parameters holds additional commandline arguments
	Parameters []string `json:"Parameters,omitempty"`
	// a map of env to run the container with
	EnvironmentVariables []string `json:"EnvironmentVariables,omitempty"`
	// working directory inside the container
	WorkingDirectory string `json:"WorkingDirectory,omitempty"`
}

// Describes a raw WASM job
type JobSpecWasm struct {
	// The module that contains the WASM code to start running.
	EntryModule StorageSpec `json:"EntryModule,omitempty"`

	// The name of the function in the EntryModule to call to run the job. For
	// WASI jobs, this will always be `_start`, but jobs can choose to call
	// other WASM functions instead. The EntryPoint must be a zero-parameter
	// zero-result function.
	EntryPoint string `json:"EntryPoint,omitempty"`

	// The arguments supplied to the program (i.e. as ARGV).
	Parameters []string `json:"Parameters,omitempty"`

	// The variables available in the environment of the running program.
	EnvironmentVariables map[string]string `json:"EnvironmentVariables,omitempty"`

	// TODO #880: Other WASM modules whose exports will be available as imports
	// to the EntryModule.
	ImportModules []StorageSpec `json:"ImportModules,omitempty"`
}

type S3StorageSpec struct {
	Bucket         string `json:"Bucket,omitempty"`
	Key            string `json:"Key,omitempty"`
	ChecksumSHA256 string `json:"Checksum,omitempty"`
	VersionID      string `json:"VersionID,omitempty"`
	Endpoint       string `json:"Endpoint,omitempty"`
	Region         string `json:"Region,omitempty"`
}

// StorageSpec represents some data on a storage engine. Storage engines are
// specific to particular execution engines, as different execution engines
// will mount data in different ways.
type StorageSpec struct {
	// StorageSource is the abstract source of the data. E.g. a storage source
	// might be a URL download, but doesn't specify how the execution engine
	// does the download or what it will do with the downloaded data.
	StorageSource StorageSourceType `json:"StorageSource,omitempty"`

	// Name of the spec's data, for reference.
	Name string `json:"Name,omitempty" example:"job-9304c616-291f-41ad-b862-54e133c0149e-host-QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL"` //nolint:lll

	// The unique ID of the data, where it makes sense (for example, in an
	// IPFS storage spec this will be the data's CID).
	// NOTE: The below is capitalized to match IPFS & IPLD (even though it's out of golang fmt)
	CID string `json:"CID,omitempty" example:"QmTVmC7JBD2ES2qGPqBNVWnX1KeEPNrPGb7rJ8cpFgtefe"`

	// Source URL of the data
	URL string `json:"URL,omitempty"`

	S3 *S3StorageSpec `json:"S3,omitempty"`

	// URL of the git Repo to clone
	Repo string `json:"Repo,omitempty"`

	// The path of the host data if we are using local directory paths
	SourcePath string `json:"SourcePath,omitempty"`

	// Allow write access for locally mounted inputs
	ReadWrite bool `json:"ReadWrite,omitempty"`

	// The path that the spec's data should be mounted on, where it makes
	// sense (for example, in a Docker storage spec this will be a filesystem
	// path).
	Path string `json:"Path,omitempty"`

	// Additional properties specific to each driver
	Metadata map[string]string `json:"Metadata,omitempty"`
}

type ResourceUsageConfig struct {
	// https://github.com/BTBurke/k8sresource string
	CPU string `json:"CPU,omitempty"`
	// github.com/c2h5oh/datasize string
	Memory string `json:"Memory,omitempty"`
	// github.com/c2h5oh/datasize string

	Disk string `json:"Disk,omitempty"`
	GPU  string `json:"GPU"` // unsigned integer string

}

//go:generate stringer -type=Network --trimprefix=Network
type Network int

const (
	// NetworkNone specifies that the job does not require networking.
	NetworkNone Network = iota

	// NetworkFull specifies that the job requires unfiltered raw IP networking.
	NetworkFull

	// NetworkHTTP specifies that the job requires HTTP networking to certain domains.
	//
	// The model is: the job specifier submits a job with the domain(s) it will
	// need to communicate with, the compute provider uses this to make some
	// decision about the risk of the job and bids accordingly, and then at run
	// time the traffic is limited to only the domain(s) specified.
	//
	// As a command, something like:
	//
	//  bacalhau docker run —network=http —domain=crates.io —domain=github.com -i ipfs://Qmy1234myd4t4,dst=/code rust/compile
	//
	// The “risk” for the compute provider is that the job does something that
	// violates its terms, the terms of its hosting provider or ISP, or even the
	// law in its jurisdiction (e.g. accessing and spreading illegal content,
	// performing cyberattacks). So the same sort of risk as operating a Tor
	// exit node.
	//
	// The risk for the job specifier is that we are operating in an environment
	// they are paying for, so there is an incentive to hijack that environment
	// (e.g. via a compromised package download that runs a crypto miner on
	// install, and uses up all the paid-for job time). Having the traffic
	// enforced to only domains specified makes those sorts of attacks much
	// trickier and less valuable.
	//
	// The compute provider might well enforce its limits by other means, but
	// having the domains specified up front allows it to skip bidding on jobs
	// it knows will fail in its executor. So this is hopefully a better UX for
	// job specifiers who can have their job picked up only by someone who will
	// run it successfully.
	NetworkHTTP
)

var domainRegex = regexp.MustCompile(`\b([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}\b`)

func ParseNetwork(s string) (Network, error) {
	for typ := NetworkNone; typ <= NetworkHTTP; typ++ {
		if equal(typ.String(), s) {
			return typ, nil
		}
	}

	return NetworkNone, fmt.Errorf("%T: unknown type '%s'", NetworkNone, s)
}

func (n Network) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *Network) UnmarshalText(text []byte) (err error) {
	name := string(text)
	*n, err = ParseNetwork(name)
	return
}

type NetworkConfig struct {
	Type    Network  `json:"Type"`
	Domains []string `json:"Domains,omitempty"`
}

type LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	Key string `json:"Key"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator selection.Operator `json:"Operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	Values []string `json:"Values,omitempty"`
}

type TargetingMode bool

const (
	TargetAny TargetingMode = false
	TargetAll TargetingMode = true
)

func (t TargetingMode) String() string {
	if bool(t) {
		return "all"
	} else {
		return "any"
	}
}

func ParseTargetingMode(s string) (TargetingMode, error) {
	switch s {
	case "any":
		return TargetAny, nil
	case "all":
		return TargetAll, nil
	default:
		return TargetAny, fmt.Errorf(`expecting "any" or "all", not %q`, s)
	}
}

type Deal struct {
	// Whether the job should be run on any matching node (false) or all
	// matching nodes (true). If true, other fields in this struct are ignored.
	TargetingMode TargetingMode `json:"TargetingMode,omitempty"`
	// The maximum number of concurrent compute node bids that will be
	// accepted by the requester node on behalf of the client.
	Concurrency int `json:"Concurrency,omitempty"`
}

// Spec is a complete specification of a job that can be run on some
// execution provider.
type Spec struct {
	// Deprecated: use EngineSpec.
	Engine Engine `json:"Engine,omitempty"`

	EngineSpec EngineSpec `json:"EngineSpec,omitempty"`

	// there can be multiple publishers for the job

	// Deprecated: use PublisherSpec instead
	Publisher     Publisher     `json:"Publisher,omitempty"`
	PublisherSpec PublisherSpec `json:"PublisherSpec,omitempty"`

	// Deprecated: use EngineSpec.
	Docker JobSpecDocker `json:"Docker,omitempty"`
	// Deprecated: use EngineSpec.
	Wasm JobSpecWasm `json:"Wasm,omitempty"`

	// the compute (cpu, ram) resources this job requires
	Resources ResourceUsageConfig `json:"Resources,omitempty"`

	// The type of networking access that the job needs
	Network NetworkConfig `json:"Network,omitempty"`

	// How long a job can run in seconds before it is killed.
	// This includes the time required to run, verify and publish results
	Timeout int64 `json:"Timeout,omitempty"`

	// the data volumes we will read in the job
	// for example "read this ipfs cid"
	Inputs []StorageSpec `json:"Inputs,omitempty"`

	// the data volumes we will write in the job
	// for example "write the results to ipfs"
	Outputs []StorageSpec `json:"Outputs,omitempty"`

	// Annotations on the job - could be user or machine assigned
	Annotations []string `json:"Annotations,omitempty"`

	// NodeSelectors is a selector which must be true for the compute node to run this job.
	NodeSelectors []LabelSelectorRequirement `json:"NodeSelectors,omitempty"`

	// Do not track specified by the client
	DoNotTrack bool `json:"DoNotTrack,omitempty"`

	// The deal the client has made, such as which job bids they have accepted.
	Deal Deal `json:"Deal,omitempty"`
}

type PublicKey []byte

func (pk PublicKey) MarshalText() ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(pk)), nil
}

func (pk *PublicKey) UnmarshalText(text []byte) error {
	ba, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return err
	}
	*pk = ba
	return nil
}

type JobRequester struct {
	// The ID of the requester node that owns this job.
	RequesterNodeID string `json:"RequesterNodeID,omitempty" example:"QmXaXu9N5GNetatsvwnTfQqNtSeKAD6uCmarbh3LMRYAcF"`

	// The public key of the Requester node that created this job
	// This can be used to encrypt messages back to the creator
	RequesterPublicKey PublicKey `json:"RequesterPublicKey,omitempty"`
}

type Metadata struct {
	// The unique global ID of this job in the bacalhau network.
	ID string `json:"ID,omitempty" example:"92d5d4ee-3765-4f78-8353-623f5f26df08"`

	// Time the job was submitted to the bacalhau network.
	CreatedAt time.Time `json:"CreatedAt,omitempty" example:"2022-11-17T13:29:01.871140291Z"`

	// The ID of the client that created this job.
	ClientID string `json:"ClientID,omitempty" example:"ac13188e93c97a9c2e7cf8e86c7313156a73436036f30da1ececc2ce79f9ea51"`

	Requester JobRequester `json:"Requester,omitempty"`
}

type Job struct {
	APIVersion string `json:"APIVersion" example:"V1beta1"`

	// TODO this doesn't seem like it should be a part of the job as it cannot be known by a client ahead of time.
	Metadata Metadata `json:"Metadata,omitempty"`

	// The specification of this job.
	Spec Spec `json:"Spec,omitempty"`
}

// JobWithInfo is the job request + the result of attempting to run it on the network
type JobWithInfo struct {
	// Job info
	Job Job `json:"Job"`
	// The current state of the job
	State JobState `json:"State"`
}
