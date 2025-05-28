package options

func GetDefaultTimeoutOptions() data.DealTimeouts {
	return data.DealTimeouts{
		Agree: data.DealTimeout{
			Timeout:    GetDefaultServeOptionUint64("TIMEOUT_AGREE_TIME", 3600),
			Collateral: GetDefaultServeOptionUint64("TIMEOUT_AGREE_COLLATERAL", 1),
		},
		SubmitResults: data.DealTimeout{
			Timeout:    GetDefaultServeOptionUint64("TIMEOUT_SUBMIT_RESULTS_COLLATERAL", 3600),
			Collateral: GetDefaultServeOptionUint64("TIMEOUT_SUBMIT_RESULTS_COLLATERAL", 1),
		},
		JudgeResults: data.DealTimeout{
			Timeout:    GetDefaultServeOptionUint64("TIMEOUT_JUDGE_RESULTS_COLLATERAL", 3600),
			Collateral: GetDefaultServeOptionUint64("TIMEOUT_JUDGE_RESULTS_COLLATERAL", 1),
		},
		MediateResults: data.DealTimeout{
			Timeout:    GetDefaultServeOptionUint64("TIMEOUT_MEDIATE_RESULTS_COLLATERAL", 3600),
			Collateral: GetDefaultServeOptionUint64("TIMEOUT_MEDIATE_RESULTS_COLLATERAL", 1),
		},
	}
}
