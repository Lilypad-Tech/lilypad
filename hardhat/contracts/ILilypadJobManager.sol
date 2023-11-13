// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./ILilypadJobClient.sol";

// the smart contract that is responsible for handling the life
// cycle of a job - it will pay the registered solver to manage the job
// this is the simple implementation - later we can implement the whole protocol
// where the calling smart contract can decide if it wants to agree to a resource offer
// and accept results
interface ILilypadJobManager is ILilypadJobClient {
  function runJob(
    string memory module,
    string[] memory inputs,
    address payee
  ) external returns (uint256);
}
