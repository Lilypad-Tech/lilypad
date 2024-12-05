//go:build unit

package system

import "testing"

func TestGetOTelServiceName(t *testing.T) {
	tests := map[string]struct {
		input    Service
		expected string
	}{
		"solver": {
			input:    SolverService,
			expected: "solver",
		},
		"resource-provider": {
			input:    ResourceProviderService,
			expected: "resource_provider",
		},
		"job-creator": {
			input:    JobCreatorService,
			expected: "job_creator",
		},
		"mediator": {
			input:    MediatorService,
			expected: "mediator",
		},
		"default": {
			input:    DefaultService,
			expected: "default",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if result, expected := GetOTelServiceName(test.input), test.expected; result != expected {
				t.Fatalf("GetOTelServiceName(%v) returned %q; expected %q", test.input, result, expected)
			}
		})
	}
}
