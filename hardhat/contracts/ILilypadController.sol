// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadController {

  function agree(
    string memory dealId,
    SharedStructs.DealMembers memory members,
    SharedStructs.DealTimeouts memory timeouts,
    SharedStructs.DealPricing memory pricing
  ) external returns (SharedStructs.Agreement memory);

  function addResult(
    string memory dealId,
    string memory resultsId,
    string memory dataId,
    uint256 instructionCount
  ) external;

  function acceptResult(
    string memory dealId
  ) external;
  
  function checkResult(
    string memory dealId
  ) external;

  function mediationAcceptResult(
    string memory dealId
  ) external;

  function mediationRejectResult(
    string memory dealId
  ) external;

  function timeoutAgree(
    string memory dealId
  ) external;

  function timeoutSubmitResult(
    string memory dealId
  ) external;

  function timeoutJudgeResult(
    string memory dealId
  ) external;
  
  function timeoutMediateResult(
    string memory dealId
  ) external;
}
