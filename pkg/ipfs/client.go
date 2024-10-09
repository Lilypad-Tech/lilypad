package ipfs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ipfs/boxo/files"
	boxopath "github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/kubo/client/rpc"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	ma "github.com/multiformats/go-multiaddr"
)

type Client struct {
	API  coreiface.CoreAPI
	addr string
}

// NewClient creates an API client for the given ipfs node API multiaddress.
func NewClient(ctx context.Context, apiAddr string) (*Client, error) {
	addr, err := ma.NewMultiaddr(apiAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse api address '%s': %w", apiAddr, err)
	}

	// This http.Transport is the same that httpapi.NewApi would use if we weren't passing in our own http.Client
	defaultTransport := &http.Transport{
		Proxy:             http.ProxyFromEnvironment,
		DisableKeepAlives: true,
	}
	api, err := httpapi.NewApiWithClient(addr, &http.Client{
		Transport: defaultTransport,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to '%s': %w", apiAddr, err)
	}

	client := &Client{
		API:  api,
		addr: apiAddr,
	}

	return client, nil
}

func (c *Client) Get(ctx context.Context, cidString string, outputPath string) error {
	path, err := cidToPath(cidString)
	if err != nil {
	}

	node, err := c.API.Unixfs().Get(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to get path '%s': %w", path, err)
	}

	if err := files.WriteTo(node, outputPath); err != nil {
		return fmt.Errorf("failed to write to '%s': %w", outputPath, err)
	}

	return nil
}

func cidToPath(cidString string) (boxopath.Path, error) {
	c, err := cid.Decode(cidString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cid '%s': %w", cidString, err)
	}

	return boxopath.FromCid(c), nil
}
