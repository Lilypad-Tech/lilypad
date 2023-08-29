import { ethers } from 'hardhat'
import {
  getWallet,
  deployToken,
  fundAllAccountsWithTokens,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'

export async function getTokenSigners() {
  const signers = await ethers.getSigners()
  const escrow = signers[signers.length-1]
  const controller = signers[signers.length-2]
  const payee = signers[signers.length-3]
  return {
    escrow,
    controller,
    payee,
  }
}

export async function setupToken() {
  const admin = getWallet('admin')
  const token = await deployToken(
    admin,
    DEFAULT_TOKEN_SUPPLY,
  )
  return token
}

export async function setupTokenWithAutoController() {
  const admin = getWallet('admin')
  const token = await setupToken()
  const signers = await getTokenSigners()
  await token
    .connect(admin)
    .setEscrowAddress(signers.escrow.address)
  await token
    .connect(admin)
    .setControllerAddress(signers.controller.address)
  return token
}

export async function setupTokenWithFunds() {
  const token = await setupToken()
  await fundAllAccountsWithTokens(token, DEFAULT_TOKENS_PER_ACCOUNT)
  return token
}