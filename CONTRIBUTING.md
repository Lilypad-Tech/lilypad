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
# install the latest bacalhau which works with GPUs (https://github.com/bacalhau-project/bacalhau/issues/2858)
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.0.3/bacalhau_v1.0.3_linux_amd64.tar.gz
tar xfv bacalhau_v1.0.3_linux_amd64.tar.gz
mv bacalhau /usr/local/bin
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

Run the following commands in separate terminals:

```bash
./stack solver
```

```bash
./stack mediator
```

```bash
./stack resource-provider
```

```bash
./stack run cowsay
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
