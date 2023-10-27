import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'
import {
  DEFAULT_TOKEN_SUPPLY,
} from '../utils/web3'

const deployToken: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  // log the admin address
  console.log(`admin: ${admin}`)
  await deploy("LilypadToken", {
    from: admin,
    args: [
      // FIXME: change the token name and symbol
      "Lilypad Token Test",
      "LPTT",
      DEFAULT_TOKEN_SUPPLY,
    ],  
    log: true,
  })
  return true
}

deployToken.id = 'deployToken'

export default deployToken