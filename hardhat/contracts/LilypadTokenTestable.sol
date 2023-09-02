// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./LilypadToken.sol";

// a version of LilypadToken that can be called by any address
// so we can run unit tests
contract LilypadTokenTestable is LilypadToken {
  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) LilypadToken(name, symbol, initialSupply) {}

  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
