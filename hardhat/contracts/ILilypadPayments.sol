// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadPayments {

  /**
   * Init
   */

  function setTokenAddress(
    address _tokenAddress
  ) external;

  /**
   * Agreements
   */
  function agreeResourceProvider(
    SharedStructs.Deal memory deal
  ) external;

  function agreeJobCreator(
    SharedStructs.Deal memory deal
  ) external;

  /**
   * Results
   */

  function addResult(
    SharedStructs.Deal memory deal,
    uint256 resultsCollateral
  ) external;

  function acceptResult(
    SharedStructs.Deal memory deal,
    uint256 jobCost,
    uint256 resultsCollateral
  ) external;

  function challengeResult(
    SharedStructs.Deal memory deal
  ) external;

  /**
   * Mediation
   */

  function mediationAcceptResult(
    SharedStructs.Deal memory deal,
    address mediator,
    uint256 jobCost,
    uint256 resultsCollateral
  ) external;

  function mediationRejectResult(
    SharedStructs.Deal memory deal,
    address mediator,
    uint256 resultsCollateral
  ) external;

  /**
   * Timeouts
   */

  function timeoutSubmitResult(
    SharedStructs.Deal memory deal
  ) external;

  function timeoutJudgeResult(
    SharedStructs.Deal memory deal,
    uint256 resultsCollateral
  ) external;

  function timeoutMediateResult(
    SharedStructs.Deal memory deal,
    uint256 resultsCollateral
  ) external;
  
}
