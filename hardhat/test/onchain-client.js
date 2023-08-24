const hre = require('hardhat')
const chai = require('chai')
const chaiAsPromised = require('chai-as-promised')
const ethereumWaffle = require('ethereum-waffle')

chai.use(ethereumWaffle.solidity)
chai.use(chaiAsPromised)
const { expect } = chai

const {
  ethers,
} = hre

const getBalance = async (address) => {
  const amount = await ethers.provider.getBalance(address)
  const inEth = ethers.utils.formatEther(amount)
  return Math.round(parseFloat(inEth))
}

describe("Modicum", async () => {

  let modicumContract
  let examplesContract
  let adminAccount
  let solverAccount
  let mediatorAccount
  let resourceProviderAccount
  let jobCreatorAccount

  const deployModicum = async () => {
    const modicumFactory = (await ethers.getContractFactory(
      "Modicum",
      {
        signer: adminAccount,
      }
    ))
    modicumContract = await modicumFactory.deploy()
    await modicumContract.deployed()
  }

  const deployExamples = async (modicumAddress) => {
    const examplesFactory = (await ethers.getContractFactory(
      "NaiveExamplesClient",
      {
        signer: adminAccount,
      }
    ))
    examplesContract = await examplesFactory.deploy(modicumAddress)
    await examplesContract.deployed()
  }

  const deployContracts = async () => {
    await deployModicum()
    await deployExamples(modicumContract.address)
    console.log(`modicumContract: ${modicumContract.address}`)
    console.log(`examplesContract: ${examplesContract.address}`)
  }

  context('contract', () => {
    beforeEach(async () => {
      [
        adminAccount,
        solverAccount,
        mediatorAccount,
        resourceProviderAccount,
        jobCreatorAccount,
      ] = await ethers.getSigners()
      await deployContracts()
    })

    describe('modicum', () => {
      it('deploys', async () => {
        const template = `apples"oranges`
        const params = `Here's a "thing"`

        const jobSpec = await modicumContract
          .connect(jobCreatorAccount)
          .getModuleSpec(template, params)

        console.log('--------------------------------------------')
        console.log(jobSpec)

        const parsedJobSpec = JSON.parse(jobSpec)

        expect(parsedJobSpec.template).to.equal(template)
        expect(parsedJobSpec.params).to.equal(params)
      })

      it('reverts if we don\'t pay enough or there are no mediators', async () => {
        const JOB_COST = ethers.utils.parseEther("100")
        const JOB_COST_NOT_ENOUGH = ethers.utils.parseEther("99")
        const CID = "i_am_cid"

        await expect(
          examplesContract
            .connect(jobCreatorAccount)
            .runCowsay("holy cow", {
              value: JOB_COST,
            })
        ).to.be.revertedWith('No mediators provided')

        await modicumContract
          .connect(mediatorAccount)
          .registerMediator(
            1, //Architecture arch,
            0, //instructionPrice
            0, //bandwidthPrice
            0, //availabilityValue
            0  //verificationCount
          )

        await modicumContract
          .connect(adminAccount)
          .setDefaultMediators([
            mediatorAccount.address,
          ])

        await expect(
          examplesContract
            .connect(jobCreatorAccount)
            .runCowsay("holy cow", {
              value: JOB_COST,
            })
        ).to.be.revertedWith('Module not found')

        await expect(
          modicumContract
            .connect(jobCreatorAccount)
            .postJobOfferPartOne(
              'cowsay:v0.0.1',
              1,
              1,
              1,
              1,
              1,
              168933053300,
              1, {
                value: JOB_COST,
              }
            )
        ).to.be.revertedWith('Module not found')

        await modicumContract
          .connect(adminAccount)
          .setModuleCost('cowsay:v0.0.1', JOB_COST)

        await expect(
          examplesContract
            .connect(jobCreatorAccount)
            .runCowsay("holy cow", {
              value: JOB_COST_NOT_ENOUGH,
            })
        ).to.be.revertedWith('Not enough funds sent for job')

        await expect(
          modicumContract
            .connect(jobCreatorAccount)
            .postJobOfferPartOne(
              'cowsay:v0.0.1',
              1,
              1,
              1,
              1,
              1,
              168933053300,
              1, {
                value: JOB_COST_NOT_ENOUGH,
              }
            )
        ).to.be.revertedWith('Not enough funds sent for job')

        await expect(
          examplesContract
            .connect(jobCreatorAccount)
            .runCowsay("holy cow", {
              value: JOB_COST,
            })
        ).to.not.be.reverted

        await expect(
          modicumContract
            .connect(jobCreatorAccount)
            .postJobOfferPartOne(
              'cowsay:v0.0.1',
              1,
              1,
              1,
              1,
              1,
              168933053300,
              1, {
                value: JOB_COST,
              }
            )
        ).to.not.be.reverted
      })

      it('will not let us post a resource offer with too little of a deposit', async () => {
        const JOB_COST = ethers.BigNumber.from("1000")
        const DEPOSIT_MULTIPLE = ethers.BigNumber.from("10")
        await modicumContract
          .connect(adminAccount)
          .setModuleCost('cowsay:v0.0.1', JOB_COST)
        await modicumContract
          .connect(adminAccount)
          .setResourceProviderDepositMultiple(DEPOSIT_MULTIPLE)
        await modicumContract
          .connect(resourceProviderAccount)
          .registerResourceProvider(
            1, //Architecture arch,
            0, //timePerInstruction
          )
        await expect(
          modicumContract
            .connect(resourceProviderAccount)
            .postResOffer(0,0,0,0,0,0,0,0,0, {
              value: JOB_COST.mul(DEPOSIT_MULTIPLE).sub(1),
            })
        ).to.be.revertedWith("You need to deposit more money to offer this resource")

        const resourceDeposit = await modicumContract.getRequiredResourceProviderDeposit()
        await expect(
          modicumContract
            .connect(resourceProviderAccount)
            .postResOffer(0,0,0,0,0,0,0,0,0, {
              value: resourceDeposit,
            })
        ).to.not.be.reverted
      })

      it('runs a job', async () => {
        const JOB_COST = ethers.utils.parseEther("100")
        const DEPOSIT_MULTIPLE = ethers.BigNumber.from("10")
        const CID = "i_am_cid"

        const resourceProviderBalanceBefore = await getBalance(resourceProviderAccount.address)
        const jobCreatorBalanceBefore = await getBalance(jobCreatorAccount.address)

        await modicumContract
          .connect(adminAccount)
          .setModuleCost('cowsay:v0.0.1', JOB_COST)

        await modicumContract
          .connect(adminAccount)
          .setResourceProviderDepositMultiple(DEPOSIT_MULTIPLE)

        await modicumContract
          .connect(mediatorAccount)
          .registerMediator(
            1, //Architecture arch,
            0, //instructionPrice
            0, //bandwidthPrice
            0, //availabilityValue
            0  //verificationCount
          )

        await modicumContract
          .connect(adminAccount)
          .setDefaultMediators([
            mediatorAccount.address,
          ])

        await modicumContract
          .connect(resourceProviderAccount)
          .registerResourceProvider(
            1, //Architecture arch,
            0, //timePerInstruction
          )

        const postResourceOfferTrx = await modicumContract
          .connect(resourceProviderAccount)
          .postResOffer(0,0,0,0,0,0,0,0,0,{
            value: JOB_COST.mul(DEPOSIT_MULTIPLE),
          })

        const postResourceOfferReceipt = await postResourceOfferTrx.wait();
        let resourceOfferEvent

        for (let i = 0; i < postResourceOfferReceipt.logs.length; i++) {
          const parsedLog = modicumContract.interface.parseLog(postResourceOfferReceipt.logs[i])
          if(parsedLog.name == 'ResourceOfferPosted') {
            resourceOfferEvent = parsedLog
          }
        }

        expect(resourceOfferEvent).to.not.be.undefined

        const resourceOfferId = resourceOfferEvent.args.offerId

        

        const runCowsayTrx = await examplesContract
          .connect(jobCreatorAccount)
          .runCowsay("holy cow", {
            value: JOB_COST,
          })

        const runCowsayReceipt = await runCowsayTrx.wait();
        let runCowsayEvent

        for (let i = 0; i < runCowsayReceipt.logs.length; i++) {
          const parsedLog = modicumContract.interface.parseLog(runCowsayReceipt.logs[i])
          if(parsedLog.name == 'JobOfferPostedPartTwo') {
            runCowsayEvent = parsedLog
          }
        }

        expect(runCowsayEvent).to.not.be.undefined

        const jobOfferId = runCowsayEvent.args.offerId

        const postMatchTrx = await modicumContract
          .connect(solverAccount)
          .postMatch(
            jobOfferId,
            resourceOfferId,
            mediatorAccount.address, 
          )

        const postMatchReceipt = await postMatchTrx.wait();
        let postMatchEvent

        for (let i = 0; i < postMatchReceipt.logs.length; i++) {
          const parsedLog = modicumContract.interface.parseLog(postMatchReceipt.logs[i])
          if(parsedLog.name == 'Matched') {
            postMatchEvent = parsedLog
          }
        }

        expect(postMatchEvent).to.not.be.undefined
        
        const postResultTrx = await modicumContract
          .connect(resourceProviderAccount)
          .postResult(
            0,
            jobOfferId,
            0,
            "",
            CID,
            JOB_COST,
            0,
          )

        const postResultReceipt = await postResultTrx.wait();
        const postResultEvent = examplesContract.interface.parseLog(postResultReceipt.logs[1])

        expect(postResultEvent).to.not.be.undefined
        expect(postResultEvent.name).to.equal('ReceivedJobResults')
        expect(postResultEvent.args.jobID).to.equal(jobOfferId)
        expect(postResultEvent.args.cid).to.equal(CID)

        const resourceProviderBalanceAfter = await getBalance(resourceProviderAccount.address)
        const jobCreatorBalanceAfter = await getBalance(jobCreatorAccount.address)

        console.dir({
          resourceProviderBalanceBefore: resourceProviderBalanceBefore,
          resourceProviderBalanceAfter: resourceProviderBalanceAfter,
          jobCreatorBalanceBefore: jobCreatorBalanceBefore,
          jobCreatorBalanceAfter: jobCreatorBalanceAfter,
        })

        expect(jobCreatorBalanceAfter).to.equal(jobCreatorBalanceBefore - 100)
        expect(resourceProviderBalanceAfter).to.equal(resourceProviderBalanceBefore + 100)
      })

      context('verification', () => {

        const JOB_COST = ethers.utils.parseEther("100")
        const DEPOSIT_MULTIPLE = ethers.BigNumber.from("10")
        const CID = "i_am_cid"
        let resourceOfferId
        let jobOfferId
        let resultId
        let matchId

        const setupVerificationTest = async () => {
          await modicumContract
            .connect(adminAccount)
            .setModuleCost('cowsay:v0.0.1', JOB_COST)

          await modicumContract
            .connect(adminAccount)
            .setResourceProviderDepositMultiple(DEPOSIT_MULTIPLE)

          await modicumContract
            .connect(mediatorAccount)
            .registerMediator(
              1, //Architecture arch,
              0, //instructionPrice
              0, //bandwidthPrice
              0, //availabilityValue
              0  //verificationCount
            )

          await modicumContract
            .connect(adminAccount)
            .setDefaultMediators([
              mediatorAccount.address,
            ])

          await modicumContract
            .connect(resourceProviderAccount)
            .registerResourceProvider(
              1, //Architecture arch,
              0, //timePerInstruction
            )

          const postResourceOfferTrx = await modicumContract
            .connect(resourceProviderAccount)
            .postResOffer(0,0,0,0,0,0,0,0,0,{
              value: JOB_COST.mul(DEPOSIT_MULTIPLE),
            })

          const postResourceOfferReceipt = await postResourceOfferTrx.wait();
          let resourceOfferEvent

          for (let i = 0; i < postResourceOfferReceipt.logs.length; i++) {
            const parsedLog = modicumContract.interface.parseLog(postResourceOfferReceipt.logs[i])
            if(parsedLog.name == 'ResourceOfferPosted') {
              resourceOfferEvent = parsedLog
            }
          }

          expect(resourceOfferEvent).to.not.be.undefined

          resourceOfferId = resourceOfferEvent.args.offerId

          await modicumContract
            .connect(jobCreatorAccount)
            .registerJobCreator()

          await modicumContract
            .connect(jobCreatorAccount)
            .jobCreatorAddTrustedMediator(mediatorAccount.address)

          await modicumContract
            .connect(jobCreatorAccount)
            .postJobOfferPartOne(
              'cowsay:v0.0.1',
              1,
              1,
              1,
              1,
              1,
              168933053300,
              1,
              {
                value: JOB_COST,
              }
            )

          const runCowsayTrx = await modicumContract
            .connect(jobCreatorAccount)
            .postJobOfferPartTwo(
              1,
              "",
              jobCreatorAccount.address,
              0,
              0,
              `{"template":"cowsay:v0.0.1","params":"holy cow"}`,
          )

          const runCowsayReceipt = await runCowsayTrx.wait();
          let runCowsayEvent

          for (let i = 0; i < runCowsayReceipt.logs.length; i++) {
            const parsedLog = modicumContract.interface.parseLog(runCowsayReceipt.logs[i])
            if(parsedLog.name == 'JobOfferPostedPartTwo') {
              runCowsayEvent = parsedLog
            }
          }

          expect(runCowsayEvent).to.not.be.undefined

          jobOfferId = runCowsayEvent.args.offerId

          await modicumContract
            .connect(solverAccount)
            .postMatch(
              jobOfferId,
              resourceOfferId,
              mediatorAccount.address, 
            )

          const postResultTrx = await modicumContract
            .connect(resourceProviderAccount)
            .postResult(
              0,
              jobOfferId,
              0,
              "",
              CID,
              JOB_COST,
              0,
            )

          const postResultReceipt = await postResultTrx.wait();
          let postResultEvent

          for (let i = 0; i < postResultReceipt.logs.length; i++) {
            const parsedLog = modicumContract.interface.parseLog(postResultReceipt.logs[i])
            if(parsedLog.name == 'ResultPosted') {
              postResultEvent = parsedLog
            }
          }

          expect(postResultEvent).to.not.be.undefined

          resultId = postResultEvent.args.resultId
          
          const rejectTrx = await modicumContract
            .connect(jobCreatorAccount)
            .rejectResult(
              resultId,
            )

          const rejectReceipt = await rejectTrx.wait();
          let rejectEvent

          for (let i = 0; i < rejectReceipt.logs.length; i++) {
            const parsedLog = modicumContract.interface.parseLog(rejectReceipt.logs[i])
            if(parsedLog.name == 'JobAssignedForMediation') {
              rejectEvent = parsedLog
            }
          }

          expect(rejectEvent).to.not.be.undefined

          matchId = rejectEvent.args.matchId
        }

        it('handles verification for a bad actor', async () => {
          const resourceProviderBalanceBefore = await getBalance(resourceProviderAccount.address)
          const jobCreatorBalanceBefore = await getBalance(jobCreatorAccount.address)

          await setupVerificationTest()

          await modicumContract
            .connect(mediatorAccount)
            .postMediationResult(
              matchId,
              jobOfferId,
              0, // ResultStatus = Completed
              "",
              "567",
              2, // Verdict = WrongResults,
              0, // Party = ResourceProvider
            )

            const resourceProviderBalanceAfter = await getBalance(resourceProviderAccount.address)
            const jobCreatorBalanceAfter = await getBalance(jobCreatorAccount.address)

            console.dir({
              resourceProviderBalanceBefore: resourceProviderBalanceBefore,
              resourceProviderBalanceAfter: resourceProviderBalanceAfter,
              jobCreatorBalanceBefore: jobCreatorBalanceBefore,
              jobCreatorBalanceAfter: jobCreatorBalanceAfter,
            })

            expect(jobCreatorBalanceAfter).to.equal(jobCreatorBalanceBefore)
            expect(resourceProviderBalanceAfter).to.equal(resourceProviderBalanceBefore - (100 * 10))
        })

        it('handles verification for a good actor', async () => {
          const resourceProviderBalanceBefore = await getBalance(resourceProviderAccount.address)
          const jobCreatorBalanceBefore = await getBalance(jobCreatorAccount.address)

          await setupVerificationTest()

          await modicumContract
            .connect(mediatorAccount)
            .postMediationResult(
              matchId,
              jobOfferId,
              0, // ResultStatus = Completed
              "",
              "567",
              3, // Verdict = CorrectResults,
              0, // Party = ResourceProvider
            )

            const resourceProviderBalanceAfter = await getBalance(resourceProviderAccount.address)
            const jobCreatorBalanceAfter = await getBalance(jobCreatorAccount.address)

            console.dir({
              resourceProviderBalanceBefore: resourceProviderBalanceBefore,
              resourceProviderBalanceAfter: resourceProviderBalanceAfter,
              jobCreatorBalanceBefore: jobCreatorBalanceBefore,
              jobCreatorBalanceAfter: jobCreatorBalanceAfter,
            })

            expect(jobCreatorBalanceAfter).to.equal(jobCreatorBalanceBefore - 100)
            expect(resourceProviderBalanceAfter).to.equal(resourceProviderBalanceBefore + 100)
        })
      })
    })
  })
  
})