// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadStorage {

  /**
   * Users
   */

  function getUser(
    address userAddress
  ) external returns (SharedStructs.User memory);

  function updateUser(
    uint256 metadataCID,
    string memory url,
    SharedStructs.ServiceType[] memory roles,
    address[] memory trustedMediators,
    address[] memory trustedDirectories
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

  
  /**
   * Deals
   */

  function getDeal(
    uint256 dealId
  ) external returns (SharedStructs.Deal memory);

  function getDealsForParty(
    address party
  ) external returns (uint256[] memory);

  function addDeal(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 jobCollateral,
    uint256 resultsCollateral
  ) external returns (SharedStructs.Deal memory);
  
  /**
   * Agreements
   */

  function getAgreement(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);
  
  function agreeResourceProvider(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);

  function agreeJobCreator(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);
  
  /**
   * Results
   */

  function getResult(
    uint256 dealId
  ) external returns (SharedStructs.Result memory);

  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) external returns (SharedStructs.Result memory);

  function timeoutResult(
    uint256 dealId
  ) external;

  function acceptResult(
    uint256 dealId
  ) external;

  function rejectResult(
    uint256 dealId
  ) external;

  /**
   * Checkers
   */

  function hasDeal(
    uint256 dealId
  ) external returns (bool);

  function isNegotiating(
    uint256 dealId
  ) external returns (bool);

  function isAgreement(
    uint256 dealId
  ) external returns (bool);

  function isTimeout(
    uint256 dealId
  ) external returns (bool);

  function isSubmitted(
    uint256 dealId
  ) external returns (bool);

  function isAccepted(
    uint256 dealId
  ) external returns (bool);

  function isRejected(
    uint256 dealId
  ) external returns (bool);
  
}
