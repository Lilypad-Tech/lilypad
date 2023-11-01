// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Pausable.sol";
import "./ControllerOwnable.sol";

/*
  standard ERC20 token but with some additional features:

   * ControllerOwnable
     * for paying in, we use tx.origin so the actual payee must call the contract
     * for paying out - we use the Controller Ownable feature so only the payments contract
       can pay out from the escrow account
   * escrowBalanceOf
     * get the current escrow balance for an address
   * payEscrow
     * pay into the escrow account
   * refundEscrow
     * get refunded from the escrow account
   * payJob
     * reduce the "from" account by X amount
     * actually pay that amount to the "to" address
   * slashEscrow
     * reduce the "slashed" account by X amount

  the escrow functions are designed to be called by the payments contract

   * deploy this contract as admin
   * deploy the payments contract as admin and pass this address to it
   * update the ControllerOwnable address of this contract to be the payments contract
  
  now, only the payments contract can call the escrow functions that pay out

 */
contract LilypadToken is ControllerOwnable, ERC20, ERC20Pausable {

  // keep track of the current escrow balance for each address
  mapping(address => uint256) private escrowBalances;

  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) ERC20(name, symbol) {
    _mint(msg.sender, initialSupply);
  }

  function escrowBalanceOf(
    address _address
  ) public view returns (uint256) {
    return escrowBalances[_address];
  }

  // money being paid into the escrow account
  function payEscrow(
    uint256 amount
  ) public returns (bool) {
    // it's important we use tx.origin and not msg.sender here
    // msg.sender will be the payments contract
    // tx.origin will be the user who called the controller -> payments -> token
    // i.e. the account that is actually paying into the escrow address
    _transfer(tx.origin, address(this), amount);
    escrowBalances[tx.origin] += amount;
    return true;
  }

  // money being paid back from the escrow account
  function refundEscrow(
    address toAddress,
    uint256 amount
  ) public onlyController returns (bool) {
    require(toAddress != address(0), "LilypadToken: toAddress cannot be zero address");
    require(escrowBalances[toAddress] >= amount, "LilypadToken: not enough funds in escrow");
    escrowBalances[toAddress] -= amount;
    _transfer(address(this), toAddress, amount);
    return true;
  }

  // pay the RP account from the JC escrow account
  function payJob(
    address fromAddress,
    address toAddress,
    uint256 amount
  ) public onlyController returns (bool) {
    require(escrowBalances[fromAddress] >= amount, "LilypadToken: not enough funds in escrow");
    escrowBalances[fromAddress] -= amount;
    _transfer(address(this), toAddress, amount);
    return true;
  }

  // the given party has been slashed so the money stays in the contract
  // TODO: what should happen to slashed funds?
  // at the moment we move them to the owner address so they are not locked
  function slashEscrow(
    address slashedAddress,
    uint256 amount
  ) public onlyController returns (bool) {
    require(escrowBalances[slashedAddress] >= amount, "LilypadToken: not enough funds in escrow");
    escrowBalances[slashedAddress] -= amount;
    _transfer(address(this), owner(), amount);
    return true;
  }

  function pause() public onlyOwner {
      _pause();
  }

  function unpause() public onlyOwner {
      _unpause();
  }

  function _beforeTokenTransfer(address from, address to, uint256 amount) internal override(ERC20, ERC20Pausable) {
      super._beforeTokenTransfer(from, to, amount);
  }
}
