package solver

import (
	"context"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/metric"
)

type metrics struct {
	dealNegotiating       metric.Int64Gauge
	dealAgreed            metric.Int64Gauge
	resultsSubmitted      metric.Int64Gauge
	resultsAccepted       metric.Int64Gauge
	resultsRejected       metric.Int64Gauge
	resultsChecked        metric.Int64Gauge
	mediationAccepted     metric.Int64Gauge
	mediationRejected     metric.Int64Gauge
	timeoutSubmitResults  metric.Int64Gauge
	timeoutJudgeResults   metric.Int64Gauge
	timeoutMediateResults metric.Int64Gauge
	jobOfferCancelled     metric.Int64Gauge
}

type dealStats struct {
	stateCounts stateCounts
}

type stateCounts struct {
	dealNegotiating       int64
	dealAgreed            int64
	resultsSubmitted      int64
	resultsAccepted       int64
	resultsRejected       int64
	resultsChecked        int64
	mediationAccepted     int64
	mediationRejected     int64
	timeoutSubmitResults  int64
	timeoutJudgeResults   int64
	timeoutMediateResults int64
	jobOfferCancelled     int64
}

func newMetrics(meter metric.Meter) (*metrics, error) {
	// Deal states
	dealNegotiating, err := meter.Int64Gauge(
		"solver.deal.state.deal_negotiating",
		metric.WithDescription("Number of deals in a DealNegotiating state."),
	)
	if err != nil {
		return nil, err
	}
	dealAgreed, err := meter.Int64Gauge(
		"solver.deal.state.deal_agreed",
		metric.WithDescription("Number of deals in a DealAgreed state."),
	)
	if err != nil {
		return nil, err
	}
	resultsSubmitted, err := meter.Int64Gauge(
		"solver.deal.state.results_submitted",
		metric.WithDescription("Number of deals in a ResultsSubmitted state."),
	)
	if err != nil {
		return nil, err
	}
	resultsAccepted, err := meter.Int64Gauge(
		"solver.deal.state.results_accepted",
		metric.WithDescription("Number of deals in a ResultsAccepted state."),
	)
	if err != nil {
		return nil, err
	}
	resultsRejected, err := meter.Int64Gauge(
		"solver.deal.state.results_rejected",
		metric.WithDescription("Number of deals in a ResultsRejected state."),
	)
	if err != nil {
		return nil, err
	}
	resultsChecked, err := meter.Int64Gauge(
		"solver.deal.state.results_checked",
		metric.WithDescription("Number of deals in a ResultsChecked state."),
	)
	if err != nil {
		return nil, err
	}
	mediationAccepted, err := meter.Int64Gauge(
		"solver.deal.state.mediation_accepted",
		metric.WithDescription("Number of deals in a MediationAccepted state."),
	)
	if err != nil {
		return nil, err
	}
	mediationRejected, err := meter.Int64Gauge(
		"solver.deal.state.mediation_rejected",
		metric.WithDescription("Number of deals in a MediationRejected state."),
	)
	if err != nil {
		return nil, err
	}
	timeoutSubmitResults, err := meter.Int64Gauge(
		"solver.deal.state.timeout_submit_results",
		metric.WithDescription("Number of deals in a TimeoutSubmitResults state."),
	)
	if err != nil {
		return nil, err
	}
	timeoutJudgeResults, err := meter.Int64Gauge(
		"solver.deal.state.timeout_judge_results",
		metric.WithDescription("Number of deals in a TimeoutJudgeResults state."),
	)
	if err != nil {
		return nil, err
	}
	timeoutMediateResults, err := meter.Int64Gauge(
		"solver.deal.state.timeout_mediate_results",
		metric.WithDescription("Number of deals in a TimeoutMediateResults state."),
	)
	if err != nil {
		return nil, err
	}
	jobOfferCancelled, err := meter.Int64Gauge(
		"solver.deal.state.job_offer_cancelled",
		metric.WithDescription("Number of deals in a JobOfferCancelled state."),
	)
	if err != nil {
		return nil, err
	}

	return &metrics{
		dealNegotiating,
		dealAgreed,
		resultsSubmitted,
		resultsAccepted,
		resultsRejected,
		resultsChecked,
		mediationAccepted,
		mediationRejected,
		timeoutSubmitResults,
		timeoutJudgeResults,
		timeoutMediateResults,
		jobOfferCancelled,
	}, nil
}

func reportDealMetrics(ctx context.Context, meter metric.Meter, deals []data.DealContainer) error {
	var dealStats dealStats

	// Compute deal state counts
	for _, deal := range deals {
		switch data.GetAgreementStateString(deal.State) {
		case "DealNegotiating":
			dealStats.stateCounts.dealNegotiating += 1
		case "DealAgreed":
			dealStats.stateCounts.dealAgreed += 1
		case "ResultsSubmitted":
			dealStats.stateCounts.resultsSubmitted += 1
		case "ResultsAccepted":
			dealStats.stateCounts.resultsAccepted += 1
		case "ResultsRejected":
			dealStats.stateCounts.resultsRejected += 1
		case "ResultsChecked":
			dealStats.stateCounts.resultsChecked += 1
		case "MediationAccepted":
			dealStats.stateCounts.mediationAccepted += 1
		case "MediationRejected":
			dealStats.stateCounts.mediationRejected += 1
		case "TimeoutSubmitResults":
			dealStats.stateCounts.timeoutSubmitResults += 1
		case "TimeoutJudgeResults":
			dealStats.stateCounts.timeoutJudgeResults += 1
		case "TimeoutMediateResults":
			dealStats.stateCounts.timeoutMediateResults += 1
		case "JobOfferCancelled":
			dealStats.stateCounts.jobOfferCancelled += 1
		default:
			log.Warn().Msgf("unknown deal state ID: %d", deal.State)
		}
	}

	// Record metrics
	metrics, err := newMetrics(meter)
	if err != nil {
		log.Warn().Msgf("failed to create solver metrics: %s", err)
		return err
	}
	metrics.dealNegotiating.Record(ctx, dealStats.stateCounts.dealNegotiating)
	metrics.dealAgreed.Record(ctx, dealStats.stateCounts.dealAgreed)
	metrics.resultsSubmitted.Record(ctx, dealStats.stateCounts.resultsSubmitted)
	metrics.resultsAccepted.Record(ctx, dealStats.stateCounts.resultsAccepted)
	metrics.resultsRejected.Record(ctx, dealStats.stateCounts.resultsRejected)
	metrics.resultsChecked.Record(ctx, dealStats.stateCounts.resultsChecked)
	metrics.mediationAccepted.Record(ctx, dealStats.stateCounts.mediationAccepted)
	metrics.mediationRejected.Record(ctx, dealStats.stateCounts.mediationRejected)
	metrics.timeoutSubmitResults.Record(ctx, dealStats.stateCounts.timeoutSubmitResults)
	metrics.timeoutJudgeResults.Record(ctx, dealStats.stateCounts.timeoutJudgeResults)
	metrics.timeoutMediateResults.Record(ctx, dealStats.stateCounts.timeoutMediateResults)
	metrics.jobOfferCancelled.Record(ctx, dealStats.stateCounts.jobOfferCancelled)

	return nil
}
