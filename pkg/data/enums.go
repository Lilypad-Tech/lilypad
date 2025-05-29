package data

import (
	"errors"
	"fmt"
)

// ServiceType corresponds to ServiceType in TypeScript
var ServiceType = []string{
	"Solver",
	"Mediator",
	"Directory",
	"ResourceProvider",
	"JobCreator",
}

// AgreementState corresponds to AgreementState in TypeScript
// TimeoutSubmitResults, TimeoutJudgeResults, and TimeoutMediateResults are no longer used
// but we are stuck with them because of the enum. We plan to remove it: https://github.com/Lilypad-Tech/lilypad/issues/559
var AgreementState = []string{
	"DealNegotiating",
	"DealAgreed",
	"ResultsSubmitted",
	"ResultsAccepted",
	"ResultsChecked",
	"MediationAccepted",
	"MediationRejected",
	"TimeoutSubmitResults",
	"TimeoutJudgeResults",
	"TimeoutMediateResults",
	"JobOfferCancelled",
	"JobTimedOut",
	"Uploading",
	"Downloading",
	"TimeoutMatch",
	"TimeoutExecution",
	"TimeoutDownload",
}

// PaymentReason corresponds to PaymentReason in TypeScript
var PaymentReason = []string{
	"PaymentCollateral",
	"ResultsCollateral",
	"TimeoutCollateral",
	"JobPayment",
	"MediationFee",
}

// PaymentDirection corresponds to PaymentDirection in TypeScript
var PaymentDirection = []string{
	"PaidIn",
	"PaidOut",
	"Refunded",
	"Slashed",
}

// GetTypeIndex corresponds to getTypeIndex in TypeScript
func GetTypeIndex(name string, arr []string, itemType string) (uint8, error) {
	for i, t := range arr {
		if t == itemType {
			return uint8(i), nil
		}
	}
	return 0, errors.New(fmt.Sprintf("no %s of type %s", name, itemType))
}

// GetServiceType corresponds to getServiceType in TypeScript
func GetServiceType(itemType string) (uint8, error) {
	return GetTypeIndex("ServiceType", ServiceType, itemType)
}

// GetAgreementState corresponds to getAgreementState in TypeScript
func GetAgreementState(itemType string) (uint8, error) {
	return GetTypeIndex("AgreementState", AgreementState, itemType)
}

func GetAgreementStateString(itemType uint8) string {
	return AgreementState[itemType]
}

func GetAgreementStateIndex(itemType string) uint8 {
	index, _ := GetAgreementState(itemType)
	return index
}

func GetDefaultAgreementState() uint8 {
	return GetAgreementStateIndex("DealNegotiating")
}

func IsActiveAgreementState(itemType uint8) bool {
	return itemType == GetAgreementStateIndex("DealNegotiating") || itemType == GetAgreementStateIndex("DealAgreed")
}

func IsTerminalAgreementState(itemType uint8) bool {
	return itemType == GetAgreementStateIndex("JobOfferCancelled") ||
		itemType == GetAgreementStateIndex("TimeoutMatch") ||
		itemType == GetAgreementStateIndex("TimeoutExecution") ||
		itemType == GetAgreementStateIndex("TimeoutDownload") ||
		itemType == GetAgreementStateIndex("JobTimedOut") ||
		itemType == GetAgreementStateIndex("ResultsAccepted") ||
		itemType == GetAgreementStateIndex("MediationAccepted") ||
		itemType == GetAgreementStateIndex("MediationRejected")
}

func IsTimeoutAgreementState(state uint8) bool {
	return state == GetAgreementStateIndex("TimeoutMatch") ||
		state == GetAgreementStateIndex("TimeoutExecution") ||
		state == GetAgreementStateIndex("TimeoutDownload") ||
		state == GetAgreementStateIndex("JobTimedOut")
}

// GetPaymentReason corresponds to getPaymentReason in TypeScript
func GetPaymentReason(itemType string) (uint8, error) {
	return GetTypeIndex("PaymentReason", PaymentReason, itemType)
}

// GetPaymentDirection corresponds to getPaymentDirection in TypeScript
func GetPaymentDirection(itemType string) (uint8, error) {
	return GetTypeIndex("PaymentDirection", PaymentDirection, itemType)
}

func main() {
	// Sample usage
	index, err := GetServiceType("Solver")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Index:", index)
	}
}
