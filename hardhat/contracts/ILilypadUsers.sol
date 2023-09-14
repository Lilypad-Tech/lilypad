// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadUsers {

  /**
   * Users
   */

  function getUser(
    address userAddress
  ) external returns (SharedStructs.User memory);

  function updateUser(
    string memory metadataCID,
    string memory url,
    SharedStructs.ServiceType[] memory roles
  ) external returns (SharedStructs.User memory);

  // add the given user to a list of service types
  // that can be easily discovered (e.g. how to list the solvers)
  function addUserToList(
    SharedStructs.ServiceType serviceType
  ) external;

  function removeUserFromList(
    SharedStructs.ServiceType serviceType
  ) external;

  function showUsersInList(
    SharedStructs.ServiceType serviceType
  ) external returns (address[] memory);
  
}
