import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployController: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  await deploy("LilypadController", {
    from: admin,
    args: [],
    log: true,
  })

  const storageContract = await deployments.get('LilypadStorage')
  const paymentsContract = await deployments.get('LilypadPayments')
  
  await execute(
    'LilypadController',
    {
      from: admin,
      log: true,
    },
    'initialize',
    storageContract.address,
    paymentsContract.address,
  )
  return true
}

deployController.id = 'deployController'

export default deployController