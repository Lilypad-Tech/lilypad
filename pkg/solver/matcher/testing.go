//go:build unit || bench

package matcher

import "github.com/lilypad-tech/lilypad/v2/pkg/data"

type matchTestCase struct {
	name           string
	resourceOffer  data.ResourceOffer
	jobOffer       data.JobOffer
	expectedResult string
	shouldMatch    bool
}

func getMatchTestCases() []matchTestCase {
	// Base configurations
	baseServices := data.ServiceConfig{
		Solver:   "oranges",
		Mediator: []string{"apples"},
	}

	baseResourceOffer := data.ResourceOffer{
		Spec: data.MachineSpec{
			CPU:  1000,
			GPU:  1000,
			RAM:  1024,
			Disk: 2924295844659,
			GPUs: []data.GPUSpec{{
				Name:   "NVIDIA RTX 4090",
				Vendor: "NVIDIA",
				VRAM:   24576,
			}},
		},
		DefaultPricing: data.DealPricing{
			InstructionPrice: 10,
		},
		Mode:     data.FixedPrice,
		Services: baseServices,
	}

	baseJobOffer := data.JobOffer{
		Spec: data.MachineSpec{
			CPU:  1000,
			GPU:  1000,
			RAM:  1024,
			Disk: 2924295844659,
			GPUs: []data.GPUSpec{},
		},
		Mode:     data.MarketPrice,
		Services: baseServices,
	}

	// Module configs
	cowsayModuleConfig := data.ModuleConfig{
		Name: "cowsay",
		Repo: "https://github.com/Lilypad-Tech/lilypad-module-cowsay",
		Hash: "v0.0.4",
		Path: "/lilypad_module.json.tmpl",
	}

	lilysayModuleConfig := data.ModuleConfig{
		Name: "lilysay",
		Repo: "https://github.com/Lilypad-Tech/lilypad-module-lilysay",
		Hash: "v0.5.2",
		Path: "/lilypad_module.json.tmpl",
	}

	testCases := []matchTestCase{
		// Matching cases
		{
			name:           "Basic match",
			resourceOffer:  baseResourceOffer,
			jobOffer:       baseJobOffer,
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name: "VRAM match when job creator specifies VRAM",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.GPUs = []data.GPUSpec{{VRAM: 24576}}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.GPUs = []data.GPUSpec{{VRAM: 20000}}
				return j
			}(),
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name: "VRAM match when job creator does not specify VRAM",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.GPUs = []data.GPUSpec{{VRAM: 24576}}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				return j
			}(),
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name: "Disk space match when job creator specifies disk space",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.Disk = 2924295844659 // Bytes
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.Disk = 1000000000000
				return j
			}(),
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name: "Disk space match when job creator does not specify disk space",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.Disk = 2924295844659 // Bytes
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.Disk = 0 // zero-value
				return j
			}(),
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name: "Resource provider supports module",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				moduleID, _ := data.GetModuleID(cowsayModuleConfig)
				r.Modules = []string{moduleID}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Module = cowsayModuleConfig
				return j
			}(),
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name:          "Fixed price can afford",
			resourceOffer: baseResourceOffer,
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Mode = data.FixedPrice
				j.Pricing.InstructionPrice = 11
				return j
			}(),
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},
		{
			name: "Different mediators with one matching mediator",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Services.Mediator = []string{"apples2", "apples"}
				return r
			}(),
			jobOffer:       baseJobOffer,
			expectedResult: "*matcher.offersMatched",
			shouldMatch:    true,
		},

		// Mismatch cases
		{
			name:          "CPU mismatch",
			resourceOffer: baseResourceOffer,
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.CPU = 2000
				return j
			}(),
			expectedResult: "*matcher.cpuMismatch",
			shouldMatch:    false,
		},
		{
			name:          "GPU mismatch",
			resourceOffer: baseResourceOffer,
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.GPU = 2000
				return j
			}(),
			expectedResult: "*matcher.gpuMismatch",
			shouldMatch:    false,
		},
		{
			name:          "RAM mismatch",
			resourceOffer: baseResourceOffer,
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.RAM = 2048
				return j
			}(),
			expectedResult: "*matcher.ramMismatch",
			shouldMatch:    false,
		},
		{
			name: "VRAM requested is more than resource offer VRAM",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.GPUs = []data.GPUSpec{{VRAM: 24576}}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.GPUs = []data.GPUSpec{{VRAM: 40000}}
				return j
			}(),
			expectedResult: "*matcher.vramMismatch",
			shouldMatch:    false,
		},
		{
			name: "VRAM requested but resource offer has none",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.GPUs = []data.GPUSpec{}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.GPUs = []data.GPUSpec{{VRAM: 8192}}
				return j
			}(),
			expectedResult: "*matcher.vramMismatch",
			shouldMatch:    false,
		},
		{
			name: "Disk space requested is more than resource offer disk space",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Spec.Disk = 2924295844659 // Bytes
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Spec.Disk = 4000000000000
				return j
			}(),
			expectedResult: "*matcher.diskSpaceMismatch",
			shouldMatch:    false,
		},
		{
			name: "Resource provider does not support module",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				moduleID, _ := data.GetModuleID(cowsayModuleConfig)
				r.Modules = []string{moduleID}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Module = lilysayModuleConfig
				return j
			}(),
			expectedResult: "*matcher.moduleMismatch",
			shouldMatch:    false,
		},
		{
			name: "Resource provider using unimplemented market pricing",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Mode = data.MarketPrice
				return r
			}(),
			jobOffer:       baseJobOffer,
			expectedResult: "*matcher.marketPriceUnavailable",
			shouldMatch:    false,
		},
		{
			name:          "Fixed price too expensive",
			resourceOffer: baseResourceOffer,
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Mode = data.FixedPrice
				j.Pricing.InstructionPrice = 9
				return j
			}(),
			expectedResult: "*matcher.priceMismatch",
			shouldMatch:    false,
		},
		{
			name: "Mismatched mediators",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Services.Mediator = []string{"apples2"}
				return r
			}(),
			jobOffer:       baseJobOffer,
			expectedResult: "*matcher.mediatorMismatch",
			shouldMatch:    false,
		},
		{
			name: "Empty mediators",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Services.Mediator = []string{}
				return r
			}(),
			jobOffer: func() data.JobOffer {
				j := baseJobOffer
				j.Services.Mediator = []string{}
				return j
			}(),
			expectedResult: "*matcher.mediatorMismatch",
			shouldMatch:    false,
		},
		{
			name: "Mismatched solver",
			resourceOffer: func() data.ResourceOffer {
				r := baseResourceOffer
				r.Services.Solver = "pears"
				return r
			}(),
			jobOffer:       baseJobOffer,
			expectedResult: "*matcher.solverMismatch",
			shouldMatch:    false,
		},
	}

	return testCases
}
