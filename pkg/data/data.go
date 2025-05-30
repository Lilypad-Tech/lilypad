package data

import (
	"encoding/json"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data/bacalhau"
)

// used by resource providers to describe their resources
// use by job offers to describe their requirements
// when used by resource providers - these are absolute values
// when used by job offers - these are minimum requirements
type MachineSpec struct {
	// Milli-GPU
	// Whilst it's unlikely that partial GPU's make sense
	// let's not use a float and fix the precision to 1/1000
	GPU int `json:"gpu"`

	GPUs []GPUSpec `json:"gpus"`

	// Milli-CPU
	CPU int `json:"cpu"`

	// Megabytes
	RAM int `json:"ram"`

	// Disk space available
	Disk int `json:"disk"`
}

type GPUSpec struct {
	Name string `json:"name"`

	Vendor string `json:"vendor"`

	VRAM int `json:"vram"`
}

// this is what is loaded from the template file in the git repo
type Module struct {
	Author      ModuleAuthor `json:"author"`
	Description string       `json:"description"`

	// An image URL
	Image string      `json:"image"`
	Model ModuleModel `json:"model"`

	// Expected input files
	InputFiles InputFiles `json:"inputFiles"`

	// the min spec that this module requires
	// e.g. does this module need a GPU?
	// the module file itself will contain this spec
	// and so the module will need to be downloaded
	// and executed for this spec to be known
	Machine MachineSpec `json:"machine"`

	// the bacalhau job spec
	Job bacalhau.Job `json:"job"`
}

type ModuleAuthor struct {
	Name string `json:"name"`

	// An Avatar URL
	Avatar  string `json:"avatar"`
	Address string `json:"address"`
}

type ModuleModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size string `json:"size"`
}

type InputFiles struct {
	Required []string `json:"required"`
	Optional []string `json:"optional"`
}

// describes a workload to be run
// this pins a go-template.yaml file
// that is a bacalhau job spec
type ModuleConfig struct {

	// used for the shortcuts
	// this is in the modules package
	// where we keep a map of named modules
	// and their versions onto the
	// repo, hash and path below
	Name string `json:"name"`

	// needs to be a http url for a git repo
	// we must be able to clone it without credentials
	Repo string `json:"repo"`
	// the git hash to pin the module
	// we will 'git checkout' this hash
	Hash string `json:"hash"`
	// once the checkout has been done
	// this is the path to the module template
	// within the repo
	Path string `json:"path"`
}

type Result struct {
	// this is the cid of the result where ID is set to empty string
	ID     string `json:"id"`
	DealID string `json:"deal_id"`
	// the CID of the actual results
	DataID           string `json:"data_id"`
	Error            string `json:"error"`
	InstructionCount uint64 `json:"instruction_count"`
}

// Provides compatibility for older clients that expect the results_id field
func (r Result) MarshalJSON() ([]byte, error) {
	// TODO(bgins) Remove when older clients have been deprecated

	// Create an auxiliary type to avoid recursively calling json.Marshal
	// https://stackoverflow.com/a/23046869
	type ResultAux Result

	// Add results_id field to the existing Result fields and marshal
	return json.Marshal(struct {
		ResultAux
		ResultsID string `json:"results_id"`
	}{
		ResultAux: ResultAux(r),
		ResultsID: r.DataID,
	})
}

// Provides compatibility for newer clients that expect the data_id field
func (r *Result) UnmarshalJSON(data []byte) error {
	// TODO(bgins) Remove when older clients have been deprecated

	// Create an auxiliary type to avoid recursively calling json.Unmarshal
	// https://stackoverflow.com/a/52433660
	type ResultAux Result

	// Unmarshal into auxiliary type
	var aux ResultAux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Cast the auxilliary type to Result
	*r = Result(aux)

	// Create a raw map to capture the results_id field
	var rawMap map[string]interface{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	// Check if results_id exists and assign it to DataID if so
	if resultsID, ok := rawMap["results_id"]; ok {
		if strID, ok := resultsID.(string); ok {
			r.DataID = strID
		}
	}

	return nil
}

// MarketPrice means - get me the best deal
// job creators will do this by default i.e. "just buy me the cheapest"
// FixedPrice means - take it or leave it
// resource creators will do this by default i.e. "this is my price"
type PricingMode string

const (
	MarketPrice PricingMode = "MarketPrice"
	FixedPrice  PricingMode = "FixedPrice"
)

// the mediator and directory services that are trusted
// by the RP and JC - the solver will find an intersection
// of these and attach them to the deal
type ServiceConfig struct {
	Solver   string   `json:"solver" toml:"solver"`
	Mediator []string `json:"mediator" toml:"mediator"`
	APIHost  string   `json:"api_host" toml:"api_host"`
}

type TargetConfig struct {
	Address string `json:"address" toml:"address"`
}

// posted to the solver by a job creator
type JobOffer struct {
	// this is the cid of the job offer where ID is set to empty string
	ID string `json:"id"`
	// Checked for recency on the solver and the
	// starting point for job run times
	CreatedAt int `json:"created_at"`
	// Prevents offers with duplicate IDs
	Nonce string `json:"nonce"`
	// the address of the job creator
	JobCreator string `json:"job_creator"`
	// the actual module that is being offered
	// this must hash to the ModuleID above
	Module ModuleConfig `json:"module"`
	// the spec required by the module
	// this will have been hoisted from the module itself
	Spec MachineSpec `json:"spec"`
	// the user inputs to the module
	// these values will power the go template
	Inputs map[string]string `json:"inputs"`
	// Expected input files
	InputFiles InputFiles `json:"inputFiles"`
	// tells the solver how to match these prices
	// for JC this will normally be MarketPrice
	Mode PricingMode `json:"mode"`
	// the offered price and timeouts
	Pricing  DealPricing  `json:"pricing"`
	Timeouts DealTimeouts `json:"timeouts"`

	// which parties are trusted by the job creator
	Services ServiceConfig `json:"trusted_parties"`

	// which node(s) (if any) to target
	Target TargetConfig `json:"target"`
}

// this is what the solver keeps track of so we can know
// what the current state of the deal is
type JobOfferContainer struct {
	ID         string   `json:"id"`
	DealID     string   `json:"deal_id"`
	JobCreator string   `json:"job_creator"`
	State      uint8    `json:"state"`
	JobOffer   JobOffer `json:"job_offer"`
}

// posted to the solver by a resource provider
type ResourceOffer struct {
	// this is the cid of the resource offer where ID is set to empty string
	ID string `json:"id"`
	// this is basically a nonce so we don't have one ID pointing at multiple offers
	CreatedAt int `json:"created_at"`
	// the address of the resource provider
	ResourceProvider string `json:"resource_provider"`
	// allows a resource provider to manage multiple offers
	// that are essentially the same
	Index int `json:"index"`
	// the spec being offered
	Spec MachineSpec `json:"spec"`
	// the module ID's that this resource provider can run
	// an empty list means ALL modules
	Modules []string `json:"modules"`
	// tells the solver how to match these prices
	// for RP this will normally be FixedPrice
	// we expect the default pricing to be filled in
	Mode PricingMode `json:"mode"`
	// the default pricing for this resource offer
	// i.e. this is for any module
	DefaultPricing  DealPricing  `json:"default_pricing"`
	DefaultTimeouts DealTimeouts `json:"default_timeouts"`
	// the pricing for each module
	// this allows a resource provider to charge more
	// for certain modules
	ModulePricing  map[string]DealPricing  `json:"module_pricing"`
	ModuleTimeouts map[string]DealTimeouts `json:"module_timeouts"`

	// which parties are trusted by the resource provider
	Services ServiceConfig `json:"trusted_parties"`
}

// this is what the solver keeps track of so we can know
// what the current state of the deal is
type ResourceOfferContainer struct {
	ID               string        `json:"id"`
	DealID           string        `json:"deal_id"`
	ResourceProvider string        `json:"resource_provider"`
	State            uint8         `json:"state"`
	ResourceOffer    ResourceOffer `json:"resource_offer"`
}

type DealMembers struct {
	Solver           string   `json:"solver"`
	JobCreator       string   `json:"job_creator"`
	ResourceProvider string   `json:"resource_provider"`
	Mediators        []string `json:"mediators"`
}

type DealTimeout struct {
	Timeout    uint64 `json:"timeout"`
	Collateral uint64 `json:"collateral"`
}

type DealTimeouts struct {
	Agree          DealTimeout `json:"agree"`
	SubmitResults  DealTimeout `json:"submit_results"`
	JudgeResults   DealTimeout `json:"judge_results"`
	MediateResults DealTimeout `json:"mediate_results"`
}

type DealPricing struct {
	InstructionPrice          uint64 `json:"instruction_price"`
	PaymentCollateral         uint64 `json:"payment_collateral"`
	ResultsCollateralMultiple uint64 `json:"results_collateral_multiple"`
	MediationFee              uint64 `json:"mediation_fee"`
}

// represents a solver decision
// the solver keeps track of "no" decisions to avoid trying to repeatedly match
// things it's already decided it can't match
type MatchDecision struct {
	JobOffer      string `json:"job_offer"`
	ResourceOffer string `json:"resource_offer"`
	Deal          string `json:"deal"`
	Result        bool   `json:"result"`
}

// this is the struct that will have it's ID taken and used
// as the reference for what both parties agreed to
// the solver will publish this deal to the directory
type Deal struct {
	// this is the cid of the deal where ID is set to empty string
	ID            string        `json:"id"`
	Members       DealMembers   `json:"members"`
	Pricing       DealPricing   `json:"pricing"`
	Timeouts      DealTimeouts  `json:"timeouts"`
	JobOffer      JobOffer      `json:"job_offer"`
	ResourceOffer ResourceOffer `json:"resource_offer"`
}

// we keep track of tx ids on behalf of resource providers
// and job creators - we use these to "marK" a deal as having
// had a transaction submitted but that tx has not yet been included
// in a block - therefore, we let the job creator & resource provider
// update these at will
type DealTransactionsJobCreator struct {
	Agree                string `json:"agree"`
	AcceptResult         string `json:"accept_result"`
	CheckResult          string `json:"check_result"`
	TimeoutAgree         string `json:"timeout_agree"`
	TimeoutSubmitResult  string `json:"timeout_submit_result"`
	TimeoutMediateResult string `json:"timeout_mediate_result"`
}

type DealTransactionsResourceProvider struct {
	Agree                string `json:"agree"`
	AddResult            string `json:"add_result"`
	TimeoutAgree         string `json:"timeout_agree"`
	TimeoutJudgeResult   string `json:"timeout_judge_result"`
	TimeoutMediateResult string `json:"timeout_mediate_result"`
}

type DealTransactionsMediator struct {
	MediationAcceptResult string `json:"mediation_accept_result"`
	MediationRejectResult string `json:"mediation_reject_result"`
}

type DealTransactions struct {
	ResourceProvider DealTransactionsResourceProvider `json:"resource_provider"`
	JobCreator       DealTransactionsJobCreator       `json:"job_creator"`
	Mediator         DealTransactionsMediator         `json:"mediator"`
}

type DealContainer struct {
	ID string `json:"id"`
	// Set at match time.
	// Used to compute match and execution durations.
	CreatedAt int `json:"created_at"`
	// Set when job outputs are uploaded by the resource provider.
	// Used to compute execution and output retrieval durations.
	UploadAt int `json:"upload_at"`
	// Set when job outputs are downloaded by the job creator.
	// Used to compute output retrieval and job run durations.
	DownloadAt       int              `json:"download_at"`
	JobCreator       string           `json:"job_creator"`
	ResourceProvider string           `json:"resource_provider"`
	JobOffer         string           `json:"job_offer"`
	ResourceOffer    string           `json:"resource_offer"`
	State            uint8            `json:"state"`
	Deal             Deal             `json:"deal"`
	Transactions     DealTransactions `json:"transactions"`
	Mediator         string           `json:"mediator"`
}

type MinerHashRate struct {
	ID       string  `json:"id"`
	Address  string  `json:"address"`
	Date     int64   `json:"date"`
	Hashrate float64 `json:"hashrate"`
}
