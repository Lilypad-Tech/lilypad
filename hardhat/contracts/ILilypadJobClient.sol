// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

// a smart contract that is running a job
// it will call the IOnChainJobManager with it's job
// which will end up with the resultsAdded function being called by the manager
interface ILilypadJobClient {
  function submitResults(
    uint256 id,
    string memory dealId,
    string memory dataId
  ) external;
}
