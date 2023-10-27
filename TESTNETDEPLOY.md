# Testnet Deployment

## Create Seven New Accounts

Follow the README.md in the `generate_accts` directory to create seven new accounts.

Update the following environment variables in `.env`:
```
export ADMIN_ADDRESS=
export ADMIN_PRIVATE_KEY=
export FAUCET_ADDRESS=
export FAUCET_PRIVATE_KEY=
export SOLVER_ADDRESS=
export SOLVER_PRIVATE_KEY=
export MEDIATOR_ADDRESS=
export MEDIATOR_PRIVATE_KEY=
export RESOURCE_PROVIDER_ADDRESS=
export RESOURCE_PROVIDER_PRIVATE_KEY=
export JOB_CREATOR_ADDRESS=
export JOB_CREATOR_PRIVATE_KEY=
export DIRECTORY_ADDRESS=
export DIRECTORY_PRIVATE_KEY=
```

## Create a new Infura Project

Create a new Infura project and update the following environment variable in `.env`:
```
export INFURA_KEY=
```

## Fund the Seven New Accounts

Fund the seven new accounts with testnet ETH using MetaMask.

## Deploy Contracts
