import hre, { ethers } from 'hardhat'
import bluebird from 'bluebird'
import {
  BigNumberish,
  Signer,
  AddressLike
} from 'ethers'
import { Account } from './types'
import {
  ACCOUNTS,
  getAccount,
} from './accounts'
import {
  LilypadToken,
  LilypadPayments,
  LilypadStorage,
  LilypadController,
} from '../typechain-types'

/*

  CONSTANTS

*/
export const AMOUNT_TO_FUND = ethers.parseEther('10000')

// a million tokens in total
export const DEFAULT_TOKEN_SUPPLY = ethers.parseEther('1000000')

// each service gets 1000 tokens
export const DEFAULT_TOKENS_PER_ACCOUNT = ethers.parseEther('1000')

/*

  WALLET UTILS

*/
export const getWallet = (name: string) => {
  const account = getAccount(name)
  return new ethers.Wallet(account.privateKey, ethers.provider)
}

export const getAddress = (name: string) => {
  const account = getAccount(name)
  return account.address
}

export const getRandomWallet = () => {
  return ethers.Wallet.createRandom()
}

export const transfer = async (fromAccount: Account, toAccount: Account, amount: BigNumberish) => {
  const signer = new hre.ethers.Wallet(fromAccount.privateKey)

  const tx = await signer.sendTransaction({
    to: toAccount.address,
    value: amount,
  })
  await tx.wait()

  console.log(`Moved ${amount} from ${fromAccount.name} (${fromAccount.address}) to ${toAccount.name} (${toAccount.address}) - ${tx.hash}.`)
}

/*

  DEPLOYMENT

*/
export async function deployContract<T extends any>(
  name: string,
  signer: Signer,
  args: any[] = [],
): Promise<T>{
  const factory = await ethers.getContractFactory(
    name,
    signer,
  )
  const contract = await factory.deploy(...args) as unknown as T
  return contract
}

export async function connectContract<T extends any>(
  name: string,
): Promise<T>{
  const deployment = await hre.deployments.get(name)
  const factory = await hre.ethers.getContractFactory(name)
  const contract = factory.attach(deployment.address) as unknown as T
  return contract
}

export async function getContractAddress(
  name: string,
): Promise<AddressLike>{
  const deployment = await hre.deployments.get(name)
  return deployment.address
}


/*

  TOKEN

*/
export async function deployToken(
  signer: Signer,
  tokenSupply: BigNumberish = DEFAULT_TOKEN_SUPPLY,
) {
  return deployContract<LilypadToken>('LilypadToken', signer, [
    'LilyPad',
    'LLY',
    tokenSupply,
  ])
}

export async function connectToken() {
  return connectContract<LilypadToken>('LilypadToken')
}

export async function getTokenAddress() {
  return getContractAddress('LilypadToken')
}

export async function fundTokens(
  tokenContract: LilypadToken,
  address: AddressLike,
  amount: BigNumberish = DEFAULT_TOKENS_PER_ACCOUNT,
) {
  await tokenContract
    .connect(getWallet('admin'))
    .transfer(address, amount)
}

export async function fundAccountsWithTokens(
  tokenContract: LilypadToken,
  amount: BigNumberish = DEFAULT_TOKENS_PER_ACCOUNT,
) {
  await bluebird.mapSeries(ACCOUNTS, async (account) => {
    await fundTokens(tokenContract, account.address, amount)
  })
}

/*

  PAYMENTS

*/
export async function deployPayments(
  signer: Signer,
  tokenAddress: AddressLike,
) {
  const payments = await deployContract<LilypadPayments>('LilypadPayments', signer)
  await payments
    .connect(signer)
    .initialize(tokenAddress)
  return payments
}


/*

  STORAGE

*/
export async function deployStorage(signer: Signer) {
  return deployContract<LilypadStorage>('LilypadStorage', signer)
}

export async function connectStorage() {
  return connectContract<LilypadStorage>('LilypadStorage')
}

export async function getStorageAddress() {
  return getContractAddress('LilypadStorage')
}



/*

  CONTROLLER

*/
export async function deployController(
  signer: Signer,
  storageAddress: AddressLike,
  paymentsAddress: AddressLike,
) {
  const controller = await deployContract<LilypadController>('LilypadController', signer)
  await controller
    .connect(signer)
    .initialize(storageAddress, paymentsAddress)
  return controller
}

export async function connectController() {
  return connectContract<LilypadController>('LilypadController')
}

export async function getControllerAddress() {
  return getContractAddress('LilypadController')
}
