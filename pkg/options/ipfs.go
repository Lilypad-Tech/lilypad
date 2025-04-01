package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/v2/pkg/ipfs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func GetDefaultIPFSOptions() ipfs.IPFSOptions {
	return ipfs.IPFSOptions{
		Addr: GetDefaultServeOptionString("IPFS_CONNECT", ""),
	}
}

func AddIPFSCliFlags(cmd *cobra.Command, options *ipfs.IPFSOptions) {
	cmd.PersistentFlags().StringVar(
		&options.Addr, "ipfs-connect", options.Addr,
		`The IPFS multiaddress to connect to (IPFS_CONNECT)`,
	)
}

func ProcessIPFSOptions(options ipfs.IPFSOptions, network string) (ipfs.IPFSOptions, error) {
	config, err := getConfig(network)
	if err != nil {
		log.Error().Msgf("failed to load config for network %s: ", err)
		return options, err
	}

	// Apply configs when environment variables or command line options are not used
	if options.Addr == "" {
		options.Addr = config.IPFSOptions.Addr
	}

	return options, nil
}

func CheckIPFSOptions(options ipfs.IPFSOptions) error {
	if len(options.Addr) == 0 {
		return fmt.Errorf("No IPFS multiaddress specified - please use IPFS_CONNECT or --ipfs-connect")
	}
	return nil
}
