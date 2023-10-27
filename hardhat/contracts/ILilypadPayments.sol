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
    string memory dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) external;

  function agreeJobCreator(
    string memory dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) external;

  /**
   * Results
   */

  function addResult(
    string memory dealId,
    address resourceProvider,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) external;

  function acceptResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) external;

  function checkResult(
    string memory dealId,
    address jobCreator,
    uint256 timeoutCollateral,
    uint256 mediationFee
  ) external;

  /**
   * Mediation
   */

  function mediationAcceptResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 jobCost,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) external;

  function mediationRejectResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) external;

  /**
   * Timeouts
   */

  function timeoutAgreeRefundResourceProvider(
    string memory dealId,
    address resourceProvider,
    uint256 timeoutCollateral
  ) external;

  function timeoutAgreeRefundJobCreator(
    string memory dealId,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) external;

  function timeoutSubmitResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 timeoutCollateral
  ) external;

  function timeoutJudgeResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 resultsCollateral,
    uint256 timeoutCollateral
  ) external;

  function timeoutMediateResult(
    string memory dealId,
    address resourceProvider,
    address jobCreator,
    uint256 paymentCollateral,
    uint256 resultsCollateral,
    uint256 mediationFee
  ) external;
  
}
