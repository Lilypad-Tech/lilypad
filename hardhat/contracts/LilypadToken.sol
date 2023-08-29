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
  address private escrowAddress;

  // the address that is allowed to manage the escrow balance
  address private controllerAddress;

  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) ERC20(name, symbol) {
    _mint(msg.sender, initialSupply);
  }

  function setEscrowAddress(address _escrowAddress) public onlyOwner {
    require(_escrowAddress != address(0), "LilypadToken: escrow address must be defined");
    escrowAddress = _escrowAddress;
  }

  function setControllerAddress(address _controllerAddress) public onlyOwner {
    require(_controllerAddress != address(0), "LilypadToken: controller address must be defined");
    controllerAddress = _controllerAddress;
  }

  // so we can pay in money to our payments
  function _msgSender() internal view override returns (address) {
    return tx.origin;
  }

  function payinEscrow(
    address fromAddress,
    uint256 amount
  ) public returns (bool) {
    require(fromAddress != address(0), "LilypadToken: fromAddress cannot be zero address");
    require(_msgSender() == fromAddress, "LilypadToken: only message sender can call payinEscrow");
    _transfer(fromAddress, escrowAddress, amount);
    return true;
  }

  function payoutEscrow(
    address toAddress,
    uint256 amount
  ) public returns (bool) {
    require(toAddress != address(0), "LilypadToken: toAddress cannot be zero address");
    require(msg.sender == controllerAddress, "LilypadToken: only controller can call payoutEscrow");
    _transfer(escrowAddress, toAddress, amount);
    return true;
  }
}
