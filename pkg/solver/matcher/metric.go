package matcher

import "go.opentelemetry.io/otel/metric"

type metrics struct {
	jobOffers      metric.Int64Gauge
	resourceOffers metric.Int64Gauge
	deals          metric.Int64Gauge
}

func newMetrics(meter metric.Meter) (*metrics, error) {
	jobOffers, err := meter.Int64Gauge(
		"solver.matcher.job_offers",
		metric.WithDescription("Total number of job offers."),
	)
	if err != nil {
		return nil, err
	}

	resourceOffers, err := meter.Int64Gauge(
		"solver.matcher.resource_offers",
		metric.WithDescription("Total number of resource offers."),
	)
	if err != nil {
		return nil, err
	}

	deals, err := meter.Int64Gauge(
		"solver.matcher.deals",
		metric.WithDescription("Total number of deals."),
	)
	if err != nil {
		return nil, err
	}

	return &metrics{
		jobOffers:      jobOffers,
		resourceOffers: resourceOffers,
		deals:          deals,
	}, nil
}
