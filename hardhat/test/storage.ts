import {
  time,
  loadFixture,
} from '@nomicfoundation/hardhat-toolbox/network-helpers'
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
  LilypadStorage,
} from '../typechain-types'

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

    const rpCID = ethers.getBigInt(123)
    const rpURL = "abc"

    const jcCID = ethers.getBigInt(456)
    const jcURL = "def"

    const addUsers = async (storage: LilypadStorage) => {
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
    }

    it("Should be able to add and get users", async function () {
      const storage = await loadFixture(setupStorage)
      
      await addUsers(storage)
      
      const rp = await storage.getUser(getAddress('resource_provider'))
      const jc = await storage.getUser(getAddress('job_creator'))

      expect(rp.metadataCID).to.equal(rpCID)
      expect(rp.url).to.equal(rpURL)
      expect(rp.roles).to.deep.equal([getServiceType('ResourceProvider')])
      expect(rp.trustedMediators).to.deep.equal([getAddress('mediator')])
      expect(rp.trustedDirectories).to.deep.equal([getAddress('directory')])

      expect(jc.metadataCID).to.equal(jcCID)
      expect(jc.url).to.equal(jcURL)
      expect(jc.roles).to.deep.equal([getServiceType('JobCreator')])
      expect(jc.trustedMediators).to.deep.equal([getAddress('mediator')])
      expect(jc.trustedDirectories).to.deep.equal([getAddress('directory')])
    })

  })

})