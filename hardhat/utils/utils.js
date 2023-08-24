const hre = require('hardhat')

const {
  deployments,
} = hre

// amount is in wei
const transfer = async (fromAccount, toAccount, amount) => {
  const signer = new hre.ethers.Wallet(fromAccount.privateKey, hre.ethers.provider)

  const tx = await signer.sendTransaction({
    to: toAccount.address,
    value: amount,
  })
  await tx.wait()

  console.log(`Moved ${amount} from ${fromAccount.name} (${fromAccount.address}) to ${toAccount.name} (${toAccount.address}) - ${tx.hash}.`)
}

const getContract = async (name) => {
  const deployment = await deployments.get(name)
  const Factory = await ethers.getContractFactory(name)
  return Factory.attach(deployment.address)
}

module.exports = {
  transfer,
  getContract,
}