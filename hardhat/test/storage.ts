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

    it("Should be able to update an existing user", async function () {
      const storage = await loadFixture(setupStorage)
      await addUsers(storage)

      const newCID = ethers.getBigInt(456)
      const newURL = "def"

      expect(await storage
        .connect(getWallet('job_creator'))
        .updateUser(
          newCID,
          newURL,
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
      
      const jc = await storage.getUser(getAddress('job_creator'))

      expect(jc.metadataCID).to.equal(newCID)
      expect(jc.url).to.equal(newURL)
    })

    it("Should be able to list users of a certain role", async function () {
      const storage = await loadFixture(setupStorage)
      await addUsers(storage)

      expect(await storage.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([])

      expect(await storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await storage.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([
          getAddress('job_creator'),
        ])
    })

    it("Should deny adding a user to a list without the role", async function () {
      const storage = await loadFixture(setupStorage)
      await addUsers(storage)

      await expect(storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('ResourceProvider')
        )
      ).to.be.revertedWith('User must have that role')
    })

    it("Should deny adding a user to a list twice", async function () {
      const storage = await loadFixture(setupStorage)
      await addUsers(storage)

      expect(await storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      await expect(storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is already part of that list')
    })

    it("Should allow a user to remove themselves from a list", async function () {
      const storage = await loadFixture(setupStorage)
      await addUsers(storage)

      expect(await storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await storage
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await storage.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([])
    })

    it("Should deny removing a user from a list if they have not added themselves", async function () {
      const storage = await loadFixture(setupStorage)
      await addUsers(storage)

      await expect(storage
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is not part of that list')
    })

  })

})