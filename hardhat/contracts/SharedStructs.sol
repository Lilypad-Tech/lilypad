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

  // a Deal forms the agreement between both parties
  // both parties must have called "agree_deal" with the exact
  // same parameters before the deal is considered valid
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
    uint256 jobCollateral;

    // how much collateral the RP will post to attest it's results are correct
    uint256 resultsCollateral;
  }

  // so we keep the deal details seperated from the fact that both parties have agreed
  // to the deal - this represents a 2 sided agreement and records the timestamp
  // of when both parties agreed to the deal
  struct Agreement {
    uint256 resourceProviderAgreedAt;
    uint256 jobCreatorAgreedAt;
  }

  // what the RP submits back once having run the job
  struct Result {
    // the id of the deal that this result is for
    uint256 dealId;

    // the CID of the results on IPFS (and directory service)
    uint256 resultsId;

    // how many instructions were executed by the RP
    uint256 instructionCount;
  }
}
