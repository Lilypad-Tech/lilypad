package options

import (
	"embed"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/web3"
)

//go:embed configs
var fs embed.FS

type Config struct {
	Web3          web3.Web3Options   `toml:"web3"`
	ServiceConfig data.ServiceConfig `toml:"services"`
}

// TODO(bgins) Check for user-defined config files
func getConfig(network string) (*Config, error) {
	var config Config

	path := fmt.Sprintf("configs/%s.toml", network)
	config_toml, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Network %s does not exist", network)
	}

	_, err = toml.Decode(string(config_toml), &config)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse config file")
	}

	return &config, nil
}
