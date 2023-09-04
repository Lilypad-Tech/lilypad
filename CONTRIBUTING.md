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

### geth

We need to start geth, move funds to our admin account and then fund the various other accounts.

```bash
./stack geth
./stack fund-admin
./stack fund-services
./stack balances
```

### compile contracts

```bash
./stack compile-contracts
```

### deploy contracts

```bash
./stack deploy-contracts
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