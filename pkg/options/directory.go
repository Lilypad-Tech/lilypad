package options

import (
	"github.com/bacalhau-project/lilypad/pkg/directory"
	"github.com/spf13/cobra"
)

func NewDirectoryOptions() directory.DirectoryOptions {
	return directory.DirectoryOptions{
		Server: GetDefaultServerOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
}

func AddDirectoryCliFlags(cmd *cobra.Command, options *directory.DirectoryOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
	AddServerCliFlags(cmd, &options.Server)
}

func CheckDirectoryOptions(options directory.DirectoryOptions) error {
	err := CheckWeb3Options(options.Web3, true)
	if err != nil {
		return err
	}
	err = CheckServerOptions(options.Server)
	if err != nil {
		return err
	}
	return nil
}

func ProcessDirectoryOptions(options directory.DirectoryOptions) (directory.DirectoryOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckDirectoryOptions(options)
}
