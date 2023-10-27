// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadStorage {
  
  /**
   * Deals
   */

  function getDeal(
    string memory dealId
  ) external returns (SharedStructs.Deal memory);

  function getDealsForParty(
    address party
  ) external returns (uint256[] memory);

  function ensureDeal(
    string memory dealId,
    SharedStructs.DealMembers memory members,
    SharedStructs.DealTimeouts memory timeouts,
    SharedStructs.DealPricing memory pricing
  ) external returns (SharedStructs.Deal memory);

  /**
   * Agreements
   */

  function getAgreement(
    string memory dealId
  ) external returns (SharedStructs.Agreement memory);
  
  function agreeResourceProvider(
    string memory dealId
  ) external returns (SharedStructs.Agreement memory);

  function agreeJobCreator(
    string memory dealId
  ) external returns (SharedStructs.Agreement memory);
  
  /**
   * Post Results
   */

  function getResult(
    string memory dealId
  ) external returns (SharedStructs.Result memory);

  function addResult(
    string memory dealId,
    string memory resultsId,
    string memory dataId,
    uint256 instructionCount
  ) external returns (SharedStructs.Result memory);

  /**
   * Judge Results
   */

  function acceptResult(
    string memory dealId
  ) external;

  function checkResult(
    string memory dealId
  ) external;

  /**
   * Mediation
   */
  
  function mediationAcceptResult(
    string memory dealId
  ) external;

  function mediationRejectResult(
    string memory dealId
  ) external;

  /**
   * Timeouts
   */
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

  /**
   * Costings
   */

  function getJobCost(
    string memory dealId
  ) external returns (uint256);

  function getResultsCollateral(
    string memory dealId
  ) external returns (uint256);
  
  /**
   * Checkers
   */

  function hasDeal(
    string memory dealId
  ) external returns (bool);

  function isState(
    string memory dealId,
    SharedStructs.AgreementState state
  ) external returns (bool);
  
}
