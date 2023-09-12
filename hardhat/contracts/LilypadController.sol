// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ILilypadStorage.sol";
import "./ILilypadPayments.sol";

contract LilypadController is Ownable, Initializable {

  /**
   * Types
   */
  
  address private storageAddress;
  address private paymentsAddress;

  ILilypadStorage private storageContract;
  ILilypadPayments private paymentsContract;
  
  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(address _storageAddress, address _paymentsAddress) public initializer {
    setStorageAddress(_storageAddress);
    setPaymentsAddress(_paymentsAddress);
  }

  function setStorageAddress(address _storageAddress) public onlyOwner {
    require(_storageAddress != address(0), "Storage address must be defined");
    storageAddress = _storageAddress;
    storageContract = ILilypadStorage(storageAddress);
  }

  function setPaymentsAddress(address _paymentsAddress) public onlyOwner {
    require(_paymentsAddress != address(0), "Payments address must be defined");
    paymentsAddress = _paymentsAddress;
    paymentsContract = ILilypadPayments(_paymentsAddress);
  }

  /**
   * Agreements
   */

  // * create the deal if not exists
  // * otherwise compare the values to ensure 2 sided agreement
  // * check the RP or JC is calling this
  // * if RP:
  //    * mark the deal as RP agreed
  //    * pay in the timeout collateral
  // * if JC:
  //    * mark the deal as JC agreed
  //    * pay in the payment collateral and timeout collateral
  // * if both sides have agreed then mark the deal as agreed
  // * emit the event
  function agree(
    uint256 dealId,
    SharedStructs.DealMembers memory members,
    SharedStructs.DealTimeouts memory timeouts,
    SharedStructs.DealPricing memory pricing
  ) public returns (SharedStructs.Agreement memory) {
    SharedStructs.Deal memory deal = storageContract.ensureDeal(
      dealId,
      members,
      timeouts,
      pricing
    );
    bool isResourceProvider = tx.origin == deal.members.resourceProvider;
    bool isJobCreator = tx.origin == deal.members.jobCreator;
    require(isResourceProvider || isJobCreator, "Only RP or JC can agree to deal");

    if(isResourceProvider) {
      storageContract.agreeResourceProvider(dealId);
      paymentsContract.agreeResourceProvider(
        dealId,
        deal.members.resourceProvider,
        deal.timeouts.submitResults.collateral
      );
    }
    else if(isJobCreator) {
      storageContract.agreeJobCreator(dealId);
      paymentsContract.agreeJobCreator(
        dealId,
        deal.members.jobCreator,
        deal.pricing.paymentCollateral,
        // the JC paus the judge results collateral
        deal.timeouts.judgeResults.collateral
      );
    }
    return storageContract.getAgreement(dealId);
  }

  /**
   * Results
   */

  // * check the RP is calling this
  // * mark the deal as results submitted
  // * calculate the cost of the job
  // * calculate the job collateral based on the multiple
  // * work out the difference between the timeout and results collateral
  // * pay the difference into / out of the contract to the RP
  // * emit the event
  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.DealAgreed), "Deal is not in DealAgreed state");
    require(!_hasSubmitResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.resourceProvider == tx.origin, "Only RP can add results");

    storageContract.addResult(
      dealId,
      resultsId,
      instructionCount
    );

    // how many multiple of the job cost must the RP put up as collateral
    // we need to do this after having added the result otherwise
    // we don't know the instruction count
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    paymentsContract.addResult(
      dealId,
      deal.members.resourceProvider,
      resultsCollateral,
      // this is the RP adding a results so they get their submit results timeout collateral back
      deal.timeouts.submitResults.collateral
    );
  }

  // * check the JC is calling this
  // * check we are in Submitted state
  // * mark the deal as results accepted
  // * calculate the cost of the job
  // * deduct the cost of the job from the JC payment collateral
  // * pay the RP the cost of the job
  // * refund the RP the results collateral
  // * refund the JC the job collateral minus the cost
  // * refund the JC the timeout collateral
  function acceptResult(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal is not in ResultsSubmitted state");
    require(!_hasJudgeResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.jobCreator == tx.origin, "Only JC can accept result");
    
    uint256 jobCost = storageContract.getJobCost(dealId);
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.acceptResult(dealId);
    paymentsContract.acceptResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      jobCost,
      deal.pricing.paymentCollateral,
      resultsCollateral,
      // this is the JC judging their result so they get their timeout collateral back
      deal.timeouts.judgeResults.collateral
    );
  }

  // * check the JC is calling this
  // * check we are in Submitted state
  // * check the mediator is in the list of RP trusted mediators
  // * mark the deal as results checked
  // * charge the JC the mediation fee
  // * refund the JC the timeout collateral
  // * emit the Mediation event so the mediator kicks in
  function checkResult(
    uint256 dealId,
    address mediator
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal is not in ResultsSubmitted state");
    require(!_hasJudgeResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.jobCreator == tx.origin, "Only JC can check result");

    // this function will require that the mediator is in the RP's list of trusted mediators
    storageContract.checkResult(dealId, mediator);
    paymentsContract.checkResult(
      dealId,
      deal.members.jobCreator,
      // this is the JC judging their result so they get their timeout collateral back
      deal.timeouts.judgeResults.collateral,
      deal.pricing.mediationFee
    );
  }

  /**
   * Mediation
   */

  // the mediator calls this to say that the resource provider did the correct job
  // * check the state is ResultsChecked
  // * check the mediator is calling this
  // * mark the deal as mediation accepted
  // * refund the JC what is left from the payment collateral (if any)
  // * pay the RP the cost of the job
  // * refund the RP the results collateral
  // * pay the mediator for mediating
  function mediationAcceptResult(
    uint256 dealId
  ) public {
    require(_canMediateResult(dealId), "Cannot mediate result");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);

    uint256 jobCost = storageContract.getJobCost(dealId);
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.mediationAcceptResult(dealId);
    paymentsContract.mediationAcceptResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      agreement.mediator,
      jobCost,
      deal.pricing.paymentCollateral,
      resultsCollateral,
      deal.pricing.mediationFee
    );
  }

  // the mediator calls this to say that the resource provider did the bad job
  // * check the state is ResultsChecked
  // * check the mediator is calling this
  // * mark the deal as mediation rejected
  // * refund the JC their payment collateral
  // * slash the RP's results collateral
  // * pay the mediator for mediating
  function mediationRejectResult(
    uint256 dealId
  ) public {
    require(_canMediateResult(dealId), "Cannot mediate result");

    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);

    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.mediationRejectResult(dealId);
    paymentsContract.mediationRejectResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      agreement.mediator,
      deal.pricing.paymentCollateral,
      resultsCollateral,
      deal.pricing.mediationFee
    );
  }

  function _canMediateResult(
    uint256 dealId 
  ) private returns (bool) {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsChecked), "Deal is not in ResultsChecked state");
    require(!_hasMediateResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    require(agreement.mediator != address(0), "No mediator has been selected");
    require(agreement.mediator == tx.origin, "Only mediator can accept or reject result");
    return true;
  }

  /**
   * Timeouts
   */

  function timeoutAgree(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.DealNegotiating), "Deal is not in DealNegotiating state");
    require(_hasAgreeTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.jobCreator == tx.origin || deal.members.resourceProvider == tx.origin, "Only JC or RP can call timeoutAgree");
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    storageContract.timeoutAgree(dealId);
    if (agreement.resourceProviderAgreedAt > 0) {
      // this is an RP refund
      paymentsContract.timeoutAgreeRefundResourceProvider(
        dealId,
        deal.members.resourceProvider,
        deal.timeouts.submitResults.collateral
      );
    } else if (agreement.jobCreatorAgreedAt > 0) {
      // this is an JC refund
      paymentsContract.timeoutAgreeRefundJobCreator(
        dealId,
        deal.members.jobCreator,
        deal.pricing.paymentCollateral,
        deal.timeouts.submitResults.collateral
      );
    }
  }

  // the job creator calls this after the timeout has passed and there are no results submitted
  // * check the JC is calling this
  // * mark the deal as timedout
  // * pay back the JC's job collateral
  // * slash the RP's results collateral
  // * emit the event
  function timeoutSubmitResult(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.DealAgreed), "Deal is not in DealAgreed state");
    require(_hasSubmitResultsTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.jobCreator == tx.origin, "Only JC can refund submit results timeout");

    storageContract.timeoutSubmitResult(dealId);
    paymentsContract.timeoutSubmitResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      deal.pricing.paymentCollateral,
      deal.timeouts.submitResults.collateral
    );
  }

  // the resource provider calls this after the timeout has passed after submitting results
  // and the job creator has not yet submitted their judgement on those results
  // * check the RP is calling this
  // * mark the deal as timedout
  // * pay back the RP's results collateral
  // * pay the RP the cost of the job
  // * slash the JC's timeout collateral
  // * slash the JC's job collateral
  // * emit the event
  function timeoutJudgeResult(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal is not in ResultsSubmitted state");
    require(_hasJudgeResultsTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.resourceProvider == tx.origin, "Only RP can refund judge results timeout");

    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.timeoutJudgeResult(dealId);    
    paymentsContract.timeoutJudgeResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      resultsCollateral,
      deal.timeouts.judgeResults.collateral
    );
  }

  // either the JC or RP call this after the timeout has passed after results being checked
  // this refunds both the payment and results collateral to both the JC and RP
  // * check the RP or JC is calling this
  // * mark the deal as timedout
  // * pay back the RP's results collateral
  // * pay back the JC's paymnet collateral
  // * emit the event
  function timeoutMediateResult(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsChecked), "Deal is not in ResultsChecked state");
    require(_hasMediateResultsTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.resourceProvider == tx.origin || deal.members.jobCreator == tx.origin, "Only RP or JC can refund mediate results timeout");

    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.timeoutMediateResult(dealId);
    paymentsContract.timeoutMediateResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      deal.pricing.paymentCollateral,
      resultsCollateral,
      deal.pricing.mediationFee
    );
  }

  /**
   * Timeout checkers
   */

  function _hasAgreeTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    // if we've not created the deal then we cannot have timed out
    if(agreement.dealCreatedAt == 0) {
      return false;
    }
    // if we've agreed the deal then we cannot have timed out
    if(agreement.dealAgreedAt > 0) {
      return false;
    }
    return block.timestamp > agreement.dealCreatedAt + deal.timeouts.submitResults.timeout;
  }

  function _hasSubmitResultsTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    // if we've not agreed then we cannot have timed out
    if(agreement.dealAgreedAt == 0) {
      return false;
    }
    // if we have submitted results then we cannot have timed out
    if(agreement.resultsSubmittedAt > 0) {
      return false;
    }
    return block.timestamp > agreement.dealAgreedAt + deal.timeouts.submitResults.timeout;
  }

  function _hasJudgeResultsTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    // if we've not submitted results then we cannot have timed out
    if(agreement.resultsSubmittedAt == 0) {
      return false;
    }
    // if we have accepted results then we cannot have timed out
    if(agreement.resultsAcceptedAt > 0) {
      return false;
    }
    // if we have accepted results then we cannot have timed out
    if(agreement.resultsCheckedAt > 0) {
      return false;
    }
    return block.timestamp > agreement.resultsSubmittedAt + deal.timeouts.judgeResults.timeout;
  }

  function _hasMediateResultsTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    // if we've not checked results then we cannot have timed out
    if(agreement.resultsCheckedAt == 0) {
      return false;
    }
    return block.timestamp > agreement.resultsCheckedAt + deal.timeouts.mediateResults.timeout;
  }

}
