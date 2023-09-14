import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployUsers: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  const sharedStructs = await deployments.get('SharedStructs')
  await deploy("LilypadUsers", {
    from: admin,
    args: [],
    log: true,
    libraries: {
      SharedStructs: sharedStructs.address,
    },
  })
  await execute(
    'LilypadUsers',
    {
      from: admin,
      log: true,
    },
    'initialize'
  )
  return true
}

deployUsers.id = 'deployUsers'

export default deployUsers