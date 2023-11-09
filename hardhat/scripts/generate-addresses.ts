import Wallet from 'ethereumjs-wallet'


const generate = (name: string) => {
  const wallet = Wallet.generate()
  console.log(`export PRIVATE_KEY_${name}=${wallet.getPrivateKeyString()}`)
  console.log(`export ADDRESS_${name}=${wallet.getAddressString()}`)
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
