// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

library SharedStructs {

  enum ServiceType {
    Solver,
    Mediator,
    ResourceProvider,
    JobCreator
  }

  enum AgreementState {

    // the two parties have not yet both agreed to the deal
    DealNegotiating,

    // both parties have agreed
    DealAgreed,

    // results have been submitted by the RP
    ResultsSubmitted,

    // the JC has accepted the results
    ResultsAccepted,

    // the JC has checked the results
    ResultsChecked,

    // a mediator has accepted the results
    MediationAccepted,

    // a mediator has rejected the results
    MediationRejected,

    // this means the counter-party did not agree to the deal in time
    TimeoutAgree,

    // this means the RP did not submit results in time
    TimeoutSubmitResults,

    // this means the JC did not accept or reject results in time
    TimeoutJudgeResults,

    // this means the mediator did not accept or submit judgement in time
    TimeoutMediateResults

  }

  // we map addresses onto infomation about the user
  struct User {
    address userAddress;
    // the CID of information for this user
    string metadataCID;
    string url;
    ServiceType[] roles;
  }

  // the various addresses involved in runnig a deal
  struct DealMembers {
    // the address of the solver service that the RP and JC have agreed to use
    address solver;
    // the addresses of the RP and JC that have agreed to this deal
    address jobCreator;
    address resourceProvider;
    // the list of mediators that the RP and JC have agreed to use
    address[] mediators;
  }

  // a timeout represents the agreed amount of time and the penalty
  // that is applied if the timeout is exceeded  
  struct DealTimeout {
    uint256 timeout;
    uint256 collateral;
  }
  
  // the various forms of timeout a deal can have
  struct DealTimeouts { 
    DealTimeout agree;
    DealTimeout submitResults;
    DealTimeout judgeResults;
    DealTimeout mediateResults;
  }

  // configure the cost of a deal
  struct DealPricing {
  // agreed price per instruction
    uint256 instructionPrice;

    // the collateral that the JC has put up to pay for the job
    // the final cost of the job will be deducted from this
    uint256 paymentCollateral;

    // how much collateral the RP will post to attest it's results are correct
    // this is a multiple of the cost of the job which is known at results
    // submission time
    uint256 resultsCollateralMultiple;

    // how much is the JC willing to pay the mediator to resolve disputes
    uint256 mediationFee;
  }

  // a Deal forms the information that is agreed between both parties
  // both parties must have called "agree_deal" with the exact
  // same parameters before the deal is considered valid
  // a Deal is immutable - nothing about it can be updated
  struct Deal {
    // the CID of the Deal document on IPFS (and directory service)
    // this contains the job spec, the job offer and the resource offer
    string dealId;

    // who is participating in this deal
    DealMembers members;
    
    // the timeout settings for a deal
    DealTimeouts timeouts;

    // the pricing settings for a deal
    DealPricing pricing;    
  }

  // what the RP submits back once having run the job
  // this is also immutable
  struct Result {
    // the id of the deal that this result is for
    string dealId;

    // the CID of the results on IPFS (and directory service)
    // NOTE - this is not the CID of actual results data rather
    // of the JSON object reporting that data
    string resultsId;

    // this is the actual lower level data CID
    string dataId;

    // how many instructions were executed by the RP
    uint256 instructionCount;
  }

  // an agreement keeps track of the state of a deal and it's fields can be mutated
  struct Agreement {
    // the current state of the agreement
    AgreementState state;

    uint256 resourceProviderAgreedAt;
    uint256 jobCreatorAgreedAt;
    uint256 dealCreatedAt;
    uint256 dealAgreedAt;

    uint256 resultsSubmittedAt;
    uint256 resultsAcceptedAt;
    uint256 resultsCheckedAt;

    uint256 mediationAcceptedAt;
    uint256 mediationRejectedAt;

    uint256 timeoutAgreeAt;
    uint256 timeoutSubmitResultsAt;
    uint256 timeoutJudgeResultsAt;
    uint256 timeoutMediateResultsAt;
  }

  struct JobOffer {
    uint256 id;
    
    // this is the contract that will be triggered
    // once there are some results
    address calling_contract;

    // this is the address that is paying for the job
    // they must have called approve on the token contract
    // and granted the solver address the number of tokens
    // required
    address payee;

    // the job spec
    string module;
    string[] inputs;
  }
}
