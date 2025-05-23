package matcher

import (
	"fmt"
	"strings"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
)

type matchResult interface {
	matched() bool
	message() string
	attributes() []attribute.KeyValue
}

type offersMatched struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ offersMatched) matched() bool   { return true }
func (_ offersMatched) message() string { return "offers matched" }
func (result offersMatched) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
	}
}

type cpuMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ cpuMismatch) matched() bool   { return false }
func (_ cpuMismatch) message() string { return "did not match CPU" }
func (result cpuMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.Int("match_result.job_offer.spec.cpu", result.jobOffer.Spec.CPU),
		attribute.Int("match_result.resource_offer.spec.cpu", result.resourceOffer.Spec.CPU),
	}
}

type gpuMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ gpuMismatch) matched() bool   { return false }
func (_ gpuMismatch) message() string { return "did not match GPU" }
func (result gpuMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.Int("match_result.job_offer.spec.gpu", result.jobOffer.Spec.GPU),
		attribute.Int("match_result.resource_offer.spec.gpu", result.resourceOffer.Spec.GPU),
	}
}

type ramMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ ramMismatch) matched() bool   { return false }
func (_ ramMismatch) message() string { return "did not match RAM" }
func (result ramMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.Int("match_result.job_offer.spec.ram", result.jobOffer.Spec.RAM),
		attribute.Int("match_result.resource_offer.spec.ram", result.resourceOffer.Spec.RAM),
	}
}

type vramMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ vramMismatch) matched() bool   { return false }
func (_ vramMismatch) message() string { return "did not match VRAM" }
func (result vramMismatch) attributes() []attribute.KeyValue {
	var resourceOfferVRAMs []int
	for _, gpu := range result.resourceOffer.Spec.GPUs {
		resourceOfferVRAMs = append(resourceOfferVRAMs, gpu.VRAM)
	}

	var jobOfferVRAMS []int
	for _, gpu := range result.jobOffer.Spec.GPUs {
		jobOfferVRAMS = append(jobOfferVRAMS, gpu.VRAM)
	}

	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.IntSlice("match_result.resource_offer.spec.gpus.vram", resourceOfferVRAMs),
		attribute.IntSlice("match_result.job_offer.spec.gpus.vram", jobOfferVRAMS),
	}
}

type diskSpaceMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ diskSpaceMismatch) matched() bool   { return false }
func (_ diskSpaceMismatch) message() string { return "did not match disk space" }
func (result diskSpaceMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.Int("match_result.job_offer.spec.disk", result.jobOffer.Spec.Disk),
		attribute.Int("match_result.resource_offer.spec.disk", result.resourceOffer.Spec.Disk),
	}
}

type moduleIDError struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
	err           error
}

func (_ moduleIDError) matched() bool   { return false }
func (_ moduleIDError) message() string { return "error computing module ID" }
func (result moduleIDError) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.String("match_result.err", result.err.Error()),
		attribute.String("match_result.job_offer.module.repo", result.jobOffer.Module.Repo),
		attribute.String("match_result.job_offer.module.hash", result.jobOffer.Module.Hash),
	}
}

type moduleMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
	moduleID      string
}

func (_ moduleMismatch) matched() bool   { return false }
func (_ moduleMismatch) message() string { return "resource provider does not provide module" }
func (result moduleMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.String("match_result.module_id", result.moduleID),
		attribute.StringSlice("match_result.resource_offer.modules", result.resourceOffer.Modules),
		attribute.String("match_result.job_offer.module.repo", result.jobOffer.Module.Repo),
		attribute.String("match_result.job_offer.module.hash", result.jobOffer.Module.Hash),
	}
}

type marketPriceUnavailable struct {
	resourceOffer data.ResourceOffer
}

func (_ marketPriceUnavailable) matched() bool { return false }
func (_ marketPriceUnavailable) message() string {
	return "no support for market priced resource offers"
}
func (result marketPriceUnavailable) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.String("match_result.resource_offer.mode", string(result.resourceOffer.Mode)),
	}
}

type priceMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
	// moduleID      string
}

func (_ priceMismatch) matched() bool { return false }
func (_ priceMismatch) message() string {
	return "fixed price job offer cannot afford resource offer"
}
func (result priceMismatch) attributes() []attribute.KeyValue {
	// If the module instruction price is not specified, this lookup will use the zero-value of 0
	// moduleInstructionPrice := result.resourceOffer.ModulePricing[result.moduleID].InstructionPrice

	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.Int("match_result.job_offer.pricing.instruction_price", int(result.jobOffer.Pricing.InstructionPrice)),
		// attribute.Int("match_result.resource_offer.module_pricing.instruction_price", int(moduleInstructionPrice)),
		attribute.Int("match_result.resource_offer.default_pricing.instruction_price", int(result.resourceOffer.DefaultPricing.InstructionPrice)),
		attribute.String("match_result.job_offer.mode", string(result.jobOffer.Mode)),
		attribute.String("match_result.resource_offer.mode", string(result.resourceOffer.Mode)),
		// attribute.String("match_result.module_id", result.moduleID),
		attribute.StringSlice("match_result.resource_offer.modules", result.resourceOffer.Modules),
		attribute.String("match_result.job_offer.module.repo", result.jobOffer.Module.Repo),
		attribute.String("match_result.job_offer.module.hash", result.jobOffer.Module.Hash),
	}
}

// TODO(bgins) Rename to validator
type mediatorMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ mediatorMismatch) matched() bool   { return false }
func (_ mediatorMismatch) message() string { return "no matching mutual mediators" }
func (result mediatorMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.StringSlice("match_result.job_offer.services.mediator", result.jobOffer.Services.Mediator),
		attribute.StringSlice("match_result.resource_offer.services.mediator", result.resourceOffer.Services.Mediator),
	}
}

type solverMismatch struct {
	resourceOffer data.ResourceOffer
	jobOffer      data.JobOffer
}

func (_ solverMismatch) matched() bool   { return false }
func (_ solverMismatch) message() string { return "no matching solver" }
func (result solverMismatch) attributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("match_result", fmt.Sprintf("%T", result)),
		attribute.Bool("match_result.matched", result.matched()),
		attribute.String("match_result.message", result.message()),
		attribute.String("match_result.job_offer.services.solver", result.jobOffer.Services.Solver),
		attribute.String("match_result.resource_offer.services.solver", result.resourceOffer.Services.Solver),
	}
}

// the most basic of matchers
// basically just check if the resource offer >= job offer cpu, gpu & ram
// if the job offer is zero then it will match any resource offer
func matchOffers(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) matchResult {
	if resourceOffer.Spec.CPU < jobOffer.Spec.CPU {
		return cpuMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}
	if resourceOffer.Spec.GPU < jobOffer.Spec.GPU {
		return gpuMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}
	if resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
		return ramMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	// Skip VRAM check when job offer does not request VRAM
	if len(jobOffer.Spec.GPUs) > 0 {
		// Mismatch if job offer requests VRAM but resource provider has none
		if len(resourceOffer.Spec.GPUs) == 0 {
			return vramMismatch{
				jobOffer:      jobOffer,
				resourceOffer: resourceOffer,
			}
		}

		// Mismatch if job offer largest VRAM is greater than resource offer largest VRAM
		largestResourceOfferVRAM := getLargestVRAM(resourceOffer.Spec.GPUs)
		largestJobOfferVRAM := getLargestVRAM(jobOffer.Spec.GPUs)
		if largestResourceOfferVRAM < largestJobOfferVRAM {
			return vramMismatch{
				jobOffer:      jobOffer,
				resourceOffer: resourceOffer,
			}
		}
	}

	if resourceOffer.Spec.Disk < jobOffer.Spec.Disk {
		return diskSpaceMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	// moduleID, err := data.GetModuleID(jobOffer.Module)
	// if err != nil {
	// 	return moduleIDError{
	// 		jobOffer:      jobOffer,
	// 		resourceOffer: resourceOffer,
	// 		err:           err,
	// 	}
	// }

	// if the resource provider has specified modules then check them
	// if len(resourceOffer.Modules) > 0 {
	// 	// if the resourceOffer.Modules array does not contain the moduleID then we don't match
	// 	hasModule := false
	// 	for _, module := range resourceOffer.Modules {
	// 		if module == moduleID {
	// 			hasModule = true
	// 			break
	// 		}
	// 	}

	// 	if !hasModule {
	// 		return moduleMismatch{
	// 			jobOffer:      jobOffer,
	// 			resourceOffer: resourceOffer,
	// 			moduleID:      moduleID,
	// 		}
	// 	}
	// }

	// we don't currently support market priced resource offers
	if resourceOffer.Mode == data.MarketPrice {
		return marketPriceUnavailable{
			resourceOffer: resourceOffer,
		}
	}

	// if both are fixed price then we filter out "cannot afford"
	if resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice {
		if resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice {
			return priceMismatch{
				jobOffer:      jobOffer,
				resourceOffer: resourceOffer,
				// moduleID:      moduleID,
			}
		}
	}

	// TODO(bgins) Rename to validator
	mutualMediators := data.GetMutualServices(resourceOffer.Services.Mediator, jobOffer.Services.Mediator)
	if len(mutualMediators) == 0 {
		return mediatorMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	if resourceOffer.Services.Solver != jobOffer.Services.Solver {
		return solverMismatch{
			jobOffer:      jobOffer,
			resourceOffer: resourceOffer,
		}
	}

	return offersMatched{
		jobOffer:      jobOffer,
		resourceOffer: resourceOffer,
	}
}

func getLargestVRAM(gpus []data.GPUSpec) int {
	largestVRAM := 0
	for _, gpu := range gpus {
		if gpu.VRAM > largestVRAM {
			largestVRAM = gpu.VRAM
		}
	}
	return largestVRAM
}

func logMatch(result matchResult, log *zerolog.Logger) {
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
			Str("pricing mode", string(r.resourceOffer.Mode)).
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
