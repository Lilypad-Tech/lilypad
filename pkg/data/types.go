package data

import (
	"math/big"
)

// used by resource providers to describe their resources
// use by job offers to describe their requirements
// when used by resource providers - these are absolute values
// when used by job offers - these are minimum requirements
type Spec struct {
	// Milli-GPU
	// Whilst it's unlikely that partial GPU's make sense
	// let's not use a float and fix the precision to 1/1000
	GPU int `json:"gpu"`

	// Milli-CPU
	CPU int `json:"cpu"`

	// Megabytes
	RAM int `json:"ram"`
}

// describes a workload to be run
// this pins a go-template.yaml file
// that is a bacalhau job spec
type Module struct {
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
	// modules decide which resources they need
	Spec Spec `json:"spec"`
}

type Result struct {
	// this is the cid of the result where ID is set to empty string
	ID     string `json:"id"`
	DealID string `json:"deal_id"`
	// the CID of the actual results
	DataID           string  `json:"results_id"`
	InstructionCount big.Int `json:"instruction_count"`
}

// posted to the solver by a job creator
type JobOffer struct {
	// this is the cid of the job offer where ID is set to empty string
	ID string `json:"id"`
	// the address of the job creator
	JobCreator string `json:"job_creator"`
	// this is the CID of the Module description
	ModuleID string `json:"module_id"`
	// the actual module that is being offered
	// this must hash to the ModuleID above
	Module Module `json:"module"`
	// the user inputs to the module
	// these values will power the go template
	Inputs map[string]string `json:"inputs"`
}

// posted to the solver by a resource provider
type ResourceOffer struct {
	// this is the cid of the resource offer where ID is set to empty string
	ID string `json:"id"`
	// the address of the job creator
	ResourceProvider string `json:"resource_provider"`
	// allows a resource provider to manage multiple offers
	// that are essentially the same
	Index int `json:"index"`
	// the spec being offered
	Spec Spec `json:"spec"`
	// the module ID's that this resource provider can run
	// an empty list means ALL modules
	Modules []string `json:"modules"`
}

// represents the cost of a job
// this is both the bid and ask in the 2 sided marketplace
// job creators will attach this to their job offers
// and resource providers will attach this to their resource offers
// the solvers job is to propose the most efficient match
// for the job creator
type Deal struct {
	// this is the cid of the deal where ID is set to empty string
	ID                        string        `json:"id"`
	ResourceProvider          string        `json:"resource_provider"`
	JobCreator                string        `json:"job_creator"`
	InstructionPrice          big.Int       `json:"instruction_price"`
	Timeout                   big.Int       `json:"timeout"`
	TimeoutCollateral         big.Int       `json:"timeout_collateral"`
	PaymentCollateral         big.Int       `json:"payment_collateral"`
	ResultsCollateralMultiple big.Int       `json:"results_collateral_multiple"`
	MediationFee              big.Int       `json:"mediation_fee"`
	JobOffer                  JobOffer      `json:"job_offer"`
	ResourceOffer             ResourceOffer `json:"resource_offer"`
}
