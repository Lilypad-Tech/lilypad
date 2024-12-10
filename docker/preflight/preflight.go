package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pelletier/go-toml"
)

const ERC20TokenABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"}]`

type Module struct {
	ModuleId string `json:"ModuleId"`
	Image    string `json:"Image,omitempty"`
}

type GPUStats struct {
	Name        string
	GUID        string
	Utilization int
	MemoryUsed  int
	TotalMemory int
}

func pullAllowList() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a GitHub organization with module-allowlist")
		fmt.Println("Usage: preflight <github_organization>")
		fmt.Println("Example: preflight lilypad-tech")
		fmt.Println("You may fork https://github.com/Lilypad-Tech/module-allowlist and use your own organization")
		os.Exit(1)
	}
	allLilypadURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/module-allowlist/main/allowlist.json", os.Args[1])
	resp, err := http.Get(allLilypadURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch allowlist. specify github organization with module-allowlist")
		fmt.Println("Usage: preflight <github_organization>")
		fmt.Println("Example: preflight lilypad-tech")
		fmt.Println("You may fork https://github.com/Lilypad-Tech/module-allowlist and use your own organization")
		log.Fatalf("Failed to fetch allowlist: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var modules []Module
	if err := json.Unmarshal(body, &modules); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	for i, module := range modules {
		if strings.HasPrefix(module.ModuleId, "http") {
			moduleUrl := strings.Replace(module.ModuleId, "github.com", "raw.githubusercontent.com", 1) + "/main/lilypad_module.json.tmpl"
			respModule, err := http.Get(moduleUrl)
			if err != nil {
				log.Printf("Failed to fetch module: %v", err)
				continue
			}
			defer respModule.Body.Close()

			bodyModule, err := ioutil.ReadAll(respModule.Body)
			if err != nil {
				log.Printf("Failed to read module response body: %v", err)
				continue
			}

			re := regexp.MustCompile(`"Image":\s*"([^"]+)"`)
			matches := re.FindStringSubmatch(string(bodyModule))
			if len(matches) > 1 {
				modules[i].Image = matches[1]
			} else {
				log.Printf("Failed to find image in module: %s", moduleUrl)
			}
		}
	}

	for _, module := range modules {
		if module.Image != "" {
			fmt.Println("Module ModuleId:", module.ModuleId)
			fmt.Println("Module Image:", module.Image)
			cmd := exec.Command("docker", "pull", module.Image)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				log.Fatalf("Failed to get stdout: %v", err)
			}
			stderr, err := cmd.StderrPipe()
			if err != nil {
				log.Fatalf("Failed to get stderr: %v", err)
			}

			if err := cmd.Start(); err != nil {
				log.Fatalf("Failed to start command: %v", err)
			}
			go func() {
				reader := bufio.NewReader(stdout) // Change to stderr for progress
				var output string
				for {
					char, err := reader.ReadByte()
					if err != nil {
						break
					}

					// Print raw byte for real-time display
					fmt.Print(string(char))

					output += string(char)
					// Handle both CR and NL as line endings
					if char == '\r' || char == '\n' {
						// Strip ANSI escape codes for parsing
						cleanOutput := stripAnsi(output)
						printProgress(cleanOutput)
						output = ""
					}
				}
			}()

			// Helper function to strip ANSI escape codes

			go func() {
				scanner := bufio.NewScanner(stderr)
				for scanner.Scan() {
					line := scanner.Text()
					fmt.Println(line)
					printProgress(line)
				}
			}()

			if err := cmd.Wait(); err != nil {
				log.Fatalf("Failed to pull Docker image %s: %v", module.Image, err)
			} else {
				log.Printf("Successfully pulled Docker image: %s\n", module.Image)
			}
		}

	}

}
func stripAnsi(str string) string {
	ansi := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return ansi.ReplaceAllString(str, "")
}
func printProgress(line string) {
	// Example line: "Downloading [==================> ]  50.32MB/100.64MB"
	re := regexp.MustCompile(`Downloading.*?(\d+\.\d+)(MB|GB)/(\d+\.\d+)(MB|GB)`)
	matches := re.FindStringSubmatch(line)

	if matches == nil {
		return
	}
	fmt.Print(len(matches))
	if len(matches) > 1 {
		current, err1 := strconv.ParseFloat(matches[1], 64)
		total, err2 := strconv.ParseFloat(matches[3], 64)
		unit := matches[2]

		if err1 == nil && err2 == nil {
			if unit == "GB" {
				current *= 1024
				total *= 1024
			}

			fmt.Printf("Progress: %.2fMB/%.2fMB\n", current, total)
		}
	}
}

func getGPUStats() ([]GPUStats, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,gpu_uuid,utilization.gpu,memory.used,memory.total", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var stats []GPUStats
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Split(line, ", ")
		utilization, _ := strconv.Atoi(fields[2])
		memoryUsed, _ := strconv.Atoi(fields[3])
		totalMemory, _ := strconv.Atoi(fields[4])
		stats = append(stats, GPUStats{
			Name:        fields[0],
			GUID:        fields[1],
			Utilization: utilization,
			MemoryUsed:  memoryUsed,
			TotalMemory: totalMemory,
		})
	}
	return stats, nil
}

func getMemoryUtilization() int {
	stats, err := getGPUStats()
	if err != nil {
		log.Fatalf("Error getting GPU stats: %v\n", err)
		// return 100
		// log.Fatalf("Failed to start command
	}

	for _, stat := range stats {
		memoryUtilization := float64(stat.MemoryUsed) / float64(stat.TotalMemory) * 100
		fmt.Printf("GPU %s (%s): Memory Utilization: %.2f%%\n",
			stat.Name, stat.GUID, memoryUtilization)
		if int(memoryUtilization) > 50 {
			log.Fatalf("Memory Utilization is greater than 50%")
		} else {
			return int(memoryUtilization)
		}
	}
	log.Fatalf("Np GPU stats found")
	return 100
}

func getCoreUtilization() int {
	stats, err := getGPUStats()
	if err != nil {
		log.Fatalf("Error getting GPU stats: %v\n", err)
		return 100
	}

	for _, stat := range stats {
		fmt.Printf("GPU %s (%s): Core Utilization: %d%%\n",
			stat.Name, stat.GUID, stat.Utilization)
		// return stat.Utilization
		if int(stat.Utilization) > 50 {
			log.Fatalf("Core Utilization is greater than 50%")
		} else {
			return int(stat.Utilization)
		}
	}
	log.Fatalf("Np GPU stats found")
	return 100
}

func getBalances() {
	privateKeyHex := os.Getenv("WEB3_PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("PRIVATE_KEY environment variable not set")
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Printf("Public Address: %s\n", walletAddress.Hex())

	// client, err := ethclient.Dial("wss://rpc.ankr.com/arbitrum_sepolia")
	client, err := ethclient.Dial(os.Getenv("WEB3_RPC_URL"))

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	balance, err := client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve the balance: %v", err)
	}

	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	fmt.Printf("ETH Balance: %s\n", ethBalance.Text('f', 18))
	if ethBalance.Cmp(big.NewFloat(.01)) < 0 {
		log.Fatalf("Insufficient ETH balance: minimum 0.01 ETH required")
	}
	tokenAddress := common.HexToAddress(os.Getenv("WEB3_TOKEN_ADDRESS")) //"0x0352485f8a3cB6d305875FaC0C40ef01e0C06535")

	parsedABI, err := abi.JSON(strings.NewReader(ERC20TokenABI))
	if err != nil {
		log.Fatalf("Failed to parse token ABI: %v", err)
	}
	token := bind.NewBoundContract(tokenAddress, parsedABI, client, client, client)

	var tokenBalance = new(big.Int)
	var result []interface{}
	err = token.Call(&bind.CallOpts{}, &result, "balanceOf", walletAddress)
	if err == nil && len(result) > 0 {
		tokenBalance = result[0].(*big.Int)
	}
	if err != nil {
		log.Fatalf("Failed to retrieve the token balance: %v", err)
	}

	tokenBalanceFloat := new(big.Float).Quo(new(big.Float).SetInt(tokenBalance), big.NewFloat(1e18))

	fmt.Printf("ERC-20 Token Balance: %s\n", tokenBalanceFloat.Text('f', 18))
	if tokenBalanceFloat.Cmp(big.NewFloat(10)) < 0 {
		log.Fatalf("LP Token balance is less than 10")
	}
}
func updateBashrc(privateKey string) error {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	// Open .bashrc in append mode
	f, err := os.OpenFile(home+"/.bashrc", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .bashrc: %v", err)
	}
	defer f.Close()

	// Write export line
	exportLine := fmt.Sprintf("\nexport WEB3_PRIVATE_KEY=\"%s\"\n", privateKey)
	if _, err := f.WriteString(exportLine); err != nil {
		return fmt.Errorf("failed to write to .bashrc: %v", err)
	}

	// Source .bashrc
	cmd := exec.Command("bash", "-c", "source ~/.bashrc")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to source .bashrc: %v", err)
	}

	return nil
}
func readPrivKeyFromBashrc() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	content, err := os.ReadFile(home + "/.bashrc")
	if err != nil {
		log.Fatalf("Failed to read .bashrc: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "export WEB3_PRIVATE_KEY=") {
			value := strings.TrimPrefix(line, "export WEB3_PRIVATE_KEY=")
			// Remove quotes if present
			return strings.Trim(value, "\"'")
		}
	}
	return ""
}
func main() {
	fmt.Println("Lilypad Pre-Flight Check")
	// cmd := exec.Command("bash", "-c", "source ~/.bashrc")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalf("Failed to source .bashrc: %v", err)
	// }
	privKey := readPrivKeyFromBashrc()
	if privKey == "" {
		privKey = os.Getenv("WEB3_PRIVATE_KEY")
	} else {
		os.Setenv("WEB3_PRIVATE_KEY", privKey)
		fmt.Println("WEB3_PRIVATE_KEY found in ~/.bashrc")
	}
	fmt.Println("Private Key:", privKey)
	if privKey == "" {
		// Generate new private key
		key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			log.Fatalf("Failed to generate private key: %v", err)
		}

		// Convert to hex string without 0x prefix
		privKey = hex.EncodeToString(key.D.Bytes())

		// Set env var
		os.Setenv("WEB3_PRIVATE_KEY", privKey)

		err = updateBashrc(privKey)
		if err != nil {
			log.Fatalf("Failed to update bashrc: %v", err)
		}
		// Write to .env file
		f, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open .env file: %v", err)
		}
		defer f.Close()

		if _, err := f.WriteString(fmt.Sprintf("WEB3_PRIVATE_KEY=%s\n", privKey)); err != nil {
			log.Fatalf("Failed to write to .env file: %v", err)
		}
	}
	repo := "lilypad_network"
	if len(os.Args) > 1 {
		repo = os.Args[len(os.Args)-1]
	}
	if len(os.Args) > 1 && os.Args[1] == "docker" {
		cmd := exec.Command("docker-compose",
			"-f", "docker-compose-pre-flight.yml",
			"-p", os.Args[2]+"-"+os.Args[3],
			"up",
			"--build")
		// ,
		// "--build")
		cmd.Env = append(os.Environ(), fmt.Sprintf("ORG=%s", os.Args[2]))

		cmd.Env = append(cmd.Env, fmt.Sprintf("ORG=%s", repo))
		//https://raw.githubusercontent.com/arsenum/lilypad_network/refs/heads/main/config.toml
		// Parse TOML
		envVars := getVarsFromRepo(repo)
		cmd.Env = append(cmd.Env, envVars...)
		for _, envVar := range cmd.Env {
			fmt.Println("ENV Var:", envVar)

		}

		// if err != nil {
		// 	fmt.Errorf("failed to parse TOML: %v", err)
		// }

		// // Convert TOML to env vars and append
		// for _, k := range tree.Keys() {
		// 	value := tree.Get(k)
		// 	strValue := tomlValueToString(value)
		// 	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, strValue))
		// }

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to run docker-compose: %v", err)
		}
		return
	} else {
		envVars := getVarsFromRepo(repo)
		for _, envVar := range envVars {
			fmt.Println("ENV Var:", envVar)
			os.Setenv(strings.Split(envVar, "=")[0], strings.Split(envVar, "=")[1])

		}
		getMemoryUtilization()
		getCoreUtilization()
		getBalances()
		// pullAllowList()
	}
}

func getVarsFromRepo(repo string) []string {
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/refs/heads/main/config.toml", os.Args[len(os.Args)-2], repo))

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Errorf("failed to download TOML: %v", err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("failed to read TOML content: %v", err)
		}

		tree, err := toml.LoadBytes(content)
		envVars := strings.Split(tomlValueToString(tree, ""), "\n")
		return envVars
	}
	return nil
}
func tomlValueToString(value interface{}, parentKey string) string {
	switch v := value.(type) {
	case *toml.Tree:
		var envVars []string
		for _, k := range v.Keys() {
			key := k
			if parentKey != "" {
				key = parentKey + "_" + k
			}
			value := v.Get(k)
			if subTree, ok := value.(*toml.Tree); ok {
				envVars = append(envVars, tomlValueToString(subTree, strings.ToUpper(key)))
			} else {
				envVars = append(envVars, fmt.Sprintf("%s=%v", strings.ToUpper(key), value))
			}
		}
		return strings.Join(envVars, "\n")
	default:
		if parentKey != "" {
			return fmt.Sprintf("%s=%v", strings.ToUpper(parentKey), v)
		}
		return fmt.Sprintf("%v", v)
	}
}

// Flags:
//       --api-host string                            The api host to connect to (API_HOST)
//       --bacalhau-api-host string                   The api hostname for the bacalhau cluster to run jobs (default "localhost")
//       --bacalhau-api-port string                   The api port for the bacalhau cluster to run jobs (default "1234")
//       --cuda-block-size int                        Cuda block size (CUDA_BLOCK_SIZE) (default 1024)
//       --cuda-grid-size int                         Cuda grid size (sm*2) (CUDA_GRID_SIZE) (default 128)
//       --cuda-hash-per-thread int                   Cuda hash per threads (CUDA_HASH_PER_THREAD) (default 1000)
//       --disable-pow                                Disable pow mining (DISABLE_POW)
//       --disable-telemetry                          Disable telemetry (DISABLE_TELEMETRY)
//   -h, --help                                       help for resource-provider
//       --ipfs-connect string                        The IPFS multiaddress to connect to (IPFS_CONNECT)
//       --num-worker int                             Pow worker number (NUM_WORKER)
//       --offer-count int                            How many machines will we offer using the cpu, ram and gpu settings (OFFER_COUNT). (default 1)
//       --offer-cpu int                              How many milli-cpus to offer the network (OFFER_CPU). (default 1000)
//       --offer-gpu int                              How many milli-gpus to offer the network (OFFER_GPU).
//       --offer-modules stringArray                  The modules you are willing to run (OFFER_MODULES).
//       --offer-ram int                              How many megabytes of RAM to offer the network (OFFER_RAM). (default 1024)
//       --pricing-instruction-price uint             The price per instruction to offer (PRICING_INSTRUCTION_PRICE) (default 1)
//       --pricing-mediation-fee uint                 The mediation fee (PRICING_MEDIATION_FEE) (default 1)
//       --pricing-mode string                        set pricing mode (MarketPrice/FixedPrice) (default "FixedPrice")
//       --pricing-payment-collateral uint            The payment collateral (PRICING_PAYMENT_COLLATERAL) (default 2)
//       --pricing-results-collateral-multiple uint   The results collateral multiple (PRICING_RESULTS_COLLATERAL_MULTIPLE) (default 2)
//       --service-mediators strings                  The mediators we trust (SERVICE_MEDIATORS)
//       --service-solver string                      The solver to connect to (SERVICE_SOLVER)
//       --telemetry-token string                     The token to auth with telemetry service (TELEMETRY_TOKEN)
//       --telemetry-url string                       The telemetry endpoint to connect to (TELEMETRY_URL)
//       --timeout-agree-collateral uint              The collateral to timeout a deal when agreeing (TIMEOUT_AGREE_COLLATERAL) (default 1)
//       --timeout-agree-time uint                    The number of seconds to timeout a deal when agreeing (TIMEOUT_AGREE_TIME) (default 3600)
//       --timeout-judge-results-collateral uint      The collateral to timeout a deal when judging results (TIMEOUT_JUDGE_RESULTS_COLLATERAL) (default 1)
//       --timeout-judge-results-time uint            The number of seconds to timeout a deal when judging results (TIMEOUT_JUDGE_RESULTS_TIME) (default 3600)
//       --timeout-mediate-results-collateral uint    The collateral to timeout a deal when mediating results (TIMEOUT_MEDIATE_RESULTS_COLLATERAL) (default 1)
//       --timeout-mediate-results-time uint          The number of seconds to timeout a deal when mediating results (TIMEOUT_MEDIATE_RESULTS_TIME) (default 3600)
//       --timeout-submit-results-collateral uint     The collateral to timeout a deal when submitting results (TIMEOUT_SUBMIT_RESULTS_COLLATERAL) (default 1)
//       --timeout-submit-results-time uint           The number of seconds to timeout a deal when submitting results (TIMEOUT_SUBMIT_RESULTS_TIME) (default 3600)
//       --web3-chain-id int                          The chain id for the web3 RPC server (WEB3_CHAIN_ID).
//       --web3-controller-address string             The address of the controller contract (WEB3_CONTROLLER_ADDRESS).
//       --web3-payments-address string               The address of the payments contract (WEB3_PAYMENTS_ADDRESS).
//       --web3-pow-address string                    The address of the pow contract (WEB3_POW_ADDRESS).
//       --web3-private-key string                    The private key to use for signing web3 transactions (WEB3_PRIVATE_KEY).
//       --web3-rpc-url string                        The URL of the web3 RPC server (WEB3_RPC_URL).
//       --web3-storage-address string                The address of the storage contract (WEB3_STORAGE_ADDRESS).
//       --web3-token-address string                  The address of the token contract (WEB3_TOKEN_ADDRESS).
//       --web3-users-address string                  The address of the users contract (WEB3_USERS_ADDRESS).
