const hre = require('hardhat')
const chai = require('chai')
const chaiAsPromised = require('chai-as-promised')
const ethereumWaffle = require('ethereum-waffle')

chai.use(ethereumWaffle.solidity)
chai.use(chaiAsPromised)
const { expect } = chai

const {
  ethers,
} = hre

const {
  deployStorage,
} = require('./utils')

context('LilypadStorage', () => {

  let storageContract
  let adminAccount
  let resourceProviderAccount
  let jobCreatorAccount

  beforeEach(async () => {
    [
      adminAccount,
      resourceProviderAccount,
      jobCreatorAccount,
    ] = await ethers.getSigners()
    storageContract = await deployStorage(adminAccount)
  })

  describe('contract', () => {
    it('deploys', async () => {
      const testDeal = await storageContract
        .connect(jobCreatorAccount)
        .getDeal(0)
      expect(testDeal.dealId).to.equal(ethers.BigNumber.from(0))
    })

    // it('adds and gets a deal', async () => {

    //   const dealID = ethers.BigNumber.from(1)

    //   await expect(
    //     storageContract
    //       .connect(jobCreatorAccount)
    //       .addDeal("holy cow", {
    //         value: JOB_COST,
    //       })
    //   ).to.be.revertedWith('No mediators provided')

    //   const testDeal = await storageContract
    //     .connect(jobCreatorAccount)
    //     .getDeal(0)
    //   expect(testDeal.dealId).to.equal(ethers.BigNumber.from(0))
    // })
  })
})
  