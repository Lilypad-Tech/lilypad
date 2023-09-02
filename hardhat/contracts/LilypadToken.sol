// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/*
  standard ERC20 token but with some additional features:

   * Ownable
   * payinEscrow
   * payoutEscrow

  the escrow functions are designed to be called by the payments contract

   * deploy this contract as admin
   * deploy the payments contract as admin and pass this address to it
   * update the owner of this contract to be the payments contract
  
  now, only the payments contract can call the escrow functions on behalf of address(this)
  
  no other address other than address(this) (i.e. this contract) are affected by these functions

  in other words - these two functions give our payments contract the ability to
  pay into and pay out of the escrow balance regardless of who has signed the tx

 */
contract LilypadToken is Ownable, ERC20 {

  mapping(address => uint256) private escrowBalances;

  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) ERC20(name, symbol) {
    _mint(msg.sender, initialSupply);
  }

  function payinEscrow(
    uint256 amount
  ) public returns (bool) {
    // it's important we use tx.origin and not msg.sender here
    // msg.sender will be the payments contract
    // tx.origin will be the user who called the controller -> payments -> token
    // i.e. the account that is actually paying into the escrow address
    _transfer(tx.origin, address(this), amount);
    return true;
  }

  function payoutEscrow(
    address toAddress,
    uint256 amount
  ) public onlyOwner returns (bool) {
    require(toAddress != address(0), "LilypadToken: toAddress cannot be zero address");
    _transfer(address(this), toAddress, amount);
    return true;
  }
}
