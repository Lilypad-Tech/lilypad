/*
#dev setup 
install open zeplin upgrade plugin (2.5?  latest version not compatible)

#for all contracts
Add OnableUpgradable from Open Zeplin to all contracts
move property initializers to initialize

#hardhat migration initial setup
Deploy create proxy
Get Contract from factory
UpdateProxy with factory contract 

#hardhat migration updates
get proxy address for each contract from lilypad controller
get contract from factory
update proxy with factory contract
*/
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'
import { ethers,upgrades  } from 'hardhat'
const upgradePowV2: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
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
    const LilypadPowContract = await ethers.getContractFactory("LilypadPow");
    const proxy = await upgrades.deployProxy(LilypadPowContract, {initializer: 'initialize'});
    await execute(
        'LilypadController',
        {
          from: admin,
          log: true,
        },
        'setPowAddress',
        proxy.target
      )
      return true  
}
upgradePowV2.id = 'upgradePowV2'
export default upgradePowV2