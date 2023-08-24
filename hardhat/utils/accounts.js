const ethers = require('ethers')

const loadEnv = (name, defaultValue) => {
  return process.env[name] || defaultValue
}

const loadPrivateKey = (name, defaultValue) => {
  return loadEnv(`PRIVATE_KEY_${name.toUpperCase()}`, defaultValue)
}

const loadAddress = (name, defaultValue) => {
  return loadEnv(`ADDRESS_${name.toUpperCase()}`, defaultValue)
}

const AMOUNT_TO_FUND = ethers.utils.parseEther('10000')

const ACCOUNTS = [{
  name: 'admin',
  address: loadAddress('admin', '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266'),
  privateKey: loadPrivateKey('admin', '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80')
}, {
  name: 'faucet',
  address: loadAddress('faucet', '0x70997970C51812dc3A010C7d01b50e0d17dc79C8'),
  privateKey: loadPrivateKey('faucet', '0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d')
}, {
  name: 'solver',
  address: loadAddress('solver', '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC'),
  privateKey: loadPrivateKey('solver', '0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a')
}, {
  name: 'mediator',
  address: loadAddress('mediator', '0x90F79bf6EB2c4f870365E785982E1f101E93b906'),
  privateKey: loadPrivateKey('mediator', '0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6')
}, {
  name: 'resource_provider',
  address: loadAddress('resource_provider', '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65'),
  privateKey: loadPrivateKey('resource_provider', '0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a')
}, {
  name: 'job_creator',
  address: loadAddress('job_creator', '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc'),
  privateKey: loadPrivateKey('job_creator', '0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba')
}, {
  name: 'directory',
  address: loadAddress('directory', '0x976EA74026E726554dB657fA54763abd0C3a0aa9'),
  privateKey: loadPrivateKey('directory', '0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e')
}]

// map of account name -> account
const NAMED_ACCOUNTS = ACCOUNTS.reduce((all, acc) => {
  all[acc.name] = acc
  return all
}, {})

// map of account name -> account address
const ACCOUNT_ADDRESSES = ACCOUNTS.reduce((all, acc) => {
  all[acc.name] = acc.address
  return all
}, {})

// flat list of private keys in order
const PRIVATE_KEYS = ACCOUNTS.map(acc => acc.privateKey)

const getAccount = (name) => {
  const account = NAMED_ACCOUNTS[name]
  if (!account) {
    throw new Error(`Unknown account ${name}`)
  }
  return account
}

const getWallet = (accountName) => {
  const account = getAccount(accountName)
  return new ethers.Wallet(account.privateKey)
}

module.exports = {
  AMOUNT_TO_FUND,
  ACCOUNTS,
  NAMED_ACCOUNTS,
  ACCOUNT_ADDRESSES,
  PRIVATE_KEYS,
  loadEnv,
  loadPrivateKey,
  loadAddress,
  getAccount,
  getWallet,
}