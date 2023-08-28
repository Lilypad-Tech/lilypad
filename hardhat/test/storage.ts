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

  const rpCID = ethers.getBigInt(123)
  const rpURL = "abc"

  const jcCID = ethers.getBigInt(456)
  const jcURL = "def"

  const dealID = ethers.getBigInt(10)

  async function setupStorage() {
    const admin = getWallet('admin')
    const storage = await deployStorage(admin)
    return storage
  }

  async function setupStorageWithUsers() {
    const storage = await setupStorage()
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
    return storage
  }

  async function setupStorageWithUsersAndDeal() {
    const storage = await setupStorageWithUsers()
    expect(await storage
      .connect(getWallet('admin'))
      .addDeal(
        dealID,
        getAddress('resource_provider'),
        getAddress('job_creator'),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
      )
    ).to.not.be.reverted
    return storage
  }

  describe("Deployment", () => {

    it("Should have deployed with empty state", async function () {
      const storage = await loadFixture(setupStorage)
      expect(await storage.hasDeal(ethers.getBigInt(0))).to.equal(false)
    })

  })

  describe("Users", () => {

    it("Should be able to add and get users", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
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
      const storage = await loadFixture(setupStorageWithUsers)
      
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
      const storage = await loadFixture(setupStorageWithUsers)

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
      const storage = await loadFixture(setupStorageWithUsers)

      await expect(storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('ResourceProvider')
        )
      ).to.be.revertedWith('User must have that role')
    })

    it("Should deny adding a user to a list twice", async function () {
      const storage = await loadFixture(setupStorageWithUsers)

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
      const storage = await loadFixture(setupStorageWithUsers)

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
      const storage = await loadFixture(setupStorageWithUsers)

      await expect(storage
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is not part of that list')
    })

  })

  describe("Deals", () => {

    it("Should be able to add a deal and then read it", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDeal)

      const deal = await storage.getDeal(dealID)

      expect(deal.dealId).to.equal(dealID)
      expect(deal.resourceProvider).to.equal(getAddress('resource_provider'))
      expect(deal.jobCreator).to.equal(getAddress('job_creator'))
      expect(deal.instructionPrice).to.equal(ethers.getBigInt(1))
      expect(deal.timeout).to.equal(ethers.getBigInt(1))
      expect(deal.timeoutCollateral).to.equal(ethers.getBigInt(1))
      expect(deal.jobCollateral).to.equal(ethers.getBigInt(1))
      expect(deal.resultsCollateralMultiple).to.equal(ethers.getBigInt(1))

    })

    it("Should be able to see deals for specific parties", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDeal)
      expect(await storage.getDealsForParty(getAddress('resource_provider')))
        .to.deep.equal([dealID])
      expect(await storage.getDealsForParty(getAddress('job_creator')))
        .to.deep.equal([dealID])
      expect(await storage.getDealsForParty(getAddress('mediator')))
        .to.deep.equal([])
    })

    it("Should deny adding a deal the second time", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDeal)
      
      await expect(storage
        .connect(getWallet('admin'))
        .addDeal(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
        )
      ).to.be.revertedWith('Deal already exists')
    })

    it("Should deny adding a deal as not the owner", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .addDeal(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
        )
      ).to.be.revertedWith('Ownable: caller is not the owner')
    })

  })

})