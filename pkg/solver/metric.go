package solver

import (
	"context"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/system"
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
	jobTimedOut           metric.Int64Gauge
}

type jobStats struct {
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
	jobTimedOut           int64
}

func newMetrics(meter metric.Meter) (*metrics, error) {
	// Deal states
	dealNegotiating, err := meter.Int64Gauge(
		"solver.deal.state.deal_negotiating",
		metric.WithDescription("Number of jobs in a DealNegotiating state."),
	)
	if err != nil {
		return nil, err
	}
	dealAgreed, err := meter.Int64Gauge(
		"solver.deal.state.deal_agreed",
		metric.WithDescription("Number of jobs in a DealAgreed state."),
	)
	if err != nil {
		return nil, err
	}
	resultsSubmitted, err := meter.Int64Gauge(
		"solver.deal.state.results_submitted",
		metric.WithDescription("Number of jobs in a ResultsSubmitted state."),
	)
	if err != nil {
		return nil, err
	}
	resultsAccepted, err := meter.Int64Gauge(
		"solver.deal.state.results_accepted",
		metric.WithDescription("Number of jobs in a ResultsAccepted state."),
	)
	if err != nil {
		return nil, err
	}
	resultsRejected, err := meter.Int64Gauge(
		"solver.deal.state.results_rejected",
		metric.WithDescription("Number of jobs in a ResultsRejected state."),
	)
	if err != nil {
		return nil, err
	}
	resultsChecked, err := meter.Int64Gauge(
		"solver.deal.state.results_checked",
		metric.WithDescription("Number of jobs in a ResultsChecked state."),
	)
	if err != nil {
		return nil, err
	}
	mediationAccepted, err := meter.Int64Gauge(
		"solver.deal.state.mediation_accepted",
		metric.WithDescription("Number of jobs in a MediationAccepted state."),
	)
	if err != nil {
		return nil, err
	}
	mediationRejected, err := meter.Int64Gauge(
		"solver.deal.state.mediation_rejected",
		metric.WithDescription("Number of jobs in a MediationRejected state."),
	)
	if err != nil {
		return nil, err
	}
	timeoutSubmitResults, err := meter.Int64Gauge(
		"solver.deal.state.timeout_submit_results",
		metric.WithDescription("Number of jobs in a TimeoutSubmitResults state."),
	)
	if err != nil {
		return nil, err
	}
	timeoutJudgeResults, err := meter.Int64Gauge(
		"solver.deal.state.timeout_judge_results",
		metric.WithDescription("Number of jobs in a TimeoutJudgeResults state."),
	)
	if err != nil {
		return nil, err
	}
	timeoutMediateResults, err := meter.Int64Gauge(
		"solver.deal.state.timeout_mediate_results",
		metric.WithDescription("Number of jobs in a TimeoutMediateResults state."),
	)
	if err != nil {
		return nil, err
	}
	jobOfferCancelled, err := meter.Int64Gauge(
		"solver.job_offer.state.job_offer_cancelled",
		metric.WithDescription("Number of jobs in a JobOfferCancelled state."),
	)
	if err != nil {
		return nil, err
	}
	jobTimedOut, err := meter.Int64Gauge(
		"solver.job_offer.state.job_timed_out",
		metric.WithDescription("Number of jobs in a JobTimedOut state."),
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
		jobTimedOut,
	}, nil
}

func reportJobMetrics(ctx context.Context, meter metric.Meter, log *system.ServiceLogger, deals []data.DealContainer, jobOffers []data.JobOfferContainer) error {
	var jobStats jobStats

	// Compute deal state counts
	for _, deal := range deals {
		switch data.GetAgreementStateString(deal.State) {
		case "DealNegotiating":
			jobStats.stateCounts.dealNegotiating += 1
		case "DealAgreed":
			jobStats.stateCounts.dealAgreed += 1
		case "ResultsSubmitted":
			jobStats.stateCounts.resultsSubmitted += 1
		case "ResultsAccepted":
			jobStats.stateCounts.resultsAccepted += 1
		case "ResultsRejected":
			jobStats.stateCounts.resultsRejected += 1
		case "ResultsChecked":
			jobStats.stateCounts.resultsChecked += 1
		case "MediationAccepted":
			jobStats.stateCounts.mediationAccepted += 1
		case "MediationRejected":
			jobStats.stateCounts.mediationRejected += 1
		case "TimeoutSubmitResults":
			jobStats.stateCounts.timeoutSubmitResults += 1
		case "TimeoutJudgeResults":
			jobStats.stateCounts.timeoutJudgeResults += 1
		case "TimeoutMediateResults":
			jobStats.stateCounts.timeoutMediateResults += 1
		default:
			log.Trace("untracked deal state ID", deal.State)
		}
	}

	// Cancelled states may only exist in the job offer
	for _, offer := range jobOffers {
		switch data.GetAgreementStateString(offer.State) {
		case "JobOfferCancelled":
			jobStats.stateCounts.jobOfferCancelled += 1
		case "JobTimedOut":
			jobStats.stateCounts.jobTimedOut += 1
		default:
			log.Trace("job metrics skipped offer state ID", offer.State)
		}
	}

	// Record metrics
	metrics, err := newMetrics(meter)
	if err != nil {
		log.Warn("failed to create solver metrics", err)
		return err
	}
	metrics.dealNegotiating.Record(ctx, jobStats.stateCounts.dealNegotiating)
	metrics.dealAgreed.Record(ctx, jobStats.stateCounts.dealAgreed)
	metrics.resultsSubmitted.Record(ctx, jobStats.stateCounts.resultsSubmitted)
	metrics.resultsAccepted.Record(ctx, jobStats.stateCounts.resultsAccepted)
	metrics.resultsRejected.Record(ctx, jobStats.stateCounts.resultsRejected)
	metrics.resultsChecked.Record(ctx, jobStats.stateCounts.resultsChecked)
	metrics.mediationAccepted.Record(ctx, jobStats.stateCounts.mediationAccepted)
	metrics.mediationRejected.Record(ctx, jobStats.stateCounts.mediationRejected)
	metrics.timeoutSubmitResults.Record(ctx, jobStats.stateCounts.timeoutSubmitResults)
	metrics.timeoutJudgeResults.Record(ctx, jobStats.stateCounts.timeoutJudgeResults)
	metrics.timeoutMediateResults.Record(ctx, jobStats.stateCounts.timeoutMediateResults)
	metrics.jobOfferCancelled.Record(ctx, jobStats.stateCounts.jobOfferCancelled)
	metrics.jobTimedOut.Record(ctx, jobStats.stateCounts.jobTimedOut)

	return nil
}
