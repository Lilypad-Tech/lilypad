import hre, { ethers } from 'hardhat'
import { BigNumberish } from 'ethers'
import { Account } from './types'

// amount is in wei
const transfer = async (fromAccount: Account, toAccount: Account, amount: BigNumberish) => {
  const signer = new hre.ethers.Wallet(fromAccount.privateKey, hre.ethers.provider)

  const tx = await signer.sendTransaction({
    to: toAccount.address,
    value: amount,
  })
  await tx.wait()

  console.log(`Moved ${amount} from ${fromAccount.name} (${fromAccount.address}) to ${toAccount.name} (${toAccount.address}) - ${tx.hash}.`)
}

const getContract = async (name: string) => {
  const deployment = await hre.deployments.get(name)
  const Factory = await hre.ethers.getContractFactory(name)
  return Factory.attach(deployment.address)
}

module.exports = {
  transfer,
  getContract,
}