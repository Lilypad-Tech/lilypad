# Local development

## Prerequisites

- [Golang](https://go.dev/doc/install)
- [node](https://nodejs.org/en/download/package-manager)
- [Docker](https://docs.docker.com/engine/install/)

## Getting started

Running the Lilypad application locally depends on the `.local.dev` file for secrets injection. In this file are a series of private keys (with no funds on them) that are ONLY meant to be used for testing this app locally. You are free to replace these keys with your own if you wish; however, be warned that the `.local.dev` file is not included in the `.gitignore` so you must be vigilant to not commit this file in your PRs. We are not responsible for lost funds as a result of you posting your private keys on your commits/prs. So unless you have a very good reason to do so, leave the `.local.dev` file unchanged.

A minimal local Lilypad network consists of the following pieces of infrastructure:

- One blockchain node
- One IPFS node
- One bacalhau node
- One postgres database
- One solver service
- One job creator service
- One resource provider service

Order matters because the `solver`, `job creator` and `resource provider` will right away try to connect to the `blockchain node`. First, the `solver` will update the state of a known smart contract to publish a URL where other services can connect to it. Then, the `job creator`, and `resource provider` will fetch from the `blockchain` the URL for the `solver` and try to connect to it.

## Minimal local setup

### 1. Base services

The first step is to start the 3rd party services. This consists of a blockchain node (nitro), an IPFS node (kubo), a postgres database and a bacalhau node. All of the services can be started via docker compose (see: `./docker/docker-compose.base.yml`). The commands to start these services are:

```sh
./stack compose-init # This will clean the chain and boot it
./stack compose-services
```

A helper script is in place to verify balances on the accounts: `cd hardhat && npx hardhat run scripts/balances.ts --network dev`

### 2. Solver service

This process can be executed directly if Golang has been installed or in a docker container. The commands are `./stack solver`,`./stack solver-docker-build` and `./stack solver-docker-run` respectively. The `solver` service will output a log line that reads that "the solver has been registered successfully" or "the solver already exists". It is best to wait for this output before starting the services that will try to connect to the `solver`.

### 3. Job creator

This process can be executed directly if Golang has been installed or in a docker container. The commands are `./stack job-creator`,`./stack job-creator-docker-build` and `./stack job-creator-docker-run` respectively. The `job-creator` service's main function is to listen to events from the blockchain to execute jobs and when it receives such an event it will relay the payload to the `solver`. So think about the `job-creator` as the "on-chain solver".

### 4. Resource provider

For the time being this process has to be executed directly and needs Golang to be installed. This is the command to execute the service: `./stack resource-provider`. If you have a GPU you can use the following flag to use it: `./stack resource-provider --offer-gpu 1`

## Using Docker Compose

An alternative to the above for running the local stack is to use [Docker Compose](https://docs.docker.com/compose/) to run all of the services (including lilypad services contained in this repo).

Benefits of using Docker Compose include:

- Start/stop the full stack with a single command.
- Runs the docker images for all services (i.e. "closer to prod")

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

Run the Go unit tests with `./stack unit-tests` and the Hardhat unit tests with `./stack unit-tests-hardhat`.

Run the integration tests with `./stack integration-tests`. The integration tests expect all parts of the stack are running, except the request to run a job. See [Using Docker Compose](./LOCAL_DEVELOPMENT.md#using-docker-compose) to run the stack.

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
