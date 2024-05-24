import { HardhatUserConfig } from 'hardhat/config'
import '@typechain/hardhat'
import '@nomicfoundation/hardhat-toolbox'
import '@nomicfoundation/hardhat-ethers'
import '@nomicfoundation/hardhat-chai-matchers'
import 'hardhat-deploy'
import * as dotenv from 'dotenv'

import {
  ACCOUNT_ADDRESSES,
  getAccount,
} from './utils/accounts'

const ENV_FILE = process.env.DOTENV_CONFIG_PATH || '../.env'
dotenv.config({ path: ENV_FILE })

const NETWORK = process.env.NETWORK || "geth";
const NETWORK_URL = process.env.WEB3_HTTP_URL || 'http://localhost:8545';

const INFURA_KEY = process.env.INFURA_KEY || "";

const config: HardhatUserConfig = {
  solidity: '0.8.21',
  defaultNetwork: NETWORK,
  namedAccounts: ACCOUNT_ADDRESSES,
  networks: {
    hardhat: {},
    geth: {
      url: NETWORK_URL,
      chainId: 1337,
      accounts: [getAccount('admin').privateKey]
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_KEY}`,
      accounts: [getAccount('admin').privateKey],
    },
    arbitrumOne: {
      chainId: 42161,
      url: 'https://arb1.arbitrum.io/rpc',
      accounts: [getAccount('admin').privateKey],
    },
    arbitrumNova: {
      chainId: 42170,
      url: 'https://nova.arbitrum.io/rpc',
      accounts: [getAccount('admin').privateKey],
    },
    arbitrumSepolia: {
      url: 'https://sepolia-rollup.arbitrum.io/rpc',
      chainId: 421614,
      accounts: [getAccount('admin').privateKey],
    },
    local_l2: {
      url: 'http://localhost:8547',
      chainId: 412346,
      accounts: [getAccount('admin').privateKey],
    },

  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};

module.exports = config
