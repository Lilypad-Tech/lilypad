// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";

contract LilypadStorage is Ownable, Initializable {

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
    uint256 metadataCID,
    string memory url,
    SharedStructs.ServiceType[] memory roles,
    address[] memory trustedMediators,
    address[] memory trustedDirectories
  ) public returns (SharedStructs.User memory) {
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

  /**
   * Deals
   */

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

  function addDeal(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 jobCollateral,
    uint256 resultsCollateralMultiple
  ) public onlyOwner returns (SharedStructs.Deal memory) {
    require(!hasDeal(dealId), "Deal already exists");
    require(resourceProvider != address(0), "Resource provider must be defined");
    require(jobCreator != address(0), "Job creator must be defined");
    deals[dealId] = SharedStructs.Deal(
      dealId,
      resourceProvider,
      jobCreator,
      instructionPrice,
      timeout,
      timeoutCollateral,
      jobCollateral,
      resultsCollateralMultiple
    );
    agreements[dealId].state = SharedStructs.AgreementState.Negotiating;
    dealsForParty[resourceProvider].push(dealId);
    dealsForParty[jobCreator].push(dealId);
    return deals[dealId];
  }

  /**
   * Agreements
   */
  
  function getAgreement(
    uint256 dealId
  ) public view returns (SharedStructs.Agreement memory) {
    return agreements[dealId];
  }

  function agreeResourceProvider(
    uint256 dealId
  ) public onlyOwner returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].resourceProviderAgreedAt == 0, "resource provider has already agreed");
    agreements[dealId].resourceProviderAgreedAt = block.timestamp;
    _maybeAgreeDeal(dealId);
    return agreements[dealId];
  }

  // only the controller contract can call this
  function agreeJobCreator(
    uint256 dealId
  ) public onlyOwner returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].jobCreatorAgreedAt == 0, "job creator has already agreed");
    agreements[dealId].jobCreatorAgreedAt = block.timestamp;
    _maybeAgreeDeal(dealId);
    return agreements[dealId];
  }

  function _maybeAgreeDeal(
    uint256 dealId
  ) private {
    if(isAgreement(dealId)) {
      agreements[dealId].dealAgreedAt = block.timestamp;
      agreements[dealId].state = SharedStructs.AgreementState.Agreed;
    }
  }

  /**
   * Results
   */

  function getResult(
    uint256 dealId
  ) public view returns (SharedStructs.Result memory) {
    return results[dealId];
  }

  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) public onlyOwner returns (SharedStructs.Result memory) {
    require(isAgreement(dealId), "Deal not in Agree state");
    agreements[dealId].resultsSubmittedAt = block.timestamp;
    agreements[dealId].state = SharedStructs.AgreementState.Submitted;
    results[dealId] = SharedStructs.Result(
      dealId,
      resultsId,
      instructionCount
    );
    return results[dealId];
  }

  function timeoutResult(
    uint256 dealId
  ) public onlyOwner {
    require(isAgreement(dealId), "Deal not in Agree state");
    agreements[dealId].timedOutAt = block.timestamp;
    agreements[dealId].state = SharedStructs.AgreementState.Timeout;
  }

  function acceptResult(
    uint256 dealId
  ) public onlyOwner {
    require(isSubmitted(dealId), "Deal not in Submitted state");
    agreements[dealId].resultsAcceptedAt = block.timestamp;
    agreements[dealId].state = SharedStructs.AgreementState.Accepted;
  }

  function rejectResult(
    uint256 dealId
  ) public onlyOwner {
    require(isSubmitted(dealId), "Deal not in Submitted state");
    agreements[dealId].resultsRejectedAt = block.timestamp;
    agreements[dealId].state = SharedStructs.AgreementState.Rejected;
  }

  /**
   * Checkers
   */

  function hasDeal(
    uint256 dealId
  ) public view returns (bool) {
    return getDeal(dealId).dealId != 0;
  }

  function isNegotiating(
    uint256 dealId
  ) public view returns (bool) {
    if(!hasDeal(dealId)) return false;
    return agreements[dealId].state == SharedStructs.AgreementState.Negotiating;
  }

  function isAgreement(
    uint256 dealId
  ) public view returns (bool) {
    if(!hasDeal(dealId)) return false;
    return agreements[dealId].state == SharedStructs.AgreementState.Agreed;
  }

  function isTimeout(
    uint256 dealId
  ) public view returns (bool) {
    if(!hasDeal(dealId)) return false;
    return agreements[dealId].state == SharedStructs.AgreementState.Timeout;
  }

  function isSubmitted(
    uint256 dealId
  ) public view returns (bool) {
    if(!hasDeal(dealId)) return false;
    return agreements[dealId].state == SharedStructs.AgreementState.Submitted;
  }

  function isAccepted(
    uint256 dealId
  ) public view returns (bool) {
    if(!hasDeal(dealId)) return false;
    return agreements[dealId].state == SharedStructs.AgreementState.Accepted;
  }

  function isRejected(
    uint256 dealId
  ) public view returns (bool) {
    if(!hasDeal(dealId)) return false;
    return agreements[dealId].state == SharedStructs.AgreementState.Rejected;
  }
}
