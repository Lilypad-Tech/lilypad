import { HardhatUserConfig } from 'hardhat/config'
import '@typechain/hardhat'
import '@nomicfoundation/hardhat-toolbox'
import '@nomicfoundation/hardhat-ethers'
import '@nomicfoundation/hardhat-chai-matchers'
import 'hardhat-deploy'
import * as dotenv from 'dotenv'

import {
  loadAddress,
  loadPrivateKey,
  getAllKeys,
  setAccount,
  getAllAddresses,
  ACCOUNTS
} from './utils/accounts'

//override default publically known keys by doppler config
for (const acc of ACCOUNTS) {
  if (process.env[acc.name.toUpperCase()+'_PRIVATE_KEY'] && process.env[acc.name.toUpperCase()+'_ADDRESS']) {
    acc.address = loadAddress(acc.name, process.env[acc.name+'_ADDRESS']?.toString() ?? '')
    acc.privateKey = loadPrivateKey(acc.name, process.env[acc.name+'_PRIVATE_KEY']?.toString() ?? '')
    setAccount(acc)
  }
}

const PRIVATE_KEYS = getAllKeys()
const ACCOUNT_ADDRESSES = getAllAddresses()

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
      url: 'http://localhost:8545',
      chainId: 1337,
      accounts: PRIVATE_KEYS,
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_KEY}`,
      accounts: PRIVATE_KEYS,
    },
    local_l2: {
      url: 'http://localhost:8547',
      chainId: 412346,
      accounts: PRIVATE_KEYS,
    },
    arbitrumSepolia: {
      url: 'https://sepolia-rollup.arbitrum.io/rpc',
      chainId: 421614,
      accounts: PRIVATE_KEYS
    },
    arbitrumOne: {
      chainId: 42161,
      url: 'https://arb1.arbitrum.io/rpc',
      accounts: PRIVATE_KEYS
    },
    arbitrumNova: {
      chainId: 42170,
      url: 'https://nova.arbitrum.io/rpc',
      accounts: PRIVATE_KEYS
    },
    mumbai: {
      chainId: 80001,
      url: `https://polygon-testnet.public.blastapi.io`,
      accounts: PRIVATE_KEYS,
    },
    polygon: {
      chainId: 137,
      url: `https://rpc.ankr.com/polygon`,
      accounts: PRIVATE_KEYS,
    },
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};

module.exports = config