// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ILilypadToken.sol";

import "@openzeppelin/contracts/utils/Strings.sol";
import "hardhat/console.sol";
// console.log("ensureDeal");
// console.log(Strings.toString(uint256(SharedStructs.AgreementState.DealNegotiating)));
// console.log(Strings.toString(uint256(agreements[dealId].state)));

contract LilypadPayments is Ownable, Initializable {

  /**
   * Types
   */

  // the address of the LilypadToken contract
  address private tokenAddress;
  ILilypadToken private tokenContract;

  // keep track of the current escrow balance for each address
  mapping(address => uint256) private escrowBalances;

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

  // used for debugging
  mapping(address => string) private accountNames;

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(
    address _tokenAddress
  ) public initializer {
    setTokenAddress(_tokenAddress);

    // this is only for debugging
    accountNames[address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)] = "admin";
    accountNames[address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8)] = "faucet";
    accountNames[address(0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC)] = "solver";
    accountNames[address(0x90F79bf6EB2c4f870365E785982E1f101E93b906)] = "mediator";
    accountNames[address(0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65)] = "resource_provider";
    accountNames[address(0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc)] = "job_creator";
    accountNames[address(0x976EA74026E726554dB657fA54763abd0C3a0aa9)] = "directory";
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

  function getEscrowBalance(
    address _address
  ) public view returns (uint256) {
    return escrowBalances[_address];
  }

  /**
   * Agreements
   */

  // * pay in the timeout collateral
  function agreeResourceProvider(
    uint256 dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) public {
    require(tx.origin == resourceProvider, "Can only be called by the RP");
    _pay(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.PaidIn
    );
    escrowBalances[resourceProvider] += timeoutCollateral;
  }

  // * pay in the payment collateral and timeout collateral
  function agreeJobCreator(
    uint256 dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    _pay(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.PaidIn
    );
    escrowBalances[jobCreator] += paymentCollateral;

    _pay(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.PaidIn
    );
    escrowBalances[jobCreator] += timeoutCollateral;
  }

  /**
   * Results
   */

  // * pay back the timeout collateral
  // * pay in the results collateral
  function addResult(
    uint256 dealId,
    address resourceProvider,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) public onlyOwner {
    // the refund of the timeout collateral
    _pay(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[resourceProvider] -= timeoutCollateral;

    // the payment of the job collateral
    _pay(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.PaidIn
    );
    escrowBalances[resourceProvider] += resultsCollateral;
  }

  // * pay the RP the job cost
  // * refund the RP the results collateral
  // * refund the JC the job collateral minus the job cost
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

    // pay the RP the job cost
    _pay(
      dealId,
      resourceProvider,
      jobCost,
      PaymentReason.JobPayment,
      PaymentDirection.PaidOut
    );

    // the difference between the job collateral and the job cost
    // is how much the job creator get's back
    int256 jcRefund = int256(paymentCollateral) - int256(jobCost);

    // if the job cost more than the payment collateral then we shold not go negative
    // otherwise we are paying out more than the JC has put in
    //
    // the RP is loosing out a bit here but they agreed to doing the job
    if(jcRefund > 0) {
      _pay(
        dealId,
        jobCreator,
        uint256(jcRefund),
        PaymentReason.PaymentCollateral,
        PaymentDirection.Refunded
      );
    }

    // this accounts for the job cost and the potential refund of the payment collateral
    escrowBalances[jobCreator] -= paymentCollateral;

    // pay the JC the timeout collateral
    _pay(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[jobCreator] -= timeoutCollateral;

    // pay the RP results collateral
    _pay(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[resourceProvider] -= resultsCollateral;
  }

  // * charge the JC the mediation fee
  // * refund the JC the timeout collateral
  function checkResult(
    uint256 dealId,
    address jobCreator,
    uint256 timeoutCollateral,
    uint256 mediationFee
  ) public onlyOwner {
    // the refund of the timeout collateral
    _pay(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[jobCreator] -= timeoutCollateral;

    // the payment of the mediation fee
    _pay(
      dealId,
      jobCreator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.PaidIn
    );
    escrowBalances[jobCreator] += mediationFee;
  }

  /**
   * Mediation
   */

  // * pay the RP the job cost
  // * refund the RP the results collateral
  // * refund the JC the job collateral minus the job cost
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

    // pay the RP the job cost
    _pay(
      dealId,
      resourceProvider,
      jobCost,
      PaymentReason.JobPayment,
      PaymentDirection.PaidOut
    );

    // pay the mediator the fee
    _pay(
      dealId,
      mediator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.PaidOut
    );
    escrowBalances[jobCreator] -= mediationFee;

    // the difference between the job collateral and the job cost
    // is how much the job creator get's back
    int256 jcRefund = int256(paymentCollateral) - int256(jobCost);

    // if the job cost more than the payment collateral then we shold not go negative
    // otherwise we are paying out more than the JC has put in
    //
    // the RP is loosing out a bit here but they agreed to doing the job
    if(jcRefund > 0) {
      _pay(
        dealId,
        jobCreator,
        uint256(jcRefund),
        PaymentReason.PaymentCollateral,
        PaymentDirection.Refunded
      );
    }

    // this accounts for the job cost and the potential refund of the payment collateral
    escrowBalances[jobCreator] -= paymentCollateral;

    _pay(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[resourceProvider] -= resultsCollateral;
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
    _pay(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[jobCreator] -= paymentCollateral;

    _pay(
      dealId,
      mediator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.PaidOut
    );
    escrowBalances[jobCreator] -= mediationFee;

    // slash the RP
    _pay(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Slashed
    );
    escrowBalances[resourceProvider] -= resultsCollateral;
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
    // the refund of the job collateral to the JC
    _pay(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[jobCreator] -= paymentCollateral;

    // the slashing of the timeout collateral for the RP
    _pay(
      dealId,
      resourceProvider,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Slashed
    );
    escrowBalances[resourceProvider] -= timeoutCollateral;
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
    // the refund of the results collateral to the RP
    _pay(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[resourceProvider] -= resultsCollateral;

    // the slashing of the timeout collateral for the RP
    _pay(
      dealId,
      jobCreator,
      timeoutCollateral,
      PaymentReason.TimeoutCollateral,
      PaymentDirection.Slashed
    );
    escrowBalances[jobCreator] -= timeoutCollateral;
  }

  // * pay back the RP's results collateral
  // * pay back the JC's payment collateral
  // * pay back the JC's mediation fee
  function timeoutMediateResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) public onlyOwner {
    // the refund of the results collateral to the RP
    _pay(
      dealId,
      resourceProvider,
      resultsCollateral,
      PaymentReason.ResultsCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[resourceProvider] -= resultsCollateral;

    // the refund of the payment collateral to the JC
    _pay(
      dealId,
      jobCreator,
      paymentCollateral,
      PaymentReason.PaymentCollateral,
      PaymentDirection.Refunded
    );
    escrowBalances[jobCreator] -= paymentCollateral;

    // the refund of the mediation fee to the JC
    _pay(
      dealId,
      jobCreator,
      mediationFee,
      PaymentReason.MediationFee,
      PaymentDirection.Refunded
    );
    escrowBalances[jobCreator] -= mediationFee;
  }

  /**
   * Payment utils
   */

  // this is always called by the spender of the token
  // and so even though the controller calls the payment contract
  // the token is configured to use tx.origin as the spender
  // i.e. the owner of the tokens is who is calling this
  function _payIn(
    uint256 amount
  ) private {
    require(tokenContract.balanceOf(tx.origin) >= amount, "Insufficient balance");

    console.log("PAY IN");
    console.log(accountNames[tx.origin]);
    console.log(amount);

    bool success = tokenContract.payinEscrow(amount);
    require(success, "Transfer failed");
  }

  // take X tokens from the contract's token balance and send them to the given address
  function _payOut(
    address payWho,
    uint256 amount
  ) private {
    require(tokenContract.balanceOf(tokenAddress) >= amount, "Insufficient balance");

    console.log("PAY OUT");
    console.log(accountNames[payWho]);
    console.log(amount);

    bool success = tokenContract.payoutEscrow(payWho, amount);
    require(success, "Transfer failed");
  }

  function _pay(
    uint256 dealId,
    address payee,
    uint256 amount,
    PaymentReason reason,
    PaymentDirection direction
  ) private {
    if(direction == PaymentDirection.Refunded) {
      require(escrowBalances[payee] >= amount, "Insufficient balance");
    }
    if(direction == PaymentDirection.PaidIn) {
      _payIn(amount);
    }
    else if(direction == PaymentDirection.PaidOut || direction == PaymentDirection.Refunded) {
      _payOut(payee, amount);
    }
    emit Payment(
      dealId,
      payee,
      amount,
      reason,
      direction
    );
  }
}
