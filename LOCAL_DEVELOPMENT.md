# Local development

## Prerequisites

- [Golang](https://go.dev/doc/install)
- [node](https://nodejs.org/en/download/package-manager)
- [Docker](https://docs.docker.com/engine/install/)

## Getting started

Running the Lilypad application locally depends on the `.local.dev` file for secrets injection. In this file are a series of private keys (with no funds on them) that are ONLY meant to be used for testing this app locally. You are free to replace these keys with your own if you wish; however, be warned that the `.local.dev` file is not included in the `.gitignore` so you must be vigilant to not commit this file in your PRs. We are not responsible for lost funds as a result of you posting your private keys on your commits/prs. So unless you have a very good reason to do so, leave the `.local.dev` file unchanged.

A minimal local Lilypad network consists of the following pieces of infrastructure:

- One blockchain node
- One solver service
- One job creator service
- One IPFS node
- One bacalhau node
- One resource provider service

Order matters because the `solver`, `job creator` and `resource provider` will right away try to connect to the `blockchain node`. First, the `solver` will update the state of a known smart contract to publish a URL where other services can connect to it. Then, the `job creator`, and `resource provider` will fetch from the `blockchain` the URL for the `solver` and try to connect to it.

## Minimal local setup

### 1. Blockchain node

The node can can run directly from an existing docker image. It will initialize itself to a blank state with the admin address holding the funds to deploy the smart contracts (including the Lilypad token) and fund the accounts used by the different services. The blockchain is ephemeral, meaning every time you restart the node the state will be reset, you can work around this by keeping the node active if needed.

These are the commands to run the node and boot the network: `./stack chain-clean` (the first time this won't do anything, but I find it better to get in the habit of resetting artifacts everytime), `./stack chain` to run the node and `./stack chain-boot` to fund the accounts with ETH (for gas fees), compile the contracts, add Golang bindings to use the contracts directly in go code, deploy the contracts and fund the accounts with Lilypad tokens.

#### Summary of command sequence

```sh
./stack chain-clean
./stack chain
# then in another terminal
./stack chain-boot
```

A helper script is in place to verify balances on the accounts: `cd hardhat && npx hardhat run scripts/balances.ts --network dev`

### 2. Solver service

This process can be executed directly if Golang has been installed or in a docker container. The commands are `./stack solver`,`./stack solver-docker-build` and `./stack solver-docker-run` respectively. The `solver` service will output a log line that reads that "the solver has been registered successfully" or "the solver already exists". It is best to wait for this output before starting the services that will try to connect to the `solver`.

### 3. Job creator

This process can be executed directly if Golang has been installed or in a docker container. The commands are `./stack job-creator`,`./stack job-creator-docker-build` and `./stack job-creator-docker-run` respectively. The `job-creator` service's main function is to listen to events from the blockchain to execute jobs and when it receives such an event it will relay the payload to the `solver`. So think about the `job-creator` as the "on-chain solver".

### 4. IPFS node

This process can be run directly using an IPFS Kubo binary. [Download a Kubo binary](https://dist.ipfs.tech/#kubo) for your platform and architecture. Initiliaze an repository where Kubo will store settings and internal data:

```sh
ipfs init
```

Start the IPFS node with the default settings:

```sh
ipfs daemon
```

Kubo should report an RPC API server listening on `/ip4/127.0.0.1/tcp/5001`. Our Bacalhau node stores job results on this IPFS node.

### 5. Bacalhau node

For the time being this process has to be executed directly. This means following the instructions to download their cli tool and expose it as a bin that can be used. Here's how to install the `bacalhau` tool:

#### Linux
```sh
# install the latest
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_linux_amd64.tar.gz
# extract the downloaded archive and move the `bacalhau` binary to `/usr/local/bin`
tar xfv bacalhau_v1.3.2_linux_amd64.tar.gz
mv bacalhau /usr/local/bin
```
#### Mac OS
```sh
# install the latest
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_darwin_amd64.tar.gz
# extract the downloaded archive and move the `bacalhau` binary to `/usr/local/bin`
tar xfv bacalhau_v1.3.2_darwin_amd64.tar.gz
mv bacalhau /usr/local/bin
```

Once the tool has been installed, the following command can be used to start the node: `./stack bacalhau-node`.

### 6. Resource provider

For the time being this process has to be executed directly and needs Golang to be installed. This is the command to execute the service: `./stack resource-provider`. If you have a GPU you can use the following flag to use it: `./stack resource-provider --offer-gpu 1`

#### Bacalhau and Resource provider as one component

There is ongoing work to pack together the `bacalhau` node and `resource provider` service in a docker container as these two are highly coupled and can be abstracted into one component.

## Using Docker Compose

An alternative to the above for running the local stack is to use [Docker Compose](https://docs.docker.com/compose/).

Benefits of using Docker Compose include:

- Start/stop the full stack with a single command.
- Runs the docker images for all services (i.e. "closer to prod")
- Chain and Bacalhau state are maintained between runs (in the `./data` directory).

The main drawback is, for development, you'll need to re-build the images after changes for testing.

All of the docker commands have been wrapped by `./stack` - to simplify doppler configuration, etc.

### Running the stack

_First run_: run `./stack compose-init`. This essentially runs `./stack chain-clean` and `./stack chain-boot`.

Run `./stack compose-up` to start the stack.

### Re-building images

The first time you run docker compose, it will pull / build the images for all services. If you're making code changes, you'll want to re-build the docker images with your local changes. This can be done with `./stack compose-build`.

### Stopping / shutting down

Run `./stack compose-down`.

## Running a job

Once all the services are up and running this command can be used to trigger an on-chain job: `./stack run-cowsay-onchain`

### Tests

There are two commands that can be used to run existing tests: `./stack unit-tests` and `./stack integration-tests` (bear in mind the latter expects a blockchain node to be running locally).

## Notes on tooling

Things should work right out-of-the-box, no extra configuration should be needed as Doppler provides the environment variables that are required with the current setup.

## Troubleshooting

In this section we'll address some common problems you might face when trying to boot up Lilypad locally

### Chain-boot Related issues

If you try and run `./stack compose-init` or `./stack chain-boot` and get the following error
```bash
ProviderError: failed with 51333200 gas: insufficient funds for gas * price + value: address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 have 9318991353400000000 want 10000000000000000000
    at HttpProvider.request (/Users/nshahnazarian/Development/git/lilypad/hardhat/node_modules/hardhat/src/internal/core/providers/http.ts:88:21)
    at processTicksAndRejections (node:internal/process/task_queues:95:5)
    at async HardhatEthersProvider.estimateGas (/Users/nshahnazarian/Development/git/lilypad/hardhat/node_modules/@nomicfoundation/hardhat-ethers/src/internal/hardhat-ethers-provider.ts:237:27)
    at async Wallet.populateTransaction (/Users/nshahnazarian/Development/git/lilypad/hardhat/node_modules/ethers/src.ts/providers/abstract-signer.ts:105:28)
    at async Wallet.sendTransaction (/Users/nshahnazarian/Development/git/lilypad/hardhat/node_modules/ethers/src.ts/providers/abstract-signer.ts:232:21)
    at async transferEther (/Users/nshahnazarian/Development/git/lilypad/hardhat/utils/web3.ts:61:14)
    at async /Users/nshahnazarian/Development/git/lilypad/hardhat/scripts/fund-services-ether.ts:15:5
```
This can be addressed by doing the following:
- Open your Docker Desktop app, go to `Volumes` and delete `lilypad_chain-data` as there might be stale data in the volume not allowing you to properly execute all the transactions `chain-boot` executes

### Issues running onchain cowsay

If you find that you have issues with the Job Creator not picking up your `run-cowsay-onchain` command while running the Lilypad stack through Docker, do the following:
1. Stop the Docker stack by pressing ctrl+c
2. Run the following command to clean up your Docker environment: `./stack compose-down && docker system prune -a`
3. Open your Docker Desktop app, go to `Volumes` and delete `lilypad_chain-data` as there might be stale data in the volume not allowing you to properly execute all the transactions
4. Re-run your Docker stack using: `./stack compose-build && ./stack compose-init && ./stack compose-up`
5. Re-run the onchain cowsay job: `./stack run-cowsay-onchain`
