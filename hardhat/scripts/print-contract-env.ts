import {
  getControllerAddress,
  getTokenAddress,
} from '../utils/web3'

async function main() {
  const controllerAddress = await getControllerAddress()
  const tokenAddress = await getTokenAddress()
  console.log(`export WEB3_CONTROLLER_ADDRESS=${controllerAddress}`)
  console.log(`export WEB3_TOKEN_ADDRESS=${tokenAddress}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
