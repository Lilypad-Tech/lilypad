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
  deployPayments,
  fundAllAccountsWithTokens,
  DEFAULT_TOKEN_SUPPLY,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'
import {
  ACCOUNTS,
} from '../utils/accounts'
import {
  getServiceType,
  getAgreementState,
} from '../utils/enums'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Payments", () => {

  const dealID = ethers.getBigInt(10)

  async function setupToken() {
    const admin = getWallet('admin')
    const token = await deployToken(admin, DEFAULT_TOKEN_SUPPLY)
    return token
  }

  async function setupTokenWithFunds() {
    const token = await setupToken()
    await fundAllAccountsWithTokens(token, DEFAULT_TOKENS_PER_ACCOUNT)
    return token
  }

  async function setupPayments() {
    const admin = getWallet('admin')
    const token = await setupTokenWithFunds()
    const payments = await deployPayments(admin, token.getAddress())
    return {
      token,
      payments
    }
  }

  describe("Token supply", () => {

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

  describe("Token approval", () => {

    // test that when as a job creator - we agree to a deal
    // the tokens are moved from our account to the payments contract
    // and that happens automatically because the arrangement of
    // contracts is such that the auto approval works
    it("Should auto approve the movement of tokens", async function () {

      const timeoutCollateral = ethers.getBigInt(10)

      const {
        token,
        payments,
      } = await loadFixture(setupPayments)

      // await expect(payments
      //   .connect(getWallet('admin'))
      //   .agreeResourceProvider(
      //     dealID,
      //     getAddress('resource_provider'),
      //     timeoutCollateral,
      //   )
      // )
      //   .to.emit(storage, "DealStateChange")
      //   .withArgs(dealID, getAgreementState('MediationRejected'))
      
    })

  })

  // describe("End to end", () => {

  //   it("Should run a job and payout", async function () {
  //     const {
  //       storage,
  //       token,
  //       controller,
  //     } = await loadFixture(setupContracts)


  //   })

  // })

})