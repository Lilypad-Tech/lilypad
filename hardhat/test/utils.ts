import { ethers } from 'hardhat'
import {
  getWallet,
  deployToken,
  deployPayments,
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
export async function setupTokenFixture(withFunds = true) {
  const admin = getWallet('admin')
  const token = await deployToken(
    admin,
    DEFAULT_TOKEN_SUPPLY,
  )
  if(withFunds) {
    await fundAccountsWithTokens(token, DEFAULT_TOKENS_PER_ACCOUNT)
  }
  return token
}

/*

  PAYMENTS

*/
export async function setupPaymentsFixture() {
  const admin = getWallet('admin')
  const token = await setupTokenFixture(true)
  const payments = await deployPayments(admin, token.getAddress())
  token.transferOwnership(payments.getAddress())
  return {
    token,
    payments,
  }
}