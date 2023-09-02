// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./LilypadPayments.sol";

contract LilypadPaymentsTestable is LilypadPayments {
  function checkControllerAccess() internal pure override {
    // allow the tests to call payment functions directly so
    // we can test simple balance adjustments and event emitting
  }
}
