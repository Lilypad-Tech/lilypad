// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ILilypadStorage.sol";
import "./ILilypadToken.sol";

contract LilypadController is Ownable, Initializable {

  /**
   * Types
   */
  
  address private storageAddress;
  address private tokenAddress;

  ILilypadStorage private storageContract;
  ILilypadToken private tokenContract;

  /**
   * Events
   */

  event ResourceProviderAgreed(uint256 indexed dealId);
  event JobCreatorAgreed(uint256 indexed dealId);
  event DealAgreed(uint256 indexed dealId);
  event TimeoutSubmitResults(uint256 indexed dealId);
  event TimeoutJudgeResults(uint256 indexed dealId);
  event TimeoutMediateResults(uint256 indexed dealId);
  event ResultAdded(uint256 indexed dealId);
  event ResultAccepted(uint256 indexed dealId);
  event ResultChallenged(uint256 indexed dealId, address indexed mediator);
  event Payment(
    uint256 indexed dealId,
    address payee,
    uint256 amount,
    SharedStructs.PaymentReason reason,
    SharedStructs.PaymentDirection direction
  );
  
  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(address _storageAddress, address _tokenAddress) public initializer {
    setStorageAddress(_storageAddress);
    setTokenAddress(_tokenAddress);
  }

  function setStorageAddress(address _storageAddress) public onlyOwner {
    require(_storageAddress != address(0), "Storage address must be defined");
    storageAddress = _storageAddress;
    storageContract = ILilypadStorage(storageAddress);
  }

  function setTokenAddress(address _tokenAddress) public onlyOwner {
    require(_tokenAddress != address(0), "Token address must be defined");
    tokenAddress = _tokenAddress;
    tokenContract = ILilypadToken(storageAddress);
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
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 paymentCollateral,
    uint256 resultsCollateralMultiple,
    uint256 mediationFee
  ) public returns (SharedStructs.Agreement memory) {
    SharedStructs.Deal memory deal = storageContract.ensureDeal(
      dealId,
      resourceProvider,
      jobCreator,
      instructionPrice,
      timeout,
      timeoutCollateral,
      paymentCollateral,
      resultsCollateralMultiple,
      mediationFee
    );
    bool isResourceProvider = tx.origin == deal.resourceProvider;
    bool isJobCreator = tx.origin == deal.jobCreator;
    require(isResourceProvider || isJobCreator, "Only RP or JC can agree to deal");

    if(isResourceProvider) {
      storageContract.agreeResourceProvider(dealId);
      _payIn(deal.timeoutCollateral);
      emit Payment(
        dealId,
        deal.resourceProvider,
        deal.timeoutCollateral,
        SharedStructs.PaymentReason.TimeoutCollateral,
        SharedStructs.PaymentDirection.PaidIn
      );
      emit ResourceProviderAgreed(dealId);
    }
    else if(isJobCreator) {
      storageContract.agreeJobCreator(dealId);
      _payIn(deal.paymentCollateral + deal.timeoutCollateral);
      emit Payment(
        dealId,
        deal.jobCreator,
        deal.paymentCollateral,
        SharedStructs.PaymentReason.PaymentCollateral,
        SharedStructs.PaymentDirection.PaidIn
      );
      emit Payment(
        dealId,
        deal.resourceProvider,
        deal.timeoutCollateral,
        SharedStructs.PaymentReason.TimeoutCollateral,
        SharedStructs.PaymentDirection.PaidIn
      );
      emit JobCreatorAgreed(dealId);
    }

    // both sides have agreed!
    if(storageContract.isState(dealId, SharedStructs.AgreementState.DealAgreed)) {
      emit DealAgreed(dealId);
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
  ) public returns (SharedStructs.Result memory) {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.DealAgreed), "Deal is not in DealAgreed state");
    require(!_hasSubmitResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.resourceProvider == tx.origin, "Only RP can add results");

    SharedStructs.Result memory result = storageContract.addResult(
      dealId,
      resultsId,
      instructionCount
    );

    // how many multiple of the job cost must the RP put up as collateral
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    // what is the difference between what the RP has already paid and needs to now pay?
    // the RP has paid in the timeout collateral
    // it will now be charged the results collateral
    // a positive number means we are owed money
    // a negative number means we pay the RP a refund
    int256 resultsTimeoutDiff = int256(resultsCollateral) - int256(deal.timeoutCollateral);
    
    if(resultsTimeoutDiff > 0) {
      // the RP pays us because the job collateral is higher than the timeout collateral
      _payIn(uint256(resultsTimeoutDiff));
    }
    else if(resultsTimeoutDiff < 0) {
      // we pay the RP because the job collateral is lower
      _payOut(deal.resourceProvider, uint256(resultsTimeoutDiff));
    }

    // the refund of the timeout collateral
    emit Payment(
      dealId,
      deal.resourceProvider,
      deal.timeoutCollateral,
      SharedStructs.PaymentReason.TimeoutCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // the payment of the job collateral
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentReason.ResultsCollateral,
      SharedStructs.PaymentDirection.PaidIn
    );

    emit ResultAdded(dealId);

    return result;
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
    require(deal.jobCreator == tx.origin, "Only JC can accept result");
    
    storageContract.acceptResult(dealId);

    uint256 jobCost = storageContract.getJobCost(dealId);
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);
    
    // the difference between the job collateral and the job cost
    // is how much the job creator get's back
    int256 paymentCollateralRefund = int256(deal.paymentCollateral) - int256(jobCost);

    // the job cost more than the job collateral
    // this means the RP get's less than instruction count * instruction price
    // however they took on the deal knowing how much collateral was put up
    // equally however - it's very important to zero this number to prevent
    // the JC getting money back that they didn't pay in
    if(paymentCollateralRefund < 0) {
      paymentCollateralRefund = 0;
    }

    // we pay back the remaining job collateral and timeout collateral to the job creator
    _payOut(deal.jobCreator, uint256(paymentCollateralRefund) + deal.timeoutCollateral);

    if(paymentCollateralRefund > 0) {
      emit Payment(
        dealId,
        deal.jobCreator,
        uint256(paymentCollateralRefund),
        SharedStructs.PaymentReason.PaymentCollateral,
        SharedStructs.PaymentDirection.Refunded
      );
    }

    emit Payment(
      dealId,
      deal.jobCreator,
      deal.timeoutCollateral,
      SharedStructs.PaymentReason.TimeoutCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // now we pay back the results collateral and the job payment to the RP
    _payOut(deal.resourceProvider, resultsCollateral + jobCost);

    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentReason.ResultsCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    emit Payment(
      dealId,
      deal.resourceProvider,
      jobCost,
      SharedStructs.PaymentReason.JobPayment,
      SharedStructs.PaymentDirection.PaidOut
    );
  }

  // * check the JC is calling this
  // * check we are in Submitted state
  // * check the mediator is in the list of RP trusted mediators
  // * mark the deal as results challenged
  // * charge the JC the mediation fee
  // * refund the JC the timeout collateral
  // * emit the Mediation event so the mediator kicks in
  function challengeResult(
    uint256 dealId,
    address mediator
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal is not in ResultsSubmitted state");
    require(!_hasJudgeResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.jobCreator == tx.origin, "Only JC can challenge result");

    // this function will require that the mediator is in the RP's list of trusted mediators
    storageContract.challengeResult(dealId, mediator);

    // what is the difference between what the JC has already paid and needs to now pay?
    // the JC has paid in the timeout collateral
    // it will now be charged the mediation fee
    // a positive number means we are owed money
    // a negative number means we pay the RP a refund
    int256 timeoutMediationDiff = int256(deal.timeoutCollateral) - int256(deal.mediationFee);

    if(timeoutMediationDiff > 0) {
      // the RP pays us because the job collateral is higher than the timeout collateral
      _payIn(uint256(timeoutMediationDiff));
    }
    else if(timeoutMediationDiff < 0) {
      // we pay the RP because the job collateral is lower
      _payOut(deal.resourceProvider, uint256(timeoutMediationDiff));
    }
    
    // the refund of the timeout collateral
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.timeoutCollateral,
      SharedStructs.PaymentReason.TimeoutCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // the payment of the mediation fee
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.mediationFee,
      SharedStructs.PaymentReason.MediationFee,
      SharedStructs.PaymentDirection.PaidIn
    );

    emit ResultChallenged(dealId, mediator);
  }

  /**
   * Mediation
   */

  // the mediator calls this to say that the resource provider did the correct job
  // * check the state is ResultsChallenged
  // * check the mediator is calling this
  // * mark the deal as mediation accepted
  // * refund the JC what is left from the payment collateral (if any)
  // * pay the RP the cost of the job
  // * refund the RP the results collateral
  // * pay the mediator for mediating
  function mediateAcceptResult(
    uint256 dealId
  ) public {
    require(_canMediateResult(dealId), "Cannot mediate result");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);

    uint256 jobCost = storageContract.getJobCost(dealId);
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    int256 paymentCollateralRefund = int256(deal.paymentCollateral) - int256(jobCost);

    // if there is a refund for the JC then let's pay it
    if(paymentCollateralRefund > 0) {
      _payOut(deal.jobCreator, uint256(paymentCollateralRefund));
      emit Payment(
        dealId,
        deal.jobCreator,
        uint256(paymentCollateralRefund),
        SharedStructs.PaymentReason.PaymentCollateral,
        SharedStructs.PaymentDirection.Refunded
      );
    }

    // now we pay back the results collateral and the job payment to the RP
    _payOut(deal.resourceProvider, resultsCollateral + jobCost);

    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentReason.ResultsCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    emit Payment(
      dealId,
      deal.resourceProvider,
      jobCost,
      SharedStructs.PaymentReason.JobPayment,
      SharedStructs.PaymentDirection.PaidOut
    );

    // pay the mediator
    _payOut(agreement.mediator, deal.mediationFee);

    emit Payment(
      dealId,
      agreement.mediator,
      deal.mediationFee,
      SharedStructs.PaymentReason.MediationFee,
      SharedStructs.PaymentDirection.PaidOut
    );
  }

  // the mediator calls this to say that the resource provider did the bad job
  // * check the state is ResultsChallenged
  // * check the mediator is calling this
  // * mark the deal as mediation rejected
  // * refund the JC their payment collateral
  // * slash the RP's results collateral
  // * pay the mediator for mediating
  function mediateRejectResult(
    uint256 dealId
  ) public {
    require(_canMediateResult(dealId), "Cannot mediate result");

    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);

    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    // refund the JC their payment collateral
    _payOut(deal.jobCreator, deal.paymentCollateral);

    emit Payment(
      dealId,
      deal.jobCreator,
      deal.paymentCollateral,
      SharedStructs.PaymentReason.PaymentCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // pay the mediator
    _payOut(agreement.mediator, deal.mediationFee);

    emit Payment(
      dealId,
      agreement.mediator,
      deal.mediationFee,
      SharedStructs.PaymentReason.MediationFee,
      SharedStructs.PaymentDirection.PaidOut
    );

    // slash the RP
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentReason.ResultsCollateral,
      SharedStructs.PaymentDirection.Slashed
    );

  }

  function _canMediateResult(
    uint256 dealId 
  ) private returns (bool) {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsChallenged), "Deal is not in ResultsChallenged state");
    require(!_hasMediateResultsTimedOut(dealId), "Deal has timed out");
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    require(agreement.mediator != address(0), "No mediator has been selected");
    require(agreement.mediator == tx.origin, "Only mediator can accept or reject result");
    return true;
  }


  /**
   * Timeouts
   */

  // the job creator calls this after the timeout has passed and there are no results submitted
  // * check the JC is calling this
  // * mark the deal as timedout
  // * pay back the JC's job collateral
  // * slash the RP's results collateral
  // * emit the event
  function timeoutSubmitResults(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.DealAgreed), "Deal is not in DealAgreed state");
    require(_hasSubmitResultsTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.jobCreator == tx.origin, "Only JC can refund submit results timeout");
    storageContract.timeoutSubmitResult(dealId);

    // refund the job creator
    _payOut(deal.jobCreator, deal.paymentCollateral);

    // the refund of the job collateral to the JC
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.paymentCollateral,
      SharedStructs.PaymentReason.PaymentCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // the slashing of the timeout collateral for the RP
    emit Payment(
      dealId,
      deal.resourceProvider,
      deal.timeoutCollateral,
      SharedStructs.PaymentReason.TimeoutCollateral,
      SharedStructs.PaymentDirection.Slashed
    );

    emit TimeoutSubmitResults(dealId);
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
  function timeoutJudgeResults(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal is not in ResultsSubmitted state");
    require(_hasJudgeResultsTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.resourceProvider == tx.origin, "Only RP can refund judge results timeout");
    storageContract.timeoutJudgeResult(dealId);

    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    // refund the resource provider
    _payOut(deal.resourceProvider, resultsCollateral);

    // the refund of the results collateral to the RP
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentReason.PaymentCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // the slashing of the timeout collateral for the RP
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.timeoutCollateral,
      SharedStructs.PaymentReason.TimeoutCollateral,
      SharedStructs.PaymentDirection.Slashed
    );

    emit TimeoutJudgeResults(dealId);
  }

  // either the JC or RP call this after the timeout has passed after results being challenged
  // this refunds both the payment and results collateral to both the JC and RP
  // * check the RP or JC is calling this
  // * mark the deal as timedout
  // * pay back the RP's results collateral
  // * pay back the JC's paymnet collateral
  // * emit the event
  function timeoutMediateResults(
    uint256 dealId
  ) public {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsChallenged), "Deal is not in ResultsChallenged state");
    require(_hasMediateResultsTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.resourceProvider == tx.origin || deal.jobCreator == tx.origin, "Only RP or JC can refund mediate results timeout");
    storageContract.timeoutMediateResult(dealId);

    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    // refund the resource provider
    _payOut(deal.resourceProvider, resultsCollateral);

    // refund the job creator
    _payOut(deal.jobCreator, deal.paymentCollateral);

    // the refund of the results collateral to the RP
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentReason.ResultsCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    // the refund of the payment collateral to the JC
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.paymentCollateral,
      SharedStructs.PaymentReason.PaymentCollateral,
      SharedStructs.PaymentDirection.Refunded
    );

    emit TimeoutMediateResults(dealId);
  }

  /**
   * Timeout checkers
   */

  function _hasSubmitResultsTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    return block.timestamp > agreement.dealAgreedAt + deal.timeout;
  }

  function _hasJudgeResultsTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    return block.timestamp > agreement.resultsSubmittedAt + deal.timeout;
  }

  function _hasMediateResultsTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    return block.timestamp > agreement.resultsChallengedAt + deal.timeout;
  }

  /**
   * Payments
   */

  // move tokens around inside the erc-20 contract
  function _pay(
    address from,
    address to,
    uint256 amount
  ) private {
    require(tokenContract.balanceOf(from) >= amount, "Insufficient balance");
    require(tokenContract.allowance(from, to) >= amount, "Allowance too low");
    bool success = tokenContract.transferFrom(from, to, amount);
    require(success, "Transfer failed");
  }

  // take X tokens from the tx sender and add them to the contract's token balance
  function _payIn(
    uint256 amount
  ) private {
    // approve the tokens we are about to move
    // this works because _payIn is always called as part of the user who is paying
    // into the contract
    tokenContract.approve(address(this), amount);
    _pay(
      tx.origin,
      address(this),
      amount
    );
  }

  // take X tokens from the contract's token balance and send them to the given address
  function _payOut(
    address payWho,
    uint256 amount
  ) private {
    _pay(
      address(this),
      payWho,
      amount
    );
  }
}





