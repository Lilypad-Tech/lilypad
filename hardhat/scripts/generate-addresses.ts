import Wallet from 'ethereumjs-wallet'


const generate = (name: string) => {
  const wallet = Wallet.generate()
  console.log(`export ${name}_PRIVATE_KEY=${wallet.getPrivateKeyString()}`)
  console.log(`export ${name}_ADDRESS=${wallet.getAddressString()}`)
}

async function main() {

  generate('ADMIN')
  generate('FAUCET')
  generate('SOLVER')
  generate('MEDIATOR')
  generate('RESOURCE_PROVIDER')
  generate('JOB_CREATOR')
  generate('DIRECTORY')
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
