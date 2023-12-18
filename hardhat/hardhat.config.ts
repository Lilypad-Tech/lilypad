import { HardhatUserConfig } from 'hardhat/config'
import '@typechain/hardhat'
import '@nomicfoundation/hardhat-toolbox'
import '@nomicfoundation/hardhat-ethers'
import '@nomicfoundation/hardhat-chai-matchers'
import 'hardhat-deploy'
import * as dotenv from 'dotenv'

import {
  ACCOUNT_ADDRESSES,
  PRIVATE_KEYS,
} from './utils/accounts'
import {ethers} from "hardhat";

const ENV_FILE = process.env.DOTENV_CONFIG_PATH || '../.env'
dotenv.config({ path: ENV_FILE })

const NETWORK = process.env.NETWORK || "geth";

const INFURA_KEY = process.env.INFURA_KEY || "";

const config: HardhatUserConfig = {
  solidity: '0.8.21',
  defaultNetwork: NETWORK,
  namedAccounts: ACCOUNT_ADDRESSES,
  networks: {
    hardhat: {
      accounts: [
        {
          privateKey: process.env.PRIVATE_KEY || 'beb00ab9be22a34a9c940c27d1d6bfe59db9ab9de4930c968b16724907591b3f',
          balance: `${1000000000000000000000000n}`,
        },
      ],
    },
    geth: {
      url: 'http://localhost:8545',
      chainId: 1337,
      accounts: PRIVATE_KEYS,
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_KEY}`,
      accounts: PRIVATE_KEYS,
    },
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};

module.exports = config