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
} from './fixtures'

chai.use(chaiAsPromised)
const { expect } = chai

describe("Storage", () => {

  const rpCID = ethers.getBigInt(123)
  const rpURL = "abc"

  const jcCID = ethers.getBigInt(456)
  const jcURL = "def"

  const dealID = ethers.getBigInt(10)
  const resultsID = ethers.getBigInt(11)
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

  async function setupStorageWithUsers() {
    const storage = await setupStorage()
    expect(await storage
      .connect(getWallet('resource_provider'))
      .updateUser(
        rpCID,
        rpURL,
        [
          getServiceType('ResourceProvider'),
        ],
        [
          getAddress('mediator')
        ],
        [
          getAddress('directory')
        ]
      )
    ).to.not.be.reverted
    expect(await storage
      .connect(getWallet('job_creator'))
      .updateUser(
        jcCID,
        jcURL,
        [
          getServiceType('JobCreator'),
        ],
        [
          getAddress('mediator')
        ],
        [
          getAddress('directory')
        ]
      )
    ).to.not.be.reverted
    return storage
  }

  async function setupStorageWithUsersAndDeal() {
    const storage = await setupStorageWithUsers()
    expect(await storage
      .connect(getWallet('admin'))
      .ensureDeal(
        dealID,
        getAddress('resource_provider'),
        getAddress('job_creator'),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1),
        ethers.getBigInt(1)
      )
    ).to.not.be.reverted
    return storage
  }

  async function setupStorageWithUsersAndDealAndAgreement() {
    const storage = await setupStorageWithUsersAndDeal()

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

  async function setupStorageWithUsersAndDealAndAgreementAndResult() {
    const storage = await setupStorageWithUsersAndDealAndAgreement()

    await expect(storage
      .connect(getWallet('admin'))
      .addResult(
        dealID,
        resultsID,
        instructionCount,
      )
    )
      .to.emit(storage, "DealStateChange")
      .withArgs(dealID, getAgreementState('ResultsSubmitted'))
    
    return storage
  }

  async function setupStorageWithUsersAndDealAndAgreementAndResultAndChallenge() {
    const storage = await setupStorageWithUsersAndDealAndAgreementAndResult()

    await expect(storage
      .connect(getWallet('admin'))
      .checkResult(
        dealID,
        getAddress('mediator')
      )
    )
      .to.emit(storage, "DealStateChange")
      .withArgs(dealID, getAgreementState('ResultsChecked'))
    
    return storage
  }

  describe("Deployment", () => {

    it("Should have deployed with empty state", async function () {
      const storage = await loadFixture(setupStorage)
      expect(await storage.hasDeal(ethers.getBigInt(0))).to.equal(false)
    })

  })

  describe("Users", () => {

    it("Should be able to add and get users", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
      const rp = await storage.getUser(getAddress('resource_provider'))
      const jc = await storage.getUser(getAddress('job_creator'))

      expect(rp.metadataCID).to.equal(rpCID)
      expect(rp.url).to.equal(rpURL)
      expect(rp.roles).to.deep.equal([getServiceType('ResourceProvider')])
      expect(rp.trustedMediators).to.deep.equal([getAddress('mediator')])
      expect(rp.trustedDirectories).to.deep.equal([getAddress('directory')])

      expect(jc.metadataCID).to.equal(jcCID)
      expect(jc.url).to.equal(jcURL)
      expect(jc.roles).to.deep.equal([getServiceType('JobCreator')])
      expect(jc.trustedMediators).to.deep.equal([getAddress('mediator')])
      expect(jc.trustedDirectories).to.deep.equal([getAddress('directory')])
    })

    it("Should be able to update an existing user", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
      const newCID = ethers.getBigInt(456)
      const newURL = "def"

      expect(await storage
        .connect(getWallet('job_creator'))
        .updateUser(
          newCID,
          newURL,
          [
            getServiceType('JobCreator'),
          ],
          [
            getAddress('mediator')
          ],
          [
            getAddress('directory')
          ]
        )
      ).to.not.be.reverted
      
      const jc = await storage.getUser(getAddress('job_creator'))

      expect(jc.metadataCID).to.equal(newCID)
      expect(jc.url).to.equal(newURL)
    })

    it("Should be able to list users of a certain role", async function () {
      const storage = await loadFixture(setupStorageWithUsers)

      expect(await storage.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([])

      expect(await storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await storage.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([
          getAddress('job_creator'),
        ])
    })

    it("Should deny adding a user to a list without the role", async function () {
      const storage = await loadFixture(setupStorageWithUsers)

      await expect(storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('ResourceProvider')
        )
      ).to.be.revertedWith('User must have that role')
    })

    it("Should deny adding a user to a list twice", async function () {
      const storage = await loadFixture(setupStorageWithUsers)

      expect(await storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      await expect(storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is already part of that list')
    })

    it("Should allow a user to remove themselves from a list", async function () {
      const storage = await loadFixture(setupStorageWithUsers)

      expect(await storage
        .connect(getWallet('job_creator'))
        .addUserToList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await storage
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.not.be.reverted

      expect(await storage.showUsersInList(getServiceType('JobCreator')))
        .to.deep.equal([])
    })

    it("Should deny removing a user from a list if they have not added themselves", async function () {
      const storage = await loadFixture(setupStorageWithUsers)

      await expect(storage
        .connect(getWallet('job_creator'))
        .removeUserFromList(
          getServiceType('JobCreator')
        )
      ).to.be.revertedWith('User is not part of that list')
    })

  })

  describe("Access control", () => {
    it("Can only run ensureDeal if there is a controller address set", async function () {
      const storage = await loadFixture(setupStorageNoTest)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .ensureDeal(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1)
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run ensureDeal by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .ensureDeal(
          dealID,
          getAddress('resource_provider'),
          getAddress('job_creator'),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1)
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
          ethers.getBigInt(1),
          ethers.getBigInt(1),
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run addResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .addResult(
          dealID,
          ethers.getBigInt(1),
          ethers.getBigInt(1),
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
          getAddress('mediator')
        )
      ).to.be.revertedWith('ControllerOwnable: Controller address must be defined')
    })

    it("Can only run checkResult by the controller", async function () {
      const storage = await loadFixture(setupStorageNoTestWithControllerAddress)
      
      await expect(storage
        .connect(getWallet('resource_provider'))
        .checkResult(
          dealID,
          getAddress('mediator')
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
      const storage = await loadFixture(setupStorageWithUsersAndDeal)

      const deal = await storage.getDeal(dealID)

      expect(deal.dealId).to.equal(dealID)
      expect(deal.resourceProvider).to.equal(getAddress('resource_provider'))
      expect(deal.jobCreator).to.equal(getAddress('job_creator'))
      expect(deal.instructionPrice).to.equal(ethers.getBigInt(1))
      expect(deal.timeout).to.equal(ethers.getBigInt(1))
      expect(deal.timeoutCollateral).to.equal(ethers.getBigInt(1))
      expect(deal.paymentCollateral).to.equal(ethers.getBigInt(1))
      expect(deal.resultsCollateralMultiple).to.equal(ethers.getBigInt(1))
      expect(deal.mediationFee).to.equal(ethers.getBigInt(1))

      expect(await storage.hasDeal(dealID))
        .to.equal(true)

    })

    it("Should be able to see deals for specific parties", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDeal)
      expect(await storage.getDealsForParty(getAddress('resource_provider')))
        .to.deep.equal([dealID])
      expect(await storage.getDealsForParty(getAddress('job_creator')))
        .to.deep.equal([dealID])
      expect(await storage.getDealsForParty(getAddress('mediator')))
        .to.deep.equal([])
    })

    it("Should error when the RP and JC are the same", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
      await expect(storage
        .connect(getWallet('admin'))
        .ensureDeal(
          dealID,
          getAddress('resource_provider'),
          getAddress('resource_provider'),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1)
        )
      ).to.be.revertedWith('RP and JC cannot be the same')
    })

    it("Should error when the RP is empty", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
      await expect(storage
        .connect(getWallet('admin'))
        .ensureDeal(
          dealID,
          ethers.ZeroAddress,
          getAddress('job_creator'),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1)
        )
      ).to.be.revertedWith('Resource provider must be defined')
    })

    it("Should error when the JC is empty", async function () {
      const storage = await loadFixture(setupStorageWithUsers)
      
      await expect(storage
        .connect(getWallet('admin'))
        .ensureDeal(
          dealID,
          getAddress('resource_provider'),
          ethers.ZeroAddress,
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1),
          ethers.getBigInt(1)
        )
      ).to.be.revertedWith('Job creator must be defined')
    })

    const compareErrors = [
      '',
      'RP does not match',
      'JC does not match',
      'Instruction price does not match',
      'Timeout does not match',
      'Timeout collateral does not match',
      'Payment collateral does not match',
      'Results collateral does not match',
      'Mediation fee does not match',
    ]
    const badArgs: any = {
      0: ethers.getBigInt(100),
      1: getAddress('mediator'),
      2: getAddress('mediator'),
    }
    const goodArgs: any[] = [
      dealID,
      getAddress('resource_provider'),
      getAddress('job_creator'),
      ethers.getBigInt(1),
      ethers.getBigInt(1),
      ethers.getBigInt(1),
      ethers.getBigInt(1),
      ethers.getBigInt(1),
      ethers.getBigInt(1)
    ]
    compareErrors.forEach((expectedError, i) => {
      if(i == 0) return
      const passArgs: any[] = [].concat(...goodArgs)
      passArgs[i] = badArgs[i] || ethers.getBigInt(0)
      it(`Should compare error: ${expectedError}`, async function () {
        const storage = await loadFixture(setupStorageWithUsersAndDeal)
        const connectedStorage = storage.connect(getWallet('admin')) as any
        await expect(
          connectedStorage.ensureDeal(...passArgs)
        ).to.be.revertedWith(expectedError)
      })
      
    })

  })

  describe("Agreements", () => {

    it("Should be negotiating before agreements are made", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDeal)
      expect(await storage.isState(dealID, getAgreementState('DealNegotiating')))
        .to.equal(true)
    })

    it("Should be able to agree both parties", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.state).to.equal(getAgreementState('DealAgreed'))
      expect(agreement.resourceProviderAgreedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.jobCreatorAgreedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.dealAgreedAt).to.not.equal(ethers.getBigInt(0))

      expect(await storage.isState(dealID, getAgreementState('DealAgreed')))
        .to.equal(true)
    })

    it("Should throw if we agree the second time", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)
      await expect(storage
        .connect(getWallet('admin'))
        .agreeResourceProvider(
          dealID
        )
      ).to.be.revertedWith('Resource provider has already agreed')
      await expect(storage
        .connect(getWallet('admin'))
        .agreeJobCreator(
          dealID
        )
      ).to.be.revertedWith('Job creator has already agreed')
    })
  })

  describe("Results", () => {
    it("Should be able to add and get a result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResult)

      const result = await storage.getResult(dealID)
      expect(result.dealId).to.equal(dealID)
      expect(result.resultsId).to.equal(resultsID)
      expect(result.instructionCount).to.equal(instructionCount)

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.resultsSubmittedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('ResultsSubmitted'))
    })

    it("Should throw if we try to add results and not in agreed state", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDeal)

      await expect(storage
        .connect(getWallet('admin'))
        .addResult(
          dealID,
          resultsID,
          instructionCount,
        )
      ).to.be.revertedWith('Deal not in DealAgreed state')
    })

    it("Should be able to accept a result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResult)

      await expect(storage
        .connect(getWallet('admin'))
        .acceptResult(
          dealID
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('ResultsAccepted'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.resultsAcceptedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('ResultsAccepted'))
    })

    it("Should be able to challenge a result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResult)

      await expect(storage
        .connect(getWallet('admin'))
        .checkResult(
          dealID,
          getAddress('mediator')
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('ResultsChecked'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.resultsCheckedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('ResultsChecked'))
    })

    it("Should throw if we try to accept results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .acceptResult(
          dealID
        )
      ).to.be.revertedWith('Deal not in ResultsSubmitted state')
    })

    it("Should throw if we try to challenge results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .checkResult(
          dealID,
          getAddress('mediator')
        )
      ).to.be.revertedWith('Deal not in ResultsSubmitted state')
    })

    // TODO: mediate results without a trusted mediator address

  })

  describe("Mediation", () => {
    
    it("Should be able to mediator -> accept a result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResultAndChallenge)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationAcceptResult(
          dealID
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('MediationAccepted'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.mediationAcceptedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('MediationAccepted'))
    })

    it("Should be able to mediator -> reject a result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResultAndChallenge)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationRejectResult(
          dealID
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('MediationRejected'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.mediationRejectedAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('MediationRejected'))
    })

    it("Should throw if we try to accept mediation results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationAcceptResult(
          dealID
        )
      ).to.be.revertedWith('Deal not in ResultsChecked state')
    })

    it("Should throw if we try to challenge results and not in submitted state", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .mediationRejectResult(
          dealID
        )
      ).to.be.revertedWith('Deal not in ResultsChecked state')
    })

  })

  describe("Timeouts", () => {
    it("Should be able to timeout a submit result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreement)

      await expect(storage
        .connect(getWallet('admin'))
        .timeoutSubmitResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('TimeoutSubmitResults'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.timeoutSubmitResultsAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('TimeoutSubmitResults'))
    })

    it("Should be able to timeout a judge result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResult)

      await expect(storage
        .connect(getWallet('admin'))
        .timeoutJudgeResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('TimeoutJudgeResults'))

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.timeoutJudgeResultsAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('TimeoutJudgeResults'))
    })

    it("Should be able to timeout a mediation result", async function () {
      const storage = await loadFixture(setupStorageWithUsersAndDealAndAgreementAndResultAndChallenge)

      await expect(storage
        .connect(getWallet('admin'))
        .timeoutMediateResult(
          dealID,
        )
      )
        .to.emit(storage, "DealStateChange")
        .withArgs(dealID, getAgreementState('TimeoutMediateResults'))
      

      const agreement = await storage.getAgreement(dealID)
      expect(agreement.timeoutMediateResultsAt).to.not.equal(ethers.getBigInt(0))
      expect(agreement.state).to.equal(getAgreementState('TimeoutMediateResults'))
    })

  })

})