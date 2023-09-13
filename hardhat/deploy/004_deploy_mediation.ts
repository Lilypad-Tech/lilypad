import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployMediation: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  const sharedStructs = await deployments.get('SharedStructs')
  await deploy("LilypadMediationRandom", {
    from: admin,
    args: [],
    log: true,
    libraries: {
      SharedStructs: sharedStructs.address,
    },
  })
  await execute(
    'LilypadMediationRandom',
    {
      from: admin,
      log: true,
    },
    'initialize'
  )
  return true
}

deployMediation.id = 'deployMediation'

export default deployMediation