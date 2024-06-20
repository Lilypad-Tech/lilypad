import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployPow: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  await deploy("LilypadPow", {
    from: admin,
    args: [],
    log: true,
  })
  await execute(
    'LilypadPow',
    {
      from: admin,
      log: true,
    },
    'initialize'
  )
  return true
}

deployPow.id = 'deployPow'

export default deployPow
