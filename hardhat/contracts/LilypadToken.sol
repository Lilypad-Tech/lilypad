// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/*
  standard ERC20 token but with 2 key changes

   * use tx.origin instead of msg.sender
   * `payFromEscrow` function
  
  ## tx.origin

  we use tx.origin so that actual wallet owners can call controller functions
  that pass through to the payments contract which then passes the tx through to here

  equally, this is just an ERC-20 contract where any user can move their funds
  around freely

  so - the one change we make is to use tx.origin instead of msg.sender
  this means the ORIGINAL sender of the tx is who pays into the contract
  and not the address of the calling contract (most likely the payments contract)

  in the case that a wallet owner calls a function directly on this contract
  it's the same behaviour as msg.sender

  ## payFromEscrow()

  we need a way for our controller (i.e. payments) contract to be able to pay out to users
  so we manage:
  
   * "_escrowAddress" - the address that "has" the escrow balance
   * "_controllerAddress" - the address that can manage the escrow balance

  this address will accumulate and pay out tokens to/from users

 */
contract LilypadToken is Ownable, ERC20 {

  // the address the escrow is paid into and out of
  address private _escrowAddress;

  // the address that is allowed to manage the escrow balance
  address private _controllerAddress;

  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply,
    address escrowAddress,
    address controllerAddress
  ) ERC20(name, symbol) {
    require(escrowAddress != address(0), "LilypadToken: escrowAddress cannot be zero address");
    require(controllerAddress != address(0), "LilypadToken: controllerAddress cannot be zero address");
    _mint(msg.sender, initialSupply);
    _escrowAddress = escrowAddress;
    _controllerAddress = controllerAddress;
  }

  // so we can pay in money to our payments
  function _msgSender() internal view override returns (address) {
    return tx.origin;
  }

  function payoutEscrow(
    address toAddress,
    uint256 amount
  ) public returns (bool) {
    require(msg.sender == _controllerAddress, "LilypadToken: only controller can pay from escrow");
    require(toAddress != address(0), "LilypadToken: toAddress cannot be zero address");
    _transfer(_escrowAddress, toAddress, amount);
    return true;
  }
}
