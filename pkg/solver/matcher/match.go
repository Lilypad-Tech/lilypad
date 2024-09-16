package matcher

import (
	"strings"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/rs/zerolog/log"
)

type matchResult interface {
	matched() bool
	message() string
}

type offersMatched struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ offersMatched) matched() bool   { return true }
func (_ offersMatched) message() string { return "offers matched" }

type cpuMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ cpuMismatch) matched() bool   { return false }
func (_ cpuMismatch) message() string { return "did not match CPU" }

type gpuMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ gpuMismatch) matched() bool   { return false }
func (_ gpuMismatch) message() string { return "did not match GPU" }

type ramMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ ramMismatch) matched() bool   { return false }
func (_ ramMismatch) message() string { return "did not match RAM" }

type moduleIDError struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
	err           error
}

func (_ moduleIDError) matched() bool   { return false }
func (_ moduleIDError) message() string { return "error getting module ID" }

type moduleMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ moduleMismatch) matched() bool   { return false }
func (_ moduleMismatch) message() string { return "did not match modules" }

type marketPriceUnavailable struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ marketPriceUnavailable) matched() bool { return false }
func (_ marketPriceUnavailable) message() string {
	return "do not support market priced resource offers"
}

type priceMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ priceMismatch) matched() bool { return false }
func (_ priceMismatch) message() string {
	return "fixed price job offer cannot afford resource offer"
}

type mediatorMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ mediatorMismatch) matched() bool   { return false }
func (_ mediatorMismatch) message() string { return "no matching mutual mediators" }

type solverMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ solverMismatch) matched() bool   { return false }
func (_ solverMismatch) message() string { return "no matching solver" }

// the most basic of matchers
// basically just check if the resource offer >= job offer cpu, gpu & ram
// if the job offer is zero then it will match any resource offer
func doOffersMatch(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) matchResult {
	if resourceOffer.Spec.CPU < jobOffer.Spec.CPU {
		return &cpuMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}
	if resourceOffer.Spec.GPU < jobOffer.Spec.GPU {
		return &gpuMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}
	if resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
		return &ramMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	// if the resource provider has specified modules then check them
	if len(resourceOffer.Modules) > 0 {
		moduleID, err := data.GetModuleID(jobOffer.Module)
		if err != nil {
			return &moduleIDError{
				jobOffer:      jobOffer,
				resourceOffer: resourceOffer,
				err:           err,
			}
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
			return &moduleMismatch{
				jobOffer:      jobOffer,
				resourceOffer: resourceOffer,
			}
		}
	}

	// we don't currently support market priced resource offers
	if resourceOffer.Mode == data.MarketPrice {
		return &marketPriceUnavailable{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	// if both are fixed price then we filter out "cannot afford"
	if resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice {
		if resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice {
			return &priceMismatch{
				jobOffer:      jobOffer,
				resourceOffer: resourceOffer,
			}
		}
	}

	mutualMediators := data.GetMutualServices(resourceOffer.Services.Mediator, jobOffer.Services.Mediator)
	if len(mutualMediators) == 0 {
		return &mediatorMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	if resourceOffer.Services.Solver != jobOffer.Services.Solver {
		return &solverMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	return &offersMatched{
		jobOffer:      jobOffer,
		resourceOffer: resourceOffer,
	}
}

func logMatch(result matchResult) {
	switch r := result.(type) {
	case offersMatched:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Msg(r.message())
	case cpuMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Int("resource CPU", r.resourceOffer.Spec.CPU).
			Int("job CPU", r.jobOffer.Spec.CPU).
			Msg(r.message())
	case gpuMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Int("resource GPU", r.resourceOffer.Spec.GPU).
			Int("job GPU", r.jobOffer.Spec.GPU).
			Msg(r.message())
	case ramMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Int("resource RAM", r.resourceOffer.Spec.RAM).
			Int("job RAM", r.jobOffer.Spec.RAM).
			Msg(r.message())
	case moduleIDError:
		log.Error().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Err(r.err).
			Msg(r.message())
	case moduleMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Str("modules", strings.Join(r.resourceOffer.Modules, ", ")).
			Msg(r.message())
	case marketPriceUnavailable:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Msg(r.message())
	case priceMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Msg(r.message())
	case mediatorMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Msg(r.message())
	case solverMismatch:
		log.Trace().
			Str("resource offer", r.resourceOffer.ID).
			Str("job offer", r.jobOffer.ID).
			Msg(r.message())
	default:
		log.Trace().
			Msgf("unknown decision type: %v", r)
	}
}
