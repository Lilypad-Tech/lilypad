// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

// import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";

contract LilypadStorage is Ownable {

  // a map of user address -> user
  mapping(address => SharedStructs.User) public users;

  // a map of user type => user address[]
  mapping(SharedStructs.ServiceType => address[]) public usersByType;

  // a map of deal id -> deal
  mapping(uint256 => SharedStructs.Deal) public deals;

  // a map of party -> dealid[]
  mapping(address => uint256[]) public dealsForParty;

  // a map of deal id -> agreement
  mapping(uint256 => SharedStructs.Agreement) public agreements;

  // a map of deal id -> result
  mapping(uint256 => SharedStructs.Result) public results;

  /**
   * Users
   */
  function updateUser(
    uint256 metadataCID,
    string memory url,
    SharedStructs.ServiceType[] memory roles,
    address[] memory trustedMediators,
    address[] memory trustedDirectories
  ) public onlyOwner returns (SharedStructs.User memory) {
    SharedStructs.User memory newUser = SharedStructs.User(
      tx.origin,
      metadataCID,
      url,
      roles,
      trustedMediators,
      trustedDirectories
    );
    users[tx.origin] = newUser;
    return newUser;
  }

  function getUser(
    address userAddress
  ) public view returns (SharedStructs.User memory) {
    return users[userAddress];
  }

  function addUserToList(
    SharedStructs.ServiceType serviceType
  ) public {
    require(users[tx.origin].userAddress != address(0), "User must exist");
    (, bool found) = _getUserListIndex(serviceType, tx.origin);
    require(!found, "User is already part of that list");
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

  /**
   * Deals
   */

  // only the controller contract can call this
  function addDeal(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 jobCollateral,
    uint256 resultsCollateral
  ) public onlyOwner returns (SharedStructs.Deal memory) {
    require(!hasDeal(dealId), "Deal already exists");
    require(resourceProvider != address(0), "Resource provider must be defined");
    require(jobCreator != address(0), "Job creator must be defined");
    SharedStructs.Deal memory newDeal = SharedStructs.Deal(
      dealId,
      resourceProvider,
      jobCreator,
      instructionPrice,
      timeout,
      timeoutCollateral,
      jobCollateral,
      resultsCollateral
    );
    deals[dealId] = newDeal;
    dealsForParty[resourceProvider].push(dealId);
    dealsForParty[jobCreator].push(dealId);
    return newDeal;
  }

  function getDeal(
    uint256 dealId
  ) public view returns (SharedStructs.Deal memory) {
    return deals[dealId];
  }

  function getDealsForParty(
    address party
  ) public view returns (uint256[] memory) {
    return dealsForParty[party];
  }

  function hasDeal(
    uint256 dealId
  ) public view returns (bool) {
    return getDeal(dealId).dealId != 0;
  }

  /**
   * Agreements
   */
  
  // only the controller contract can call this
  function agreeResourceProvider(
    uint256 dealId
  ) public onlyOwner returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].resourceProviderAgreedAt == 0, "resource provider has already agreed");
    agreements[dealId].resourceProviderAgreedAt = block.timestamp;
    return agreements[dealId];
  }

  // only the controller contract can call this
  function agreeJobCreator(
    uint256 dealId
  ) public onlyOwner returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].jobCreatorAgreedAt == 0, "job creator has already agreed");
    agreements[dealId].jobCreatorAgreedAt = block.timestamp;
    return agreements[dealId];
  }

  function hasAgreement(
    uint256 dealId
  ) public view returns (bool) {
    require(hasDeal(dealId), "Deal does not exist");
    return agreements[dealId].resourceProviderAgreedAt != 0 && agreements[dealId].jobCreatorAgreedAt != 0;
  }

  function getAgreement(
    uint256 dealId
  ) public view returns (SharedStructs.Agreement memory) {
    return agreements[dealId];
  }

  /**
   * Results
   */

  // only the controller contract can call this
  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) public onlyOwner returns (SharedStructs.Result memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(hasAgreement(dealId), "Agreement does not exist");
    SharedStructs.Result memory newResult = SharedStructs.Result(
      dealId,
      resultsId,
      instructionCount
    );
    results[dealId] = newResult;
    return newResult;
  }

  function hasResult(
    uint256 dealId
  ) public view returns (bool) {
    return getResult(dealId).dealId != 0;
  }

  function getResult(
    uint256 dealId
  ) public view returns (SharedStructs.Result memory) {
    return results[dealId];
  }
}
