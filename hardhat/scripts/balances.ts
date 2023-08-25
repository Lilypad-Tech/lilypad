import { ethers } from 'hardhat'
import bluebird from 'bluebird'
import {
  ACCOUNTS,
} from '../utils/accounts'

async function main() {
  await bluebird.mapSeries(ACCOUNTS, async (account) => {
    const balance = await ethers.provider.getBalance(account.address)
    console.log(`${account.name} - ${account.address} = ${ethers.formatEther(balance)}`)
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
