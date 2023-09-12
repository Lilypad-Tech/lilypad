package options

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/spf13/cobra"
)

func GetDefaultPricingMode(mode data.PricingMode) data.PricingMode {
	return data.PricingMode(GetDefaultServeOptionString("PRICING_MODE", string(mode)))
}

func GetDefaultPricingOptions() data.DealPricing {
	return data.DealPricing{
		// let's make the default price 1 ether
		InstructionPrice: GetDefaultServeOptionUint64("PRICING_INSTRUCTION_PRICE", 1),
		// 2 x ether for payment collateral (assuming modules that have a single instruction count)
		PaymentCollateral: GetDefaultServeOptionUint64("PRICING_PAYMENT_COLLATERAL", 2),
		// 2 x results collateral multiple
		ResultsCollateralMultiple: GetDefaultServeOptionUint64("PRICING_RESULTS_COLLATERAL_MULTIPLE", 2),
		// 1 ether for mediation fee
		MediationFee: GetDefaultServeOptionUint64("PRICING_MEDIATION_FEE", 1),
	}
}

func AddPricingModeCliFlags(cmd *cobra.Command, pricingMode *data.PricingMode) {
	cmd.PersistentFlags().StringVar(
		(*string)(pricingMode), "pricing-mode", string(*pricingMode),
		"set pricing mode (MarketPrice/FixedPrice)",
	)
}

func AddPricingCliFlags(cmd *cobra.Command, pricingConfig *data.DealPricing) {
	cmd.PersistentFlags().Uint64Var(
		&pricingConfig.InstructionPrice, "pricing-instruction-price", pricingConfig.InstructionPrice,
		`The price per instruction to offer (PRICING_INSTRUCTION_PRICE)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&pricingConfig.PaymentCollateral, "pricing-payment-collateral", pricingConfig.PaymentCollateral,
		`The payment collateral (PRICING_PAYMENT_COLLATERAL)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&pricingConfig.ResultsCollateralMultiple, "pricing-results-collateral-multiple", pricingConfig.ResultsCollateralMultiple,
		`The results collateral multiple (PRICING_RESULTS_COLLATERAL_MULTIPLE)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&pricingConfig.MediationFee, "pricing-mediation-fee", pricingConfig.MediationFee,
		`The mediation fee (PRICING_MEDIATION_FEE)`,
	)
}
