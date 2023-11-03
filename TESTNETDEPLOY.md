# Testnet Deployment


## Installation

First, install the required Node.js modules and generate a local `.env` file containing private keys for various services. Run the following commands:

```bash
(cd hardhat && yarn install)
./stack print-env > .env
```

## Booting the Stack

### 1 - Bacalhau

To run a Bacalhau node on the same machine as the resource provider, follow these steps:

```bash
# install the latest bacalhau which works with GPUs (https://github.com/bacalhau-project/bacalhau/issues/2858)
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.0.3/bacalhau_v1.0.3_linux_amd64.tar.gz
tar xfv bacalhau_v1.0.3_linux_amd64.tar.gz
mv bacalhau /usr/local/bin
# configure this to where you want the ipfs data to be stored
export BACALHAU_SERVE_IPFS_PATH=/tmp/lilypad/data/ipfs
# run bacalhau as both compute node and requester node
./stack bacalhau-serve
```

## Create Seven New Accounts

Follow the `README.md` in the `generate_accts` directory to create seven new accounts.

Copy `hardhat/.env.sample` to `.env` and update the following environment variables:
```
ADDRESS_ADMIN=
PRIVATE_KEY_ADMIN=
ADDRESS_FAUCET=
PRIVATE_KEY_FAUCET=
ADDRESS_SOLVER=
PRIVATE_KEY_SOLVER=
ADDRESS_MEDIATOR=
PRIVATE_KEY_MEDIATOR=
ADDRESS_RESOURCE_PROVIDER=
PRIVATE_KEY_RESOURCE_PROVIDER=
ADDRESS_JOB_CREATOR=
PRIVATE_KEY_JOB_CREATOR=
ADDRESS_DIRECTORY=
PRIVATE_KEY_DIRECTORY=
```

## Create a new Infura Project

Create a new Infura project and update the following environment variable in `hardhat/.env`:
```
INFURA_KEY=
```

Also add the infura key to the `.env` file:
```
export INFURA_KEY=
```

## Setup Hardhat

Add NETWORK to the `hardhat/.env` file:
```
NETWORK=sepolia
```

Update the following values in the `.env` file. Replace `<INFURA_KEY>` with the Infura key from above:
```
export NETWORK=sepolia
export WEB3_RPC_URL=wss://sepolia.infura.io/ws/v3/<INFURA_KEY>
export WEB3_CHAIN_ID=11155111
```

## Fund the Seven New Accounts

Fund the `admin` acccount with .7 ETH.

Fund the remaining six accounts with .1 ETH each.

```bash
./stack fund-services-ether
```

Check the balances

```bash
./stack balances
```

## Compile Contracts

```bash
./stack compile-contracts
```

## Deploy Contracts

```bash
./stack deploy-contracts
```

## Fund Services Tokens

```bash
./stack fund-services-tokens
```

### Run Services

Run the following commands in separate terminals:

```bash
./stack solver
```

Wait for the solver to start when `🟡 SOL solver registered` is logged, and then run:

```bash
./stack mediator
```

If you have a GPU, run the following command in a separate terminal window:

```bash
./stack resource-provider --offer-gpu 1
```

Otherwise, if you don't have a GPU:

```bash
./stack resource-provider
```

Run Cowsay:

```bash
./stack run cowsay:v0.0.1 -i Message="moo"
```

Run SDXL:

```bash
./stack runsdxl sdxl:v0.9-lilypad1 PROMPT="beautiful view of iceland with a record player"
```

### 4 - Run Cowsay On-Chain

Start the on-chain Job Creator:

```bash
./stack jobcreator
```

```bash
./stack run-cowsay-onchain