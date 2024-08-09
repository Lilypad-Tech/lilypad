# Lilypad Smart Contracts

Table of Contents

1. [Contracts Overview](#contracts-overview)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Deploy Smart Contracts on Arbitrum Testnet](#deploy-smart-contracts-on-arbitrum-testnet)
5. [Update Existing Smart Contracts](#update-existing-smart-contracts)
7. [Strategy for Arbitrum Mainnet Deployment](#strategy-for-arbitrum-mainnet-deployment)
8. [Running Jobs from Smart Contracts](#running-jobs-from-smart-contracts)

# Contracts Overview
## Core Contracts:
- **LilypadPayments.sol**: Manages payment flows within the ecosystem.
- **LilypadToken.sol**: Manages the ERC20-compatible Lilypad tokens.
- **LilypadStorage.sol**: Provides decentralized storage for data across the platform.
- **LilypadMediationRandom.sol**: Implements a random selection mechanism for mediators.
- **LilypadUsers.sol**: Handles user management, including registration and access control.
- **LilypadPow.sol**: Provides Proof-of-Work (PoW) functionalities.
- **LilypadOnChainJobCreator.sol**: Allows users to create and manage on-chain jobs.
- **LilypadController.sol**: Manages core functions and interactions between different components.
- **SharedStructs.sol**: Contains shared data structures used by multiple contracts.
- **ExampleClient.sol**: A demonstration contract for client interactions with Lilypad.

# Prerequisites
Ensure you have the following installed:

- Node.js (v20 or higher) and npm
- HardHat (installed globally): npm install --global hardhat
- MetaMask (or another Ethereum wallet)
- Infura API key (for interacting with Arbitrum networks)
- Etherscan API key (for verifying contracts on Etherscan)

# Installation
Step 1: Setup Environment
```bash
cp .env.sample.env
```
Update the .env file with your own environment variables.

Step 2: Install Dependencies
```bash
npm install
```

Step 3: Compile Smart Contracts
```bash
npx hardhat compile
```

Step 4: Run Tests
```bash
npx hardhat test
```

# Deploy Smart Contracts on Arbitrum Testnet
Step 1: Configure hardhat.config.js to include Arbitrum Testnet details:
```js
  networks: {
    testnet:{
      url: 'https://sepolia-rollup.arbitrum.io/rpc',
      chainId: 421614,
      accounts: [getAccount('admin').privateKey],
    },
  },
```

Step 2: Deploy contracts to Arbitrum Testnet:
```bash
npx hardhat run deploy/{{deploy_contract}}.js --network testnet
```
There are multiple deploy scripts available for different contracts. Ensure to use the appropriate deploy script for your deployment.

Step 3: After deployment, verify contracts on the Arbiscan Testnet using:
```bash
npx hardhat verify --network testnet DEPLOYED_CONTRACT_ADDRESS
```

# Update Existing Smart Contracts
Step 1: Make Contract Changes
Edit the Solidity files located in the contracts/ directory.

Step 2: Recompile Contracts
```bash
npx hardhat compile
```

Step 3: Deploy Updated Contracts
For incremental changes:
```bash
npx hardhat run deploy/{{deploy_contract}}.js --network testnet
```
There are multiple deploy scripts available for different contracts. Ensure to use the appropriate deploy script for your deployment.

Step 4: Migrate Data (If Needed)
If your contract update affects storage, ensure to include migration scripts or processes to move data from the old contract to the new one.

# Strategy for Arbitrum Mainnet Deployment

Step 1: Finalize Contract Code
Ensure all contracts are thoroughly tested and audited. This is crucial for deploying on the mainnet.

Step 2: Create a Deployment Plan
Map out the deployment process, considering dependencies between contracts and the order of deployment.

Step 3: Update HardHat Configuration
Add Arbitrum Mainnet to your hardhat.config.js:
```js
  networks: {
    mainnet:{
      url: 'https://arb1.arbitrum.io/rpc',
      chainId: 42161,
      accounts: [getAccount('admin').privateKey],
    },
  },
```

Step 4: Deploy to Arbitrum Mainnet:
```bash
npx hardhat run deploy/{{deploy_contract}}.js --network mainnet
```
There are multiple deploy scripts available for different contracts. Ensure to use the appropriate deploy script for your deployment.

Step 5: Monitor and Verify
After deployment, use tools like Etherscan or Arbiscan to monitor the contracts and ensure they are functioning as expected. Verify the contracts if necessary.

Step 6: Post-Deployment Actions
Set up any required governance or admin controls. Notify users or clients about the new deployment and any actions they need to take.

# Running Jobs from Smart Contracts

We have deloyed the [LilypadOnChainJobCreator](contracts/LilypadOnChainJobCreator.sol) contract which you can use to trigger running jobs on the Lilypad network from other smart contracts.

It works in tandem with the `lilypad jobcreator` on-chain which will watch the on-chain contract and manage jobs on behalf of contracts that submit them.

### Creating a Job

You will need to know the contract address for the on-chain job creator so we can submit transactions to it.

The production controller address is `0xF2fD1B9b262982F12446149A27d8901Ac68dcB59` and you can ask it for the address of the on-chain job creator `getJobCreatorAddress()`

Running a job involves 2 phases:

 * calling `approve` on the ERC-20 contract to allow the solver to spend your tokens
 * trigger the job via the on chain job manager

Now we know the address of the on-chain job controller - we can ask it for 3 things:

 * the address of the ERC-20 token contract - `getTokenAddress()`
 * how much the required deposit it - `getRequiredDeposit()`
 * the address of the solver that will handle running the job for us - `getControllerAddress()`

Knowing these 3 things means we can call the standard ERC-20 `approve` to allow the solver to spend our tokens on our behalf.

Now - we can call the `runJob` method of the on chain controller from another contract.  This will cause the job-creator service to kick in and do the following things:

 * check that funds have been approved for the solver
 * transfer those funds to it's wallet
 * run the job on lilypad
 * call the `submitResults` method on the on-chain job creator
 * the on-chain job creator will call the `submitResults` of the original calling contract

The following is an example on-chain smart contract:

```solidity
// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ILilypadJobManager.sol";
import "./ILilypadJobClient.sol";

contract ExampleClient is Ownable, Initializable, ILilypadJobClient {

  address private jobManagerAddress;
  ILilypadJobManager private jobManagerContract;

  mapping(uint256 => string) private jobResults;

  event JobCreated(
    uint256 id,
    string message
  );

  event JobCompleted(
    uint256 id,
    string dealId,
    string dataId
  );

  function initialize(address _jobManagerAddress) public initializer {
    setJobManagerAddress(_jobManagerAddress);
  }

  function setJobManagerAddress(address _jobManagerAddress) public onlyOwner {
    require(_jobManagerAddress != address(0), "Job manager address");
    jobManagerAddress = _jobManagerAddress;
    jobManagerContract = ILilypadJobManager(jobManagerAddress);
  }

  function getJobResult(uint256 _jobID) public view returns (string memory) {
    return jobResults[_jobID];
  }

  function runCowsay(
    string memory message
  ) public {
    string[] memory inputs = new string[](1);
    inputs[0] = string(abi.encodePacked("Message=", message));
    uint256 id = jobManagerContract.runJob(
      "cowsay:v0.0.4",
      inputs,
      msg.sender
    );

    emit JobCreated(
      id,
      message
    );
  }

  function submitResults(
    uint256 id,
    string memory dealId,
    string memory dataId
  ) public override {
    jobResults[id] = dataId;
    emit JobCompleted(
      id,
      dealId,
      dataId
    );
  }
}
```

Here is an example of a script that brings all of this together:

```typescript
import bluebird from 'bluebird'
import {
  connectToken,
  connectJobManager,
  connectExampleClient,
  getWallet,
  getAddress,
} from '../utils/web3'
import { ethers } from 'hardhat'

async function main() {
  // it's annoying to not be able to use argv but hardhat complains about it
  const message = process.env.MESSAGE || 'Hello World!'

  const token = await connectToken()
  const manager = await connectJobManager()
  const client = await connectExampleClient()

  const setRequiredDepositTx = await manager
    .connect(getWallet('solver'))
    .setRequiredDeposit(ethers.parseEther("2"))
  await setRequiredDepositTx.wait()

  const requiredDeposit = await manager.getRequiredDeposit()

  console.log(`requiredDeposit: ${Number(requiredDeposit)}`)

  const paytokensTx = await token
    .connect(getWallet('job_creator'))
    .approve(getAddress('solver'), requiredDeposit)
  await paytokensTx.wait()

  console.log(`tokens approved: ${paytokensTx.hash}`)

  const runjobTx = await client
    .connect(getWallet('job_creator'))
    .runCowsay(message)
  const receipt = await runjobTx.wait()
  if(!receipt) throw new Error(`no receipt`)

  console.log(`submitted job: ${runjobTx.hash}`)

  let jobID = 0

  receipt.logs.forEach(log => {
    const logs = client.interface.parseLog(log as any)
    if(!logs) return
    jobID = Number(logs.args[0])
  })

  console.log(`Job ID: ${jobID}`)
  console.log(`Waiting for job to be completed...`)

  let result = ''

  while(!result) {
    result = await client.getJobResult(jobID)
    if(!result) {
      await bluebird.delay(1000)
    }
  }

  console.log(`Job result: ${result}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

```
