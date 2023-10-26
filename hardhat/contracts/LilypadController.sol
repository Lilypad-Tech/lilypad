// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ILilypadController.sol";
import "./ILilypadStorage.sol";
import "./ILilypadPayments.sol";
import "./ILilypadMediation.sol";

contract LilypadController is ILilypadController, Ownable, Initializable {

  /**
   * Types
   */
  
  address private storageAddress;
  address private userAddress;
  address private paymentsAddress;
  address private mediationAddress;
  address private jobCreatorAddress;

  ILilypadStorage private storageContract;
  ILilypadPayments private paymentsContract;
  ILilypadMediationHandler private mediationContract;

  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(
    address _storageAddress,
    address _usersAddress,
    address _paymentsAddress,
    address _mediationAddress,
    address _jobCreatorAddress
  ) public initializer {
    setStorageAddress(_storageAddress);
    setUsersAddress(_usersAddress);
    setPaymentsAddress(_paymentsAddress);
    setMediationAddress(_mediationAddress);
    setJobCreatorAddress(_jobCreatorAddress);
  }

  function setStorageAddress(address _storageAddress) public onlyOwner {
    require(_storageAddress != address(0), "Storage address");
    storageAddress = _storageAddress;
    storageContract = ILilypadStorage(storageAddress);
  }

  function getStorageAddress() public view returns(address) {
    return storageAddress;
  }

  function setUsersAddress(address _usersAddress) public onlyOwner {
    require(_usersAddress != address(0), "Users address");
    userAddress = _usersAddress;
  }

  function getUsersAddress() public view returns(address) {
    return userAddress;
  }

  function setPaymentsAddress(address _paymentsAddress) public onlyOwner {
    require(_paymentsAddress != address(0), "Payments address");
    paymentsAddress = _paymentsAddress;
    paymentsContract = ILilypadPayments(_paymentsAddress);
  }

  function getPaymentsAddress() public view returns(address) {
    return paymentsAddress;
  }

  function setMediationAddress(address _mediationAddress) public onlyOwner {
    require(_mediationAddress != address(0), "Mediation address");
    mediationAddress = _mediationAddress;
    mediationContract = ILilypadMediationHandler(_mediationAddress);
  }

  function getMediationAddress() public view returns(address) {
    return mediationAddress;
  }

  function setJobCreatorAddress(address _jobCreatorAddress) public onlyOwner {
    require(_jobCreatorAddress != address(0), "JobCreator address");
    jobCreatorAddress = _jobCreatorAddress;
  }

  function getJobCreatorAddress() public view returns(address) {
    return jobCreatorAddress;
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
    string memory dealId,
    SharedStructs.DealMembers memory members,
    SharedStructs.DealTimeouts memory timeouts,
    SharedStructs.DealPricing memory pricing
  ) public override returns (SharedStructs.Agreement memory) {
    SharedStructs.Deal memory deal = storageContract.ensureDeal(
      dealId,
      members,
      timeouts,
      pricing
    );
    bool isResourceProvider = tx.origin == deal.members.resourceProvider;
    bool isJobCreator = tx.origin == deal.members.jobCreator;
    require(isResourceProvider || isJobCreator, "Only RP / JC");

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
    string memory dealId,
    string memory resultsId,
    // this is the CID of the actual data
    // otherwise onchain clients cannot know the actual data they want to fetch
    string memory dataId,
    uint256 instructionCount
  ) public override {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.DealAgreed), "DealAgreed");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.resourceProvider == tx.origin, "Only RP");

    storageContract.addResult(
      dealId,
      resultsId,
      dataId,
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
    string memory dealId
  ) public override {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "ResultsSubmitted");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.jobCreator == tx.origin, "Only JC");
    
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
    string memory dealId
  ) public override {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "ResultsSubmitted");
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    require(deal.members.jobCreator == tx.origin, "Only JC");

    // this function will require that the mediator is in the RP's list of trusted mediators
    storageContract.checkResult(dealId);
    paymentsContract.checkResult(
      dealId,
      deal.members.jobCreator,
      // this is the JC judging their result so they get their timeout collateral back
      deal.timeouts.judgeResults.collateral,
      deal.pricing.mediationFee
    );

    // trigger the mediation process by calling the contract
    mediationContract.mediationRequest(deal);
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
    string memory dealId
  ) public override {
    require(mediationAddress == _msgSender(), "Only mediation");
    require(_canMediateResult(dealId), "Cannot mediate");
    
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    uint256 jobCost = storageContract.getJobCost(dealId);
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.mediationAcceptResult(dealId);
    paymentsContract.mediationAcceptResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
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
    string memory dealId
  ) public override {
    // only the current mediation contract can call this
    require(mediationAddress == _msgSender(), "Only mediation");
    require(_canMediateResult(dealId), "Cannot mediate");

    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    uint256 resultsCollateral = storageContract.getResultsCollateral(dealId);

    storageContract.mediationRejectResult(dealId);
    paymentsContract.mediationRejectResult(
      dealId,
      deal.members.resourceProvider,
      deal.members.jobCreator,
      deal.pricing.paymentCollateral,
      resultsCollateral,
      deal.pricing.mediationFee
    );
  }

  function _canMediateResult(
    string memory dealId 
  ) private returns (bool) {
    require(storageContract.isState(dealId, SharedStructs.AgreementState.ResultsChecked), "ResultsChecked");
    return true;
  }

  /**
   * Timeouts
   */

  function timeoutAgree(
    string memory dealId
  ) public override {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    require(deal.members.jobCreator == tx.origin || deal.members.resourceProvider == tx.origin, "Only JC or RP");
    require(agreement.state == SharedStructs.AgreementState.DealNegotiating, "Not correct state");
    require(block.timestamp > agreement.dealCreatedAt + deal.timeouts.agree.timeout, "Not timed out");
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
    string memory dealId
  ) public override {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    require(deal.members.jobCreator == tx.origin, "Only JC");
    require(agreement.state == SharedStructs.AgreementState.DealAgreed, "Not correct state");
    require(block.timestamp > agreement.dealAgreedAt + deal.timeouts.submitResults.timeout, "Not timed out");
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
    string memory dealId
  ) public override {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    require(deal.members.resourceProvider == tx.origin, "Only RP");
    require(agreement.state == SharedStructs.AgreementState.ResultsSubmitted, "Not correct state");
    require(block.timestamp > agreement.resultsSubmittedAt + deal.timeouts.judgeResults.timeout, "Not timed out");
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
    string memory dealId
  ) public override {
    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);
    require(deal.members.resourceProvider == tx.origin || deal.members.jobCreator == tx.origin, "Only RP or JC");
    require(agreement.state == SharedStructs.AgreementState.ResultsChecked, "Not correct state");
    require(block.timestamp > agreement.resultsSubmittedAt + deal.timeouts.judgeResults.timeout, "Not timed out");
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
}
