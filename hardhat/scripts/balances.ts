import { ethers } from 'hardhat'
import bluebird from 'bluebird'
import {
  ACCOUNTS,
} from '../utils/accounts'
import {
  connectToken,
} from '../utils/web3'
import {
  LilypadToken,
} from '../typechain-types'

async function main() {
  let token: LilypadToken
  
  try {
    token = await connectToken()
  } catch(e) {}
  
  await bluebird.mapSeries(ACCOUNTS, async (account) => {
    const balanceEther = await ethers.provider.getBalance(account.address)
    let balanceTokens = ethers.getBigInt(0)
    if(token) {
      try {
        balanceTokens = await token.balanceOf(account.address)
      } catch(e) {}
    }
    console.log(`
${account.name} - ${account.address}
  * ETHER: ${ethers.formatEther(balanceEther)}
  * TOKENS: ${balanceTokens}`)
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
