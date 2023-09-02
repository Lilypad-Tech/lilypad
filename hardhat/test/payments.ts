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
} from './utils'

chai.use(chaiAsPromised)
const { expect } = chai

// https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat

describe.only("Payments", () => {

  const dealID = ethers.getBigInt(10)

  describe("Deals", () => {

    it("Should agree resource provider", async function () {

      const timeoutCollateral = ethers.getBigInt(10)

      const {
        token,
        payments,
      } = await loadFixture(setupPaymentsFixture)

      const balanceBefore = await token.balanceOf(getAddress('resource_provider'))
      const escrowBefore = await payments.getEscrowBalance(getAddress('resource_provider'))

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

      const balanceAfter = await token.balanceOf(getAddress('resource_provider'))
      const escrowAfter = await payments.getEscrowBalance(getAddress('resource_provider'))

      expect(balanceAfter).to.equal(balanceBefore - timeoutCollateral)
      expect(escrowAfter).to.equal(escrowBefore + timeoutCollateral)
    })

  })
})