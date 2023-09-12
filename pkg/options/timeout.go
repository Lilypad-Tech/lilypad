package options

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

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

func AddTimeoutCliFlags(cmd *cobra.Command, timeoutConfig *data.DealTimeouts) {
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.Agree.Timeout, "timeout-agree-time", timeoutConfig.Agree.Timeout,
		`The number of seconds to timeout a deal when agreeing (TIMEOUT_AGREE_TIME)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.Agree.Collateral, "timeout-agree-collateral", timeoutConfig.Agree.Collateral,
		`The collateral to timeout a deal when agreeing (TIMEOUT_AGREE_COLLATERAL)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.SubmitResults.Timeout, "timeout-submit-results-time", timeoutConfig.SubmitResults.Timeout,
		`The number of seconds to timeout a deal when submitting results (TIMEOUT_SUBMIT_RESULTS_TIME)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.SubmitResults.Collateral, "timeout-submit-results-collateral", timeoutConfig.SubmitResults.Collateral,
		`The collateral to timeout a deal when submitting results (TIMEOUT_SUBMIT_RESULTS_COLLATERAL)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.JudgeResults.Timeout, "timeout-judge-results-time", timeoutConfig.JudgeResults.Timeout,
		`The number of seconds to timeout a deal when judging results (TIMEOUT_JUDGE_RESULTS_TIME)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.JudgeResults.Collateral, "timeout-judge-results-collateral", timeoutConfig.JudgeResults.Collateral,
		`The collateral to timeout a deal when judging results (TIMEOUT_JUDGE_RESULTS_COLLATERAL)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.MediateResults.Timeout, "timeout-mediate-results-time", timeoutConfig.MediateResults.Timeout,
		`The number of seconds to timeout a deal when mediating results (TIMEOUT_MEDIATE_RESULTS_TIME)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.MediateResults.Collateral, "timeout-mediate-results-collateral", timeoutConfig.MediateResults.Collateral,
		`The collateral to timeout a deal when mediating results (TIMEOUT_MEDIATE_RESULTS_COLLATERAL)`,
	)
}
