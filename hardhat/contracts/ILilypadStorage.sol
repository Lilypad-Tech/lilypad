// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./SharedStructs.sol";

interface ILilypadStorage {

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
  ) external returns (SharedStructs.Deal memory);

  function getDeal(
    uint256 dealId
  ) external returns (SharedStructs.Deal memory);

  // return a list of deal ids for which a party has been
  // mentioned as either the RP or JC
  function getDealsForParty(
    address party
  ) external returns (uint256[] memory);

  function hasDeal(
    uint256 dealId
  ) external returns (bool);

  event DealAdded(uint256 indexed dealId);

  /**
   * Agreements
   */

  // this will check:
  //  * that the deal exists
  //  * that the sender of the tx is either rp or jc
  //  * that the rp or jc has not already agreed
  function agree(
    uint256 dealId
  ) external returns (SharedStructs.Agreement memory);

  // tells if you if both parties have agreed to the given deal
  function hasAgreement(
    uint256 dealId
  ) external returns (bool);

  // when one party agrees to a deal this is emitted
  // along with the address of the agreeing party
  // this can be used to lookup whether it's the RP or JC
  event PartyAgreed(uint256 indexed dealId, address indexed agreedParty); 

  // once both parties have agreed this is emitted
  event DealAgreed(uint256 indexed dealId);

  /**
   * Results
   */

  // results
  function addResult(
    uint256 dealId,
    uint256 resultsId,
    uint256 instructionCount
  ) external returns (SharedStructs.Result memory);

  function getResult(
    uint256 dealId
  ) external returns (SharedStructs.Result memory);

  function hasResult(
    uint256 dealId
  ) external returns (bool);

  event ResultAdded(uint256 indexed dealId);  
}
