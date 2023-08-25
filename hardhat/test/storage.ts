import {
  time,
  loadFixture,
} from '@nomicfoundation/hardhat-toolbox/network-helpers'
import { anyValue } from '@nomicfoundation/hardhat-chai-matchers/withArgs'
import chai from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { ethers } from 'hardhat'
import {
  getWallet,
  deployStorage,
} from '../utils/contracts'
import {
  SharedStructs,  
} from '../typechain-types/contracts/LilypadStorage'

chai.use(chaiAsPromised)
const { expect } = chai

describe("Storage", () => {

  async function setupStorage() {
    const admin = getWallet('admin')
    const storage = await deployStorage(admin)
    return storage
  }

  describe("Deployment", () => {

    it("Should have deployed with empty state", async function () {
      const storage = await loadFixture(setupStorage)
      expect(await storage.hasDeal(ethers.getBigInt(0))).to.equal(false)
    })

  })

})