import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployShared: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  await deploy("SharedStructs", {
    from: admin,
    log: true,
  })
  return true
}

deployShared.id = 'deployShared'

export default deployShared