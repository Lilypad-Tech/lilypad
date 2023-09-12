package options

import "github.com/bacalhau-project/lilypad/pkg/directory"

func NewDirectoryOptions() directory.DirectoryOptions {
	return directory.DirectoryOptions{
		Server: GetDefaultServerOptions(),
		Web3:   GetDefaultWeb3Options(),
	}
}
