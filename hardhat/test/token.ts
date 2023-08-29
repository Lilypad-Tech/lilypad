import {
  time,
  loadFixture,
} from '@nomicfoundation/hardhat-toolbox/network-helpers'
import chai from 'chai'
import bluebird from 'bluebird'
import chaiAsPromised from 'chai-as-promised'
import { ethers } from 'hardhat'
import {
  getWallet,
  getAddress,
  deployToken,
  getRandomWallet,
  fundAllAccountsWithTokens,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'
import {
  ACCOUNTS,
} from '../utils/accounts'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Token", () => {

  async function getSigners() {
    const signers = await ethers.getSigners()
    const escrow = signers[signers.length-1]
    const controller = signers[signers.length-2]
    const payee = signers[signers.length-3]
    return {
      escrow,
      controller,
      payee,
    }
  }

  async function setupToken() {
    const signers = await getSigners()
    const admin = getWallet('admin')
    const token = await deployToken(
      admin,
      DEFAULT_TOKEN_SUPPLY,
      signers.escrow.address,
      signers.controller.address,
    )
    return token
  }

  async function setupTokenWithFunds() {
    const token = await setupToken()
    await fundAllAccountsWithTokens(token, DEFAULT_TOKENS_PER_ACCOUNT)
    return token
  }

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

  describe("Pay out escrow", () => {

    it("Should pay out when called from the controller", async function () {
      const amount = ethers.getBigInt(100)
      const token = await loadFixture(setupToken)
      const signers = await getSigners()

      expect(await token
        .connect(getWallet('admin'))
        .transfer(
          signers.escrow.address,
          amount
        )
      ).to.not.be.reverted

      expect(await token.balanceOf(signers.escrow.address)).to.equal(amount)

      expect(await token
        .connect(signers.controller)
        .payoutEscrow(
          signers.payee.address,
          amount
        )
      ).to.not.be.reverted

      expect(await token.balanceOf(signers.escrow.address)).to.equal(0)
      expect(await token.balanceOf(signers.payee.address)).to.equal(amount)
    })

    it("Can only be run by the controller", async function () {
      const amount = ethers.getBigInt(100)
      const token = await loadFixture(setupToken)
      const signers = await getSigners()

      await expect(token
        .connect(getWallet('admin'))
        .payoutEscrow(
          signers.payee.address,
          amount
        )
      ).to.be.revertedWith('LilypadToken: only controller can pay from escrow')
    })

  })
})