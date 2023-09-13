import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployStorage: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  const sharedStructs = await deployments.get('SharedStructs')
  await deploy("LilypadStorage", {
    from: admin,
    args: [],
    log: true,
    libraries: {
      SharedStructs: sharedStructs.address,
    },
  })
  await execute(
    'LilypadStorage',
    {
      from: admin,
      log: true,
    },
    'initialize'
  )
  return true
}

deployStorage.id = 'deployStorage'

export default deployStorage