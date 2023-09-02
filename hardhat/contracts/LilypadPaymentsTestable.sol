// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./LilypadPayments.sol";

contract LilypadPaymentsTestable is LilypadPayments {
  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
