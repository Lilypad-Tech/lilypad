// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";

// as well as being ownable (i.e. our admin wallet)
// this contract has a modifier called onlyController
// it works the same way as Ownable but only the Owner
// can change the controller address so it gives us a way of
// re-pointing contracts if needed
contract ControllerOwnable is Ownable {

  // the address of the controller that is allowed to call functions
  address private controllerAddress;

  // used to "freeze" the controller address - even the admin cannot
  // change it from this point onwards
  bool private canChangeControllerAddress = true;

  modifier onlyController() {
    _checkControllerAccess();
    _;
  }

  function _checkControllerAccess() internal view virtual returns (bool) {
    require(controllerAddress != address(0), "ControllerOwnable: Controller address must be defined");
    require(_msgSender() == controllerAddress, "ControllerOwnable: Only the controller can call this method");
    return true;
  }

  function setControllerAddress(address _controllerAddress) public onlyOwner {
    require(_controllerAddress != address(0), "ControllerOwnable: Controller address must be defined");
    require(canChangeControllerAddress, "ControllerOwnable: canChangeControllerAddress is disabled");
    controllerAddress = _controllerAddress;
  }

  function getControllerAddress() public view returns (address) {
    return controllerAddress;
  }

  function disableChangeControllerAddress() public onlyOwner {
    canChangeControllerAddress = false;
  }
}
