// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

library SharedStructs {

  enum ServiceType {
    Solver,
    Mediator,
    Directory,
    ResourceProvider,
    JobCreator
  }

  enum AgreementState {
    Negotiating,
    Agreed,
    Timeout,
    Submitted,
    Accepted,
    Rejected
  }

  enum PaymentDirection {
    // money flowing into the contract
    // i.e. we GET paid
    In,
    // money flowing out of the contract
    // i.e. we PAY someone
    Out,
    // money that has not moved but is now part of the contract's "pool"
    // this money will not be recovered by whoever it was slashed from
    Slashed
  }

  enum PaymentReason {

    // the money the JC puts up to pay for the job
    // this is part of the deal the RP has to agree to
    JobCollateral, // JC -> contract.agree

    // the refund the JC gets if the results are accepted
    // this is the amount they overpaid for the job OR
    JobCollateralRefund, // contract.acceptResult -> JC

    // the refund the JC gets if the results timed out
    // this is the total amount
    JobCollateralTimeoutRefund, // contract.refundTimeout -> JC

    // the money the RP puts up to attest it's results are correct
    ResultsCollateral, // RP -> contract.addResult

    // the refund of the ResultsCollateral the RP gets if the results are accepted
    ResultsCollateralRefund, // contract.acceptResult -> RP

    // the RP list it's results collateral because the results were rejected
    ResultsCollateralSlashed, // contract.rejectResult -> contract pool

    // the money the RP gets paid for the job for running it successfully
    JobPayment, // contract.acceptResult -> RP

    // the money the RP puts up to prevent timeouts
    TimeoutCollateral, // RP -> contract.agree

    // when the results are posted in time
    // the RP gets this back
    TimeoutCollateralRefund, // contract.addResult -> RP

    // the RP lost it's timeout collateral because it didn't post results in time
    // this balance is now part of the smart contract
    TimeoutCollateralSlashed // contract.refundTimeout -> contract pool
  }

  // we map addresses onto infomation about the user
  struct User {
    address userAddress;
    // the CID of information for this user
    uint256 metadataCID;
    string url;
    ServiceType[] roles;
    address[] trustedMediators;
    address[] trustedDirectories;
  }

  // a Deal forms the information that is agreed between both parties
  // both parties must have called "agree_deal" with the exact
  // same parameters before the deal is considered valid
  // a Deal is immutable - nothing about it can be updated
  struct Deal {
    // the CID of the Deal document on IPFS (and directory service)
    // this contains the job spec, the job offer and the resource offer
    uint256 dealId;

    // the addresses of the RP and JC that have agreed to this deal
    address resourceProvider;
    address jobCreator;

    // agreed price per instruction
    uint256 instructionPrice;

    // the max time we will wait for the results
    uint256 timeout;

    // the collateral that the RP has put up to prevent timeouts
    uint256 timeoutCollateral;

    // the collateral that the JC has put up to pay for the job
    // the final cost of the job will be deducted from this
    uint256 jobCollateral;

    // how much collateral the RP will post to attest it's results are correct
    // this is a multiple of the cost of the job which is known at results
    // submission time
    uint256 resultsCollateralMultiple;
  }

  // what the RP submits back once having run the job
  // this is also immutable
  struct Result {
    // the id of the deal that this result is for
    uint256 dealId;

    // the CID of the results on IPFS (and directory service)
    uint256 resultsId;

    // how many instructions were executed by the RP
    uint256 instructionCount;
  }

  // an agreement keeps track of the state of a deal and it's fields can be mutated
  struct Agreement {
    // the current state of the agreement
    AgreementState state;
    uint256 resourceProviderAgreedAt;
    uint256 jobCreatorAgreedAt;
    uint256 dealAgreedAt;
    uint256 timedOutAt;
    uint256 resultsSubmittedAt;
    uint256 resultsAcceptedAt;
    uint256 resultsRejectedAt;
  }
}
