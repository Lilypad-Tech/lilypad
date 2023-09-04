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
  getAgreementState,
} from '../utils/enums'
import {
  getWallet,
  getAddress,
} from '../utils/web3'
import {
  ACCOUNTS,
} from '../utils/accounts'
import {
  setupControllerFixture,
} from './fixtures'
import {
  LilypadToken,
  LilypadStorage,
  LilypadController,
} from '../typechain-types'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Controller", () => {

  const dealID = ethers.getBigInt(10)
  const resultsID = ethers.getBigInt(11)
  const instructionPrice = ethers.getBigInt(10)
  const instructionCount = ethers.getBigInt(1)
  const timeout = ethers.getBigInt(100)
  const timeoutCollateral = ethers.getBigInt(10)
  const resultsCollateralMultiple = ethers.getBigInt(4)
  const resultsCollateral = ethers.getBigInt(40)
  const paymentCollateral = ethers.getBigInt(30)
  const jobCost = ethers.getBigInt(10)
  const mediationFee = ethers.getBigInt(5)

  async function getBalances(token: LilypadToken, accountName: string) {
    const tokens = await token.balanceOf(getAddress(accountName))
    const escrow = await token.escrowBalanceOf(getAddress(accountName))
    return {
      tokens,
      escrow,
    }
  }

  async function checkDeal(storage: LilypadStorage, party: string) {
    const deal = await storage.getDeal(dealID)

    expect(deal.dealId).to.equal(dealID)
    expect(deal.resourceProvider).to.equal(getAddress('resource_provider'))
    expect(deal.jobCreator).to.equal(getAddress('job_creator'))
    expect(deal.instructionPrice).to.equal(instructionPrice)
    expect(deal.timeout).to.equal(timeout)
    expect(deal.timeoutCollateral).to.equal(timeoutCollateral)
    expect(deal.paymentCollateral).to.equal(paymentCollateral)
    expect(deal.resultsCollateralMultiple).to.equal(resultsCollateralMultiple)
    expect(deal.mediationFee).to.equal(mediationFee)

    expect(await storage.hasDeal(dealID))
      .to.equal(true)

    expect(await storage.getDealsForParty(getAddress(party)))
      .to.deep.equal([dealID])
  }

  async function checkAgreement(storage: LilypadStorage, desiredState: string) {
    const agreement = await storage.getAgreement(dealID)
    expect(agreement.state).to.equal(getAgreementState(desiredState))
  }

  async function agree(controller: LilypadController, party: string) {
    return controller
      .connect(getWallet(party))
      .agree(
        dealID,
        getAddress('resource_provider'),
        getAddress('job_creator'),
        instructionPrice,
        timeout,
        timeoutCollateral,
        paymentCollateral,
        resultsCollateralMultiple,
        mediationFee,
      )
  }

  async function setupController() {
    const {
      token,
      payments,
      storage,
      controller,
    } = await setupControllerFixture({
      withFunds: true,
    })
    const tokenAddress = await token.getAddress()
    return {
      token,
      payments,
      storage,
      controller,
      tokenAddress,
    }
  }

  async function setupControllerWithDeal() {
    const ret = await setupController()
    await agree(ret.controller, 'job_creator')
    await agree(ret.controller, 'resource_provider')
    return ret
  }

  async function setupControllerWithResults() {
    const ret = await setupControllerWithDeal()
    await ret.controller
      .connect(getWallet('resource_provider'))
      .addResult(
        dealID,
        resultsID,
        instructionCount
      )
    return ret
  }

  describe("Deals", () => {

    it("Should agreeResourceProvider", async function () {
      const {
        token,
        payments,
        storage,
        controller,
        tokenAddress,
      } = await loadFixture(setupController)

      const balancesBeforeAgreeRP = await getBalances(token, 'resource_provider')

      await expect(
        agree(controller, 'resource_provider')
      )
        .to.emit(controller, 'ResourceProviderAgreed')
        .withArgs(
          dealID,
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

      const balancesAfterAgreeRP = await getBalances(token, 'resource_provider')

      expect(balancesAfterAgreeRP.tokens).to.equal(balancesBeforeAgreeRP.tokens - timeoutCollateral)
      expect(balancesAfterAgreeRP.escrow).to.equal(balancesBeforeAgreeRP.escrow + timeoutCollateral)

      await checkDeal(storage, 'resource_provider')      
      await checkAgreement(storage, 'DealNegotiating')
    })

    it("Should agreeJobCreator", async function () {
      const {
        token,
        payments,
        storage,
        controller,
        tokenAddress,
      } = await loadFixture(setupController)

      const balancesBeforeAgreeJC = await getBalances(token, 'job_creator')

      await expect(
        agree(controller, 'job_creator')
      )
        .to.emit(controller, 'JobCreatorAgreed')
        .withArgs(
          dealID,
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          dealID,
          getAddress('job_creator'),
          paymentCollateral,
          getPaymentReason('PaymentCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('job_creator'),
          tokenAddress,
          timeoutCollateral,
        )
        .to.emit(token, 'Transfer')
        .withArgs(
          getAddress('job_creator'),
          tokenAddress,
          paymentCollateral,
        )

      const balancesAfterAgreeJC = await getBalances(token, 'job_creator')

      expect(balancesAfterAgreeJC.tokens).to.equal(balancesBeforeAgreeJC.tokens - timeoutCollateral - paymentCollateral)
      expect(balancesAfterAgreeJC.escrow).to.equal(balancesBeforeAgreeJC.escrow + timeoutCollateral + paymentCollateral)

      await checkDeal(storage, 'job_creator')
      await checkAgreement(storage, 'DealNegotiating')
    })

    it("Should agree both sides", async function () {
      const {
        storage,
        controller,
      } = await loadFixture(setupController)

      await expect(
        agree(controller, 'job_creator')
      )
        .to.emit(controller, 'JobCreatorAgreed')
        .withArgs(
          dealID,
        )

      await expect(
        agree(controller, 'resource_provider')
      )
        .to.emit(controller, 'ResourceProviderAgreed')
        .withArgs(
          dealID,
        )
        .to.emit(controller, 'DealAgreed')
        .withArgs(
          dealID,
        )
        
      await checkAgreement(storage, 'DealAgreed')
    })
  })

  describe("Results", () => {
    it("Post results as RP", async function () {
      const {
        token,
        tokenAddress,
        payments,
        storage,
        controller,
      } = await loadFixture(setupControllerWithDeal)

      const balancesBeforeRP = await getBalances(token, 'resource_provider')

      await expect(controller
        .connect(getWallet('resource_provider'))
        .addResult(
          dealID,
          resultsID,
          instructionCount
        )
      )
        .to.emit(controller, 'ResultAdded')
        .withArgs(
          dealID,
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

      const balancesAfterRP = await getBalances(token, 'resource_provider')

      expect(balancesAfterRP.tokens).to.equal(balancesBeforeRP.tokens + timeoutCollateral - resultsCollateral)
      expect(balancesAfterRP.escrow).to.equal(balancesBeforeRP.escrow - timeoutCollateral + resultsCollateral)
      await checkAgreement(storage, 'ResultsSubmitted')
    })

    it("Accepts results as JC", async function () {
      const {
        token,
        tokenAddress,
        payments,
        storage,
        controller,
      } = await loadFixture(setupControllerWithResults)

      const balancesBeforeJC = await getBalances(token, 'job_creator')
      const balancesBeforeRP = await getBalances(token, 'resource_provider')

      await expect(controller
        .connect(getWallet('job_creator'))
        .acceptResult(
          dealID,
        )
      )
        .to.emit(controller, 'ResultAccepted')
        .withArgs(
          dealID,
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
      const balancesAfterJC = await getBalances(token, 'job_creator')
      const balancesAfterRP = await getBalances(token, 'resource_provider')

      expect(balancesAfterJC.tokens).to.equal(balancesBeforeJC.tokens + (paymentCollateral - jobCost) + timeoutCollateral)
      expect(balancesAfterJC.escrow).to.equal(balancesBeforeJC.escrow - timeoutCollateral - paymentCollateral)
      expect(balancesAfterRP.tokens).to.equal(balancesBeforeRP.tokens + jobCost + resultsCollateral)
      expect(balancesAfterRP.escrow).to.equal(balancesBeforeRP.escrow - resultsCollateral)

      await checkAgreement(storage, 'ResultsAccepted')
    })
  })

  describe.only("End to end", () => {

    // TODO: we need a lot more tests for the unhappy path
    it("Runs a job in the happy path", async function () {
      const {
        token,
        tokenAddress,
        payments,
        storage,
        controller,
      } = await loadFixture(setupController)

      const balancesBeforeJC = await getBalances(token, 'job_creator')
      const balancesBeforeRP = await getBalances(token, 'resource_provider')

      await agree(controller, 'job_creator')
      await agree(controller, 'resource_provider')
      await controller
        .connect(getWallet('resource_provider'))
        .addResult(
          dealID,
          resultsID,
          instructionCount
        )
      await controller
        .connect(getWallet('job_creator'))
        .acceptResult(
          dealID,
        )

      const balancesAfterJC = await getBalances(token, 'job_creator')
      const balancesAfterRP = await getBalances(token, 'resource_provider')

      expect(balancesAfterRP.tokens).to.equal(balancesBeforeRP.tokens + jobCost)
      expect(balancesAfterJC.tokens).to.equal(balancesBeforeJC.tokens - jobCost)
      
      await checkAgreement(storage, 'ResultsAccepted')
    })

  })
})