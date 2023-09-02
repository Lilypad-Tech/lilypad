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
  setupToken,
  setupTokenWithFunds,
} from './utils'

chai.use(chaiAsPromised)
const { expect } = chai

describe("Token", () => {

  describe("Initial Supply", () => {

    it("Should fund admin with initial supply", async function () {
      const token = await loadFixture(setupToken)
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

  describe("escrow", () => {

    it("Should pay in and pay out", async function () {
      const amount = ethers.getBigInt(100)
      const token = await loadFixture(setupTokenWithFunds)
      const tokenAddress = await token.getAddress()

      const getBalances = async () => {
        return bluebird.props({
          job_creator: token.balanceOf(getAddress('job_creator')),
          resource_provider: token.balanceOf(getAddress('resource_provider')),
          escrow: token.balanceOf(tokenAddress),
        })
      }

      const balances1 = await getBalances()

      expect(await token
        .connect(getWallet('job_creator'))
        .payinEscrow(amount)
      ).to.not.be.reverted

      const balances2 = await getBalances()
      expect(balances2.job_creator).to.equal(balances1.job_creator - amount)
      expect(balances2.escrow).to.equal(balances1.escrow + amount)

      expect(await token
        .connect(getWallet('admin'))
        .payoutEscrow(getAddress('resource_provider'), amount)
      ).to.not.be.reverted

      const balances3 = await getBalances()
      expect(balances3.resource_provider).to.equal(balances2.resource_provider + amount)
      expect(balances3.escrow).to.equal(balances2.escrow - amount)
    })

    it("Can only be run by the controller", async function () {
      const amount = ethers.getBigInt(100)
      const token = await loadFixture(setupToken)

      await expect(token
        .connect(getWallet('resource_provider'))
        .payoutEscrow(
          getAddress('resource_provider'),
          amount
        )
      ).to.be.revertedWith('Ownable: caller is not the owner')
    })

  })
})