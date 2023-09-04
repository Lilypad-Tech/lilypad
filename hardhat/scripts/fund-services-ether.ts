import bluebird from 'bluebird'
import {
  ACCOUNTS,
  getAccount,
} from '../utils/accounts'
import {
  transferEther,
  DEFAULT_ETHER_PER_ACCOUNT,
} from '../utils/web3'

async function main() {
  const adminAccount = getAccount('admin')
  await bluebird.mapSeries(ACCOUNTS, async (toAccount) => {
    if(toAccount.name === 'admin') return
    await transferEther(adminAccount, toAccount, DEFAULT_ETHER_PER_ACCOUNT)
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
