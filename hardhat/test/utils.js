const {
  ethers,
} = require('hardhat')

const getBalance = async (address) => {
  const amount = await ethers.provider.getBalance(address)
  const inEth = ethers.utils.formatEther(amount)
  return Math.round(parseFloat(inEth))
}

const deployStorage = async (signer) => {
  const storageFactory = (await ethers.getContractFactory(
    "LilypadStorage",
    {
      signer,
    }
  ))
  const storageContract = await storageFactory.deploy()
  await storageContract.deployed()
  return storageContract
}


module.exports = {
  getBalance,
  deployStorage,
}
