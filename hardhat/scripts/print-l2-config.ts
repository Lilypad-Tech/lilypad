import {
    ACCOUNTS,
  } from '../utils/accounts'
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

    // console.log(`export WEB3_RPC_URL=ws://localhost:8548`)
    console.log('#################################################################################################')
    console.log(`export WEB3_CONTROLLER_ADDRESS=${controllerAddress}`)
    console.log(`export WEB3_TOKEN_ADDRESS=${tokenAddress}`)
    console.log(`export WEB3_MEDIATION_ADDRESS=${mediationAddress}`)
    console.log(`export WEB3_JOBCREATOR_ADDRESS=${jobManagerAddress}`)
    console.log(`export WEB3_PAYMENTS_ADDRESS=${paymentsAddress}`)
    console.log(`export WEB3_STORAGE_ADDRESS=${storageAddress}`)
    console.log(`export WEB3_USERS_ADDRESS=${usersAddress}`)


    // loop over accounts and print env for address and private key
    ACCOUNTS.forEach((account) => {
      console.log(`export ${account.name.toUpperCase()}_ADDRESS=${account.address}`)
      console.log(`export ${account.name.toUpperCase()}_PRIVATE_KEY=${account.privateKey}`)
    })

    console.log('export SERVER_PORT=8080')
    console.log('export SERVER_URL=http://localhost:8080')
    console.log('export API_HOST=http://localhost:8002/')
    console.log('export INFURA_KEY=')
    console.log('export LOG_LEVEL=debug')
    console.log('export NETWORK=dev')
    console.log('export WEB3_RPC_URL=ws://localhost:8548')
    console.log('export WEB3_CHAIN_ID=412346')
    console.log('#################################################################################################')
    console.log("Optional: append the exports above to ~/.bashrc or ~/.bash_profile and run 'source ~/.bashrc' or 'source ~/.bash_profile' to set the environment variables.")
    // console.log("Example: code ~/.bashrc")
  }

  main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
