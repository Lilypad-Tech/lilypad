// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadStorage {

  // a new user has registered
  event UserUpdated(
    SharedStructs.ServiceType indexed userType,
    address indexed userAddress,
    uint256 indexed metadataCID
  );

  // trusted providers have been changed for a user
  event TrustedProvidersAdded(
    address indexed userAddress,
    SharedStructs.ServiceType indexed providerType,
    address[] indexed trustedProviders
  );

  // when one party agrees to a deal this is emitted
  // along with the address of the agreeing party
  // this can be used to lookup whether it's the RP or JC
  event ResourceProviderAgreed(uint256 indexed dealId, address indexed resourceProvider);
  event JobCreatorAgreed(uint256 indexed dealId, address indexed jobCreator);

  // once both parties have agreed this is emitted
  event DealAgreed(uint256 indexed dealId);

  // a result has been subnmitted by a resource provider
  event ResultAdded(uint256 indexed dealId);

  /**
   * Users
   */
  function updateUser(
    address userAddress,
    uint256 metadataCID,
    string memory url,
    SharedStructs.ServiceType[] memory roles,
    address[] memory trustedMediators,
    address[] memory trustedDirectories
  ) external returns (SharedStructs.User memory);

  function getUser(
    address userAddress
  ) external returns (SharedStructs.User memory);

  // add the given user to a list of service types
  // that can be easily discovered (e.g. how to list the solvers)
  function addUserToList(
    SharedStructs.ServiceType serviceType,
    address userAddress
  ) external;

  function removeUserFromList(
    SharedStructs.ServiceType serviceType,
    address userAddress
  ) external;

  function showUsersInList(
    SharedStructs.ServiceType serviceType
  ) external returns (address[] memory);

  
  /**
   * Deals
   */

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

  function hasDeal(
    uint256 dealId
  ) external returns (bool);

  function getDeal(
    uint256 dealId
  ) external returns (SharedStructs.Deal memory);

  // return a list of deal ids for which a party has been
  // mentioned as either the RP or JC
  function getDealsForParty(
    address party
  ) external returns (uint256[] memory);
  
  /**
   * Agreements
   */

  // this will check:
  //  * that the deal exists
  //  * that the sender of the tx is either rp or jc
  //  * that the rp or jc has not already agreed
  function agreeResourceProvider(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);

  function agreeJobCreator(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);

  // tells if you if both parties have agreed to the given deal
  function hasAgreement(
    uint256 dealId
  ) external returns (bool);

  function getAgreement(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);

  /**
   * Results
   */

  // results
  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) external returns (SharedStructs.Result memory);

  function hasResult(
    uint256 dealId
  ) external returns (bool);

  function getResult(
    uint256 dealId
  ) external returns (SharedStructs.Result memory);
}
