# Local development

## Prerequisites

- [Golang](https://go.dev/doc/install)
- [node](https://nodejs.org/en/download/package-manager)
- [Docker](https://docs.docker.com/engine/install/)
- [Doppler](https://docs.doppler.com/docs/install-cli)

## Getting started

A minimal local Lilypad network consists of the following pieces of infrastructure:

- One blockchain node
- One solver service
- One job creator service
- One bacalhau node
- One resource provider service

Order matters because the `solver`, `job creator` and `resource provider` will right away try to connect to the `blockchain node`. First, the `solver` will update the state of a known smart contract to publish a URL where other services can connect to it. Then, the `job creator`, and `resource provider` will fetch from the `blockchain` the URL for the `solver` and try to connect to it.

## Minimal local setup

### 1. Blockchain node

The node can can run directly from an existing docker image, it will initialize itself to a blank state with the admin address holding the funds to deploy the smart contracts (including the Lilypad token) and fund the accounts used by the different services. The blockchain is ephemeral, meaning every time you restart the node the state will be reset, you can work around this by keeping the node active if needed.

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

### 4. Bacalhau node

For the time being this process has to be executed directly. This means following the instructions to download their cli tool and expose it as a bin that can be used. Here's how to install the `bacalhau` tool:

```sh
# install the latest
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_linux_amd64.tar.gz
# extract the downloaded archive and move the `bacalhau` binary to `/usr/local/bin`
tar xfv bacalhau_v1.3.2_linux_amd64.tar.gz
mv bacalhau /usr/local/bin
```

Once the tool has been installed, the following command can be used to start the node: `./stack bacalhau-node`.

### 5. Resource provider

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
