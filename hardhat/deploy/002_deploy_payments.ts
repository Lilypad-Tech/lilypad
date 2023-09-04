import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployPayments: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  const tokenContract = await deployments.get('LilypadToken')
  await deploy("LilypadPayments", {
    from: admin,
    args: [],
    log: true,
  })
  await execute(
    'LilypadPayments',
    {
      from: admin,
      log: true,
    },
    'initialize',
    tokenContract.address
  )
  console.log('--------------------------------------------')
  console.log(tokenContract.address)
  const paymentsContract = await deployments.get('LilypadPayments')
  await execute(
    'LilypadToken',
    {
      from: admin,
      log: true,
    },
    'setControllerAddress',
    paymentsContract.address
  )
  console.log('--------------------------------------------')
  console.log(paymentsContract.address)
  return true
}

deployPayments.id = 'deployPayments'

export default deployPayments