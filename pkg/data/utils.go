package data

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/controller"
	"github.com/ethereum/go-ethereum/common"

	mdag "github.com/ipfs/go-merkledag"
)

// CalculateCID takes an interface, serializes it to JSON, and returns its IPFS CID
func CalculateCID(v interface{}) (string, error) {
	// Serialize the struct to JSON
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	// Create a Dag Node from the JSON bytes
	node := mdag.NodeWithData(data)

	// Compute CID of the Dag Node
	c := node.Cid()

	return c.String(), nil
}

func GetJobOfferID(offer JobOffer) (string, error) {
	offer.ID = ""
	return CalculateCID(offer)
}

func GetResourceOfferID(offer ResourceOffer) (string, error) {
	offer.ID = ""
	return CalculateCID(offer)
}

func GetDealID(deal Deal) (string, error) {
	deal.ID = ""
	return CalculateCID(deal)
}

func GetModuleID(module ModuleConfig) (string, error) {
	return CalculateCID(module)
}

func GetMutualServices(a []string, b []string) []string {
	mutual := []string{}
	for _, aParty := range a {
		for _, bParty := range b {
			if aParty == bParty {
				mutual = append(mutual, aParty)
			}
		}
	}
	return mutual
}

func GetDeal(
	jobOffer JobOffer,
	resourceOffer ResourceOffer,
) (Deal, error) {
	mutualMediators := GetMutualServices(resourceOffer.Services.Mediator, jobOffer.Services.Mediator)
	if len(mutualMediators) <= 0 {
		return Deal{}, fmt.Errorf("no mutual mediators")
	}

	if jobOffer.Services.Solver != resourceOffer.Services.Solver {
		return Deal{}, fmt.Errorf("no mutual solver")
	}

	dealData := Deal{
		Members: DealMembers{
			Solver:           jobOffer.Services.Solver,
			JobCreator:       jobOffer.JobCreator,
			ResourceProvider: resourceOffer.ResourceProvider,
			Mediators:        mutualMediators,
		},
		// TODO: this assumes marketing pricing for the client
		// this should be configurable
		Pricing: resourceOffer.DefaultPricing,
		// TODO: this assumes resource provider timeouts
		// this should be configurable
		Timeouts:      resourceOffer.DefaultTimeouts,
		JobOffer:      jobOffer,
		ResourceOffer: resourceOffer,
	}

	id, err := GetDealID(dealData)

	if err != nil {
		return dealData, err
	}

	dealData.ID = id
	return dealData, nil
}

func GetJobOfferContainer(
	jobOffer JobOffer,
) JobOfferContainer {
	return JobOfferContainer{
		ID:         jobOffer.ID,
		DealID:     "",
		JobCreator: jobOffer.JobCreator,
		State:      GetDefaultAgreementState(),
		JobOffer:   jobOffer,
	}
}

func GetResourceOfferContainer(
	resourceOffer ResourceOffer,
) ResourceOfferContainer {
	return ResourceOfferContainer{
		ID:               resourceOffer.ID,
		DealID:           "",
		ResourceProvider: resourceOffer.ResourceProvider,
		State:            GetDefaultAgreementState(),
		ResourceOffer:    resourceOffer,
	}
}

func GetDealContainer(
	deal Deal,
) DealContainer {
	return DealContainer{
		ID:               deal.ID,
		JobCreator:       deal.JobOffer.JobCreator,
		ResourceProvider: deal.ResourceOffer.ResourceProvider,
		JobOffer:         deal.JobOffer.ID,
		ResourceOffer:    deal.ResourceOffer.ID,
		State:            GetDefaultAgreementState(),
		Deal:             deal,
	}
}

func CheckResourceOffer(resourceOffer ResourceOffer) error {
	if resourceOffer.Mode == MarketPrice {
		return fmt.Errorf("resource offer mode cannot be market price")
	}

	if resourceOffer.Services.Solver == "" {
		return fmt.Errorf("resource offer must name it's solver")
	}

	if len(resourceOffer.Services.Mediator) <= 0 {
		return fmt.Errorf("resource offer must have at least one trusted mediator")
	}

	return nil
}

func CheckJobOffer(jobOffer JobOffer) error {
	if jobOffer.Services.Solver == "" {
		return fmt.Errorf("job offer must name it's solver")
	}

	if len(jobOffer.Services.Mediator) <= 0 {
		return fmt.Errorf("job offer must have at least one trusted mediator")
	}

	return nil
}

func CheckResult(result Result) error {
	if result.DataID == "" && result.Error == "" {
		return fmt.Errorf("result must have a data id")
	}
	return nil
}

func ConvertDealMembers(
	members DealMembers,
) controller.SharedStructsDealMembers {
	mediators := []common.Address{}
	for _, mediator := range members.Mediators {
		mediators = append(mediators, common.HexToAddress(mediator))
	}
	return controller.SharedStructsDealMembers{
		Solver:           common.HexToAddress(members.Solver),
		JobCreator:       common.HexToAddress(members.JobCreator),
		ResourceProvider: common.HexToAddress(members.ResourceProvider),
		Mediators:        mediators,
	}
}

func EtherToWei(etherAmount float64) *big.Int {
	ether := new(big.Float).SetFloat64(etherAmount)
	weiMultiplier := new(big.Float).SetFloat64(1e18)
	wei := new(big.Float).Mul(ether, weiMultiplier)

	weiInt := new(big.Int)
	wei.Int(weiInt)

	return weiInt
}

func ConvertDealTimeout(
	timeout DealTimeout,
	withCollateral bool,
) controller.SharedStructsDealTimeout {
	collateral := EtherToWei(0)
	if withCollateral {
		collateral = EtherToWei(float64(timeout.Collateral))
	}
	return controller.SharedStructsDealTimeout{
		Timeout:    big.NewInt(int64(timeout.Timeout)),
		Collateral: collateral,
	}
}

func ConvertDealTimeouts(
	timeouts DealTimeouts,
) controller.SharedStructsDealTimeouts {
	return controller.SharedStructsDealTimeouts{
		Agree:          ConvertDealTimeout(timeouts.Agree, false),
		SubmitResults:  ConvertDealTimeout(timeouts.SubmitResults, true),
		JudgeResults:   ConvertDealTimeout(timeouts.JudgeResults, true),
		MediateResults: ConvertDealTimeout(timeouts.MediateResults, false),
	}
}

func ConvertDealPricing(
	pricing DealPricing,
) controller.SharedStructsDealPricing {
	return controller.SharedStructsDealPricing{
		InstructionPrice:          EtherToWei(float64(pricing.InstructionPrice)),
		PaymentCollateral:         EtherToWei(float64(pricing.PaymentCollateral)),
		ResultsCollateralMultiple: big.NewInt(int64(pricing.ResultsCollateralMultiple)),
		MediationFee:              EtherToWei(float64(pricing.MediationFee)),
	}
}
