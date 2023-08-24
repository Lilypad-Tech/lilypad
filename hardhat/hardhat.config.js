require('hardhat-deploy')
require('@nomicfoundation/hardhat-toolbox')
const dotenv = require('dotenv')
const ENV_FILE = process.env.DOTENV_CONFIG_PATH || '../.env'
dotenv.config({ path: ENV_FILE })
const {
  ACCOUNT_ADDRESSES,
  PRIVATE_KEYS,
} = require('./utils/accounts')

const config = {
  solidity: '0.8.6',
  defaultNetwork: 'geth',
  namedAccounts: ACCOUNT_ADDRESSES,
  networks: {
    hardhat: {},
    geth: {
      url: 'http://localhost:8545',
      chainId: 1337,
      accounts: PRIVATE_KEYS,
    },
  },
};

module.exports = config
