package store

import (
	"github.com/lilypad-tech/lilypad/pkg/data"
)

// Adding all pkg/data types to pkg/store/sqlite/models converters

func ToJobOfferData(jobOffer data.JobOfferContainer) JobOfferData {
	jobOfferData := JobOfferData{}
	jobOfferData.CID = jobOffer.ID
	jobOfferData.DealID = jobOffer.DealID
	jobOfferData.JobCreator = jobOffer.JobCreator
	jobOfferData.State = jobOffer.State
	jobOfferData.JobOffer = jobOffer.JobOffer

	return jobOfferData
}

func ToResourceOfferData(resourceOffer data.ResourceOfferContainer) ResourceOfferData {
	resourceOfferData := ResourceOfferData{}
	resourceOfferData.CID = resourceOffer.ID
	resourceOfferData.DealID = resourceOffer.DealID
	resourceOfferData.ResourceProvider = resourceOffer.ResourceProvider
	resourceOfferData.State = resourceOffer.State
	resourceOfferData.ResourceOffer = resourceOffer.ResourceOffer

	return resourceOfferData
}

func ToDealData(deal data.DealContainer) DealData {
	dealData := DealData{}
	dealData.CID = deal.ID
	dealData.JobCreator = deal.JobCreator
	dealData.ResourceProvider = deal.ResourceProvider
	dealData.JobOffer = deal.JobOffer
	dealData.ResourceOffer = deal.ResourceOffer
	dealData.State = deal.State
	dealData.Deal = deal.Deal
	dealData.Transactions = deal.Transactions
	dealData.Mediator = deal.Mediator

	return dealData
}

func ToResultData(result data.Result) ResultData {
	resultData := ResultData{}
	resultData.CID = result.ID
	resultData.DealID = result.DealID
	resultData.DataID = result.DataID
	resultData.Error = result.Error
	resultData.InstructionCount = result.InstructionCount

	return resultData
}

func ToMatchDecisionData(matchDecision data.MatchDecision) MatchDecisionData {
	matchDecisionData := MatchDecisionData{}
	matchDecisionData.JobOffer = matchDecision.JobOffer
	matchDecisionData.ResourceOffer = matchDecision.ResourceOffer
	matchDecisionData.ResourceOffer = matchDecision.ResourceOffer
	return matchDecisionData
}

// }
//
// // represents a solver decision
// // the solver keeps track of "no" decisions to avoid trying to repeatedly match
// // things it's already decided it can't match
// type MatchDecisionData struct {
// 	gorm.Model
// 	JobOffer      string `json:"job_offer"`
// 	ResourceOffer string `json:"resource_offer"`
// 	Deal          string `json:"deal"`
// 	Result        bool   `json:"result"`
// }
