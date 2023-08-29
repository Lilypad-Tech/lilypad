// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

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
    uint256 dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) external;

  function agreeJobCreator(
    uint256 dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) external;

  /**
   * Results
   */

  function addResult(
    uint256 dealId,
    address resourceProvider,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) external;

  function acceptResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) external;

  function challengeResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 timeoutCollateral,
    uint256 mediationFee
  ) external;

  /**
   * Mediation
   */

  function mediationAcceptResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    address mediator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) external;

  function mediationRejectResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    address mediator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) external;

  /**
   * Timeouts
   */

  function timeoutSubmitResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) external;

  function timeoutJudgeResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) external;

  function timeoutMediateResult(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral
  ) external;
  
}
