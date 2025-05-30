package web3

import (
	"math/big"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/web3/bindings/controller"
)

// DealTimeout, DealTimeouts, and convert functions are retained
// for apiV1 compatibility with hardcoded values.
type dealTimeout struct {
	Timeout    uint64 `json:"timeout"`
	Collateral uint64 `json:"collateral"`
}

type dealTimeouts struct {
	Agree          dealTimeout `json:"agree"`
	SubmitResults  dealTimeout `json:"submit_results"`
	JudgeResults   dealTimeout `json:"judge_results"`
	MediateResults dealTimeout `json:"mediate_results"`
}

func convertDealTimeouts(
	timeouts dealTimeouts,
) controller.SharedStructsDealTimeouts {
	return controller.SharedStructsDealTimeouts{
		Agree:          convertDealTimeout(timeouts.Agree, false),
		SubmitResults:  convertDealTimeout(timeouts.SubmitResults, true),
		JudgeResults:   convertDealTimeout(timeouts.JudgeResults, true),
		MediateResults: convertDealTimeout(timeouts.MediateResults, false),
	}
}

func convertDealTimeout(
	timeout dealTimeout,
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
