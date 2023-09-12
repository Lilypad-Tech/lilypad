import {
  getControllerAddress,
} from '../utils/web3'

async function main() {
  const controllerAddress = await getControllerAddress()
  console.log(`export WEB3_CONTROLLER_ADDRESS=${controllerAddress}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
