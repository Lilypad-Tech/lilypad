const bluebird = require('bluebird')
const {
  ACCOUNTS,
  AMOUNT_TO_FUND,
  getAccount,
} = require('../utils/accounts')

const {
  transfer,
} = require('../utils/utils')

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
