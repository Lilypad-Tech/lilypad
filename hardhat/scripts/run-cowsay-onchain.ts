import bluebird from 'bluebird'
import {
  connectToken,
  connectJobManager,
  connectExampleClient,
} from '../utils/web3'
import { ethers } from 'hardhat'

async function main() {
  // the private key of the person with tokens
  const privateKey = process.env.WEB3_PRIVATE_KEY
  if(!privateKey) {
    console.error(`WEB3_PRIVATE_KEY env variable is required`)
    process.exit(1)
  }
  // it's annoying to not be able to use argv but hardhat complains about it
  const message = process.env.MESSAGE || 'Hello World!'
  const wallet = new ethers.Wallet(privateKey).connect(ethers.provider)

  const token = await connectToken()
  const manager = await connectJobManager()

  // this is a deloyed version of ExampleClient.sol
  // this is be replaced by whatever client contract
  // has been deployed
  const client = await connectExampleClient()

  // how many tokens do jobs currently cost?
  const requiredDeposit = await manager.getRequiredDeposit()
  // what is the address of the solver who we approve tokens to be spent by?
  const solverAddress = await manager.getControllerAddress()

  console.log(`requiredDeposit: ${Number(requiredDeposit)}`)

  const paytokensTx = await token
    .connect(wallet)
    .approve(solverAddress, requiredDeposit)
  await paytokensTx.wait()

  console.log(`tokens approved: ${paytokensTx.hash}`)

  const runjobTx = await client
    .connect(wallet)
    .runCowsay(message)
  const receipt = await runjobTx.wait()
  if(!receipt) throw new Error(`no receipt`)

  console.log(`submitted job: ${runjobTx.hash}`)
  
  let jobID = 0

  receipt.logs.forEach((log: any) => {
    const logs = client.interface.parseLog(log as any)
    if(!logs) return
    jobID = Number(logs.args[0])
  })

  console.log(`Job ID: ${jobID}`)
  console.log('--------------------------------------------')
  console.log(`Waiting for job to be completed...`)

  let result = ''

  while(!result) {
    result = await client.getJobResult(jobID)
    if(!result) {
      await bluebird.delay(1000)
    }
    console.log(`waiting for job result: ${new Date().toLocaleString()}`)
  }

  console.log(`Job result: ${result}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
