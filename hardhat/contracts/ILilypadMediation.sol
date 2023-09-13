// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

// this is our controller contract
// it will be called by the mediation handler once a decision has been reached
interface ILilypadMediationRequester {
  function mediationAcceptResult(
    string memory dealId
  ) external;

  function mediationRejectResult(
    string memory dealId
  ) external;
}

// this is the contract that will handle mediating a deal
interface ILilypadMediationHandler {
  function mediationRequest(
    SharedStructs.Deal memory deal 
  ) external; 
}
