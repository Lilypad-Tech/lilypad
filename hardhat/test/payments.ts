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
  deployToken,
  deployPayments,
  DEFAULT_TOKEN_SUPPLY,
} from '../utils/web3'
import {
  getServiceType,
  getAgreementState,
} from '../utils/enums'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Payments", () => {

  async function setupToken() {
    const admin = getWallet('admin')
    const token = await deployToken(admin, DEFAULT_TOKEN_SUPPLY)
    return token
  }

  async function setupPayments() {
    const admin = getWallet('admin')
    const token = await setupToken()
    const payments = await deployPayments(admin, token.getAddress())
    return {
      token,
      payments
    }
  }

  async function setupContracts() {
    
  }

  describe("Token supply", () => {

    it("Should initially fund accounts", async function () {
      const token = await loadFixture(setupToken)
      expect(await token.balanceOf(getAddress('admin'))).to.equal(DEFAULT_TOKEN_SUPPLY)
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