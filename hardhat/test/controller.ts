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
  deployContracts,
  DEFAULT_TOKENS_PER_ACCOUNT,
} from '../utils/web3'
import {
  getServiceType,
  getAgreementState,
} from '../utils/enums'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe("Controller", () => {

  async function setupContracts() {
    return deployContracts(getWallet('admin'))
  }

  describe("Token supply", () => {

    it("Should initially fund accounts", async function () {
      const {
        token,
      } = await loadFixture(setupContracts)
      expect(await token.balanceOf(getAddress('job_creator'))).to.equal(DEFAULT_TOKENS_PER_ACCOUNT)
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