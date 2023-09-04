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
  setupControllerFixture,
} from './fixtures'
import {
  LilypadToken,
} from '../typechain-types'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Controller", () => {

  const dealID = ethers.getBigInt(10)
  const instructionPrice = ethers.getBigInt(1)
  const timeout = ethers.getBigInt(100)
  const timeoutCollateral = ethers.getBigInt(10)
  const resultsCollateralMultiple = ethers.getBigInt(4)
  const resultsCollateral = ethers.getBigInt(40)
  const paymentCollateral = ethers.getBigInt(30)
  const jobCost = ethers.getBigInt(20)
  const mediationFee = ethers.getBigInt(5)

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
        token,
        payments,
        storage,
        controller,
        tokenAddress,
      } = await loadFixture(setupController)

      const balancesBeforeAgreeRP = await getBalances(token, 'resource_provider')

      await expect(controller
        .connect(getWallet('resource_provider'))
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

      expect(await storage.getDealsForParty(getAddress('resource_provider')))
        .to.deep.equal([dealID])
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

      await expect(controller
        .connect(getWallet('job_creator'))
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

      expect(await storage.getDealsForParty(getAddress('job_creator')))
        .to.deep.equal([dealID])
    })
  })
})