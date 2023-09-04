import {
  loadFixture,
} from '@nomicfoundation/hardhat-toolbox/network-helpers'
import chai from 'chai'
import bluebird from 'bluebird'
import chaiAsPromised from 'chai-as-promised'
import { ethers } from 'hardhat'
import {
  getPaymentReason,
  getPaymentDirection,
} from '../utils/enums'
import {
  getWallet,
  getAddress,
} from '../utils/web3'
import {
  ACCOUNTS,
} from '../utils/accounts'
import {
  setupPaymentsFixture,
} from './fixtures'
import {
  LilypadToken,
} from '../typechain-types'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Payments", () => {

  const dealID = ethers.getBigInt(10)

  function setupPayments() {
    return setupPaymentsFixture({
      testMode: true,
      withFunds: true,
    })
  }

  async function getBalances(token: LilypadToken, accountName: string) {
    const tokens = await token.balanceOf(getAddress(accountName))
    const escrow = await token.escrowBalanceOf(getAddress(accountName))
    return {
      tokens,
      escrow,
    }
  }

  describe.only("Deals", () => {

    it("Should agreeResourceProvider", async function () {

      const timeoutCollateral = ethers.getBigInt(10)

      const {
        payments,
        token,
      } = await loadFixture(setupPayments)

      const tokenAddress = await token.getAddress()
      const balanceBefore = await getBalances(token, 'resource_provider')

      await expect(payments
        .connect(getWallet('resource_provider'))
        .agreeResourceProvider(
          dealID,
          getAddress('resource_provider'),
          timeoutCollateral,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('resource_provider'),
          tokenAddress,
          timeoutCollateral,
        )

      const balanceAfter = await getBalances(token, 'resource_provider')

      console.log({
        balanceBefore,
        balanceAfter
      })

      expect(balanceAfter.tokens).to.equal(balanceBefore.tokens - timeoutCollateral)
      expect(balanceAfter.escrow).to.equal(balanceBefore.escrow + timeoutCollateral)
    })

    it.only("Should agreeJobCreator", async function () {

      const payentCollateral = ethers.getBigInt(0)
      const timeoutCollateral = ethers.getBigInt(10)

      const {
        token,
        payments,
      } = await loadFixture(setupPayments)

      const tokenAddress = await token.getAddress()
      const balanceBefore = await getBalances(token, 'job_creator')

      await expect(payments
        .connect(getWallet('job_creator'))
        .agreeJobCreator(
          dealID,
          getAddress('job_creator'),
          payentCollateral,
          timeoutCollateral,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          payentCollateral,
          getPaymentReason('PaymentCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('job_creator'),
          tokenAddress,
          payentCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('job_creator'),
          tokenAddress,
          timeoutCollateral,
        )

      const balanceAfter = await getBalances(token, 'job_creator')

      expect(balanceAfter.tokens).to.equal(balanceBefore.tokens - payentCollateral - timeoutCollateral)
      expect(balanceAfter.escrow).to.equal(balanceBefore.escrow + payentCollateral + timeoutCollateral)
    })


  })

  describe("Results", () => {

    it("Should add a result", async function () {

      const resultsCollateral = ethers.getBigInt(20)
      const timeoutCollateral = ethers.getBigInt(10)

      const {
        token,
        payments,
      } = await loadFixture(setupPayments)

      await expect(payments
        .connect(getWallet('resource_provider'))
        .agreeResourceProvider(
          dealID,
          getAddress('resource_provider'),
          timeoutCollateral,
        )
      ).to.not.be.reverted

      const balanceBefore = await token.balanceOf(getAddress('resource_provider'))
      const escrowBefore = await token.escrowBalanceOf(getAddress('resource_provider'))

      await expect(payments
        .connect(getWallet('resource_provider'))
        .addResult(
          dealID,
          getAddress('resource_provider'),
          resultsCollateral,
          timeoutCollateral,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('PaidIn'),
        )

      const balanceAfter = await token.balanceOf(getAddress('resource_provider'))
      const escrowAfter = await token.escrowBalanceOf(getAddress('resource_provider'))

      expect(balanceAfter).to.equal(balanceBefore - timeoutCollateral)
      expect(escrowAfter).to.equal(escrowBefore + timeoutCollateral)
    })

  })
})