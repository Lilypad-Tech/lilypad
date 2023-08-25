import hre, { ethers } from 'hardhat'
import {
  BigNumberish,
  Signer,
  AddressLike
} from 'ethers'
import { Account } from './types'
import {
  getAccount,
} from './accounts'
import {
  LilypadStorage,
  LilypadToken,
  LilypadController,
} from '../typechain-types'

export const AMOUNT_TO_FUND = ethers.parseEther('10000')

export const getWallet = (name: string) => {
  const account = getAccount(name)
  return new ethers.Wallet(account.privateKey, ethers.provider)
}

export const getAddress = (name: string) => {
  const account = getAccount(name)
  return account.address
}

// amount is in wei
export const transfer = async (fromAccount: Account, toAccount: Account, amount: BigNumberish) => {
  const signer = new hre.ethers.Wallet(fromAccount.privateKey)

  const tx = await signer.sendTransaction({
    to: toAccount.address,
    value: amount,
  })
  await tx.wait()

  console.log(`Moved ${amount} from ${fromAccount.name} (${fromAccount.address}) to ${toAccount.name} (${toAccount.address}) - ${tx.hash}.`)
}

export async function deployContract<T extends any>(
  name: string,
  signer: Signer,
): Promise<T>{
  const factory = await ethers.getContractFactory(
    name,
    signer,
  )
  const contract = await factory.deploy() as unknown as T
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

export async function deployStorage(signer: Signer) {
  return deployContract<LilypadStorage>('LilypadStorage', signer)
}

export async function connectStorage() {
  return connectContract<LilypadStorage>('LilypadStorage')
}

export async function getStorageAddress() {
  return getContractAddress('LilypadStorage')
}

export async function deployToken(signer: Signer) {
  return deployContract<LilypadToken>('LilypadToken', signer)
}

export async function connectToken() {
  return connectContract<LilypadToken>('LilypadToken')
}

export async function getTokenAddress() {
  return getContractAddress('LilypadToken')
}

export async function deployController(
  signer: Signer,
  storageAddress: AddressLike,
  tokenAddress: AddressLike,
) {
  const controller = await deployContract<LilypadController>('LilypadController', signer)
  await controller
    .connect(signer)
    .initialize(storageAddress, tokenAddress)
  return controller
}

export async function connectController() {
  return connectContract<LilypadController>('LilypadController')
}

export async function getControllerAddress() {
  return getContractAddress('LilypadController')
}

export async function deployContracts(
  signer: Signer,
) {
  const storage = await deployStorage(signer)
  const token = await deployToken(signer)
  const storageAddress = await getStorageAddress()
  const tokenAddress = await getTokenAddress()
  const controller = await deployController(signer, storageAddress, tokenAddress)
  return {
    storage,
    token,
    controller,
  }
}