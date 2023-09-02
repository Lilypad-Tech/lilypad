import { ethers } from 'hardhat'
import {
  getWallet,
  deployToken,
  fundAllAccountsWithTokens,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'

export async function setupToken() {
  const admin = getWallet('admin')
  const token = await deployToken(
    admin,
    DEFAULT_TOKEN_SUPPLY,
  )
  return token
}

export async function setupTokenWithFunds() {
  const token = await setupToken()
  await fundAllAccountsWithTokens(token, DEFAULT_TOKENS_PER_ACCOUNT)
  return token
}