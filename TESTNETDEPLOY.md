# Testnet Deployment

## Create Seven New Accounts

Follow the README.md in the `generate_accts` directory to create seven new accounts.

Update the following environment variables in `.env`:
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

Create a new Infura project and update the following environment variable in `.env`:
```
INFURA_KEY=
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

## Fund Services Tokens

```bash
./stack fund-services-tokens
```