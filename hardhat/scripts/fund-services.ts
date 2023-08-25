import bluebird from 'bluebird'
import {
  ACCOUNTS,
  getAccount,
} from '../utils/accounts'
import {
  transfer,
  AMOUNT_TO_FUND,
} from '../utils/utils'

async function main() {
  const adminAccount = getAccount('admin')
  await bluebird.mapSeries(ACCOUNTS, async (toAccount) => {
    if(toAccount.name === 'admin') return
    await transfer(adminAccount, toAccount, AMOUNT_TO_FUND)
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
