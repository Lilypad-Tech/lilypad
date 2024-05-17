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

This has been implemented in a docker container and two functions are already in place to build and run the container: `./stack chain-docker-build` and `./stack chain-docker-run`. The run command accepts an optional parameter for a local path where to store the blockchain data. This is handy to keep the current state of the blockchain around even after restarts on the container. The first time this commands are executed the blockchain will be in its genesis state, several things need to happen:

- Fund the admin account, which will be used to deploy the smart contracts and token.
- Compile and deploy the smart contracts.
- Create and fund the accounts with the native gas token and the Lilypad token.

To do the first thing run the command `./stack chain-fund-admin`. Then, the remaining processes can be executed with the command `./stack chain-boot`. This last command executes via hardhat, so it will need the URL of the blockchain node where it will send the transactions (this is currently hardcoded to `http://localhost:8546`, which is ok for this setup but keep in mind in case you decide to configure things differently or if you want to run this command on a deployed blockchain node).

### 2. Solver service

This process can be executed directly if Golang has been installed or in a docker container. The commands are `./stack solver`,`./stack solver-docker-build` and `./stack solver-docker-run` respectively. The `solver` service will output a log line that reads that "the solver has been registered successfully" or "the solver already exists". It is best to wait for this output before starting the services that will try to connect to the `solver`.

### 3. Job creator

This process can be executed directly if Golang has been installed or in a docker container. The commands are `./stack job-creator`,`./stack job-creator-docker-build` and `./stack job-creator-docker-run` respectively. The `job-creator` service's main function is to listen to events from the blockchain to execute jobs and when it receives such an event it will relay the payload to the `solver`. So think about the `job-creator` as the "on-chain solver".

### 4. Bacalhau node

For the time being this process has to be executed directly. This means following the instructions to download their cli tool and expose it as a bin that can be used. Here's how to install the `bacalhau` tool:

```sh
# install the latest (at the time of this write-up it was 1.3.0)
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.0/bacalhau_v1.3.0_linux_amd64.tar.gz
# extract the downloaded archive and move the `bacalhau` binary to `/usr/local/bin`
tar xfv bacalhau_v1.3.0_linux_amd64.tar.gz
mv bacalhau /usr/local/bin
```

Once the tool has been installed, the following command can be used to start the node: `./stack bacalhau-node`.

### 5. Resource provider

For the time being this process has to be executed directly and needs Golang to be installed. This is the command to execute the service: `./stack resource-provider`. If you have a GPU you can use the following flag to use it: `./stack resource-provider --offer-gpu 1`

#### Bacalhau and Resource provider as one component

There is ongoing work to pack together the `bacalhau` node and `resource provider` service in a docker container as these two are highly coupled and can be abstracted into one component.

## Running a job

Once all the services are up and running this command can be used to trigger an on-chain job: `./stack run-cowsay-onchain`

### Tests

There are two commands that can be used to run existing tests: `./stack unit-tests` and `./stack integration-tests` (bear in mind the latter expects a blockchain node to be running locally).

## Notes on tooling

Things should work right out-of-the-box, no extra configuration should be needed as Doppler provides the environment variables that are required with the current setup.
