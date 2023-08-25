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
  
  address public storageAddress;
  address public tokenAddress;

  ILilypadStorage private storageContract;
  ILilypadToken private tokenContract;

  /**
   * Events
   */

  event ResourceProviderAgreed(uint256 indexed dealId);
  event JobCreatorAgreed(uint256 indexed dealId);
  event DealAgreed(uint256 indexed dealId);
  event Timeout(uint256 indexed dealId);
  event ResultAdded(uint256 indexed dealId);
  event ResultAccepted(uint256 indexed dealId);
  event ResultRejected(uint256 indexed dealId);
  event Payment(
    uint256 indexed dealId,
    address payee,
    uint256 amount,
    SharedStructs.PaymentDirection direction,
    SharedStructs.PaymentReason reason
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

  function agree(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 jobCollateral,
    uint256 resultsCollateralMultiple
  ) public returns (SharedStructs.Agreement memory) {
    require(storageContract.isNegotiating(dealId), "Deal is not in negotiation state");
    require(resourceProvider != address(0), "RP must be defined");
    require(jobCreator != address(0), "JC must be defined");
    require(resourceProvider != jobCreator, "RP and JC cannot be the same");

    // we already have this deal and so the values must line up
    if(storageContract.hasDeal(dealId)) {
      SharedStructs.Deal memory existingDeal = storageContract.getDeal(dealId);
      require(existingDeal.resourceProvider == resourceProvider, "RP does not match");
      require(existingDeal.jobCreator == jobCreator, "JC does not match");
      require(existingDeal.instructionPrice == instructionPrice, "Instruction price does not match");
      require(existingDeal.timeout == timeout, "Timeout does not match");
      require(existingDeal.timeoutCollateral == timeoutCollateral, "Timeout collateral does not match");
      require(existingDeal.jobCollateral == jobCollateral, "Job collateral does not match");
      require(existingDeal.resultsCollateralMultiple == resultsCollateralMultiple, "Results collateral does not match");
    }
    else {
      // we don't have this deal yet so add it
      storageContract.addDeal(
        dealId,
        resourceProvider,
        jobCreator,
        instructionPrice,
        timeout,
        timeoutCollateral,
        jobCollateral,
        resultsCollateralMultiple
      );
    }

    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
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
        SharedStructs.PaymentDirection.In,
        SharedStructs.PaymentReason.TimeoutCollateral
      );
      emit ResourceProviderAgreed(dealId);
    }
    else if(isJobCreator) {
      storageContract.agreeJobCreator(dealId);
      _payIn(deal.jobCollateral);
      emit Payment(
        dealId,
        deal.jobCreator,
        deal.jobCollateral,
        SharedStructs.PaymentDirection.In,
        SharedStructs.PaymentReason.JobCollateral
      );
      emit JobCreatorAgreed(dealId);
    }

    // both sides have agreed!
    if(storageContract.isAgreement(dealId)) {
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
  // * pay the difference into / out of the contract
  // * emit the event
  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) public returns (SharedStructs.Result memory) {
    require(storageContract.isAgreement(dealId), "Deal is not in agreement state");
    require(!_hasTimedOut(dealId), "Deal has timed out");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.resourceProvider == tx.origin, "Only RP can add results");

    SharedStructs.Result memory result = storageContract.addResult(
      dealId,
      resultsId,
      instructionCount
    );

    // the cost of the job based on reported instructions * cost per instruction
    uint256 jobCost = deal.instructionPrice * result.instructionCount;

    // how many multiple of the job cost must the RP put up as collateral
    uint256 resultsCollateral = deal.resultsCollateralMultiple * jobCost;

    // what is the difference between what the RP has already paid and needs to now pay?
    // a positive number means we are owed money
    // a negative number means we pay the RP a refund
    int256 rpCollateralDiff = int256(resultsCollateral) - int256(deal.timeoutCollateral);
    
    if(rpCollateralDiff > 0) {
      // the RP pays us because the job collateral is higher than the timeout collateral
      _payIn(uint256(rpCollateralDiff));
    }
    else if(rpCollateralDiff < 0) {
      // we pay the RP because the job collateral is lower
      _payOut(deal.resourceProvider, uint256(rpCollateralDiff));
    }

    // the refund of the timeout collateral
    emit Payment(
      dealId,
      deal.resourceProvider,
      deal.timeoutCollateral,
      SharedStructs.PaymentDirection.Out,
      SharedStructs.PaymentReason.TimeoutCollateralRefund
    );

    // the payment of the job collateral
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentDirection.In,
      SharedStructs.PaymentReason.ResultsCollateral
    );

    emit ResultAdded(dealId);

    return result;
  }

  // * check the JC is calling this
  // * check we are in Submitted state
  // * mark the deal as results accepted
  // * calculate the cost of the job
  // * deduct the cost of the job from the JC collateral
  // * pay the RP the cost of the job and the results collateral
  // * pay the JC
  function acceptResult(
    uint256 dealId
  ) public {
    require(storageContract.isSubmitted(dealId), "Deal is not in submitted state");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.jobCreator == tx.origin, "Only JC can accept result");
    SharedStructs.Result memory result = storageContract.getResult(dealId);
    storageContract.acceptResult(dealId);

    uint256 jobCost = deal.instructionPrice * result.instructionCount;
    uint256 resultsCollateral = deal.resultsCollateralMultiple * jobCost;

    // we pay the RP the resultsCollateral + jobCost
    // we pay the JC the jobCollateral - jobCost
    int256 jcRefund = int256(deal.jobCollateral) - int256(jobCost);

    if(jcRefund > 0) {
      // the JC gets a refund
      _payOut(deal.jobCreator, uint256(jcRefund));

      // the refund to the JC because they overpaid
      emit Payment(
        dealId,
        deal.resourceProvider,
        resultsCollateral,
        SharedStructs.PaymentDirection.Out,
        SharedStructs.PaymentReason.JobCollateralRefund
      );
    }

    _payOut(deal.resourceProvider, jobCost + resultsCollateral);

    // the refund to the RP for their results collateral
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentDirection.Out,
      SharedStructs.PaymentReason.ResultsCollateralRefund
    );

    // the payment to the RP for running the job
    emit Payment(
      dealId,
      deal.resourceProvider,
      jobCost,
      SharedStructs.PaymentDirection.Out,
      SharedStructs.PaymentReason.JobPayment
    );
  }

  // * check the JC is calling this
  // * check we are in Submitted state
  // * mark the deal as results rejected
  // * refund the JC
  // * slash the RP

  // TODO: trigger the mediation process
  function rejectResult(
    uint256 dealId
  ) public {
    require(storageContract.isSubmitted(dealId), "Deal is not in submitted state");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.jobCreator == tx.origin, "Only JC can reject result");
    SharedStructs.Result memory result = storageContract.getResult(dealId);
    storageContract.rejectResult(dealId);

    uint256 jobCost = deal.instructionPrice * result.instructionCount;
    uint256 resultsCollateral = deal.resultsCollateralMultiple * jobCost;

    // the JC gets a refund
    _payOut(deal.jobCreator, deal.jobCollateral);

    // the refund to the JC because they overpaid
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.jobCollateral,
      SharedStructs.PaymentDirection.Out,
      SharedStructs.PaymentReason.JobCollateralRefund
    );

    // we slash the RP
    emit Payment(
      dealId,
      deal.resourceProvider,
      resultsCollateral,
      SharedStructs.PaymentDirection.Slashed,
      SharedStructs.PaymentReason.ResultsCollateralSlashed
    );
  }

  // the job creator calls this after the timeout has passed and there are no results submitted
  // https://ethereum.stackexchange.com/questions/86633/time-dependent-tests-with-hardhat
  // * check the JC is calling this
  // * mark the deal as timedout
  // * pay back the JC's job collateral
  // * emit the event
  function refundTimeout(
    uint256 dealId
  ) public {
    require(storageContract.isAgreement(dealId), "Deal is not in agreement state");
    require(_hasTimedOut(dealId), "Deal has not timed out yet");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.jobCreator == tx.origin, "Only JC can refund timeout");
    storageContract.timeoutResult(dealId);

    // actually refund the job creator
    _payOut(deal.jobCreator, deal.jobCollateral);

    // the refund of the job collateral to the JC
    emit Payment(
      dealId,
      deal.jobCreator,
      deal.jobCollateral,
      SharedStructs.PaymentDirection.Out,
      SharedStructs.PaymentReason.JobCollateralTimeoutRefund
    );

    // the slashing of the timeout collateral for the RP
    emit Payment(
      dealId,
      deal.resourceProvider,
      deal.timeoutCollateral,
      SharedStructs.PaymentDirection.Slashed,
      SharedStructs.PaymentReason.TimeoutCollateralSlashed
    );

    emit Timeout(dealId);
  }

  function _hasTimedOut(
    uint256 dealId
  ) private returns (bool) {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    return block.timestamp > agreement.dealAgreedAt + deal.timeout;
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





