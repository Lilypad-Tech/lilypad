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
} from '../utils/web3'
import {
  getServiceType,
  getAgreementState,
} from '../utils/enums'
import {
  setupStorageFixture,
  getDefaultTimeouts,
  getDefaultPricing,
  DEFAULT_VALUES,
} from './fixtures'

import {
  SharedStructs,
} from '../typechain-types/contracts/LilypadStorage'

chai.use(chaiAsPromised)
const { expect } = chai

describe("Storage", () => {

  const rpCID = "123"
  const rpURL = "abc"

  const jcCID = "123"
  const jcURL = "def"

  const dealID = "10"
  const resultsID = "11"
  const DATA_ID = "12"
  const instructionCount = ethers.getBigInt(25)

  function setupStorage() {
    return setupStorageFixture({
      testMode: true,
    })
  }

  function setupStorageNoTest() {
    return setupStorageFixture({
      testMode: false,
    })
  }

  function setupStorageNoTestWithControllerAddress() {
    return setupStorageFixture({
      testMode: false,
      controllerAddress: getAddress('mediator'),
    })
  }

  async function setupStorageAndDeal() {
    const storage = await setupStorage()

    const members: SharedStructs.DealMembersStruct = {
      solver: getAddress('solver'), 
      jobCreator: getAddress('job_creator'),
      resourceProvider: getAddress('resource_provider'),
      mediators: [getAddress('mediator')],
    }
    const timeouts = getDefaultTimeouts()
    const pricing = getDefaultPricing()

    expect(await storage
      .connect(getWallet('admin'))
      .ensureDeal(
        dealID,
        members,
        timeouts,
        pricing
      )
    ).to.not.be.reverted
    return storage
  }

  async function setupStorageAndDealAndAgreement() {
    const storage = await setupStorageAndDeal()

    expect(await storage
      .connect(getWallet('admin'))
      .agreeResourceProvider(
        dealID
      )
    ).to.not.be.reverted

    await expect(storage
      .connect(getWallet('admin'))
      .agreeJobCreator(
        dealID
      )
    )
      .to.emit(storage, "DealStateChange")
      .withArgs(dealID, getAgreementState('DealAgreed'))
    
    return storage
  }

  async function setupStorageAndDealAndAgreementAndResult() {
    const storage = await setupStorageAndDealAndAgreement()

    await expect(storage
      .connect(getWallet('admin'))
      .addResult(
        dealID,
        resultsID,
        DATA_ID,
        instructionCount,
      )
    )
      .to.emit(storage, "DealStateChange")
      .withArgs(dealID, getAgreementState('ResultsSubmitted'))
    
    return storage
  }

  async function setupStorageAndDealAndAgreementAndResultAndChallenge() {
    const storage = await setupStorageAndDealAndAgreementAndResult()

    await expect(storage
      .connect(getWallet('admin'))
      .checkResult(
        dealID,
      )
    )
      .to.emit(storage, "DealStateChange")
      .withArgs(dealID, getAgreementState('ResultsChecked'))
    
    return storage
  }

  describe("Deployment", () => {

    it("Should have deployed with empty state", async function () {
      const storage = await loadFixture(setupStorage)
      expect(await storage.hasDeal("0")).to.equal(false)
    })

  })

  describe("Access control", () => {
    it("Can only run ensureDeal if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)

      const members: SharedStructs.DealMembersStruct = {
        solver: getAddress('solver'), 
        jobCreator: getAddress('job_creator'),
        resourceProvider: getAddress('resource_provider'),
        mediators: [getAddress('mediator')],
      }
      const timeouts = getDefaultTimeouts()
      const pricing = getDefaultPricing()
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .ensureDeal(
          dealID,
          members,
          timeouts,
          pricing
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run ensureDeal by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      const members: SharedStructs.DealMembersStruct = {
        solver: getAddress('solver'), 
        jobCreator: getAddress('job_creator'),
        resourceProvider: getAddress('resource_provider'),
        mediators: [getAddress('mediator')],
      }
      const timeouts = getDefaultTimeouts()
      const pricing = getDefaultPricing()
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .ensureDeal(
          dealID,
          members,
          timeouts,
          pricing
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run agreeResourceProvider if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .agreeResourceProvider(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run agreeResourceProvider by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .agreeResourceProvider(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run agreeJobCreator if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .agreeJobCreator(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run agreeJobCreator by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .agreeJobCreator(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })


    it("Can only run addResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .addResult(
          dealID,
          "1",
          DATA_ID,
          ethers.parseEther("1"),
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run addResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .addResult(
          dealID,
          "1",
          DATA_ID,
          ethers.parseEther("1"),
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run acceptResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .acceptResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run acceptResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .acceptResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run checkResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .checkResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run checkResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .checkResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run mediationAcceptResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .mediationAcceptResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run mediationAcceptResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .mediationAcceptResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run mediationRejectResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .mediationRejectResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run mediationRejectResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .mediationRejectResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run timeoutSubmitResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .timeoutSubmitResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run timeoutSubmitResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .timeoutSubmitResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run timeoutJudgeResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .timeoutJudgeResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run timeoutJudgeResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .timeoutJudgeResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })

    it("Can only run timeoutMediateResult if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .timeoutMediateResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run timeoutMediateResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .timeoutMediateResult(
          dealID,
        )
      ).to.be.revertedWith('ControllerOwnable: Only the controller can call this method')
    })
  })

  describe("Deals", () => {

    it("Should be able to add a deal and then read it", async function () {
      const storage = await loadFixture(setupStorageAndDeal)

      const deal = await storage.getDeal(dealID)
      
      expect(deal.dealId).to.equal(dealID)
      expect(deal.members.resourceProvider).to.equal(getAddress('resource_provider'))
      expect(deal.members.jobCreator).to.equal(getAddress('job_creator'))

      expect(deal.pricing.instructionPrice).to.equal(DEFAULT_VALUES.instructionPrice)
      expect(deal.pricing.paymentCollateral).to.equal(DEFAULT_VALUES.paymentCollateral)
      expect(deal.pricing.resultsCollateralMultiple).to.equal(DEFAULT_VALUES.resultsCollateralMultiple)
      expect(deal.pricing.mediationFee).to.equal(DEFAULT_VALUES.mediationFee)

      expect(deal.timeouts.agree.timeout).to.equal(DEFAULT_VALUES.timeout)
      expect(deal.timeouts.agree.collateral).to.equal(ethers.parseEther("0"))
      expect(deal.timeouts.submitResults.timeout).to.equal(DEFAULT_VALUES.timeout)
      expect(deal.timeouts.submitResults.collateral).to.equal(DEFAULT_VALUES.timeoutCollateral)
      expect(deal.timeouts.judgeResults.timeout).to.equal(DEFAULT_VALUES.timeout)
      expect(deal.timeouts.judgeResults.collateral).to.equal(DEFAULT_VALUES.timeoutCollateral)
      expect(deal.timeouts.mediateResults.timeout).to.equal(DEFAULT_VALUES.timeout)
      expect(deal.timeouts.mediateResults.collateral).to.equal(ethers.parseEther("0"))

      expect(await storage.hasDeal(dealID))
        .to.equal(true)

    })

    it("Should be able to see deals for specific parties", async function () {
      const storage = await loadFixture(setupStorageAndDeal)
      expect(await storage.getDealsForParty(getAddress('resource_provider')))
        .to.deep.equal([dealID])
      expect(await storage.getDealsForParty(getAddress('job_creator')))
        .to.deep.equal([dealID])
      expect(await storage.getDealsForParty(getAddress('mediator')))
        .to.deep.equal([])
    })

    it("Should error when the RP and JC are the same", async function () {
      const storage = await loadFixture(setupStorage)
      const members: SharedStructs.DealMembersStruct = {
        solver: getAddress('solver'), 
        jobCreator: getAddress('resource_provider'),
        resourceProvider: getAddress('resource_provider'),
        mediators: [getAddress('mediator')],
      }
      const timeouts = getDefaultTimeouts()
      const pricing = getDefaultPricing()
      
      await expect(storage
        .connect(getWallet('admin'))
        .ensureDeal(
          dealID,
          members,
          timeouts,
          pricing
        )
      ).to.be.revertedWith('RP / JC same')
    })

    it("Should error when the RP is empty", async function () {
      const storage = await loadFixture(setupStorage)
      const members: SharedStructs.DealMembersStruct = {
        solver: getAddress('solver'), 
        jobCreator: getAddress('job_creator'),
        resourceProvider: ethers.ZeroAddress,
        mediators: [getAddress('mediator')],
      }
      const timeouts = getDefaultTimeouts()
      const pricing = getDefaultPricing()
      
      await expect(storage
        .connect(getWallet('admin'))
        .ensureDeal(
          dealID,
          members,
          timeouts,
          pricing
        )
      ).to.be.revertedWith('RP missing')
    })

    it("Should error when the JC is empty", async function () {
      const storage = await loadFixture(setupStorage)
      const members: SharedStructs.DealMembersStruct = {
        solver: getAddress('solver'), 
        jobCreator: ethers.ZeroAddress,
        resourceProvider: getAddress('resource_provider'),
        mediators: [getAddress('mediator')],
      }
      const timeouts = getDefaultTimeouts()
      const pricing = getDefaultPricing()
      await expect(storage
        .connect(getWallet('admin'))
        .ensureDeal(
          dealID,
          members,
          timeouts,
          pricing
        )
      ).to.be.revertedWith('JC missing')
    })

    // TODO: re-activate the compare the two deals tests
    // const compareErrors = [
    //   '',
    //   'RP does not match',
    //   'JC does not match',
    //   'Instruction price does not match',
    //   'Timeout does not match',
    //   'Timeout collateral does not match',
    //   'Payment collateral does not match',
    //   'Results collateral does not match',
    //   'Mediation fee does not match',
    // ]
    // const badArgs: any = {
    //   0: ethers.parseEther("100"),
    //   1: getAddress('mediator'),
    //   2: getAddress('mediator'),
    // }
    // const goodArgs: any[] = [
    //   dealID,
    //   getAddress('resource_provider'),
    //   getAddress('job_creator'),
    //   ethers.parseEther("1"),
    //   ethers.parseEther("1"),
    //   ethers.parseEther("1"),
    //   ethers.parseEther("1"),
    //   ethers.parseEther("1"),
    //   ethers.parseEther("1")
    // ]
    // compareErrors.forEach((expectedError, i) => {
    //   if(i == 0) return
    //   const passArgs: any[] = [].concat(...goodArgs)
    //   passArgs[i] = badArgs[i] || ethers.parseEther("0")
    //   it(`Should compare error: ${expectedError}`, async function () {
    //     const storage = await loadFixture(setupStorageWithUsersAndDeal)
    //     const connectedStorage = storage.connect(getWallet('admin')) as any
    //     await expect(
    //       connectedStorage.ensureDeal(...passArgs)
    //     ).to.be.revertedWith(expectedError)
    //   })
      
    // })

  })

  describe("Agreements", () => {

    it("Should be negotiating before agreements are made", async function () {
      const storage = await loadFixture(setupStorageAndDeal)
      expect(await storage.isState(dealID, getAgreementState('DealNegotiating')))
        .to.equal(true)
    })

    it("Should be able to agree both parties", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.state).to.equal(getAgreementState('DealAgreed'))
      expect(agreement.resourceProviderAgreedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.jobCreatorAgreedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.dealAgreedAt).to.not.equal(ethers.parseEther("0"))

      expect(await storage.isState(dealID, getAgreementState('DealAgreed')))
        .to.equal(true)
    })

    it("Should throw if we agree the second time", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)
      await expect(storage
        .connect(getWallet('admin'))
        .agreeResourceProvider(
          dealID
        )
      ).to.be.revertedWith('RP has already agreed')
      await expect(storage
        .connect(getWallet('admin'))
        .agreeJobCreator(
          dealID
        )
      ).to.be.revertedWith('JC has already agreed')
    })
  })

  describe("Results", () => {
    it("Should be able to add and get a result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResult)

      const result = await storage.getResult(dealID)
      expect(result.dealId).to.equal(dealID)
      expect(result.resultsId).to.equal(resultsID)
      expect(result.instructionCount).to.equal(instructionCount)

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.resultsSubmittedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('ResultsSubmitted'))
    })

    it("Should throw if we try to add results and not in agreed state", async function () {
      const storage = await loadFixture(setupStorageAndDeal)

      await expect(storage
        .connect(getWallet('admin'))
        .addResult(
          dealID,
          resultsID,
          DATA_ID,
          instructionCount,
        )
      ).to.be.revertedWith('DealAgreed')
    })

    it("Should be able to accept a result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResult)

      await expect(storage
        .connect(getWallet('admin'))
        .acceptResult(
          dealID
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('ResultsAccepted'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.resultsAcceptedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('ResultsAccepted'))
    })

    it("Should be able to challenge a result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResult)

      await expect(storage
        .connect(getWallet('admin'))
        .checkResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('ResultsChecked'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.resultsCheckedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('ResultsChecked'))
    })

    it("Should throw if we try to accept results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .acceptResult(
          dealID
        )
      ).to.be.revertedWith('ResultsSubmitted')
    })

    it("Should throw if we try to challenge results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .checkResult(
          dealID,
        )
      ).to.be.revertedWith('ResultsSubmitted')
    })

    // TODO: mediate results without a trusted mediator address

  })

  describe("Mediation", () => {
    
    it("Should be able to mediator -> accept a result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResultAndChallenge)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationAcceptResult(
          dealID
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('MediationAccepted'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.mediationAcceptedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('MediationAccepted'))
    })

    it("Should be able to mediator -> reject a result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResultAndChallenge)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationRejectResult(
          dealID
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('MediationRejected'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.mediationRejectedAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('MediationRejected'))
    })

    it("Should throw if we try to accept mediation results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationAcceptResult(
          dealID
        )
      ).to.be.revertedWith('ResultsChecked')
    })

    it("Should throw if we try to challenge results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationRejectResult(
          dealID
        )
      ).to.be.revertedWith('ResultsChecked')
    })

  })

  describe("Timeouts", () => {
    it("Should be able to timeout a submit result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .timeoutSubmitResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('TimeoutSubmitResults'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.timeoutSubmitResultsAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('TimeoutSubmitResults'))
    })

    it("Should be able to timeout a judge result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResult)

      await expect(storage
        .connect(getWallet('admin'))
        .timeoutJudgeResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('TimeoutJudgeResults'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.timeoutJudgeResultsAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('TimeoutJudgeResults'))
    })

    it("Should be able to timeout a mediation result", async function () {
      const storage = await loadFixture(setupStorageAndDealAndAgreementAndResultAndChallenge)

      await expect(storage
        .connect(getWallet('admin'))
        .timeoutMediateResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('TimeoutMediateResults'))
      

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.timeoutMediateResultsAt).to.not.equal(ethers.parseEther("0"))
      expect(agreement.state).to.equal(getAgreementState('TimeoutMediateResults'))
    })

  })

})