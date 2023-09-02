import {
  loadFixture,
} from '@nomicfoundation/hardhat-toolbox/network-helpers'
import chai from 'chai'
import bluebird from 'bluebird'
import chaiAsPromised from 'chai-as-promised'
import { ethers } from 'hardhat'
import {
  getWallet,
  getAddress,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'
import {
  ACCOUNTS,
} from '../utils/accounts'
import {
  setupTokenFixture,
} from './fixtures'

chai.use(chaiAsPromised)
const { expect } = chai

describe.only("Token", () => {

  function setupTokenWithFunds() {
    return setupTokenFixture({
      testMode: true,
      withFunds: true,
    })
  }
  function setupTokenWithoutFunds() {
    return setupTokenFixture({
      testMode: true,
      withFunds: false,
    })
  }
  
  describe("Initial Supply", () => {

    it("Should fund admin with initial supply", async function () {
      const token = await loadFixture(setupTokenWithoutFunds)
      expect(await token.balanceOf(getAddress('admin'))).to.equal(DEFAULT_TOKEN_SUPPLY)
    })

    it("Should fund all accounts with DEFAULT_TOKENS_PER_ACCOUNT", async function () {
      const token = await loadFixture(setupTokenWithFunds)
      await bluebird.mapSeries(ACCOUNTS, async (account) => {
        if(account.name === 'admin') return
        expect(await token.balanceOf(getAddress(account.name))).to.equal(DEFAULT_TOKENS_PER_ACCOUNT)
      })
    })
  })

  describe("Access control", () => {
    it("Can only run if there is a controller address set", async function () {
      const amount = ethers.getBigInt(100)
      const token = await setupTokenFixture({
        testMode: false,
        withFunds: false,
      })

      await expect(token
        .connect(getWallet('resource_provider'))
        .refundEscrow(
          getAddress('resource_provider'),
          amount
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only be run by the controller", async function () {
      const amount = ethers.getBigInt(100)
      const token = await setupTokenFixture({
        testMode: false,
        withFunds: false,
        // just something so it's set
        controllerAddress: getAddress('mediator'),
      })

      await expect(token
        .connect(getWallet('resource_provider'))
        .refundEscrow(
          getAddress('resource_provider'),
          amount
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })
  })

  describe("Escrow", () => {

    it("Should handle the end of end of running a job", async function () {
      const amount = ethers.getBigInt(100)
      const token = await loadFixture(setupTokenWithFunds)
      const tokenAddress = await token.getAddress()

      const getBalances = async () => {
        return bluebird.props({
          escrow: bluebird.props({
            job_creator: token.escrowBalanceOf(getAddress('job_creator')),
            resource_provider: token.escrowBalanceOf(getAddress('resource_provider')),
            slashed: token.escrowBalanceOf(getAddress('admin')),
          }),
          tokens: bluebird.props({
            job_creator: token.balanceOf(getAddress('job_creator')),
            resource_provider: token.balanceOf(getAddress('resource_provider')),
            escrow: token.balanceOf(tokenAddress),
            slashed: token.balanceOf(getAddress('admin')),
          })
          
        })
      }

      const balances1 = await getBalances()

      expect(await token
        .connect(getWallet('job_creator'))
        .payEscrow(amount)
      ).to.not.be.reverted

      const balances2 = await getBalances()
      expect(balances2.tokens.job_creator).to.equal(balances1.tokens.job_creator - amount)
      expect(balances2.escrow.job_creator).to.equal(balances1.escrow.job_creator + amount)
      expect(balances2.tokens.escrow).to.equal(amount)

      expect(await token
        .connect(getWallet('admin'))
        .payJob(
          getAddress('job_creator'),
          getAddress('resource_provider'),
          amount
        )
      ).to.not.be.reverted

      const balances3 = await getBalances()
      expect(balances3.tokens.resource_provider).to.equal(balances2.tokens.resource_provider + amount)
      expect(balances3.escrow.job_creator).to.equal(balances2.escrow.job_creator - amount)
      expect(balances3.tokens.escrow).to.equal(0)
    })

    

  })
})