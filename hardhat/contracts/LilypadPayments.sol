// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ILilypadToken.sol";

contract LilypadPayments is Ownable, Initializable {

  /**
   * Types
   */
  address private tokenAddress;
  ILilypadToken private tokenContract;

  /**
   * Enums
   */
  enum PaymentReason {

    // the money the JC puts up to pay for the job
    PaymentCollateral,

    // the money the RP puts up to attest it's results are correct
    ResultsCollateral,

    // the money the RP, JC and Mediator all put up to prevent timeouts
    TimeoutCollateral,

    // the money the RP gets paid for the job for running it successfully
    JobPayment,

    // the money the JC pays the mediator for resolving a dispute
    MediationFee
  }

  enum PaymentDirection {

    // money flowing into the contract
    // i.e. we GET paid
    PaidIn,

    // money paid out to services
    // i.e. we are PAYING
    PaidOut,

    // collateral that is locked up being refunded
    Refunded,
    
    // collateral that is locked up being slashed
    Slashed
  }

  /**
   * Events
   */
  event Payment(
    uint256 indexed dealId,
    address payee,
    uint256 amount,
    PaymentReason reason,
    PaymentDirection direction
  );

  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(address _tokenAddress) public initializer {
    setTokenAddress(_tokenAddress);
  }

  function setTokenAddress(address _tokenAddress) public onlyOwner {
    require(_tokenAddress != address(0), "Token address must be defined");
    tokenAddress = _tokenAddress;
    tokenContract = ILilypadToken(_tokenAddress);
  }

  /**
   * Controller handlers
   * 
   * these methods are called by the controller to wrap various payment
   * scenarios - hence they are all onlyOwner
   */

  /**
   * Agreements
   */

  // * pay in the timeout collateral
  function agreeResourceProvider(
    uint256 dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) public onlyOwner {
    _payIn(timeoutCollateral);
    emit Payment(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.PaidIn
    );
  }

  // * pay in the payment collateral and timeout collateral
  function agreeJobCreator(
    uint256 dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    _payIn(paymentCollateral + timeoutCollateral);
    emit Payment(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.PaidIn
    );
    emit Payment(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.PaidIn
    );
  }

  /**
   * Results
   */

  // * calculate the cost of the job
  // * calculate the job collateral based on the multiple
  // * work out the difference between the timeout and results collateral
  // * pay the difference into / out of the contract to the RP
  function addResult(
    uint256 dealId,
    address resourceProvider,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    // what is the difference between what the RP has already paid and needs to now pay?
    // the RP has paid in the timeout collateral
    // it will now be charged the results collateral
    // a positive number means we are owed money
    // a negative number means we pay the RP a refund
    int256 resultsTimeoutDiff = int256(resultsCollateral) - int256(timeoutCollateral);
    
    if(resultsTimeoutDiff > 0) {
      // the RP pays us because the job collateral is higher than the timeout collateral
      _payIn(uint256(resultsTimeoutDiff));
    }
    else if(resultsTimeoutDiff < 0) {
      // we pay the RP because the job collateral is lower
      _payOut(resourceProvider, uint256(resultsTimeoutDiff));
    }

    // the refund of the timeout collateral
    emit Payment(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Refunded
    );

    // the payment of the job collateral
    emit Payment(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.PaidIn
    );
  }

  // * calculate the cost of the job
  // * deduct the cost of the job from the JC payment collateral
  // * pay the RP the cost of the job
  // * refund the RP the results collateral
  // * refund the JC the job collateral minus the cost
  // * refund the JC the timeout collateral
  function acceptResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    // the difference between the job collateral and the job cost
    // is how much the job creator get's back
    int256 paymentCollateralRefund = int256(paymentCollateral) - int256(jobCost);

    // the job cost more than the job collateral
    // this means the RP get's less than instruction count * instruction price
    // however they took on the deal knowing how much collateral was put up
    // equally however - it's very important to zero this number to prevent
    // the JC getting money back that they didn't pay in
    if(paymentCollateralRefund < 0) {
      paymentCollateralRefund = 0;
    }

    // we pay back the remaining job collateral and timeout collateral to the job creator
    _payOut(jobCreator, uint256(paymentCollateralRefund) + timeoutCollateral);

    if(paymentCollateralRefund > 0) {
      emit Payment(
        dealId,
        jobCreator,
        uint256(paymentCollateralRefund),
        PaymentReason.PaymentCollateral,
        PaymentDirection.Refunded
      );
    }

    emit Payment(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Refunded
    );

    // now we pay back the results collateral and the job payment to the RP
    _payOut(resourceProvider, resultsCollateral + jobCost);

    emit Payment(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Refunded
    );

    emit Payment(
      dealId,
      resourceProvider,
      jobCost,
      PaymentReason.JobPayment,
      PaymentDirection.PaidOut
    );
  }

  // * charge the JC the mediation fee
  // * refund the JC the timeout collateral
  function challengeResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 timeoutCollateral,
    uint256 mediationFee
  ) public onlyOwner {
    // what is the difference between what the JC has already paid and needs to now pay?
    // the JC has paid in the timeout collateral
    // it will now be charged the mediation fee
    // a positive number means we are owed money
    // a negative number means we pay the RP a refund
    int256 timeoutMediationDiff = int256(timeoutCollateral) - int256(mediationFee);

    if(timeoutMediationDiff > 0) {
      // the RP pays us because the job collateral is higher than the timeout collateral
      _payIn(uint256(timeoutMediationDiff));
    }
    else if(timeoutMediationDiff < 0) {
      // we pay the RP because the job collateral is lower
      _payOut(resourceProvider, uint256(timeoutMediationDiff));
    }
    
    // the refund of the timeout collateral
    emit Payment(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Refunded
    );

    // the payment of the mediation fee
    emit Payment(
      dealId,
      jobCreator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.PaidIn
    );
  }

  /**
   * Mediation
   */

  // * refund the JC what is left from the payment collateral (if any)
  // * pay the RP the cost of the job
  // * refund the RP the results collateral
  // * pay the mediator for mediating
  function mediationAcceptResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    address mediator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) public onlyOwner {
    int256 paymentCollateralRefund = int256(paymentCollateral) - int256(jobCost);

    // if there is a refund for the JC then let's pay it
    if(paymentCollateralRefund > 0) {
      _payOut(jobCreator, uint256(paymentCollateralRefund));
      emit Payment(
        dealId,
        jobCreator,
        uint256(paymentCollateralRefund),
        PaymentReason.PaymentCollateral,
        PaymentDirection.Refunded
      );
    }

    // now we pay back the results collateral and the job payment to the RP
    _payOut(resourceProvider, resultsCollateral + jobCost);

    emit Payment(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Refunded
    );

    emit Payment(
      dealId,
      resourceProvider,
      jobCost,
      PaymentReason.JobPayment,
      PaymentDirection.PaidOut
    );

    // pay the mediator
    _payOut(mediator, mediationFee);

    emit Payment(
      dealId,
      mediator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.PaidOut
    );
  }

  // * refund the JC their payment collateral
  // * slash the RP's results collateral
  // * pay the mediator for mediating
  function mediationRejectResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    address mediator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) public onlyOwner {
    // refund the JC their payment collateral
    _payOut(jobCreator, paymentCollateral);

    emit Payment(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );

    // pay the mediator
    _payOut(mediator, mediationFee);

    emit Payment(
      dealId,
      mediator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.PaidOut
    );

    // slash the RP
    emit Payment(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Slashed
    );
  }

  /**
   * Timeouts
   */

  // * pay back the JC's job collateral
  // * slash the RP's results collateral
  function timeoutSubmitResults(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    // refund the job creator
    _payOut(jobCreator, paymentCollateral);

    // the refund of the job collateral to the JC
    emit Payment(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );

    // the slashing of the timeout collateral for the RP
    emit Payment(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Slashed
    );
  }

  // * pay back the RP's results collateral
  // * pay the RP the cost of the job
  // * slash the JC's timeout collateral
  // * slash the JC's job collateral
  function timeoutJudgeResults(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    // refund the resource provider
    _payOut(resourceProvider, resultsCollateral);

    // the refund of the results collateral to the RP
    emit Payment(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );

    // the slashing of the timeout collateral for the RP
    emit Payment(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Slashed
    );
  }

  // * pay back the RP's results collateral
  // * pay back the JC's paymnet collateral
  function timeoutMediateResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral
  ) public onlyOwner {
    // refund the resource provider
    _payOut(resourceProvider, resultsCollateral);

    // refund the job creator
    _payOut(jobCreator, paymentCollateral);

    // the refund of the results collateral to the RP
    emit Payment(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Refunded
    );

    // the refund of the payment collateral to the JC
    emit Payment(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );
  }

  /**
   * Payment utils
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

  // take X tokens from the tx sender and add them to the owners token balance
  function _payIn(
    uint256 amount
  ) private {
    // approve the tokens we are about to move
    // this works because _payIn is always called as part of the user who is paying
    // into the contract
    tokenContract.approve(owner(), amount);
    _pay(
      tx.origin,
      owner(),
      amount
    );
  }

  // take X tokens from the contract's token balance and send them to the given address
  function _payOut(
    address payWho,
    uint256 amount
  ) private {
    _pay(
      owner(),
      payWho,
      amount
    );
  }
}
