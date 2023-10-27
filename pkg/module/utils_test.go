package module

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

func TestPrepareModule(t *testing.T) {
	text, err := PrepareModule(data.ModuleConfig{
		Name: "cowsay:v0.0.2",
	})

	assert.NoError(t, err, "Should not return an error")
	assert.Contains(t, text, "cowsay", "Should contain the message")
	fmt.Printf("%s\n", text)
}

func TestLoadModule(t *testing.T) {
	module, err := LoadModule(data.ModuleConfig{
		Name: "cowsay:v0.0.2",
	}, map[string]string{
		"Message": "Hello, world!",
	})

	assert.NoError(t, err, "Should not return an error")
	assert.Equal(t, "grycap/cowsay@sha256:fad516b39e3a587f33ce3dbbb1e646073ef35e0b696bcf9afb1a3e399ce2ab0b", module.Job.Spec.Docker.Image)
	assert.Equal(t, "Hello, world!", module.Job.Spec.Docker.Entrypoint[1])
}

// TestSubst: [subst] can correctly substitute json encoded values into the template string.
func TestSubst(t *testing.T) {
	format := "Hello, %s!"
	inputs := []string{"hiro"}
	expectedOutput := "Hello, hiro!"

	jsonEncodedInputs := make([]string, 0, len(inputs))

	for _, input := range inputs {
		inputJ, err := json.Marshal(input)
		if err != nil {
			t.Fatalf("json marshall failed %v", err)
		}
		jsonEncodedInputs = append(jsonEncodedInputs, string(inputJ))
	}
	t.Logf("jsonEncodedInputs -%v %d", jsonEncodedInputs, len(jsonEncodedInputs))

	actualOutput := subst(format, jsonEncodedInputs...)

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, actualOutput)
	}
}
