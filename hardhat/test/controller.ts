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
  getDefaultTimeouts,
  getDefaultPricing,
  DEFAULT_VALUES,
  DEAL_ID,
  RESULTS_ID,
  DATA_ID,
} from './fixtures'
import {
  LilypadToken,
  LilypadStorage,
  LilypadController,
} from '../typechain-types'
import {
  SharedStructs,
} from '../typechain-types/contracts/LilypadStorage'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe("Controller", () => {
  const {
    instructionPrice,
    instructionCount,
    resultsCollateralMultiple,
    resultsCollateral,
    paymentCollateral,
    jobCost,
    mediationFee,
    timeout,
    timeoutCollateral,
  } = DEFAULT_VALUES

  async function getBalances(token: LilypadToken, accountName: string) {
    const tokens = await token.balanceOf(getAddress(accountName))
    const escrow = await token.escrowBalanceOf(getAddress(accountName))
    return {
      tokens,
      escrow,
    }
  }

  async function checkDeal(storage: LilypadStorage, party: string) {
    const deal = await storage.getDeal(DEAL_ID)

    expect(deal.dealId).to.equal(DEAL_ID)
    expect(deal.members.resourceProvider).to.equal(getAddress('resource_provider'))
    expect(deal.members.jobCreator).to.equal(getAddress('job_creator'))
    expect(deal.pricing.instructionPrice).to.equal(instructionPrice)
    expect(deal.pricing.paymentCollateral).to.equal(paymentCollateral)
    expect(deal.pricing.resultsCollateralMultiple).to.equal(resultsCollateralMultiple)
    expect(deal.pricing.mediationFee).to.equal(mediationFee)

    expect(deal.timeouts.agree.timeout).to.equal(timeout)
    expect(deal.timeouts.agree.collateral).to.equal(ethers.parseEther("0"))
    expect(deal.timeouts.submitResults.timeout).to.equal(timeout)
    expect(deal.timeouts.submitResults.collateral).to.equal(timeoutCollateral)
    expect(deal.timeouts.judgeResults.timeout).to.equal(timeout)
    expect(deal.timeouts.judgeResults.collateral).to.equal(timeoutCollateral)
    expect(deal.timeouts.mediateResults.timeout).to.equal(timeout)
    expect(deal.timeouts.mediateResults.collateral).to.equal(ethers.parseEther("0"))
    
    expect(await storage.hasDeal(DEAL_ID))
      .to.equal(true)

    expect(await storage.getDealsForParty(getAddress(party)))
      .to.deep.equal([DEAL_ID])
  }

  async function checkAgreement(storage: LilypadStorage, desiredState: string) {
    const agreement = await storage.getAgreement(DEAL_ID)
    expect(agreement.state).to.equal(getAgreementState(desiredState))
  }

  async function agree(controller: LilypadController, party: string) {

    const members: SharedStructs.DealMembersStruct = {
      solver: getAddress('solver'), 
      jobCreator: getAddress('job_creator'),
      resourceProvider: getAddress('resource_provider'),
      mediators: [getAddress('mediator')],
    }
    const timeouts = getDefaultTimeouts()
    const pricing = getDefaultPricing()

    return controller
      .connect(getWallet(party))
      .agree(
        DEAL_ID,
        members,
        timeouts,
        pricing,
      )
  }

  async function setupController() {
    const {
      token,
      payments,
      storage,
      users,
      mediation,
      controller,
    } = await setupControllerFixture({
      withFunds: true,
    })
    const tokenAddress = await token.getAddress()
    return {
      token,
      payments,
      storage,
      users,
      mediation,
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
        DEAL_ID,
        RESULTS_ID,
        DATA_ID,
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
        users,
        controller,
        tokenAddress,
      } = await loadFixture(setupController)

      const balancesBeforeAgreeRP = await getBalances(token, 'resource_provider')

      await expect(
        agree(controller, 'resource_provider')
      )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
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
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
          getAddress('job_creator'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('PaidIn'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
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
      ).to.not.be.reverted

      await expect(
        agree(controller, 'resource_provider')
      )
        .to.emit(storage, 'DealStateChange')
        .withArgs(
          DEAL_ID,
          getAgreementState('DealAgreed')
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
          DEAL_ID,
          RESULTS_ID,
          DATA_ID,
          instructionCount
        )
      )
        .to.emit(storage, 'DealStateChange')
        .withArgs(
          DEAL_ID,
          getAgreementState('ResultsSubmitted')
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
          getAddress('resource_provider'),
          timeoutCollateral,
          getPaymentReason('TimeoutCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
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
          DEAL_ID,
        )
      )
        .to.emit(storage, 'DealStateChange')
        .withArgs(
          DEAL_ID,
          getAgreementState('ResultsAccepted')
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
          getAddress('resource_provider'),
          jobCost,
          getPaymentReason('JobPayment'),
          getPaymentDirection('PaidOut'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
          getAddress('resource_provider'),
          resultsCollateral,
          getPaymentReason('ResultsCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
          getAddress('job_creator'),
          paymentCollateral - jobCost,
          getPaymentReason('PaymentCollateral'),
          getPaymentDirection('Refunded'),
        )
        .to.emit(payments, 'Payment')
        .withArgs(
          DEAL_ID,
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

  describe("End to end", () => {

    it("Runs a job in the happy path", async function () {
      const {
        token,
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
          DEAL_ID,
          RESULTS_ID,
          DATA_ID,
          instructionCount
        )
      await controller
        .connect(getWallet('job_creator'))
        .acceptResult(
          DEAL_ID,
        )

      const balancesAfterJC = await getBalances(token, 'job_creator')
      const balancesAfterRP = await getBalances(token, 'resource_provider')

      expect(balancesAfterRP.tokens).to.equal(balancesBeforeRP.tokens + jobCost)
      expect(balancesAfterJC.tokens).to.equal(balancesBeforeJC.tokens - jobCost)
      
      await checkAgreement(storage, 'ResultsAccepted')
    })

    it("Runs a job in the mediation OK path", async function () {
      const {
        token,
        storage,
        users,
        controller,
        mediation,
      } = await loadFixture(setupController)

      const balancesBeforeJC = await getBalances(token, 'job_creator')
      const balancesBeforeRP = await getBalances(token, 'resource_provider')
      const balancesBeforeMediator = await getBalances(token, 'mediator')

      await users
        .connect(getWallet('resource_provider'))
        .updateUser(
          "1",
          "",
          [],
        )

      await agree(controller, 'job_creator')
      await agree(controller, 'resource_provider')
      await controller
        .connect(getWallet('resource_provider'))
        .addResult(
          DEAL_ID,
          RESULTS_ID,
          DATA_ID,
          instructionCount
        )
      await controller
        .connect(getWallet('job_creator'))
        .checkResult(
          DEAL_ID,
        )
      await mediation
        .connect(getWallet('mediator'))
        .mediationAcceptResult(
          DEAL_ID,
        )

      const balancesAfterJC = await getBalances(token, 'job_creator')
      const balancesAfterRP = await getBalances(token, 'resource_provider')
      const balancesAfterMediator = await getBalances(token, 'mediator')

      expect(balancesAfterRP.tokens).to.equal(balancesBeforeRP.tokens + jobCost)
      expect(balancesAfterJC.tokens).to.equal(balancesBeforeJC.tokens - jobCost - mediationFee)
      expect(balancesAfterMediator.tokens).to.equal(balancesBeforeMediator.tokens + mediationFee)

      await checkAgreement(storage, 'MediationAccepted')
    })

    it("Runs a job in the mediation not OK path", async function () {
      const {
        token,
        storage,
        users,
        mediation,
        controller,
      } = await loadFixture(setupController)

      const balancesBeforeJC = await getBalances(token, 'job_creator')
      const balancesBeforeRP = await getBalances(token, 'resource_provider')
      const balancesBeforeMediator = await getBalances(token, 'mediator')
      const balancesBeforeAdmin = await getBalances(token, 'admin')

      await users
        .connect(getWallet('resource_provider'))
        .updateUser(
          "1",
          "",
          [],
        )

      await agree(controller, 'job_creator')
      await agree(controller, 'resource_provider')
      await controller
        .connect(getWallet('resource_provider'))
        .addResult(
          DEAL_ID,
          RESULTS_ID,
          DATA_ID,
          instructionCount
        )
      await controller
        .connect(getWallet('job_creator'))
        .checkResult(
          DEAL_ID,
        )
      await mediation
        .connect(getWallet('mediator'))
        .mediationRejectResult(
          DEAL_ID,
        )

      const balancesAfterJC = await getBalances(token, 'job_creator')
      const balancesAfterRP = await getBalances(token, 'resource_provider')
      const balancesAfterMediator = await getBalances(token, 'mediator')
      const balancesAfterAdmin = await getBalances(token, 'admin')

      expect(balancesAfterRP.tokens).to.equal(balancesBeforeRP.tokens - resultsCollateral)
      expect(balancesAfterJC.tokens).to.equal(balancesBeforeJC.tokens - mediationFee)
      expect(balancesAfterMediator.tokens).to.equal(balancesBeforeMediator.tokens + mediationFee)
      expect(balancesAfterAdmin.tokens).to.equal(balancesBeforeAdmin.tokens + resultsCollateral)

      await checkAgreement(storage, 'MediationRejected')
    })

    it("Revert agree on deal when token contract is paused", async function () {
      const {
        token,
        controller,
      } = await loadFixture(setupController)
      await token.connect(getWallet('admin')).pause()
      await expect(agree(controller, 'job_creator')).to.be.revertedWith('ERC20Pausable: token transfer while paused')
    })

  })
})