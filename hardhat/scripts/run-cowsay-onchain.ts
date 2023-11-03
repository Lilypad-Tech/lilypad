import bluebird from 'bluebird'
import {
  connectToken,
  connectJobManager,
  connectExampleClient,
  getWallet,
  getAddress,
} from '../utils/web3'
import { ethers } from 'hardhat'

async function main() {
  // it's annoying to not be able to use argv but hardhat complains about it
  const message = process.env.MESSAGE || 'Hello World!'

  const token = await connectToken()
  const manager = await connectJobManager()
  const client = await connectExampleClient()

  const setRequiredDepositTx = await manager
    .connect(getWallet('solver'))
    .setRequiredDeposit(ethers.parseEther("2"))
  await setRequiredDepositTx.wait()

  const requiredDeposit = await manager.getRequiredDeposit()

  console.log(`requiredDeposit: ${Number(requiredDeposit)}`)

  const paytokensTx = await token
    .connect(getWallet('job_creator'))
    .approve(getAddress('solver'), requiredDeposit)
  await paytokensTx.wait()

  console.log(`tokens approved: ${paytokensTx.hash}`)

  const runjobTx = await client
    .connect(getWallet('job_creator'))
    .runCowsay(message)
  const receipt = await runjobTx.wait()
  if(!receipt) throw new Error(`no receipt`)

  console.log(`submitted job: ${runjobTx.hash}`)
  
  let jobID = 0

  receipt.logs.forEach(log => {
    const logs = client.interface.parseLog(log as any)
    if(!logs) return
    jobID = Number(logs.args[0])
  })

  console.log(`Job ID: ${jobID}`)
  console.log(`Waiting for job to be completed...`)

  let result = ''

  while(!result) {
    result = await client.getJobResult(jobID)
    if(!result) {
      await bluebird.delay(1000)
    }
  }

  console.log(`Job result: ${result}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
