//go:build unit

package matcher

import (
	"testing"

	"github.com/lilypad-tech/lilypad/pkg/data"
)

func TestMatchOffers(t *testing.T) {
	services := data.ServiceConfig{
		Solver:   "oranges",
		Mediator: []string{"apples"},
	}

	basicResourceOffer := data.ResourceOffer{
		Spec: data.MachineSpec{
			CPU: 1000,
			GPU: 1000,
			RAM: 1024,
			GPUs: []data.GPUSpec{
				{
					Name:   "NVIDIA RTX 4090",
					Vendor: "NVIDIA",
					VRAM:   24576, // MB
				},
			},
		},
		DefaultPricing: data.DealPricing{
			InstructionPrice: 10,
		},
		Mode:     data.FixedPrice,
		Services: services,
	}

	basicJobOffer := data.JobOffer{
		Spec: data.MachineSpec{
			CPU:  1000,
			GPU:  1000,
			RAM:  1024,
			GPUs: []data.GPUSpec{},
		},
		Mode:     data.MarketPrice,
		Services: services,
	}

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

	sdxlModuleConfig := data.ModuleConfig{
		Name: "",
		Repo: "https://github.com/Lilypad-Tech/lilypad-module-sdxl",
		Hash: "v0.9-lilypad1",
		Path: "/lilypad_module.json.tmpl",
	}

	testCases := []struct {
		name          string
		resourceOffer func(offer data.ResourceOffer) data.ResourceOffer
		jobOffer      func(offer data.JobOffer) data.JobOffer
		shouldMatch   bool
	}{
		{
			name: "Basic match",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "CPU mismatch",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Spec.CPU = 2000
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "GPU mismatch",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Spec.GPU = 2000
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "RAM mismatch",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Spec.GPU = 2048
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "VRAM match when job creator specifies VRAM",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Spec.GPUs = []data.GPUSpec{{VRAM: 24576}}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Module = sdxlModuleConfig
				offer.Spec.GPUs = []data.GPUSpec{{VRAM: 20000}}
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "VRAM match when job creator does not specify VRAM",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Spec.GPUs = []data.GPUSpec{{VRAM: 24576}}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Module = sdxlModuleConfig
				offer.Spec.GPUs = []data.GPUSpec{}
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "VRAM requested is more than resource offer VRAM",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Spec.GPUs = []data.GPUSpec{{VRAM: 24576}}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Module = sdxlModuleConfig
				offer.Spec.GPUs = []data.GPUSpec{{VRAM: 40000}}
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "VRAM requested but resource offer has none",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Spec.GPUs = []data.GPUSpec{}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Module = sdxlModuleConfig
				offer.Spec.GPUs = []data.GPUSpec{{VRAM: 49152}}
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Resource provider supports module",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				moduleID, _ := data.GetModuleID(cowsayModuleConfig)
				offer.Modules = []string{moduleID}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Module = cowsayModuleConfig
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "Resource provider does not support module",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				moduleID, _ := data.GetModuleID(cowsayModuleConfig)
				offer.Modules = []string{moduleID}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Module = lilysayModuleConfig
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Empty mediators",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Mediator = []string{}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Services.Mediator = []string{}
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Fixed price - too expensive",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Mode = data.FixedPrice
				offer.Pricing.InstructionPrice = 9
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Fixed price - can afford",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				offer.Mode = data.FixedPrice
				offer.Pricing.InstructionPrice = 11
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "Resource provider using unimplemented market pricing",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Mode = data.MarketPrice
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Mismatched mediators",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Mediator = []string{"apples2"}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Different mediators with one matching mediator",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Mediator = []string{"apples2", "apples"}
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "Different solver",
			resourceOffer: func(offer data.ResourceOffer) data.ResourceOffer {
				offer.Services.Solver = "pears"
				return offer
			},
			jobOffer: func(offer data.JobOffer) data.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := matchOffers(tc.resourceOffer(basicResourceOffer), tc.jobOffer(basicJobOffer))
			if result.matched() != tc.shouldMatch {
				t.Errorf("Expected match to be %v, but got %v", tc.shouldMatch, result)
			}
		})
	}
}
