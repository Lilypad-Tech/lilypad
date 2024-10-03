package bacalhau

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	// "lilymetrics"

	// "context"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"

	// kubo "github.com/ipfs/kubo"
	// "github.com/ipfs/kubo/cmd/ipfs/kubo"

	config "github.com/ipfs/kubo/config"
	core "github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/repo/fsrepo"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/data/bacalhau"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"

	// "github.com/lilypad-tech/lilypad/pkg/lilymetrics"

	// "github.com/lilypad-tech/lilypad/pkg/lilymetrics"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"

	// "github.com/ipfs/kubo/core/coreapi"
	// coreapii "github.com/ipfs/go-ipfs/core/coreapi"
	coreapi "github.com/ipfs/kubo/core/coreapi"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	iface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/node/libp2p"

	// corepath "github.com/ipfs/kubo/core/corepath"

	"github.com/ipfs/kubo/plugin/loader"
	// "github.com/ipfs/kubo/repo/fsrepo"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
	ApiPort string
}

type BacalhauExecutor struct {
	Options        BacalhauExecutorOptions
	bacalhauEnv    []string
	bacalhauClient BacalhauClient
}

var (
	node *core.IpfsNode
	n    iface.CoreAPI
)

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

var loadPluginsOnce sync.Once

// CreateRepo Create a Temporary Repo
func CreateRepo() (string, error) {
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return "", onceErr
	}

	repoPath, err := os.MkdirTemp("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(io.Discard, 2048)
	// var testIdentity = config.Identity{
	// 	PeerID:  "QmNgdzLieYi8tgfo2WfTUzNVH5hQK9oAYGVf6dxN12NrHt",
	// 	PrivKey: "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",
	// }
	//cfg, err := config.InitWithIdentity(testIdentity)
	if err != nil {
		return "", err
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

// createNode Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (*core.IpfsNode, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	return core.NewNode(ctx, nodeOptions)
}

type APIHandler struct {
	api coreiface.CoreAPI
}

func NewAPIHandler(api coreiface.CoreAPI) *APIHandler {
	return &APIHandler{api: api}
}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/v0/id":
		h.handleID(w, r)
	case "/api/v0/add":
		h.handleAdd(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

type IDResponse struct {
	ID           string   `json:"ID"`
	PublicKey    string   `json:"PublicKey"`
	Addresses    []string `json:"Addresses"`
	AgentVersion string   `json:"AgentVersion"`
	Protocols    []string `json:"Protocols"`
}

func (h *APIHandler) handleID(w http.ResponseWriter, r *http.Request) {
	// id, err := h.api.Key().Self(n.)
	//id := node.Identity
	resp := IDResponse{
		ID:           node.Identity.String(),
		PublicKey:    "CAESIIzbqSqqlEOcZBlOLXS94+h2efs3blK47fsgMd9H6Okj", //node.Identity.PublicKey.String(),
		Addresses:    []string{"/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWKJDWrpxAD2Aa4FgdPVXVPVAHRjHBahHaiLfLjiao6EyQ"},
		AgentVersion: "kubo/0.28.0/",
		Protocols: []string{"/ipfs/bitswap",

			"/ipfs/bitswap/1.0.0",
			"/ipfs/bitswap/1.1.0",
			"/ipfs/bitswap/1.2.0",
			"/ipfs/id/1.0.0",
			"/ipfs/id/push/1.0.0",
			"/ipfs/lan/kad/1.0.0",
			"/ipfs/ping/1.0.0",
			"/libp2p/autonat/1.0.0",
			"/libp2p/circuit/relay/0.2.0/stop",
			"/libp2p/dcutr",
			"/x/"},
	}
	// node.Peerstore.PeersWithAddrs()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (h *APIHandler) handleAdd(w http.ResponseWriter, r *http.Request) {
	// Implement your logic to handle the `/api/v0/add` endpoint
	// ...
}
func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {
	client, err := newClient(options)
	if err != nil {
		// return nil, err
		fmt.Println("Error:", err)
	}
	bacalhauEnv := []string{
		fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost),
		fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	s, e := CreateRepo()
	if e != nil {
		fmt.Println("Error:", e)
	}
	// var err error
	node, err = createNode(context.Background(), s)

	// coreapi.NewCoreAPI(node)
	//apiServer, err := coreapi.NewCoreAPI(node)
	// apiServer, err := coreapii.NewCoreAPIFromNode(node)
	if err != nil {
		// Handle error
	}

	// Obtain the CoreAPI instance
	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		// Handle error
	}

	// Create an APIHandler instance
	handler := NewAPIHandler(api)

	// Start the HTTP server on port 5001
	fmt.Println("Starting server on :5001")
	go http.ListenAndServe(":5001", handler)
	// http.HandleFunc("/api/v0/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Implement your API request handling logic here
	// 	// You can use the `api` instance to interact with the IPFS node's API
	// 	// and return the appropriate response to the client.
	// 	// For example:
	// 	id, err := api.ID()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// err = apiServer.Run(ctx, coreiface.ServeOptionsHTTP, addr)
	// if err != nil {
	// 	// Handle error
	// }
	// apiServer.Swarm().s
	// Start the API server
	// err = apiServer.Start(context.Background())
	// if err != nil {
	// 	// Handle error
	// }

	// cfg := &core.BuildCfg{
	// 	Online:    true,
	// 	Permanent: true,
	// }
	// cfg.Host = libp2p.DefaultHostOption
	// Set the ListenAddr to ":5001" to make the node listen on port 5001
	// cfg.Host =core.  //libp2p.DefaultHostOption
	//cfg.Host = libp2p.
	//.("/ip4/0.0.0.0/tcp/5001")
	// cfg.ExtraOpts
	// kubo.(context.Background(), cfg)
	// err := daemon.Daemon(context.Background(), cfg).Start()
	// // daemon, err := kubo.InitDaemon(context.Background())
	// // err = daemon.Start(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// r := kubo.Start(kubo.BuildDefaultEnv)
	// fmt.Println("result ", r)
	valueMap := make(map[string]bool)

	// Assign values to the map
	valueMap["daemon"] = true
	// cfg1 := &core.BuildCfg{
	// 	Online:    true,
	// 	Permanent: true,
	// 	Host:      libp2p.DefaultHostOption,
	// 	Routing:   libp2p.DHTOption,
	// 	ExtraOpts: valueMap,
	// }

	// cfg := &core.BuildCfg{
	// 	Online:    true,
	// 	Permanent: true,
	// 	Host: func(ctx context.Context, id peer.ID, ps peer.Peerstore, bwr metrics.Reporter, fs *pnet.PSK) (host.Host, error) {
	// 		// h := host.Host(libp2p.New(ctx, libp2p.Identity(id), libp2p.Peerstore(ps), libp2p.BandwidthReporter(bwr), libp2p.PrivateNetwork(fs), libp2p.ListenAddrStrings("/ip4/")...))
	// 		libp2p.Host(ctx, id, ps, bwr, fs)
	// 		return h, nil
	// 		// return libp2p.Host(
	// 		// h, _ := libp2p.Host(ctx, id, ps, bwr, fs)
	// 		// return h.Host, nil
	// 		// // h.Host.Network().ListenAddresses(["/ip4/0.0.0.0/tcp/5001"])
	// 		// // h.Host.Net
	// 		// return h.Host, nil
	// 		// // return libp2p.New(ctx,
	// 		// // 	libp2p.Identity(id),
	// 		// // 	libp2p.Peerstore(ps),
	// 		// // 	libp2p.BandwidthReporter(bwr),
	// 		// // 	libp2p.PrivateNetwork(fs),
	// 		// // 	libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/5001"),
	// 		// // )
	// 	},
	// }

	//ctx := context.Background()
	// var testIdentity = config.Identity{
	// 	PeerID:  "QmNgdzLieYi8tgfo2WfTUzNVH5hQK9oAYGVf6dxN12NrHt",
	// 	PrivKey: "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",
	// }

	// good := []*config.Config{
	// 	{

	// 		Identity: testIdentity,
	// 		Addresses: config.Addresses{
	// 			Swarm: []string{"/ip4/0.0.0.0/tcp/4001", "/ip4/0.0.0.0/udp/4001/quic-v1"},
	// 			API:   []string{"/ip4/0.0.0.0/tcp/5001"},
	// 		},
	// 	},

	// 	{
	// 		Identity: testIdentity,
	// 		Addresses: config.Addresses{
	// 			Swarm: []string{"/ip4/0.0.0.0/tcp/4001", "/ip4/0.0.0.0/udp/4001/quic-v1"},
	// 			API:   []string{"/ip4/127.0.0.1/tcp/8000"},
	// 		},
	// 	},
	// }
	// cfg, err := config.Init(io.Discard, 2048)
	// r := &repo.Mock{
	// 	// C: *good[1],
	// 	C: *cfg,
	// 	D: syncds.MutexWrap(datastore.NewMapDatastore()),
	// }
	// node, err := core.NewNode(context.Background(), &core.BuildCfg{Online: true,
	// 	Permanent: true,
	// 	Repo:      r})
	// node, err := core.NewNode(context.Background(), cfg1)
	fmt.Println("ID " + node.Identity.String())
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}

	// Start the node
	// if err := node.Start(); err != nil {
	// 	fmt.Println("Error starting node:", err)
	// 	// return
	// }
	//{} core.ListenAddrOption("/ip4/0.0.0.0/tcp/5001")

	// node, err := core.NewNode(context.Background(), &core.BuildCfg{
	// 	Online: true,

	// })
	n, err = coreapi.NewCoreAPI(node)

	if err != nil {
		fmt.Println("Error initializing IPFS node:", err)
		// return "", err
	}
	// defer node.Close()

	log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)
	return &BacalhauExecutor{
		Options:        options,
		bacalhauEnv:    bacalhauEnv,
		bacalhauClient: *client,
	}, nil
}
func (executor *BacalhauExecutor) Id() (string, error) {
	nodeIdCmd := exec.Command(
		"bacalhau",
		"id",
	)
	nodeIdCmd.Env = executor.bacalhauEnv

	runOutputRaw, err := nodeIdCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error calling get id results %s, %s", err.Error(), runOutputRaw)
	}

	splitOutputs := strings.Split(strings.Trim(string(runOutputRaw), " \t\n"), "\n")
	runOutput := splitOutputs[len(splitOutputs)-1]

	var idResult struct {
		ID       string
		ClientID string
	}
	err = json.Unmarshal([]byte(runOutput), &idResult)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling job JSON %s %s", err.Error(), runOutputRaw)
	}

	return idResult.ID, nil
}
func (executor *BacalhauExecutor) Id2() (string, error) {
	nodeIdCmd := exec.Command(
		"bacalhau",
		"id",
	)
	nodeIdCmd.Env = executor.bacalhauEnv

	_, err := nodeIdCmd.CombinedOutput()
	if err != nil {
		// Handle error
	}

	// Obtain the CoreAPI instance
	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		// Handle error
	}

	// Create an APIHandler instance
	handler := NewAPIHandler(api)

	// Start the HTTP server on port 5001
	fmt.Println("Starting server on :5001")
	go http.ListenAndServe(":5001", handler)
	// http.HandleFunc("/api/v0/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Implement your API request handling logic here
	// 	// You can use the `api` instance to interact with the IPFS node's API
	// 	// and return the appropriate response to the client.
	// 	// For example:
	// 	id, err := api.ID()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// err = apiServer.Run(ctx, coreiface.ServeOptionsHTTP, addr)
	// if err != nil {
	// 	// Handle error
	// }
	// apiServer.Swarm().s
	// Start the API server
	// err = apiServer.Start(context.Background())
	// if err != nil {
	// 	// Handle error
	// }

	// cfg := &core.BuildCfg{
	// 	Online:    true,
	// 	Permanent: true,
	// }
	// cfg.Host = libp2p.DefaultHostOption
	// Set the ListenAddr to ":5001" to make the node listen on port 5001
	// cfg.Host =core.  //libp2p.DefaultHostOption
	//cfg.Host = libp2p.
	//.("/ip4/0.0.0.0/tcp/5001")
	// cfg.ExtraOpts
	// kubo.(context.Background(), cfg)
	// err := daemon.Daemon(context.Background(), cfg).Start()
	// // daemon, err := kubo.InitDaemon(context.Background())
	// // err = daemon.Start(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// r := kubo.Start(kubo.BuildDefaultEnv)
	// fmt.Println("result ", r)
	valueMap := make(map[string]bool)

	// Assign values to the map
	valueMap["daemon"] = true
	// cfg1 := &core.BuildCfg{
	// 	Online:    true,
	// 	Permanent: true,
	// 	Host:      libp2p.DefaultHostOption,
	// 	Routing:   libp2p.DHTOption,
	// 	ExtraOpts: valueMap,
	// }

	// cfg := &core.BuildCfg{
	// 	Online:    true,
	// 	Permanent: true,
	// 	Host: func(ctx context.Context, id peer.ID, ps peer.Peerstore, bwr metrics.Reporter, fs *pnet.PSK) (host.Host, error) {
	// 		// h := host.Host(libp2p.New(ctx, libp2p.Identity(id), libp2p.Peerstore(ps), libp2p.BandwidthReporter(bwr), libp2p.PrivateNetwork(fs), libp2p.ListenAddrStrings("/ip4/")...))
	// 		libp2p.Host(ctx, id, ps, bwr, fs)
	// 		return h, nil
	// 		// return libp2p.Host(
	// 		// h, _ := libp2p.Host(ctx, id, ps, bwr, fs)
	// 		// return h.Host, nil
	// 		// // h.Host.Network().ListenAddresses(["/ip4/0.0.0.0/tcp/5001"])
	// 		// // h.Host.Net
	// 		// return h.Host, nil
	// 		// // return libp2p.New(ctx,
	// 		// // 	libp2p.Identity(id),
	// 		// // 	libp2p.Peerstore(ps),
	// 		// // 	libp2p.BandwidthReporter(bwr),
	// 		// // 	libp2p.PrivateNetwork(fs),
	// 		// // 	libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/5001"),
	// 		// // )
	// 	},
	// }

	//ctx := context.Background()
	// var testIdentity = config.Identity{
	// 	PeerID:  "QmNgdzLieYi8tgfo2WfTUzNVH5hQK9oAYGVf6dxN12NrHt",
	// 	PrivKey: "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",
	// }

	// good := []*config.Config{
	// 	{

	// 		Identity: testIdentity,
	// 		Addresses: config.Addresses{
	// 			Swarm: []string{"/ip4/0.0.0.0/tcp/4001", "/ip4/0.0.0.0/udp/4001/quic-v1"},
	// 			API:   []string{"/ip4/0.0.0.0/tcp/5001"},
	// 		},
	// 	},

	// 	{
	// 		Identity: testIdentity,
	// 		Addresses: config.Addresses{
	// 			Swarm: []string{"/ip4/0.0.0.0/tcp/4001", "/ip4/0.0.0.0/udp/4001/quic-v1"},
	// 			API:   []string{"/ip4/127.0.0.1/tcp/8000"},
	// 		},
	// 	},
	// }
	// cfg, err := config.Init(io.Discard, 2048)
	// r := &repo.Mock{
	// 	// C: *good[1],
	// 	C: *cfg,
	// 	D: syncds.MutexWrap(datastore.NewMapDatastore()),
	// }
	// node, err := core.NewNode(context.Background(), &core.BuildCfg{Online: true,
	// 	Permanent: true,
	// 	Repo:      r})
	// node, err := core.NewNode(context.Background(), cfg1)
	fmt.Println("ID " + node.Identity.String())
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}

	// Start the node
	// if err := node.Start(); err != nil {
	// 	fmt.Println("Error starting node:", err)
	// 	// return
	// }
	//{} core.ListenAddrOption("/ip4/0.0.0.0/tcp/5001")

	// node, err := core.NewNode(context.Background(), &core.BuildCfg{
	// 	Online: true,

	// })
	n, err = coreapi.NewCoreAPI(node)

	if err != nil {
		fmt.Println("Error initializing IPFS node:", err)
		// return "", err
	}
	// defer node.Close()

	// log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)
	// client, err := newClient(options)
	// return &BacalhauExecutor{
	// 	Options:        options,
	// 	bacalhauEnv:    bacalhauEnv,
	// 	bacalhauClient: *client,
	// }, nil
	return "", nil
}

// Checks that Bacalhau is installed, correctly configured, and available
func (executor *BacalhauExecutor) IsAvailable() (bool, error) {
	isAlive, err := executor.bacalhauClient.alive()
	if !isAlive || err != nil {
		return false, fmt.Errorf("Bacalhau is not currently available. Please ensure that Bacalhau is running, then try again. %w", err)
	}

	// Check that we have the right version of Bacalhau
	version, err := executor.bacalhauClient.getVersion()
	if err != nil {
		return false, err
	}
	// TODO: we may want to relax this
	if version.GitVersion != "v1.3.2" {
		return false, errors.New("Bacalhau version must be v1.3.2")
	}

	return true, nil
}

func (executor *BacalhauExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	var specs []data.MachineSpec
	result, err := executor.bacalhauClient.getNodes()
	if err != nil {
		return specs, err
	}

	for _, node := range result.Nodes {
		spec := data.MachineSpec{
			CPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.CPU) * 1000, // convert float to "milli-CPU"
			RAM:  int(node.Info.ComputeNodeInfo.MaxCapacity.Memory),
			GPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.GPU),
			Disk: int(node.Info.ComputeNodeInfo.MaxCapacity.Disk),
		}
		for _, gpu := range node.Info.ComputeNodeInfo.MaxCapacity.GPUs {
			spec.GPUs = append(spec.GPUs, data.GPUSpec{
				Name:   gpu.Name,
				Vendor: string(gpu.Vendor),
				VRAM:   int(gpu.Memory),
			})
		}
		specs = append(specs, spec)
	}
	return specs, nil
}
func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	id, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	resultsDir, err := executor.copyJobResults(deal.ID, id)
	if err != nil {
		return nil, err
	}

	jobState, err := executor.getJobState(deal.ID, id)
	if err != nil {
		return nil, err
	}

	if len(jobState.State.Executions) <= 0 {
		return nil, fmt.Errorf("no executions found for job %s", id)
	}

	if jobState.State.State != bacalhau.JobStateCompleted {
		return nil, fmt.Errorf("job %s did not complete successfully: %s", id, jobState.State.State.String())
	}

	// TODO: we should think about WASM and instruction count here
	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       jobState.State.Executions[0].PublishedResult.CID,
		InstructionCount: 1,
	}

	return results, nil
}

// run the bacalhau job and return the job ID
func (executor *BacalhauExecutor) getJobID(
	deal data.DealContainer,
	module data.Module,
) (string, error) {
	// get a JSON string of the job
	jsonBytes, err := json.Marshal(module.Job)
	if err != nil {
		return "", fmt.Errorf("error getting job JSON for deal %s -> %s", deal.ID, err.Error())
	}
	bacalhauJobSpecDir, err := system.EnsureDataDir(filepath.Join("bacalhau-job-specs", deal.ID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder for job specs %s -> %s", deal.ID, err.Error())
	}
	jobPath := filepath.Join(bacalhauJobSpecDir, "job.json")
	err = system.WriteFile(jobPath, jsonBytes)
	if err != nil {
		return "", fmt.Errorf("error writing job JSON %s -> %s", deal.ID, err.Error())
	}
	// jobPath = "/home/arsen/lilypad/job.json"
	jobBytes, err := ioutil.ReadFile(jobPath)
	if err != nil {
		return "", fmt.Errorf("error reading job file %s -> %s", deal.ID, err.Error())
	}
	fmt.Println("jobPath ", string(jobBytes))
	// jobPath = "/home/arsen/lilypad/job.json"
	jobBytes, err := ioutil.ReadFile(jobPath)
	if err != nil {
		return "", fmt.Errorf("error reading job file %s -> %s", deal.ID, err.Error())
	}
	fmt.Println("jobPath ", string(jobBytes))

	runCmd := exec.Command(
		"bacalhau",
		"create",
		"--id-only",
		"--wait",
		jobPath,
	)
	runCmd.Env = executor.bacalhauEnv

	runOutput, err := runCmd.CombinedOutput()
	fmt.Println("runOutput " + string(runOutput))
	runOutput, err := runCmd.CombinedOutput()
	fmt.Println("runOutput " + string(runOutput))
	if err != nil {
		return "", fmt.Errorf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutput)
		return "", fmt.Errorf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutput)
	}

	id := strings.TrimSpace(string(runOutput))
	fmt.Printf("Got bacalhau job ID: %s\n", id)

	return id, nil
}

// func downloadDirectory(node *core.IpfsNode, cid path.Path, p string) error {
func downloadDirectory(n iface.CoreAPI, cid path.Path, p string) error {
	// Get UnixfsAPI
	//n, err := coreapi.NewCoreAPI(node)

	//unixfs := coreapi.UnixfsAPI()
	// Get directory contents
	unixfs := n.Unixfs()
	dirEntries, err := unixfs.Ls(context.Background(), cid)
	// dirEntries, err := unixfs.Ls(context.Background(), cid)
	if err != nil {
		return err
	}

	for entry := range dirEntries {
		// Construct path for the current entry
		entryPath := p + "/" + entry.Name
		// s := entry.Cid

		if entry.Type == iface.TDirectory {
			// If it's a directory, create corresponding local directory and download its contents
			if err := os.MkdirAll(entryPath, 0755); err != nil {
				return err
			}
			if err := downloadDirectory(n, path.FromCid(entry.Cid), entryPath); err != nil {
				return err
			}
		} else {
			// If it's a file, download it
			file, err := unixfs.Get(context.Background(), path.FromCid(entry.Cid))
			if err != nil {
				return err
			}
			defer file.Close()

			// Read file content
			// content, err := ioutil.ReadAll(file)
			f, ok := file.(files.File)
			if !ok {
				// handle error, file is not a regular file but could be a directory
			}
			content, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}

			// Write content to local file
			if err := ioutil.WriteFile(entryPath, content, 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
func (executor *BacalhauExecutor) copyJobResults(dealID string, jobID string) (string, error) {
	// time.Sleep(5 * time.Second)
	// span := lilymetrics.Trace(context.Background())
	// defer span.End()
	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, dealID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder of results %s -> %s", dealID, err.Error())
	}

	//bacalhau list --id-filter 2f5b845a-6ad3-41f1-969c-18a4209082e7 --output json
	info := exec.Command("bacalhau", "list", "--id-filter", jobID, "--output", "json")
	info.Env = executor.bacalhauEnv
	output, err := info.Output()
	fmt.Println("bacalhau list output ", string(output))

	var jobData []JobData
	err = json.Unmarshal([]byte(output), &jobData)
	if err != nil {
		fmt.Println("Error:", err)

	}
	if len(jobData) == 0 {
		return "", fmt.Errorf("no job data found for job %s check that bacalhau is running", jobID)

	}
	// ctx := context.Background()

	// Initialize IPFS node
	// node, err := core.NewNode(context.Background(), &core.BuildCfg{
	// 	Online: true,
	// })
	// n, err := coreapi.NewCoreAPI(node)

	// if err != nil {
	// 	fmt.Println("Error initializing IPFS node:", err)
	// 	return "", err
	// }
	// defer node.Close()

	// Example CID of the directory you want to download
	cidStr := jobData[0].State.Executions[0].PublishedResults.CID //"QmR5Jr3UJ1eRrbcfVRHX9wXam9JY5i9a1kj5oZTjK3xUN9"

	// Parse CID string
	// cid, err := coreapi.ParsePath(cidStr)
	cidObject, err := cid.Decode(cidStr)
	cid := path.FromCid(cidObject)
	if err != nil {
		fmt.Println("Error parsing CID:", err)

	}

	// Download the directory recursively
	err = downloadDirectory(n, cid, resultsDir)
	if err != nil {
		fmt.Println("Error downloading directory:", err)

	}
	fmt.Println("Directory downloaded successfully.")
	return resultsDir, nil

	//old code
	fmt.Println("cid ", jobData[0].State.Executions[0].PublishedResults.CID)
	exec.Command("ipfs", "get", jobData[0].State.Executions[0].PublishedResults.CID, "-o", resultsDir).Run()
	return resultsDir, nil
	copyCmdText := fmt.Sprintf("bacalhau get %s --output-dir %s --ipfs-swarm-addrs==/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWAW3hUQmCnzRdja86pAy9S5dBFXMRvsVF2vyWvAD3vnj1 ", jobID, resultsDir)
	log.Debug().Msgf("executor.bacalhauEnv: %s", executor.bacalhauEnv)
	log.Debug().Msgf("Executing command: %s", copyCmdText) // Log the command before execution for debugging
	copyResultsCmd := exec.Command("bacalhau", "get", jobID, "--output-dir", resultsDir)
	copyResultsCmd.Env = executor.bacalhauEnv
	fmt.Println("solved job")
	stdout, err := copyResultsCmd.StdoutPipe()
	// if err != nil {
	// 	// log.Fatal(err)
	// }

	if err := copyResultsCmd.Start(); err != nil {
		//log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		r := scanner.Text()
		fmt.Println("bacalhau get result ", r)
		// if strings.Contains(r, "Job result:") {
		// 	words := strings.Fields(r)
		// 	lastWord := words[len(words)-1]
		// 	fmt.Println("JOB " + lastWord)
		// 	cmdstd := exec.Command("cat", "/tmp/lilypad/data/downloaded-files/"+lastWord+"/stdout")

		// 	output, err := cmdstd.Output()
		// 	fmt.Println(string(output))
		// 	server.BroadcastToNamespace("/", "reply", output, "data")
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		return true
		// 	}

		// 	break
		// } else {
		// server.BroadcastToNamespace("/", "reply", r, "data")
		// }

	}

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = copyResultsCmd.CombinedOutput()
	// if err != nil {
	// 	return "", fmt.Errorf("error copying results %s -> %s, command executed: %s", dealID, err.Error(), copyCmdText)
	// }
	// wait for the results to be copied
	return resultsDir, nil

}

func (executor *BacalhauExecutor) getJobState(dealID string, jobID string) (*bacalhau.JobWithInfo, error) {
	// span := lilymetrics.Trace(context.Background())
	// defer span.End()
	describeCmd := exec.Command(
		"bacalhau",
		"describe",
		"--json",
		jobID,
	)
	describeCmd.Env = executor.bacalhauEnv

	output, err := describeCmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error calling describe command results %s -> %s", dealID, err.Error())
	}

	var job bacalhau.JobWithInfo
	err = json.Unmarshal(output, &job)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling job JSON %s -> %s", dealID, err.Error())
	}

	return &job, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)

type JobMetadata struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	ClientID  string    `json:"ClientID"`
	Requester struct {
		RequesterNodeID    string `json:"RequesterNodeID"`
		RequesterPublicKey string `json:"RequesterPublicKey"`
	} `json:"Requester"`
}

type JobSpec struct {
	Engine        string `json:"Engine"`
	Verifier      string `json:"Verifier"`
	Publisher     string `json:"Publisher"`
	PublisherSpec struct {
		Type string `json:"Type"`
	} `json:"PublisherSpec"`
	Docker struct {
		Image      string   `json:"Image"`
		Entrypoint []string `json:"Entrypoint"`
	} `json:"Docker"`
}

type JobResources struct {
	GPU string `json:"GPU"`
}

type JobNetwork struct {
	Type string `json:"Type"`
}

type JobDeal struct {
	Concurrency int `json:"Concurrency"`
}

type Job struct {
	APIVersion string       `json:"APIVersion"`
	Metadata   JobMetadata  `json:"Metadata"`
	Spec       JobSpec      `json:"Spec"`
	Resources  JobResources `json:"Resources"`
	Network    JobNetwork   `json:"Network"`
	Timeout    int          `json:"Timeout"`
	Deal       JobDeal      `json:"Deal"`
}

type JobExecution struct {
	JobID              string `json:"JobID"`
	NodeId             string `json:"NodeId"`
	ComputeReference   string `json:"ComputeReference"`
	State              string `json:"State"`
	AcceptedAskForBid  bool   `json:"AcceptedAskForBid"`
	VerificationResult struct {
		Complete bool `json:"Complete"`
		Result   bool `json:"Result"`
	} `json:"VerificationResult"`
	PublishedResults struct {
		StorageSource string `json:"StorageSource"`
		Name          string `json:"Name"`
		CID           string `json:"CID"`
	} `json:"PublishedResults"`
	RunOutput struct {
		Stdout          string `json:"stdout"`
		Stdouttruncated bool   `json:"stdouttruncated"`
		Stderr          string `json:"stderr"`
		Stderrtruncated bool   `json:"stderrtruncated"`
		ExitCode        int    `json:"exitCode"`
		RunnerError     string `json:"runnerError"`
	} `json:"RunOutput"`
	Version    int       `json:"Version"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

type JobState struct {
	JobID      string         `json:"JobID"`
	Executions []JobExecution `json:"Executions"`
	State      string         `json:"State"`
	Version    int            `json:"Version"`
	CreateTime time.Time      `json:"CreateTime"`
	UpdateTime time.Time      `json:"UpdateTime"`
	TimeoutAt  time.Time      `json:"TimeoutAt"`
}

type JobData struct {
	Job   Job      `json:"Job"`
	State JobState `json:"State"`
}
type Log struct {
	id        string
	timestamp string
	details   string
}
type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
type Update struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
type Data struct {
	Dealid   string `json:"dealid"`
	State    string `json:"state"`
	Metadata string `json:"metadata"`
}
type JobUpdate struct {
	ID         string `json:"id"`
	ModuleID   string `json:"module_id"`
	JobID      string `json:"job_id"`
	Status     string `json:"status"`
	TimeUpdate string `json:"time_update"`
	Details    string `json:"details"`
	TimeStart  string `json:"time_start"`
	// Message     string `json:"message"`
}
type DbJob struct {
	id          string
	module_id   string
	job_id      string
	status      string
	time_update string
	details     string
	time_start  string
}

type JobMetadata struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	ClientID  string    `json:"ClientID"`
	Requester struct {
		RequesterNodeID    string `json:"RequesterNodeID"`
		RequesterPublicKey string `json:"RequesterPublicKey"`
	} `json:"Requester"`
}

type JobSpec struct {
	Engine        string `json:"Engine"`
	Verifier      string `json:"Verifier"`
	Publisher     string `json:"Publisher"`
	PublisherSpec struct {
		Type string `json:"Type"`
	} `json:"PublisherSpec"`
	Docker struct {
		Image      string   `json:"Image"`
		Entrypoint []string `json:"Entrypoint"`
	} `json:"Docker"`
}

type JobResources struct {
	GPU string `json:"GPU"`
}

type JobNetwork struct {
	Type string `json:"Type"`
}

type JobDeal struct {
	Concurrency int `json:"Concurrency"`
}

type Job struct {
	APIVersion string       `json:"APIVersion"`
	Metadata   JobMetadata  `json:"Metadata"`
	Spec       JobSpec      `json:"Spec"`
	Resources  JobResources `json:"Resources"`
	Network    JobNetwork   `json:"Network"`
	Timeout    int          `json:"Timeout"`
	Deal       JobDeal      `json:"Deal"`
}

type JobExecution struct {
	JobID              string `json:"JobID"`
	NodeId             string `json:"NodeId"`
	ComputeReference   string `json:"ComputeReference"`
	State              string `json:"State"`
	AcceptedAskForBid  bool   `json:"AcceptedAskForBid"`
	VerificationResult struct {
		Complete bool `json:"Complete"`
		Result   bool `json:"Result"`
	} `json:"VerificationResult"`
	PublishedResults struct {
		StorageSource string `json:"StorageSource"`
		Name          string `json:"Name"`
		CID           string `json:"CID"`
	} `json:"PublishedResults"`
	RunOutput struct {
		Stdout          string `json:"stdout"`
		Stdouttruncated bool   `json:"stdouttruncated"`
		Stderr          string `json:"stderr"`
		Stderrtruncated bool   `json:"stderrtruncated"`
		ExitCode        int    `json:"exitCode"`
		RunnerError     string `json:"runnerError"`
	} `json:"RunOutput"`
	Version    int       `json:"Version"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

type JobState struct {
	JobID      string         `json:"JobID"`
	Executions []JobExecution `json:"Executions"`
	State      string         `json:"State"`
	Version    int            `json:"Version"`
	CreateTime time.Time      `json:"CreateTime"`
	UpdateTime time.Time      `json:"UpdateTime"`
	TimeoutAt  time.Time      `json:"TimeoutAt"`
}

type JobData struct {
	Job   Job      `json:"Job"`
	State JobState `json:"State"`
}
type Log struct {
	id        string
	timestamp string
	details   string
}
type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
type Update struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
type Data struct {
	Dealid   string `json:"dealid"`
	State    string `json:"state"`
	Metadata string `json:"metadata"`
}
type JobUpdate struct {
	ID         string `json:"id"`
	ModuleID   string `json:"module_id"`
	JobID      string `json:"job_id"`
	Status     string `json:"status"`
	TimeUpdate string `json:"time_update"`
	Details    string `json:"details"`
	TimeStart  string `json:"time_start"`
	// Message     string `json:"message"`
}
type DbJob struct {
	id          string
	module_id   string
	job_id      string
	status      string
	time_update string
	details     string
	time_start  string
}
