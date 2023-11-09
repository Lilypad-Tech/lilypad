import {
  getControllerAddress,
  getTokenAddress,
  getMediationAddress,
  getJobManagerAddress,
  getPaymentsAddress,
  getStorageAddress,
  getUsersAddress
} from '../utils/web3'

async function main() {
  const controllerAddress = await getControllerAddress()
  const tokenAddress = await getTokenAddress()
  const mediationAddress = await getMediationAddress()
  const jobManagerAddress = await getJobManagerAddress()
  const paymentsAddress = await getPaymentsAddress()
  const storageAddress = await getStorageAddress()
  const usersAddress = await getUsersAddress()
  
  console.log(`export WEB3_RPC_URL=ws://localhost:8546`)
  console.log(`export WEB3_CONTROLLER_ADDRESS=${controllerAddress}`)
  console.log(`export WEB3_TOKEN_ADDRESS=${tokenAddress}`)
  console.log(`export WEB3_MEDIATION_ADDRESS=${mediationAddress}`)
  console.log(`export WEB3_JOBCREATOR_ADDRESS=${jobManagerAddress}`)
  console.log(`export WEB3_PAYMENTS_ADDRESS=${paymentsAddress}`)
  console.log(`export WEB3_STORAGE_ADDRESS=${storageAddress}`)
  console.log(`export WEB3_USERS_ADDRESS=${usersAddress}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
