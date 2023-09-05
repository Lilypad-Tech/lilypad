import {
  getTokenAddress,
  getPaymentsAddress,
  getStorageAddress,
  getControllerAddress,
} from '../utils/web3'

async function main() {
  const tokenAddress = await getTokenAddress()
  const paymentsAddress = await getPaymentsAddress()
  const storageAddress = await getStorageAddress()
  const controllerAddress = await getControllerAddress()
  
  console.log(`export WEB3_CONTROLLER_ADDRESS=${controllerAddress}`)
  console.log(`export WEB3_PAYMENTS_ADDRESS=${paymentsAddress}`)
  console.log(`export WEB3_STORAGE_ADDRESS=${storageAddress}`)
  console.log(`export WEB3_TOKEN_ADDRESS=${tokenAddress}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
