// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ControllerOwnable.sol";
import "./SharedStructs.sol";
import "./ILilypadMediation.sol";

contract LilypadMediationRandom is ControllerOwnable, Initializable {

  // keep track of which mediators were choosen for which deals
  mapping(string => address) private mediators;

  event MediationRequested(
    string dealId,
    address mediator
  );

  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize() public initializer {
    
  }

  // this is called by the controller contract
  function mediationRequest(
    SharedStructs.Deal memory deal
  ) public onlyController {
    uint randomIndex = uint(keccak256(abi.encodePacked(block.timestamp, deal.dealId))) % deal.members.mediators.length;
    address mediator = deal.members.mediators[randomIndex];
    require(mediator != address(0), "mediator cannot be 0x0");
    mediators[deal.dealId] = mediator;
    emit MediationRequested(deal.dealId, mediator);
  }

  function getMediator(
    string memory dealId
  ) public view returns(address) {
    return mediators[dealId];
  }

  // call the controller contract as a ILilypadMediationRequester
  function mediationAcceptResult(
    string memory dealId
  ) public {
    // check the tx.origin is the same mediator that was picked
    require(mediators[dealId] != address(0), "mediator cannot be 0x0");
    require(mediators[dealId] == tx.origin, "tx.origin must be the mediator");
    // call the controller contract
    ILilypadMediationRequester(getControllerAddress()).mediationAcceptResult(dealId);
  }

  // call the controller contract as a ILilypadMediationRequester
  function mediationRejectResult(
    string memory dealId
  ) public {
    // check the tx.origin is the same mediator that was picked
    require(mediators[dealId] != address(0), "mediator cannot be 0x0");
    require(mediators[dealId] == tx.origin, "tx.origin must be the mediator");
    // call the controller contract
    ILilypadMediationRequester(getControllerAddress()).mediationRejectResult(dealId);
  }
}
