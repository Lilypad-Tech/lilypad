import {
  getControllerAddress,
  getTokenAddress,
  getMediationAddress,
  getJobManagerAddress,
} from '../utils/web3'

async function main() {
  const controllerAddress = await getControllerAddress()
  const tokenAddress = await getTokenAddress()
  const mediationAddress = await getMediationAddress()
  const jobManagerAddress = await getJobManagerAddress()
  
  console.log(`export WEB3_CONTROLLER_ADDRESS=${controllerAddress}`)
  console.log(`export WEB3_TOKEN_ADDRESS=${tokenAddress}`)
  console.log(`export WEB3_MEDIATION_ADDRESS=${mediationAddress}`)
  console.log(`export WEB3_JOBCREATOR_ADDRESS=${jobManagerAddress}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
