// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ControllerOwnable.sol";

contract LilypadUsers is Ownable, Initializable {

  // a map of user address -> user
  mapping(address => SharedStructs.User) private users;

  // a map of user type => user address[]
  mapping(SharedStructs.ServiceType => address[]) private usersByType;

  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize() public initializer {
    
  }

  /**
   * Users
   */

  function getUser(
    address userAddress
  ) public view returns (SharedStructs.User memory) {
    return users[userAddress];
  }

  function updateUser(
    string memory metadataCID,
    string memory url,
    SharedStructs.ServiceType[] memory roles
  ) public returns (SharedStructs.User memory) {
    SharedStructs.User memory newUser = SharedStructs.User(
      tx.origin,
      metadataCID,
      url,
      roles
    );
    users[tx.origin] = newUser;
    return newUser;
  }

  function addUserToList(
    SharedStructs.ServiceType serviceType
  ) public {
    require(users[tx.origin].userAddress != address(0), "User must exist");
    (, bool found) = _getUserListIndex(serviceType, tx.origin);
    require(!found, "User is already part of that list");
    require(_doesUserHaveRole(serviceType, tx.origin), "User must have that role");
    usersByType[serviceType].push(tx.origin);
  }

  function removeUserFromList(
    SharedStructs.ServiceType serviceType
  ) public {
    require(users[tx.origin].userAddress != address(0), "User must exist");
    (uint256 index, bool found) = _getUserListIndex(serviceType, tx.origin);
    require(found, "User is not part of that list");
    for (uint256 i = index; i < usersByType[serviceType].length - 1; i++) {
      usersByType[serviceType][i] = usersByType[serviceType][i + 1];
    }
    usersByType[serviceType].pop();
  }

  function showUsersInList(
    SharedStructs.ServiceType serviceType
  ) public view returns (address[] memory) {
    return usersByType[serviceType];
  }

  // returns the index of the user found in the service list
  // it returns 0 and false if the user is not found
  function _getUserListIndex(
    SharedStructs.ServiceType serviceType,
    address userAddress
  ) private view returns (uint256, bool) {
    uint256 ret = 0;
    bool found = false;
    for (uint256 i = 0; i < usersByType[serviceType].length; i++) {
      if (usersByType[serviceType][i] == userAddress) {
        ret = i;
        found = true;
        break;
      }
    }
    return (ret, found);
  }

  function _doesUserHaveRole(
    SharedStructs.ServiceType serviceType,
    address userAddress
  ) private view returns (bool) {
    bool found = false;
    for (uint256 i = 0; i < users[userAddress].roles.length; i++) {
      if (users[userAddress].roles[i] == serviceType) {
        found = true;
        break;
      }
    }
    return found;
  }
}
