# Development

We need the following installed:

 * docker
 * node.js
 * golang

## install

Get the node modules installed and generate a local .env file with private keys for the various services.

```bash
(cd hardhat && yarn install)
./stack print-env > .env
```

## boot stack

### bacalhau

We need a bacalhau node running on the same machine as the resource provider.

Here is how we got bacalhau up and running:

```bash
# install the latest bacalhau
curl -sL https://get.bacalhau.org/install.sh | bash
# configure this to where you want the ipfs data to be stored
export BACALHAU_SERVE_IPFS_PATH=/tmp/lilypad/data/ipfs
# run bacalhau as both compute node and requester node
bacalhau serve --node-type compute,requester --peer none --private-internal-ipfs=false --job-selection-accept-networked
```

### lilypad

```bash
./stack boot
```

This does the following things:

#### geth

We need to start geth, move funds to our admin account and then fund the various other accounts.

```bash
./stack geth
./stack fund-admin
./stack fund-services-ether
./stack balances
```

Geth is now running and each of our services has some ether allocated.

#### compile contracts

```bash
./stack compile-contracts
```

This will also generate go bindings in `pkg/contract/bindings/contracts`

#### deploy contracts

```bash
./stack deploy-contracts
```

Contracts are now deployed - no services have any tokens though

#### fund tokens

```bash
./stack fund-services-tokens
./stack balances
```

### run services

First we need to export the env that will configure the web3 client with the contract addresses:

```bash
source .env
eval $(./stack print-local-dev-env)
```

Then we export the private key for the service we want to run:

```bash
export WEB3_PRIVATE_KEY=$SOLVER_PRIVATE_KEY
```

Then we run the service:

```bash
go run . solver
```


## stop stack

To stop geth running at any time:

```bash
./stack geth-stop
```

To reset the geth data (i.e. a complete restart):

```bash
./stack clean
```

NOTE: you will need to re-run the `fund-admin` and `fund-services` commands after a `clean`.

## unit tests

Run the smart contract unit tests:

```bash
./stack unit-tests
```

## re-generate go bindings

When you change the smart contracts the go bindings in `pkg/contract/bindings/contracts` will need regenerating.

```bash
./stack compile-contracts
```