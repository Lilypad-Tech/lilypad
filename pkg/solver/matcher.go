package solver

import (
	"fmt"
	"sort"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
)

type ListOfResourceOffers []data.ResourceOffer

func (a ListOfResourceOffers) Len() int { return len(a) }
func (a ListOfResourceOffers) Less(i, j int) bool {
	return a[i].DefaultPricing.InstructionPrice < a[j].DefaultPricing.InstructionPrice
}
func (a ListOfResourceOffers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// the most basic of matchers
// basically just check if the resource offer >= job offer cpu, gpu & ram
// if the job offer is zero then it will match any resource offer
func doOffersMatch(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) bool {
	fmt.Printf("resourceOffer --------------------------------------\n")
	spew.Dump(resourceOffer)
	fmt.Printf("jobOffer --------------------------------------\n")
	spew.Dump(jobOffer)
	if resourceOffer.Spec.CPU < jobOffer.Spec.CPU {
		return false
	}
	if resourceOffer.Spec.GPU < jobOffer.Spec.GPU {
		return false
	}
	if resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
		return false
	}

	// if the resource provider has specified modules then check them
	if len(resourceOffer.Modules) > 0 {
		moduleID, err := data.GetModuleID(jobOffer.Module)
		if err != nil {
			return false
		}
		// if the resourceOffer.Modules array does not contain the moduleID then we don't match
		hasModule := false
		for _, module := range resourceOffer.Modules {
			if module == moduleID {
				hasModule = true
				break
			}
		}

		if !hasModule {
			return false
		}
	}

	// if both are fixed price then we filter out "cannot afford"
	if resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice {
		if resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice {
			return false
		}
	}

	mutualMediators := getMutualTrustedParties(resourceOffer.TrustedParties.Mediator, jobOffer.TrustedParties.Mediator)
	if len(mutualMediators) == 0 {
		return false
	}

	mutualDirectories := getMutualTrustedParties(resourceOffer.TrustedParties.Directory, jobOffer.TrustedParties.Directory)
	if len(mutualDirectories) == 0 {
		return false
	}

	fmt.Printf("mutualMediators --------------------------------------\n")
	spew.Dump(mutualMediators)
	fmt.Printf("mutualDirectories --------------------------------------\n")
	spew.Dump(mutualDirectories)
	return true
}

func getDeals(controller *SolverController) ([]data.Deal, error) {
	deals := []data.Deal{}

	resourceOffers, err := controller.store.GetResourceOffers(store.GetResourceOffersQuery{
		NotMatched: true,
	})
	if err != nil {
		return nil, err
	}

	jobOffers, err := controller.store.GetJobOffers(store.GetJobOffersQuery{
		NotMatched: true,
	})
	if err != nil {
		return nil, err
	}

	// loop over job offers
	for _, jobOffer := range jobOffers {

		// loop over resource offers
		matchingResourceOffers := []data.ResourceOffer{}
		for _, resourceOffer := range resourceOffers {
			if doOffersMatch(resourceOffer.ResourceOffer, jobOffer.JobOffer) {
				matchingResourceOffers = append(matchingResourceOffers, resourceOffer.ResourceOffer)
			}
		}

		log.Info().
			Int("resourceOffers", len(resourceOffers)).
			Msgf("MATCHING")

		// yay - we've got some matching resource offers
		// let's choose the cheapest one
		if len(matchingResourceOffers) > 0 {
			// now let's order the matching resource offers by price
			sort.Sort(ListOfResourceOffers(matchingResourceOffers))

			cheapestResourceOffer := matchingResourceOffers[0]

			deal, err := getDeal(jobOffer.JobOffer, cheapestResourceOffer)

			if err != nil {
				return nil, err
			}

			deals = append(deals, deal)
		}
	}

	log.Info().
		Int("jobOffers", len(jobOffers)).
		Int("resourceOffers", len(resourceOffers)).
		Int("deals", len(deals)).
		Msgf("Solver solving")

	return deals, nil
}
