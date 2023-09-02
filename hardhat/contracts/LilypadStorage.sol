// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ControllerOwnable.sol";

// import "@openzeppelin/contracts/utils/Strings.sol";
// import "hardhat/console.sol";
// console.log("ensureDeal");
// console.log(Strings.toString(uint256(SharedStructs.AgreementState.DealNegotiating)));
// console.log(Strings.toString(uint256(agreements[dealId].state)));

contract LilypadStorage is ControllerOwnable, Initializable {

  // the address that is allowed to be the msg.sender for the payment functions
  address private controllerAddress;

  // used to cut off upgrades for the remote contract
  bool private canChangeControllerAddress = true;

  // a map of user address -> user
  mapping(address => SharedStructs.User) private users;

  // a map of user type => user address[]
  mapping(SharedStructs.ServiceType => address[]) private usersByType;

  // a map of deal id -> deal
  mapping(uint256 => SharedStructs.Deal) private deals;

  // a map of party -> dealid[]
  mapping(address => uint256[]) private dealsForParty;

  // a map of deal id -> agreement
  mapping(uint256 => SharedStructs.Agreement) private agreements;

  // a map of deal id -> result
  mapping(uint256 => SharedStructs.Result) private results;

  // a map of deal id -> result
  mapping(uint256 => SharedStructs.Result) private mediations;

  event DealStateChange(
    uint256 indexed dealId,
    SharedStructs.AgreementState indexed state
  );

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

  function ensureDeal(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 paymentCollateral,
    uint256 resultsCollateralMultiple,
    uint256 mediationFee
  ) public onlyController returns (SharedStructs.Deal memory) {
    require(isState(dealId, SharedStructs.AgreementState.DealNegotiating), "Deal is not in DealNegotiating state");
    require(resourceProvider != address(0), "Resource provider must be defined");
    require(jobCreator != address(0), "Job creator must be defined");
    require(resourceProvider != jobCreator, "RP and JC cannot be the same");

    if(hasDeal(dealId)) {
      SharedStructs.Deal memory existingDeal = getDeal(dealId);
      require(existingDeal.resourceProvider == resourceProvider, "RP does not match");
      require(existingDeal.jobCreator == jobCreator, "JC does not match");
      require(existingDeal.instructionPrice == instructionPrice, "Instruction price does not match");
      require(existingDeal.timeout == timeout, "Timeout does not match");
      require(existingDeal.timeoutCollateral == timeoutCollateral, "Timeout collateral does not match");
      require(existingDeal.paymentCollateral == paymentCollateral, "Payment collateral does not match");
      require(existingDeal.resultsCollateralMultiple == resultsCollateralMultiple, "Results collateral does not match");
      require(existingDeal.mediationFee == mediationFee, "Mediation fee does not match");
    }
    else {
      deals[dealId] = SharedStructs.Deal(
        dealId,
        resourceProvider,
        jobCreator,
        instructionPrice,
        timeout,
        timeoutCollateral,
        paymentCollateral,
        resultsCollateralMultiple,
        mediationFee
      );
      dealsForParty[resourceProvider].push(dealId);
      dealsForParty[jobCreator].push(dealId);
    }

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
  ) public onlyController returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].resourceProviderAgreedAt == 0, "Resource provider has already agreed");
    agreements[dealId].resourceProviderAgreedAt = block.timestamp;
    _maybeAgreeDeal(dealId);
    return agreements[dealId];
  }

  function agreeJobCreator(
    uint256 dealId
  ) public onlyController returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].jobCreatorAgreedAt == 0, "Job creator has already agreed");
    agreements[dealId].jobCreatorAgreedAt = block.timestamp;
    _maybeAgreeDeal(dealId);
    return agreements[dealId];
  }

  /**
   * Post Results
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
  ) public onlyController returns (SharedStructs.Result memory) {
    require(isState(dealId, SharedStructs.AgreementState.DealAgreed), "Deal not in DealAgreed state");
    agreements[dealId].resultsSubmittedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.ResultsSubmitted);
    results[dealId] = SharedStructs.Result(
      dealId,
      resultsId,
      instructionCount
    );
    return results[dealId];
  }

  /**
   * Judge Results
   */

  function acceptResult(
    uint256 dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal not in ResultsSubmitted state");
    agreements[dealId].resultsAcceptedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.ResultsAccepted);
  }

  function checkResult(
    uint256 dealId,
    address mediator
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal not in ResultsSubmitted state");
    SharedStructs.Deal memory deal = getDeal(dealId);
    require(_hasTrustedMediator(deal.resourceProvider, mediator), "Resource provider does not trust that mediator");
    agreements[dealId].resultsChallengedAt = block.timestamp;
    agreements[dealId].mediator = mediator;
    _changeAgreementState(dealId, SharedStructs.AgreementState.ResultsChallenged);

  }

  /**
   * Mediation
   */

  function mediationAcceptResult(
    uint256 dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsChallenged), "Deal not in ResultsChallenged state");
    agreements[dealId].mediationAcceptedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.MediationAccepted);
  }

  function mediationRejectResult(
    uint256 dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsChallenged), "Deal not in ResultsChallenged state");
    agreements[dealId].mediationRejectedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.MediationRejected);
  }

  /**
   * Timeouts
   */

  // called because the JC waited too long for a result to be submitted
  // and wants it's money back
  function timeoutSubmitResult(
    uint256 dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.DealAgreed), "Deal not in DealAgreed state");
    agreements[dealId].timeoutSubmitResultsAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutSubmitResults);
  }

  // called because the RP waited too long for a judgement of it's results
  // and wants it's money back
  function timeoutJudgeResult(
    uint256 dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "Deal not in ResultsSubmitted state");
    agreements[dealId].timeoutJudgeResultsAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutJudgeResults);
  }
  // called because the RP or JC waited too long for a mediation of it's results
  // and both want their money back
  function timeoutMediateResult(
    uint256 dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsChallenged), "Deal not in ResultsChallenged state");
    agreements[dealId].timeoutMediateResultsAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutMediateResults);
  }

  /**
   * Costings
   */

  function getJobCost(
    uint256 dealId
  ) public view returns (uint256) {
    return deals[dealId].instructionPrice * results[dealId].instructionCount;
  }

  function getResultsCollateral(
    uint256 dealId
  ) public view returns (uint256) {
    return deals[dealId].resultsCollateralMultiple * getJobCost(dealId);
  }

  /**
   * Checkers
   */

  function hasDeal(
    uint256 dealId
  ) public view returns (bool) {
    return getDeal(dealId).dealId != 0;
  }

  function isState(
    uint256 dealId,
    SharedStructs.AgreementState state
  ) public view returns (bool) {
    // if we don't have a deal, we should check against DealNegotiating
    // as this is the default state - otherwise it's impossible to check
    // for isState('DealNegotiating')
    if(!hasDeal(dealId)) {
      return state == SharedStructs.AgreementState.DealNegotiating;
    }
    return agreements[dealId].state == state;
  }

  /**
   * Utils
   */

  function _maybeAgreeDeal(
    uint256 dealId
  ) private {
    if(agreements[dealId].resourceProviderAgreedAt != 0 && agreements[dealId].jobCreatorAgreedAt != 0) {
      agreements[dealId].dealAgreedAt = block.timestamp;
      _changeAgreementState(dealId, SharedStructs.AgreementState.DealAgreed);
    }
  }

  function _changeAgreementState(
    uint256 dealId,
    SharedStructs.AgreementState state
  ) private {
    agreements[dealId].state = state;
    emit DealStateChange(dealId, state);
  }

  function _hasTrustedMediator(
    address userId,
    address mediator
  ) private view returns (bool) {
    bool found = false;
    for (uint256 i = 0; i < users[userId].trustedMediators.length; i++) {
      if (users[userId].trustedMediators[i] == mediator) {
        found = true;
        break;
      }
    }
    return found;
  }
}
