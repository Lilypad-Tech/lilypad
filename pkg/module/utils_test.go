package module

import (
	"fmt"
	"testing"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/stretchr/testify/assert"
)

func TestPrepareModule(t *testing.T) {
	text, err := PrepareModule(data.ModuleConfig{
		Name:    "cowsay",
		Version: "v0.0.1",
	})

	assert.NoError(t, err, "Should not return an error")
	assert.Contains(t, text, "cowsay", "Should contain the message")
	fmt.Printf("%s\n", text)
}

func TestLoadModule(t *testing.T) {
	module, err := LoadModule(data.ModuleConfig{
		Name:    "cowsay",
		Version: "v0.0.1",
	}, map[string]string{
		"Message": "Hello, world!",
	})

	assert.NoError(t, err, "Should not return an error")
	assert.Equal(t, module.Job.Spec.Docker.Image, "grycap/cowsay:latest")
	assert.Equal(t, module.Job.Spec.Docker.Entrypoint[1], "Hello, world!")
}
