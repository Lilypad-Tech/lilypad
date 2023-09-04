import {
  AddressLike
} from 'ethers'
import {
  getWallet,
  deployToken,
  deployPayments,
  deployStorage,
  fundAccountsWithTokens,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'

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
  return {
    token,
    payments,
  }
}

/*

  TOKEN

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