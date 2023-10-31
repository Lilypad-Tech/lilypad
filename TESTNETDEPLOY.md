# Testnet Deployment

### bacalhau

We need a bacalhau node running on the same machine as the resource provider.

Here is how we got bacalhau up and running:

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

Follow the README.md in the `generate_accts` directory to create seven new accounts.

Copy hardhat/.env.sample to .env and update the following environment variables:
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

set defaultNetwork to `sepolia` in `hardhat.config.js`

update stack script to use `sepolia` network
export NETWORK=${NETWORK:="sepolia"}

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

## Does this need to happen a second time to get the contract/bindings?

```bash
./stack go-bindings
```

## Fund Services Tokens

```bash
./stack fund-services-tokens
```

## Environment Variables

update `stack` with new RPC URL in print-local-dev-env():
```bash
echo "export WEB3_RPC_URL=wss://sepolia.infura.io/ws/v3/$INFURA_KEY" 
```

### Run Services

Run the following commands in separate terminals:

```bash
./stack solver --web3-chain-id 11155111 --web3-payments-address 0x01B18F94B61253ba63b810ddA371eA54bbACbdC6 --web3-storage-address 0xC5a58D6BDbdB66c50ecD795C5456E1f6ADc52dD9 --web3-users-address 0x11F3f6e51B822c4f0FF8955510f81B6654a9BD0C
```

```bash
./stack mediator --web3-chain-id 11155111 --web3-payments-address 0x01B18F94B61253ba63b810ddA371eA54bbACbdC6 --web3-storage-address 0xC5a58D6BDbdB66c50ecD795C5456E1f6ADc52dD9 --web3-users-address 0x11F3f6e51B822c4f0FF8955510f81B6654a9BD0C
```

```bash
./stack resource-provider --web3-chain-id 11155111 --web3-payments-address 0x01B18F94B61253ba63b810ddA371eA54bbACbdC6 --web3-storage-address 0xC5a58D6BDbdB66c50ecD795C5456E1f6ADc52dD9 --web3-users-address 0x11F3f6e51B822c4f0FF8955510f81B6654a9BD0C
```

```bash
./stack run cowsay:v0.0.1 -i Message="moo"
```