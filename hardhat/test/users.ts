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
} from '../utils/web3'
import {
  getServiceType,
  getAgreementState,
} from '../utils/enums'
import {
  setupStorageFixture,
  getDefaultTimeouts,
  getDefaultPricing,
  DEFAULT_VALUES,
  setupUsersFixture,
} from './fixtures'

import {
  SharedStructs,
} from '../typechain-types/contracts/LilypadStorage'

chai.use(chaiAsPromised)
const { expect } = chai

describe("Users", () => {

  const rpCID = "123"
  const rpURL = "abc"

  const jcCID = "123"
  const jcURL = "def"

  const dealID = "10"
  const resultsID = "11"
  const instructionCount = ethers.getBigInt(25)

  function setupUsers() {
    return setupUsersFixture({})
  }

  async function setupTestWithUsers() {
    const users = await setupUsers()

    expect(await users
      .connect(getWallet('resource_provider'))
      .updateUser(
        rpCID,
        rpURL,
        [
          getServiceType('ResourceProvider'),
        ],
      )
    ).to.not.be.reverted

    expect(await users
      .connect(getWallet('job_creator'))
      .updateUser(
        jcCID,
        jcURL,
        [
          getServiceType('JobCreator'),
        ],
      )
    ).to.not.be.reverted
    return users
  }

  describe("Users", () => {

    it("Should be able to add and get users", async function () {
      const users = await loadFixture(setupTestWithUsers)
      
      const rp = await users.getUser(getAddress('resource_provider'))
      const jc = await users.getUser(getAddress('job_creator'))

      expect(rp.metadataCID).to.equal(rpCID)
      expect(rp.url).to.equal(rpURL)
      expect(rp.roles).to.deep.equal([getServiceType('ResourceProvider')])

      expect(jc.metadataCID).to.equal(jcCID)
      expect(jc.url).to.equal(jcURL)
      expect(jc.roles).to.deep.equal([getServiceType('JobCreator')])
    })

    it("Should be able to update an existing user", async function () {
      const storage = await loadFixture(setupTestWithUsers)
      
      const newCID = "456"
      const newURL = "def"

      expect(await storage
        .connect(getWallet('job_creator'))
        .updateUser(
          newCID,
          newURL,
          [
            getServiceType('JobCreator'),
          ],
        )
      ).to.not.be.reverted
      
      const jc = await storage.getUser(getAddress('job_creator'))

      expect(jc.metadataCID).to.equal(newCID)
      expect(jc.url).to.equal(newURL)
    })

    it("Should be able to list users of a certain role", async function () {
      const users = await loadFixture(setupTestWithUsers)

      expect(await users.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([])

      expect(await users
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await users.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([
          getAddress('job_creator'),
        ])
    })

    it("Should deny adding a user to a list without the role", async function () {
      const users = await loadFixture(setupTestWithUsers)

      await expect(users
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('ResourceProvider')
        )
      ).to.be.revertedWith('User must have that role')
    })

    it("Should deny adding a user to a list twice", async function () {
      const users = await loadFixture(setupTestWithUsers)

      expect(await users
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      await expect(users
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is already part of that list')
    })

    it("Should allow a user to remove themselves from a list", async function () {
      const users = await loadFixture(setupTestWithUsers)

      expect(await users
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await users
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await users.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([])
    })

    it("Should deny removing a user from a list if they have not added themselves", async function () {
      const users = await loadFixture(setupTestWithUsers)

      await expect(users
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is not part of that list')
    })

  })

})