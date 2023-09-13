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
  mapping(string => SharedStructs.Deal) private deals;

  // a map of party -> dealid[]
  mapping(address => string[]) private dealsForParty;

  // a map of deal id -> agreement
  mapping(string => SharedStructs.Agreement) private agreements;

  // a map of deal id -> result
  mapping(string => SharedStructs.Result) private results;

  // a map of deal id -> result
  mapping(string => SharedStructs.Result) private mediations;

  event DealStateChange(
    string indexed dealId,
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

  /**
   * Deals
   */

  function getDeal(
    string memory dealId
  ) public view returns (SharedStructs.Deal memory) {
    return deals[dealId];
  }

  function getDealsForParty(
    address party
  ) public view returns (string[] memory) {
    return dealsForParty[party];
  }

  function checkDealMembers(
    SharedStructs.DealMembers memory members
  ) private pure {
    require(members.resourceProvider != address(0), "RP missing");
    require(members.jobCreator != address(0), "JC missing");
    require(members.solver != address(0), "Solver missing");
    require(members.mediators.length > 0, "Mediators <= 0");
    require(members.resourceProvider != members.jobCreator, "RP / JC same");
  }

  function checkTimeouts(
    SharedStructs.DealTimeouts memory timeouts
  ) private pure {
    // the cost of the agree timeout cannot be > 0 because the whole point is
    // one party has not paid anything into the contract is what has timed out
    require(timeouts.agree.collateral == 0, "Agree deposit 0");
    // the same is true of the mediation timeout - it's cost cannot be zero
    require(timeouts.mediateResults.collateral == 0, "Mediate deposit 0");
  }

  function compareDealMembers(
    SharedStructs.DealMembers memory members1,
    SharedStructs.DealMembers memory members2
  ) private pure {
    require(members1.resourceProvider == members2.resourceProvider, "RP");
    require(members1.jobCreator == members2.jobCreator, "JC");
    require(members1.solver == members2.solver, "Solver");
    require(members1.mediators.length == members2.mediators.length, "Mediators");
    for (uint256 i = 0; i < members1.mediators.length; i++) {
      require(members1.mediators[i] == members2.mediators[i], "Mediator");
    }
  }

  function compareDealTimeout(
    SharedStructs.DealTimeout memory timeout1,
    SharedStructs.DealTimeout memory timeout2
  ) private pure {
    require(timeout1.timeout == timeout2.timeout, "Timeout");
    require(timeout1.collateral == timeout2.collateral, "Collateral");
  }
  
  function compareDealTimeouts(
    SharedStructs.DealTimeouts memory timeouts1,
    SharedStructs.DealTimeouts memory timeouts2
  ) private pure {
    compareDealTimeout(timeouts1.agree, timeouts2.agree);
    compareDealTimeout(timeouts1.submitResults, timeouts2.submitResults);
    compareDealTimeout(timeouts1.judgeResults, timeouts2.judgeResults);
    compareDealTimeout(timeouts1.mediateResults, timeouts2.mediateResults);
  }

  function compareDealPricing(
    SharedStructs.DealPricing memory pricing1,
    SharedStructs.DealPricing memory pricing2
  ) private pure {
    require(pricing1.instructionPrice == pricing2.instructionPrice, "Price");
    require(pricing1.paymentCollateral == pricing2.paymentCollateral, "Payment");
    require(pricing1.resultsCollateralMultiple == pricing2.resultsCollateralMultiple, "Results");
    require(pricing1.mediationFee == pricing2.mediationFee, "Mediation");
  }
  
  function ensureDeal(
    string memory dealId,
    SharedStructs.DealMembers memory members,
    SharedStructs.DealTimeouts memory timeouts,
    SharedStructs.DealPricing memory pricing
  ) public onlyController returns (SharedStructs.Deal memory) {
    require(isState(dealId, SharedStructs.AgreementState.DealNegotiating), "DealNegotiating");
    checkDealMembers(members);
    checkTimeouts(timeouts);
    if(hasDeal(dealId)) {
      SharedStructs.Deal memory existingDeal = getDeal(dealId);
      compareDealMembers(existingDeal.members, members);
      compareDealTimeouts(existingDeal.timeouts, timeouts);
      compareDealPricing(existingDeal.pricing, pricing);
    }
    else {
      deals[dealId] = SharedStructs.Deal(
        dealId,
        members,
        timeouts,
        pricing
      );
      dealsForParty[members.resourceProvider].push(dealId);
      dealsForParty[members.jobCreator].push(dealId);
    }
    return deals[dealId];
  }

  /**
   * Agreements
   */
  
  function getAgreement(
    string memory dealId
  ) public view returns (SharedStructs.Agreement memory) {
    return agreements[dealId];
  }

  function agreeResourceProvider(
    string memory dealId
  ) public onlyController returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].resourceProviderAgreedAt == 0, "RP has already agreed");
    agreements[dealId].resourceProviderAgreedAt = block.timestamp;
    _maybeAgreeDeal(dealId);
    return agreements[dealId];
  }

  function agreeJobCreator(
    string memory dealId
  ) public onlyController returns (SharedStructs.Agreement memory) {
    require(hasDeal(dealId), "Deal does not exist");
    require(agreements[dealId].jobCreatorAgreedAt == 0, "JC has already agreed");
    agreements[dealId].jobCreatorAgreedAt = block.timestamp;
    _maybeAgreeDeal(dealId);
    return agreements[dealId];
  }

  /**
   * Post Results
   */

  function getResult(
    string memory dealId
  ) public view returns (SharedStructs.Result memory) {
    return results[dealId];
  }

  function addResult(
    string memory dealId,
    string memory resultsId,
    uint256 instructionCount
  ) public onlyController returns (SharedStructs.Result memory) {
    require(isState(dealId, SharedStructs.AgreementState.DealAgreed), "DealAgreed");
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
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "ResultsSubmitted");
    agreements[dealId].resultsAcceptedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.ResultsAccepted);
  }

  function checkResult(
    string memory dealId,
    address mediator
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "ResultsSubmitted");
    agreements[dealId].resultsCheckedAt = block.timestamp;
    agreements[dealId].mediator = mediator;
    _changeAgreementState(dealId, SharedStructs.AgreementState.ResultsChecked);

  }

  /**
   * Mediati:
   */

  function mediationAcceptResult(
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsChecked), "ResultsChecked");
    agreements[dealId].mediationAcceptedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.MediationAccepted);
  }

  function mediationRejectResult(
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsChecked), "ResultsChecked");
    agreements[dealId].mediationRejectedAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.MediationRejected);
  }

  /**
   * Timeouts
   */

  // called because one party submitted a deal and the other party
  // did not agree in time
  function timeoutAgree(
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.DealNegotiating), "DealNegotiating");
    agreements[dealId].timeoutAgreeAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutAgree);
  }

  // called because the JC waited too long for a result to be submitted
  // and wants it's money back
  function timeoutSubmitResult(
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.DealAgreed), "DealAgreed");
    agreements[dealId].timeoutSubmitResultsAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutSubmitResults);
  }

  // called because the RP waited too long for a judgement of it's results
  // and wants it's money back
  function timeoutJudgeResult(
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsSubmitted), "ResultsSubmitted");
    agreements[dealId].timeoutJudgeResultsAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutJudgeResults);
  }
  // called because the RP or JC waited too long for a mediation of it's results
  // and both want their money back
  function timeoutMediateResult(
    string memory dealId
  ) public onlyController {
    require(isState(dealId, SharedStructs.AgreementState.ResultsChecked), "ResultsChecked");
    agreements[dealId].timeoutMediateResultsAt = block.timestamp;
    _changeAgreementState(dealId, SharedStructs.AgreementState.TimeoutMediateResults);
  }

  /**
   * Costings
   */

  function getJobCost(
    string memory dealId
  ) public view returns (uint256) {
    return deals[dealId].pricing.instructionPrice * results[dealId].instructionCount;
  }

  function getResultsCollateral(
string memory dealId
  ) public view returns (uint256) {
    return deals[dealId].pricing.resultsCollateralMultiple * getJobCost(dealId);
  }

  /**
   * Checkers
   */

  function hasDeal(
    string memory dealId
  ) public view returns (bool) {
    return  bytes(getDeal(dealId).dealId).length > 0;
  }

  function isState(
    string memory dealId,
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
    string memory dealId
  ) private {
    if(agreements[dealId].resourceProviderAgreedAt != 0 && agreements[dealId].jobCreatorAgreedAt != 0) {
      agreements[dealId].dealAgreedAt = block.timestamp;
      _changeAgreementState(dealId, SharedStructs.AgreementState.DealAgreed);
    } else {
      // this is used so we can know if a party can call an agree timeout trigger
      agreements[dealId].dealCreatedAt = block.timestamp;
    }
  }

  function _changeAgreementState(
    string memory dealId,
    SharedStructs.AgreementState state
  ) private {
    agreements[dealId].state = state;
    emit DealStateChange(dealId, state);
  }
}
