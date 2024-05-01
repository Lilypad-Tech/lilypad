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

const ENV_FILE = process.env.DOTENV_CONFIG_PATH || '../.env'
dotenv.config({ path: ENV_FILE })

const NETWORK = process.env.NETWORK || "geth";

const INFURA_KEY = process.env.INFURA_KEY || "";

const config: HardhatUserConfig = {
  solidity: '0.8.21',
  defaultNetwork: NETWORK,
  namedAccounts: ACCOUNT_ADDRESSES,
  networks: {
    hardhat: {},
    geth: {
      url: 'http://127.0.0.1:8545',
      chainId: 1337,
      accounts: PRIVATE_KEYS,
      allowUnlimitedContractSize: true,
      gas: 5000000, //units of gas you are willing to pay, aka gas limit
      gasPrice:  50000000000, //gas is typically in units of gwei, but you must enter it as wei here
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_KEY}`,
      accounts: PRIVATE_KEYS,
    },
    lilypad: {
      url: 'http://127.0.0.1:9652/ext/bc/2K2MUJvsDrFwJdGdcQKDy2hS8Jxti4yrsfS3JJ1UNnSkp7nKGi/rpc',
      chainId: 42,
      accounts: PRIVATE_KEYS,
    },
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};

module.exports = config