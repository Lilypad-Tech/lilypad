import { ethers } from 'hardhat'
import {
  BigNumberish,
  AddressLike,
  Signer,
} from 'ethers'
import {
  getWallet,
  fundAccountsWithTokens,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'
import {
  LilypadToken,
  LilypadPayments,
  LilypadStorage,
  LilypadUsers,
  LilypadController,
  LilypadMediationRandom,
} from '../typechain-types'
import {
  SharedStructs,
} from '../typechain-types/contracts/LilypadStorage'

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

export async function deployToken(
  signer: Signer,
  tokenSupply: BigNumberish = DEFAULT_TOKEN_SUPPLY,
  testMode = false,
) {
  return deployContract<LilypadToken>(testMode ? 'LilypadTokenTestable' : 'LilypadToken', signer, [
    'LilyPad',
    'LLY',
    tokenSupply,
  ])
}

export async function deployPayments(
  signer: Signer,
  tokenAddress: AddressLike,
  testMode = false,
) {
  const payments = await deployContract<LilypadPayments>(testMode ? 'LilypadPaymentsTestable' : 'LilypadPayments', signer)
  await payments
    .connect(signer)
    .initialize(tokenAddress)
  return payments
}

export async function deployStorage(
  signer: Signer,
  testMode = false,
) {
  return deployContract<LilypadStorage>(testMode ? 'LilypadStorageTestable' : 'LilypadStorage', signer)
}

export async function deployUsers(
  signer: Signer
) {
  return deployContract<LilypadUsers>('LilypadUsers', signer)
}

export async function deployMediation(
  signer: Signer
) {
  return deployContract<LilypadMediationRandom>('LilypadMediationRandom', signer)
}

export async function deployController(
  signer: Signer,
  storageAddress: AddressLike,
  usersAddress: AddressLike,
  paymentsAddress: AddressLike,
  mediationAddress: AddressLike,
  jobCreatorAddress: AddressLike
) {
  const controller = await deployContract<LilypadController>('LilypadController', signer)
  await controller
    .connect(signer)
    .initialize(storageAddress, usersAddress, paymentsAddress, mediationAddress, jobCreatorAddress)
  return controller
}

/*

  FIXTURES

  these are thin wrappers around our web3 utils lib

  used by tests to prepare env for unit tests

*/

/*

  TOKEN

*/

// setup the token in test mode so we can call functions on it directly
// without the ControllerOwnable module kicking in
export async function setupTokenFixture({
  testMode = false,
  withFunds = false,
  controllerAddress,
}: {
  testMode?: boolean,
  withFunds?: boolean,
  controllerAddress?: AddressLike,
}) {
  const admin = getWallet('admin')
  const token = await deployToken(
    admin,
    DEFAULT_TOKEN_SUPPLY,
    testMode,
  )
  if(withFunds) {
    await fundAccountsWithTokens(token, DEFAULT_TOKENS_PER_ACCOUNT)
  }
  if(controllerAddress) {
    await (token as any)
      .connect(admin)
      .setControllerAddress(controllerAddress)
  }
  return token
}

/*

  PAYMENTS

*/

// setup the token in non-test mode but the payments in test mode
// then we can call functions directly on the payments contract
// and the token contract will check with ControllerOwnable
export async function setupPaymentsFixture({
  testMode = false,
  withFunds = false,
  controllerAddress,
}: {
  testMode?: boolean,
  withFunds?: boolean,
  controllerAddress?: AddressLike,
}) {
  const admin = getWallet('admin')
  const token = await setupTokenFixture({
    testMode: false,
    withFunds,
  })
  const payments = await deployPayments(admin, token.getAddress(), testMode)
  await (token as any)
    .connect(admin)
    .setControllerAddress(payments.getAddress())
  if(controllerAddress) {
    await (payments as any)
      .connect(admin)
      .setControllerAddress(controllerAddress)
  }
  return {
    token,
    payments,
  }
}

/*

  STORAGE

*/

// setup the token in test mode so we can call functions on it directly
// without the ControllerOwnable module kicking in
export async function setupStorageFixture({
  testMode = false,
  controllerAddress,
}: {
  testMode?: boolean,
  controllerAddress?: AddressLike,
}) {
  const admin = getWallet('admin')
  const storage = await deployStorage(
    admin,
    testMode,
  )
  if(controllerAddress) {
    await (storage as any)
      .connect(admin)
      .setControllerAddress(controllerAddress)
  }
  return storage
}

/*

  USERS

*/

// setup the token in test mode so we can call functions on it directly
// without the ControllerOwnable module kicking in
export async function setupUsersFixture() {
  const admin = getWallet('admin')
  const users = await deployUsers(
    admin,
  )
  return users
}

/*

  MEDIATION

*/

// setup the token in test mode so we can call functions on it directly
// without the ControllerOwnable module kicking in
export async function setupMediationFixture({
  controllerAddress,
}: {
  controllerAddress?: AddressLike,
}) {
  const admin = getWallet('admin')
  const mediation = await deployMediation(
    admin,
  )
  if(controllerAddress) {
    await (mediation as any)
      .connect(admin)
      .setControllerAddress(controllerAddress)
  }
  return mediation
}

/*

  CONTROLLER

*/

export async function setupControllerFixture({
  withFunds = false,
}: {
  withFunds?: boolean,
}) {
  const admin = getWallet('admin')
  const {
    token,
    payments,
  } = await setupPaymentsFixture({
    withFunds,
  })
  const storage = await setupStorageFixture({})
  const users = await setupUsersFixture()
  const mediation = await setupMediationFixture({})
  const paymentsAddress = await payments.getAddress()
  const storageAddress = await storage.getAddress()
  const usersAddress = await users.getAddress()
  const mediationAddress = await mediation.getAddress()

  const jobCreator = getWallet('job_creator')

  const controller = await deployController(
    admin,
    storageAddress,
    usersAddress,
    paymentsAddress,
    mediationAddress,
    jobCreator
  )
  const controllerAddress = await controller.getAddress()
  await (payments as any)
    .connect(admin)
    .setControllerAddress(controllerAddress)
  await (storage as any)
    .connect(admin)
    .setControllerAddress(controllerAddress)
  await (mediation as any)
    .connect(admin)
    .setControllerAddress(controllerAddress)
  return {
    token,
    payments,
    storage,
    users,
    mediation,
    controller,
  }
}

export const DEAL_ID = "10"
export const RESULTS_ID = "11"
export const DATA_ID = "12"

export const DEFAULT_VALUES: Record<string, bigint> = {
  instructionPrice: ethers.parseEther("10"),
  instructionCount: ethers.getBigInt("1"),
  resultsCollateralMultiple: ethers.getBigInt("4"),
  resultsCollateral: ethers.parseEther("40"),
  paymentCollateral: ethers.parseEther("30"),
  jobCost: ethers.parseEther("10"),
  mediationFee: ethers.parseEther("5"),
  timeout: ethers.getBigInt("100"),
  timeoutCollateral: ethers.parseEther("10"),
}

export function getDefaultTimeouts(
  timeout = DEFAULT_VALUES.timeout,
  collateral = DEFAULT_VALUES.timeoutCollateral,
) {
  const defaultTimeout: SharedStructs.DealTimeoutStruct = {
    timeout,
    collateral,
  }
  const defaultTimeoutNoCost: SharedStructs.DealTimeoutStruct = {
    timeout,
    collateral: ethers.parseEther("0"),
  }
  const ret: SharedStructs.DealTimeoutsStruct = {
    agree: defaultTimeoutNoCost,
    submitResults: defaultTimeout,
    judgeResults: defaultTimeout,
    mediateResults: defaultTimeoutNoCost,
  }
  return ret
}

export function getDefaultPricing(
  instructionPrice = DEFAULT_VALUES.instructionPrice,
  paymentCollateral = DEFAULT_VALUES.paymentCollateral,
  resultsCollateralMultiple = DEFAULT_VALUES.resultsCollateralMultiple,
  mediationFee = DEFAULT_VALUES.mediationFee,
) {
  const ret: SharedStructs.DealPricingStruct = {
    instructionPrice,
    paymentCollateral,
    resultsCollateralMultiple,
    mediationFee,
  }
  return ret
}