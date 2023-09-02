// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface ILilypadToken is IERC20 {
  function payinEscrow(
    uint256 amount
  ) external returns (bool);

  function payoutEscrow(
    address toAddress,
    uint256 amount
  ) external returns (bool);
}
