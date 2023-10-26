// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface ILilypadToken is IERC20 {
  function escrowBalanceOf(
    address _address
  ) external returns (uint256);

  function payEscrow(
    uint256 amount
  ) external returns (bool);

  function refundEscrow(
    address toAddress,
    uint256 amount
  ) external returns (bool);

  function payJob(
    address fromAddress,
    address toAddress,
    uint256 amount
  ) external returns (bool);

  function slashEscrow(
    address slashedAddress,
    uint256 amount
  ) external returns (bool);
}
