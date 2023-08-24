const {
  ACCOUNTS,
} = require('../utils/accounts')

async function main() {
  // loop over accounts and print env for address and private key
  ACCOUNTS.forEach((account) => {
    console.log(`export ${account.name.toUpperCase()}_ADDRESS=${account.address}`)
    console.log(`export ${account.name.toUpperCase()}_PRIVATE_KEY=${account.privateKey}`)
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
