// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

// a smart contract that is running a job
// it will call the IOnChainJobManager with it's job
// which will end up with the resultsAdded function being called by the manager
interface IOnChainJobClient {
  function resultsAdded(
    string memory dealId,
    string memory resultsId,
    string memory dataId,
    uint256 instructionCount
  ) external;
}
