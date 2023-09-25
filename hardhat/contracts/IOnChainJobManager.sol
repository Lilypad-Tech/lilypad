// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";
import "./IOnChainJobClient.sol";

// the smart contract that is responsible for handling the life
// cycle of a job - it holds the escrow payment paid by client smart contracts
// which will be moved into the controller by the solver once a match
// has been made
// clients can withdraw their escrow at any time up until the agree tx has been posted
// the controller will then call "resultsAdded" once the RP has returned results
interface IOnChainJobManager is IOnChainJobClient {
  function runJob(
    string memory module,
    SharedStructs.JobOfferInput[] memory inputs
  ) external;
}
