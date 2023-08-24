const hre = require('hardhat')
const bluebird = require('bluebird')

const {
  ACCOUNTS,
} = require('../utils/accounts')

async function main() {
  await bluebird.mapSeries(ACCOUNTS, async (account) => {
    const balance = await hre.ethers.provider.getBalance(account.address)
    console.log(`${account.name} - ${account.address} = ${ethers.utils.formatEther(balance)}`)
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
