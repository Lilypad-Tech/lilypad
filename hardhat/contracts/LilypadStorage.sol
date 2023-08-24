// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "hardhat/console.sol";
import "./SharedStructs.sol";

contract LilypadStorage is Ownable {

  // a map of deal id -> deal
  mapping(uint256 => SharedStructs.Deal) public deals;

  // a map of party -> dealid[]
  mapping(address => uint256[]) public dealsForParty;

  // a map of deal id -> agreement
  mapping(uint256 => SharedStructs.Agreement) public agreements;

  // a map of deal id -> result
  mapping(uint256 => SharedStructs.Result) public results;

  event DealAdded(uint256 indexed dealId);
  event PartyAgreed(uint256 indexed dealId, address indexed agreedParty);
  event DealAgreed(uint256 indexed dealId);
  event ResultAdded(uint256 indexed dealId);

  /**
   * Deals
   */

  function addDeal(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeoutCollateral,
    uint256 jobCollateral,
    uint256 resultsCollateral
  ) public returns (SharedStructs.Deal memory) {
    require(!hasDeal(dealId), "Deal already exists");
    require(resourceProvider != address(0), "Resource provider must be defined");
    require(jobCreator != address(0), "Job creator must be defined");
    SharedStructs.Deal memory newDeal = SharedStructs.Deal(
      dealId,
      resourceProvider,
      jobCreator,
      instructionPrice,
      timeoutCollateral,
      jobCollateral,
      resultsCollateral
    );
    deals[dealId] = newDeal;
    dealsForParty[resourceProvider].push(dealId);
    dealsForParty[jobCreator].push(dealId);
    emit DealAdded(dealId);
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
  function hasAgreement(
    uint256 dealId
  ) public view returns (bool) {
    require(hasDeal(dealId), "Deal does not exist");
    SharedStructs.Agreement memory agreement = agreements[dealId];
    return agreement.resourceProviderAgreedAt != 0 && agreement.jobCreatorAgreedAt != 0;
  }

  function agree(
    uint256 dealId
  ) public returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(!hasAgreement(dealId), "Deal has already been agreed");
    SharedStructs.Deal memory deal = getDeal(dealId);
    SharedStructs.Agreement memory agreement = agreements[dealId];
    bool isResourceProvider = tx.origin == deal.resourceProvider;
    bool isJobCreator = tx.origin == deal.jobCreator;
    require(isResourceProvider || isJobCreator, "Only RP or JC can agree to deal");
    if(isResourceProvider) {
      agreement.resourceProviderAgreedAt = block.timestamp;
    }
    else if(isJobCreator) {
      agreement.jobCreatorAgreedAt = block.timestamp;
    }
    agreements[dealId] = agreement;
    return agreement;
  }

  /**
   * Results
   */

  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) public returns (SharedStructs.Result memory) {
    require(hasDeal(dealId), "Deal does not exist");
    SharedStructs.Result memory newResult = SharedStructs.Result(
      dealId,
      resultsId,
      instructionCount
    );
    results[dealId] = newResult;
    return newResult;
  }

  function getResult(
    uint256 dealId
  ) public view returns (SharedStructs.Result memory) {
    return results[dealId];
  }

  function hasResult(
    uint256 dealId
  ) public view returns (bool) {
    return getResult(dealId).dealId != 0;
  }
}
