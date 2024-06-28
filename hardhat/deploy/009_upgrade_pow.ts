import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'
const { ethers, upgrades } = require("hardhat");
const updatePowV2: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
    const { deployments, getNamedAccounts } = hre
    const { deploy, execute } = deployments
    const {
      admin,
    } = await getNamedAccounts()
    const instance = await deploy("LilypadPow", {
      from: admin,
      args: [],
      log: true,
    })
    await execute(
      'LilypadController',
      {
        from: admin,
        log: true,
      },
      'setPowAddress',
      instance.address
    )
    return true;
}
updatePowV2.id = 'updatePowV2'
export default updatePowV2
