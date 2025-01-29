package matcher

import (
	"context"
	"errors"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func GetMatchingDeals(
	ctx context.Context,
	db store.SolverStore,
	updateJobOfferState func(string, string, uint8) (*data.JobOfferContainer, error),
	tracer trace.Tracer,
	meter metric.Meter,
) ([]data.Deal, error) {
	metrics, err := newMetrics(meter)
	ctx, span := tracer.Start(ctx, "get_matching_deals")
	defer span.End()

	deals := []data.Deal{}

	// Get resource offers oldest first
	span.AddEvent("db.get_resource_offers.start")
	resourceOffers, err := db.GetResourceOffers(store.GetResourceOffersQuery{
		NotMatched:       true,
		OrderOldestFirst: true,
	})
	if err != nil {
		span.SetStatus(codes.Error, "get resource offers failed")
		span.RecordError(err)
		return nil, err
	}
	span.SetAttributes(attribute.KeyValue{
		Key:   "resource_offers",
		Value: attribute.StringSliceValue(data.GetResourceOfferContainerIDs(resourceOffers)),
	})
	span.AddEvent("db.get_resource_offers.done")

	// Get job offers oldest first
	span.AddEvent("db.get_job_offers.start")
	jobOffers, err := db.GetJobOffers(store.GetJobOffersQuery{
		NotMatched:       true,
		OrderOldestFirst: true,
	})
	if err != nil {
		span.SetStatus(codes.Error, "get job offers failed")
		span.RecordError(err)
		return nil, err
	}
	span.SetAttributes(attribute.KeyValue{
		Key:   "job_offers",
		Value: attribute.StringSliceValue(data.GetJobOfferContainerIDs(jobOffers)),
	})
	span.AddEvent("db.get_resource_offers.done")

	// Loop over job offers attempting to match with resource offers
	for _, jobOffer := range jobOffers {
		// Check for targeted jobs
		if jobOffer.JobOffer.Target.Address != "" {
			deal, err := getTargetedDeal(ctx, db, jobOffer, updateJobOfferState, tracer)
			if err != nil {
				return nil, err
			}

			if deal != nil {
				deals = append(deals, *deal)
			}
			continue
		}

		// Build list of matching resource offers
		matchingResourceOffers := []data.ResourceOffer{}
		for _, resourceOffer := range resourceOffers {
			_, matchSpan := tracer.Start(ctx, "match",
				trace.WithAttributes(attribute.String("job_offer.id", jobOffer.ID),
					attribute.String("resource_offer.id", resourceOffer.ID)),
			)

			matchSpan.AddEvent("db.get_match_decision.start")
			decision, err := db.GetMatchDecision(resourceOffer.ID, jobOffer.ID)
			if err != nil {
				matchSpan.SetStatus(codes.Error, "unable to retrieve match decision")
				matchSpan.RecordError(err)
				return nil, err
			}
			matchSpan.AddEvent("db.get_match_decision.done")

			// If a decision exists it means we've already tried to match the offers and should not try again
			if decision != nil {
				matchSpan.AddEvent("decision_already_checked",
					trace.WithAttributes(attribute.Bool("decision.result", decision.Result)))
				matchSpan.End()
				continue
			}

			// Check for a match
			matchSpan.AddEvent("match_offers.start")
			result := matchOffers(resourceOffer.ResourceOffer, jobOffer.JobOffer)
			logMatch(result)
			matchSpan.AddEvent("match_offers.done", trace.WithAttributes(result.attributes()...))

			if result.matched() {
				// Match found, add it to the matched list
				matchingResourceOffers = append(matchingResourceOffers, resourceOffer.ResourceOffer)
				matchSpan.AddEvent("append_match",
					trace.WithAttributes(attribute.KeyValue{
						Key:   "matching_resource_offers",
						Value: attribute.StringSliceValue(data.GetResourceOfferIDs(matchingResourceOffers)),
					}))
			} else {
				// No match, add a decision without a deal ID and a false result
				matchSpan.AddEvent("add_match_decision.start")
				_, err := db.AddMatchDecision(resourceOffer.ID, jobOffer.ID, "", false)
				if err != nil {
					matchSpan.SetStatus(codes.Error, "unable to record mismatch decision")
					matchSpan.RecordError(err)
					return nil, err
				}
				matchSpan.AddEvent("add_match_decision.done")
			}

			matchSpan.End()
		}

		// yay - we've got some matching resource offers
		// Now let's match the by pricing and age preference.
		if len(matchingResourceOffers) > 0 {
			var selectedResourceOffer data.ResourceOffer
			if jobOffer.JobOffer.Mode == data.MarketPrice {
				// For market price jobs, select the cheapest offer, breaking ties with age
				selectedResourceOffer = matchingResourceOffers[0]
				for _, offer := range matchingResourceOffers[1:] {
					if isCheaperOrOlder(offer, selectedResourceOffer) {
						selectedResourceOffer = offer
					}
				}
			} else {
				// For fixed price jobs, select the first (oldest) offer that meets the price requirement
				for _, offer := range matchingResourceOffers {
					if offer.DefaultPricing.InstructionPrice <= jobOffer.JobOffer.Pricing.InstructionPrice {
						selectedResourceOffer = offer
						break
					}
				}
			}

			// Selected resource offer ID will be set when we selected a resource offer.
			// Otherwise it will be a zero-value empty string.
			if selectedResourceOffer.ID != "" {
				span.AddEvent("get_deal.start",
					trace.WithAttributes(
						attribute.String("selected_resource_offer", selectedResourceOffer.ID),
						attribute.KeyValue{
							Key:   "matching_resource_offers",
							Value: attribute.StringSliceValue(data.GetResourceOfferIDs(matchingResourceOffers)),
						}))

				// Make a deal with the job offer and selected resource offer
				deal, err := data.GetDeal(jobOffer.JobOffer, selectedResourceOffer)
				if err != nil {
					span.SetStatus(codes.Error, "unable to get deal")
					span.RecordError(err)
					return nil, err
				}
				span.AddEvent("get_deal.done", trace.WithAttributes(attribute.String("deal.id", deal.ID)))

				// Add match decisions for all matching offers
				for _, matchingOffer := range matchingResourceOffers {
					addDealID := ""
					if selectedResourceOffer.ID == matchingOffer.ID {
						addDealID = deal.ID
					}

					// All match decisions had matching resource offers, set the result to true.
					// The match decision has a deal ID if it's resource offer was selected.
					span.AddEvent("add_match_decision.start")
					_, err := db.AddMatchDecision(matchingOffer.ID, jobOffer.ID, addDealID, true)
					if err != nil {
						span.SetStatus(codes.Error, "unable to add match decision")
						span.RecordError(err)
						return nil, err
					}
					span.AddEvent("add_match_decision.done")
				}

				// Add deal to overall deals list
				deals = append(deals, deal)
				span.AddEvent("append_deal",
					trace.WithAttributes(attribute.KeyValue{
						Key:   "deals",
						Value: attribute.StringSliceValue(data.GetDealIDs(deals)),
					}))
			}
		}
	}

	// Record metrics
	metrics.jobOffers.Record(ctx, int64(len(jobOffers)))
	metrics.resourceOffers.Record(ctx, int64(len(resourceOffers)))
	metrics.deals.Record(ctx, int64(len(deals)))

	log.Debug().
		Int("jobOffers", len(jobOffers)).
		Int("resourceOffers", len(resourceOffers)).
		Int("deals", len(deals)).
		Msgf(system.GetServiceString(system.SolverService, "Solver solving"))

	return deals, nil
}

// Returns true if first offer is cheaper, or has same price but is older
func isCheaperOrOlder(a, b data.ResourceOffer) bool {
	priceA := a.DefaultPricing.InstructionPrice
	priceB := b.DefaultPricing.InstructionPrice

	if priceA != priceB {
		return priceA < priceB // Choose cheaper price
	}
	return a.CreatedAt < b.CreatedAt // Choose older offer
}

// See if our jobOffer targets a specific address. If so, we will create a deal automatically
// with the matcing resourceOffer.
func getTargetedDeal(
	ctx context.Context,
	db store.SolverStore,
	jobOffer data.JobOfferContainer,
	updateJobOfferState func(string, string, uint8) (*data.JobOfferContainer, error),
	tracer trace.Tracer,
) (*data.Deal, error) {
	ctx, span := tracer.Start(ctx, "get_targeted_deal",
		trace.WithAttributes(attribute.String("job_offer.target.address", jobOffer.JobOffer.Target.Address)))
	defer span.End()

	span.AddEvent("db.get_resource_offer_by_address.start")
	resourceOffer, err := db.GetResourceOfferByAddress(jobOffer.JobOffer.Target.Address)
	if err != nil {
		return nil, err
	}

	// We don't have a resource provider for this address
	if resourceOffer == nil {
		log.Trace().
			Str("job offer", jobOffer.ID).
			Str("target address", jobOffer.JobOffer.Target.Address).
			Msgf("No resource provider found for address")
		span.SetStatus(codes.Error, "no resource provider found for address")
		span.RecordError(errors.New("no resource provider found for address"))

		updateJobOfferState(jobOffer.ID, "", data.GetAgreementStateIndex("JobOfferCancelled"))
		return nil, nil
	}
	span.AddEvent("db.get_resource_offer_by_address.found", trace.WithAttributes(attribute.String("resource_offer.id", resourceOffer.ID)))

	span.AddEvent("get_deal.start")
	deal, err := data.GetDeal(jobOffer.JobOffer, resourceOffer.ResourceOffer)
	if err != nil {
		span.SetStatus(codes.Error, "get deal failed")
		span.RecordError(err)
		return nil, err
	}
	span.AddEvent("get_deal.done", trace.WithAttributes(attribute.String("deal.id", deal.ID)))

	return &deal, nil
}
