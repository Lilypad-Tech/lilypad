// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ILilypadToken.sol";
import "./ControllerOwnable.sol";
import "./ILilypadPayments.sol";

// import "@openzeppelin/contracts/utils/Strings.sol";
// import "hardhat/console.sol";
// console.log("ensureDeal");
// console.log(Strings.toString(uint256(SharedStructs.AgreementState.DealNegotiating)));
// console.log(Strings.toString(uint256(agreements[dealId].state)));

contract LilypadPayments is ControllerOwnable, Initializable {

  /**
   * Types
   */

  // the address of the LilypadToken contract
  address private tokenAddress;
  ILilypadToken private tokenContract;

  // used to cut off upgrades for the remote contract
  bool private canChangeTokenAddress = true;

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
    string dealId,
    address payee,
    uint256 amount,
    PaymentReason reason,
    PaymentDirection direction
  );

  /**
   * Init
   */

  // used for debugging
  mapping(address => string) private accountNames;

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(
    address _tokenAddress
  ) public initializer {
    setTokenAddress(_tokenAddress);

    // this is only for debugging
    // accountNames[address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)] = "admin";
    // accountNames[address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8)] = "faucet";
    // accountNames[address(0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC)] = "solver";
    // accountNames[address(0x90F79bf6EB2c4f870365E785982E1f101E93b906)] = "mediator";
    // accountNames[address(0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65)] = "resource_provider";
    // accountNames[address(0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc)] = "job_creator";
    // accountNames[address(0x976EA74026E726554dB657fA54763abd0C3a0aa9)] = "directory";
  }

  // the LilypadToken we are calling payinEscrow and payoutEscrow on
  function setTokenAddress(address _tokenAddress) public onlyOwner {
    require(_tokenAddress != address(0), "LilepadPayments: Token address must be defined");
    require(canChangeTokenAddress, "LilypadToken: canChangeTokenAddress is disabled");
    tokenAddress = _tokenAddress;
    tokenContract = ILilypadToken(_tokenAddress);
  }

  function getTokenAddress() public view returns(address) {
    return tokenAddress;
  }

  // set for canChangePaymentsAddress
  function disableChangeTokenAddress() public onlyOwner {
    canChangeTokenAddress = false;
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
    string memory dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) public onlyController {
    // we check this here to double check who we are about to charge (the RP)
    // is who signed the TX and so we can take the money
    require(tx.origin == resourceProvider, "LilypadPayments: Can only be called by the RP");
    _payEscrow(
      dealId,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  // * pay in the payment collateral and timeout collateral
  function agreeJobCreator(
    string memory dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == jobCreator, "LilypadPayments: Can only be called by the JC");
    _payEscrow(
      dealId,
      paymentCollateral,
      PaymentReason.PaymentCollateral
    );
    _payEscrow(
      dealId,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  /**
   * Results
   */

  // * pay back the timeout collateral
  // * pay in the results collateral
  function addResult(
    string memory dealId,
    address resourceProvider,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == resourceProvider, "LilypadPayments: Can only be called by the RP");
    _payEscrow(
      dealId,
      resultsCollateral,
      PaymentReason.ResultsCollateral
    );
    _refundEscrow(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  // * pay the RP the job cost
  // * refund the RP the results collateral
  // * refund the JC the job collateral minus the job cost
  // * refund the JC the timeout collateral
  function acceptResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == jobCreator, "LilypadPayments: Can only be called by the JC");

    // what if the final job cost is more than the payment collateral?
    // well - we have to cap the job cost at that collateral
    // true - the RP has lost money but they agreed to the deal
    uint256 actualPayment = jobCost;
    uint256 jcRefund = 0;
    if(jobCost > paymentCollateral) {
      actualPayment = paymentCollateral;
    } else {
      jcRefund = paymentCollateral - jobCost;
    }

    // pay the RP the actualPayment
    _payOut(
      dealId,
      jobCreator,
      resourceProvider,
      actualPayment,
      PaymentReason.JobPayment
    );

    // if the job cost more than the payment collateral then we shold not go negative
    // otherwise we are paying out more than the JC has put in
    //
    // the RP is loosing out a bit here but they agreed to doing the job
    if(jcRefund > 0) {
      _refundEscrow(
        dealId,
        jobCreator,
        jcRefund,
        PaymentReason.PaymentCollateral
      );
    }

    // refund the JC timeout collateral
    _refundEscrow(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );

    // refund the RP results collateral
    _refundEscrow(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral
    );
  }

  // * charge the JC the mediation fee
  // * refund the JC the timeout collateral
  function checkResult(
    string memory dealId,
    address jobCreator,
    uint256 timeoutCollateral,
    uint256 mediationFee
  ) public onlyController {
    require(tx.origin == jobCreator, "LilypadPayments: Can only be called by the JC");
    
    // the refund of the timeout collateral
    _refundEscrow(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );

    // the payment of the mediation fee
    _payEscrow(
      dealId,
      mediationFee,
      PaymentReason.MediationFee
    );
  }

  /**
   * Mediation
   */

  // * pay the RP the job cost
  // * refund the RP the results collateral
  // * refund the JC the job collateral minus the job cost
  // * pay the mediator for mediating
  function mediationAcceptResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) public onlyController {
    uint256 actualPayment = jobCost;
    uint256 jcRefund = 0;
    if(jobCost > paymentCollateral) {
      actualPayment = paymentCollateral;
    } else {
      jcRefund = paymentCollateral - jobCost;
    }
    
    // pay the RP the job cost from the JC
    _payOut(
      dealId,
      jobCreator,
      resourceProvider,
      actualPayment,
      PaymentReason.JobPayment
    );

    // pay the mediator the fee from the JC
    _payOut(
      dealId,
      jobCreator,
      tx.origin,
      mediationFee,
      PaymentReason.MediationFee
    );

    // if the job cost more than the payment collateral then we shold not go negative
    // otherwise we are paying out more than the JC has put in
    //
    // the RP is loosing out a bit here but they agreed to doing the job
    if(jcRefund > 0) {

      // refund the JC the diff between payment collateral and job cost
      _refundEscrow(
        dealId,
        jobCreator,
        jcRefund,
        PaymentReason.PaymentCollateral
      );
    }

    // refund the RP the results collateral
    _refundEscrow(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral
    );
  }

  // * refund the JC their payment collateral
  // * slash the RP's results collateral
  // * pay the mediator for mediating
  function mediationRejectResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) public onlyController {
    // refund the JC their payment collateral
    _refundEscrow(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral
    );

    // pay the mediator the fee from the JC
    _payOut(
      dealId,
      jobCreator,
      tx.origin,
      mediationFee,
      PaymentReason.MediationFee
    );

    // slash the RP
    _slashEscrow(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral
    );
  }

  /**
   * Timeouts
   */
  function timeoutAgreeRefundResourceProvider(
    string memory dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == resourceProvider, "LilypadPayments: Can only be called by the RP");
    // the refund of the job collateral to the JC
    _refundEscrow(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  function timeoutAgreeRefundJobCreator(
    string memory dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == jobCreator, "LilypadPayments: Can only be called by the JC");
    // the refund of the job collateral to the JC
    _refundEscrow(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral
    );

    // the refund of the job collateral to the JC
    _refundEscrow(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  // * pay back the JC's job collateral
  // * pay back the JC's timeout collateral
  // * slash the RP's results collateral
  function timeoutSubmitResults(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == jobCreator, "LilypadPayments: Can only be called by the JC");
    // the refund of the job collateral to the JC
    _refundEscrow(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral
    );

    // the refund of the job collateral to the JC
    _refundEscrow(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
    
    // the slashing of the timeout collateral for the RP
    _slashEscrow(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  // * pay back the RP's results collateral
  // * pay the RP the cost of the job
  // * slash the JC's timeout collateral
  // * slash the JC's job collateral
  function timeoutJudgeResults(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyController {
    require(tx.origin == resourceProvider, "LilypadPayments: Can only be called by the RP");
    // the refund of the results collateral to the RP
    _refundEscrow(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral
    );

    // the slashing of the timeout collateral for the RP
    _slashEscrow(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral
    );
  }

  // * pay back the RP's results collateral
  // * pay back the JC's payment collateral
  // * pay back the JC's mediation fee
  function timeoutMediateResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) public onlyController {
    require(tx.origin == resourceProvider || tx.origin == jobCreator, "LilypadPayments: Can only be called by the RP or JC");
    // the refund of the results collateral to the RP
    _refundEscrow(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral
    );

    // the refund of the payment collateral to the JC
    _refundEscrow(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral
    );

    // the refund of the mediation fee to the JC
    _refundEscrow(
      dealId,
      jobCreator,
      mediationFee,
      PaymentReason.MediationFee
    );
  }

  /**
   * Payment utils
   */


  function _payEscrow(
    string memory dealId,
    uint256 amount,
    PaymentReason reason
  ) private {
    // we check they have that much in their token balance before moving to tokens to us
    require(tokenContract.balanceOf(tx.origin) >= amount, "LilypadPayments: Insufficient balance");

    // console.log("_payEscrow");
    // console.log(accountNames[tx.origin]);
    // console.log(amount);

    bool success = tokenContract.payEscrow(amount);
    require(success, "LilypadPayments: Pay escrow failed");

    emit Payment(
      dealId,
      tx.origin,
      amount,
      reason,
      PaymentDirection.PaidIn
    );
  }

  function _refundEscrow(
    string memory dealId,
    address toAddress,
    uint256 amount,
    PaymentReason reason
  ) private {
    // console.log("_refundEscrow");
    // console.log(accountNames[toAddress]);
    // console.log(amount);

    bool success = tokenContract.refundEscrow(toAddress, amount);
    require(success, "LilypadPayments: Refund escrow failed");

    emit Payment(
      dealId,
      toAddress,
      amount,
      reason,
      PaymentDirection.Refunded
    );
  }

  function _payOut(
    string memory dealId,
    address fromAddress,
    address toAddress,
    uint256 amount,
    PaymentReason reason
  ) private {
    // console.log("_payJob");
    // console.log(accountNames[fromAddress]);
    // console.log(accountNames[toAddress]);
    // console.log(amount);

    bool success = tokenContract.payJob(fromAddress, toAddress, amount);
    require(success, "LilypadPayments: Pay job failed");

    emit Payment(
      dealId,
      toAddress,
      amount,
      reason,
      PaymentDirection.PaidOut
    );
  }

  function _slashEscrow(
    string memory dealId,
    address slashedAddress,
    uint256 amount,
    PaymentReason reason
  ) private {
    // console.log("_slashEscrow");
    // console.log(accountNames[slashedAddress]);
    // console.log(amount);

    bool success = tokenContract.slashEscrow(slashedAddress, amount);
    require(success, "LilypadPayments: Slash escrow failed");

    emit Payment(
      dealId,
      slashedAddress,
      amount,
      reason,
      PaymentDirection.Slashed
    );
  }
}
