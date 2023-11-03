import {
  ACCOUNTS,
} from '../utils/accounts'

async function main() {
  // loop over accounts and print env for address and private key
  ACCOUNTS.forEach((account) => {
    console.log(`export ${account.name.toUpperCase()}_ADDRESS=${account.address}`)
    console.log(`export ${account.name.toUpperCase()}_PRIVATE_KEY=${account.privateKey}`)
  })
  console.log('export INFURA_KEY=')
  console.log('export LOG_LEVEL=info')
  console.log('export NETWORK=geth')
  console.log('export WEB3_RPC_URL=ws://localhost:8546')
  console.log('export WEB3_CHAIN_ID=1337')
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
