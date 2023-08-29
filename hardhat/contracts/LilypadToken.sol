// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract LilypadToken is ERC20 {

  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) ERC20(name, symbol) {
    _mint(msg.sender, initialSupply);
  }

  /**
   * utils
   */

  // we override this because various functions will be called via the controller
  // and we want to know who the original sender was to move THEIR tokens
  // not the controller contract's tokens
  function _msgSender() internal view override returns (address) {
    return tx.origin;
  }
}
