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

describe("Payments", () => {

  const dealID = ethers.getBigInt(10)
  const timeoutCollateral = ethers.getBigInt(10)
  const resultsCollateral = ethers.getBigInt(40)
  const paymentCollateral = ethers.getBigInt(30)
  const jobCost = ethers.getBigInt(20)
  const mediationFee = ethers.getBigInt(5)

  async function setupPayments() {
    const {
      payments,
      token,
    } = await setupPaymentsFixture({
      testMode: true,
      withFunds: true,
    })
    const tokenAddress = await token.getAddress()
    return {
      payments,
      token,
      tokenAddress,
    }
  }

  // get's the escrow setup to the stage that the agreement has been made
  // and now we are waiting for the results
  async function setupPaymentsWithAgreement() {
    const {
      payments,
      token,
      tokenAddress,
    } = await setupPayments()

    await payments
      .connect(getWallet('resource_provider'))
      .agreeResourceProvider(
        dealID,
        getAddress('resource_provider'),
        timeoutCollateral,
      )

    await payments
      .connect(getWallet('job_creator'))
      .agreeJobCreator(
        dealID,
        getAddress('job_creator'),
        paymentCollateral,
        timeoutCollateral,
      )
  
    return {
      payments,
      token,
      tokenAddress,
    }
  }

  async function setupPaymentsWithResults() {
    const {
      payments,
      token,
      tokenAddress,
    } = await setupPaymentsWithAgreement()

    await payments
      .connect(getWallet('resource_provider'))
      .addResult(
        dealID,
        getAddress('resource_provider'),
        resultsCollateral,
        timeoutCollateral,
      )

    return {
      payments,
      token,
      tokenAddress,
    }
  }

  async function setupPaymentsWithResultsCheck() {
    const {
      payments,
      token,
      tokenAddress,
    } = await setupPaymentsWithResults()

    await payments
      .connect(getWallet('job_creator'))
      .checkResult(
        dealID,
        getAddress('job_creator'),
        timeoutCollateral,
        mediationFee,
      )

    return {
      payments,
      token,
      tokenAddress,
    }
  }

  async function getBalances(token: LilypadToken, accountName: string) {
    const tokens = await token.balanceOf(getAddress(accountName))
    const escrow = await token.escrowBalanceOf(getAddress(accountName))
    return {
      tokens,
      escrow,
    }
  }

  describe("Deals", () => {

    it("Should agreeResourceProvider", async function () {
      const {
        payments,
        token,
        tokenAddress,
      } = await loadFixture(setupPayments)

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

      expect(balanceAfter.tokens).to.equal(balanceBefore.tokens - timeoutCollateral)
      expect(balanceAfter.escrow).to.equal(balanceBefore.escrow + timeoutCollateral)
    })

    it("Should agreeJobCreator", async function () {
      const {
        token,
        payments,
        tokenAddress,
      } = await loadFixture(setupPayments)

      const balanceBefore = await getBalances(token, 'job_creator')

      await expect(payments
        .connect(getWallet('job_creator'))
        .agreeJobCreator(
          dealID,
          getAddress('job_creator'),
          paymentCollateral,
          timeoutCollateral,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          paymentCollateral,
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
          paymentCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('job_creator'),
          tokenAddress,
          timeoutCollateral,
        )

      const balanceAfter = await getBalances(token, 'job_creator')

      expect(balanceAfter.tokens).to.equal(balanceBefore.tokens - paymentCollateral - timeoutCollateral)
      expect(balanceAfter.escrow).to.equal(balanceBefore.escrow + paymentCollateral + timeoutCollateral)
    })
  })

  describe("Results", () => {

    it("Should add a result", async function () {
      const {
        token,
        payments,
        tokenAddress,
      } = await loadFixture(setupPaymentsWithAgreement)

      const balanceBefore = await getBalances(token, 'resource_provider')

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
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          resultsCollateral,
          getPaymentReason('ResultsCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('resource_provider'),
          tokenAddress,
          resultsCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('resource_provider'),
          timeoutCollateral,
        )

      const balanceAfter = await getBalances(token, 'resource_provider')

      expect(balanceAfter.tokens).to.equal(balanceBefore.tokens + timeoutCollateral - resultsCollateral)
      expect(balanceAfter.escrow).to.equal(balanceBefore.escrow - timeoutCollateral + resultsCollateral)
    })

    it("Should accept a result", async function () {
      const {
        token,
        payments,
        tokenAddress,
      } = await loadFixture(setupPaymentsWithResults)

      const balanceBeforeJC = await getBalances(token, 'job_creator')
      const balanceBeforeRP = await getBalances(token, 'resource_provider')

      await expect(payments
        .connect(getWallet('job_creator'))
        .acceptResult(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          jobCost,
          paymentCollateral,
          resultsCollateral,
          timeoutCollateral,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          jobCost,
          getPaymentReason('JobPayment'),
          getPaymentDirection('PaidOut'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          resultsCollateral,
          getPaymentReason('ResultsCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          paymentCollateral - jobCost,
          getPaymentReason('PaymentCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('resource_provider'),
          jobCost,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('resource_provider'),
          resultsCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('job_creator'),
          paymentCollateral - jobCost,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('job_creator'),
          timeoutCollateral,
        )

      const balanceAfterJC = await getBalances(token, 'job_creator')
      const balanceAfterRP = await getBalances(token, 'resource_provider')

      expect(balanceAfterJC.tokens).to.equal(balanceBeforeJC.tokens + (paymentCollateral - jobCost) + timeoutCollateral)
      expect(balanceAfterJC.escrow).to.equal(balanceBeforeJC.escrow - timeoutCollateral - paymentCollateral)
      expect(balanceAfterRP.tokens).to.equal(balanceBeforeRP.tokens + jobCost + resultsCollateral)
      expect(balanceAfterRP.escrow).to.equal(balanceBeforeRP.escrow - resultsCollateral)
    })

    it("Should check a result", async function () {
      const {
        token,
        payments,
        tokenAddress,
      } = await loadFixture(setupPaymentsWithResults)

      const balanceBeforeJC = await getBalances(token, 'job_creator')

      await expect(payments
        .connect(getWallet('job_creator'))
        .checkResult(
          dealID,
          getAddress('job_creator'),
          timeoutCollateral,
          mediationFee,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          mediationFee,
          getPaymentReason('MediationFee'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('job_creator'),
          timeoutCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('job_creator'),
          tokenAddress,
          mediationFee,
        )
        
      const balanceAfterJC = await getBalances(token, 'job_creator')

      expect(balanceAfterJC.tokens).to.equal(balanceBeforeJC.tokens + timeoutCollateral - mediationFee)
      expect(balanceAfterJC.escrow).to.equal(balanceBeforeJC.escrow - timeoutCollateral + mediationFee)
    })

  })

  describe.only("Mediation", () => {

    it("Should accept mediation results", async function () {
      const {
        token,
        payments,
        tokenAddress,
      } = await loadFixture(setupPaymentsWithResultsCheck)

      const balanceBeforeJC = await getBalances(token, 'job_creator')
      const balanceBeforeRP = await getBalances(token, 'resource_provider')
      const balanceBeforeMediator = await getBalances(token, 'mediator')

      await expect(payments
        .connect(getWallet('mediator'))
        .mediationAcceptResult(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          getAddress('mediator'),
          jobCost,
          paymentCollateral,
          resultsCollateral,
          mediationFee,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          jobCost,
          getPaymentReason('JobPayment'),
          getPaymentDirection('PaidOut'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          resultsCollateral,
          getPaymentReason('ResultsCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          paymentCollateral - jobCost,
          getPaymentReason('PaymentCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('mediator'),
          mediationFee,
          getPaymentReason('MediationFee'),
          getPaymentDirection('PaidOut'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('resource_provider'),
          jobCost,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('resource_provider'),
          resultsCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('job_creator'),
          paymentCollateral - jobCost,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('mediator'),
          mediationFee,
        )

      const balanceAfterJC = await getBalances(token, 'job_creator')
      const balanceAfterRP = await getBalances(token, 'resource_provider')
      const balanceAfterMediator = await getBalances(token, 'mediator')

      expect(balanceAfterJC.tokens).to.equal(balanceBeforeJC.tokens + (paymentCollateral - jobCost))
      expect(balanceAfterJC.escrow).to.equal(balanceBeforeJC.escrow - paymentCollateral - mediationFee)
      expect(balanceAfterRP.tokens).to.equal(balanceBeforeRP.tokens + jobCost + resultsCollateral)
      expect(balanceAfterRP.escrow).to.equal(balanceBeforeRP.escrow - resultsCollateral)
      expect(balanceAfterMediator.tokens).to.equal(balanceBeforeMediator.tokens + mediationFee)
    })

    it.only("Should reject mediation results", async function () {
      const {
        token,
        payments,
        tokenAddress,
      } = await loadFixture(setupPaymentsWithResultsCheck)

      const balanceBeforeJC = await getBalances(token, 'job_creator')
      const balanceBeforeRP = await getBalances(token, 'resource_provider')
      const balanceBeforeMediator = await getBalances(token, 'mediator')

      await expect(payments
        .connect(getWallet('mediator'))
        .mediationRejectResult(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          getAddress('mediator'),
          paymentCollateral,
          resultsCollateral,
          mediationFee,
        )
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('resource_provider'),
          resultsCollateral,
          getPaymentReason('ResultsCollateral'),
          getPaymentDirection('Slashed'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          paymentCollateral,
          getPaymentReason('PaymentCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('mediator'),
          mediationFee,
          getPaymentReason('MediationFee'),
          getPaymentDirection('PaidOut'),
        )
        // this is the RP results collateral being slashed
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('admin'),
          resultsCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('job_creator'),
          paymentCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          tokenAddress,
          getAddress('mediator'),
          mediationFee,
        )

      const balanceAfterJC = await getBalances(token, 'job_creator')
      const balanceAfterRP = await getBalances(token, 'resource_provider')
      const balanceAfterMediator = await getBalances(token, 'mediator')

      expect(balanceAfterJC.tokens).to.equal(balanceBeforeJC.tokens + paymentCollateral)
      expect(balanceAfterJC.escrow).to.equal(balanceBeforeJC.escrow - paymentCollateral - mediationFee)
      expect(balanceAfterRP.tokens).to.equal(balanceBeforeRP.tokens)
      expect(balanceAfterRP.escrow).to.equal(balanceBeforeRP.escrow - resultsCollateral)
      expect(balanceAfterMediator.tokens).to.equal(balanceBeforeMediator.tokens + mediationFee)
    })

  })
})