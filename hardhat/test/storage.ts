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
  getAddress,
  deployStorage,
} from '../utils/web3'
import {
  getServiceType,
} from '../utils/enums'
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

  describe("Users", () => {

    it("Should be able to add and list users", async function () {
      const storage = await loadFixture(setupStorage)
      
      const rpCID = ethers.getBigInt(123)
      const rpURL = "abc"

      const jcCID = ethers.getBigInt(456)
      const jcURL = "def"

      // register the resource provider
      expect(await storage
        .connect(getWallet('resource_provider'))
        .updateUser(
          rpCID,
          rpURL,
          [
            getServiceType('ResourceProvider'),
          ],
          [
            getAddress('mediator')
          ],
          [
            getAddress('directory')
          ]
        )
      ).to.not.be.reverted

      expect(await storage
        .connect(getWallet('job_creator'))
        .updateUser(
          jcCID,
          jcURL,
          [
            getServiceType('JobCreator'),
          ],
          [
            getAddress('mediator')
          ],
          [
            getAddress('directory')
          ]
        )
      ).to.not.be.reverted

      const rp = await storage.getUser(getAddress('resource_provider'))
      const jc = await storage.getUser(getAddress('job_creator'))

      expect(rp.metadataCID).to.equal(rpCID)
      expect(rp.url).to.equal(rpURL)

      expect(jc.metadataCID).to.equal(jcCID)
      expect(jc.url).to.equal(jcURL)
    })

  })

})