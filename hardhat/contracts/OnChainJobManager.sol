// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./IOnChainJobManager.sol";
import "./IOnChainJobClient.sol";
import "./ControllerOwnable.sol";
import "./ILilypadToken.sol";

contract OnChainJobManager is IOnChainJobManager, ControllerOwnable, Initializable {

  // keep track of which contract has called which job
  // mapping(string => SharedStructs.Result) private results;

  address private tokenAddress;
  ILilypadToken private tokenContract;

  // simplistic way of the solver updating how much things currently cost
  // this is based on market pricing - the solver will post the cheapest
  // TODO: allow more detail here - for example, per module pricing
  uint256 public defaultModuleCost;

  // keep track of the escrow currently being held on behalf of
  // different contracts
  mapping(address => uint256) private escrow;

  event JobAdded(
    string module,
    SharedStructs.JobOfferInput[] inputs
  );

  /**
   * Init
   */

  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(
    address _tokenAddress
  ) public initializer {
    setTokenAddress(_tokenAddress);
  }

  function setTokenAddress(address _tokenAddress) public onlyOwner {
    require(_tokenAddress != address(0), "Token address");
    tokenAddress = _tokenAddress;
    tokenContract = ILilypadToken(tokenAddress);
  }

  function setDefaultModuleCost(uint256 cost) public onlyController {
    require(cost > 0, "Min cost");
    defaultModuleCost = cost;
  }

  // called by clients - transfer tokens and then emit event so solver
  // can kick in with a job offer on behalf of the submitting contract
  function runJob(
    string memory module,
    SharedStructs.JobOfferInput[] memory inputs
  ) public override {
    // move X tokens to the solver (i.e. the controller address)
    // emit the event so the solver can kick in and run the job on behalf
    // of the on chain client
    tokenContract.payOnChainManager(getControllerAddress(), defaultModuleCost);
  }

  // this is called by the solver once we've got results out of the controller
  // it will call the "resultsAdded" function on the original client contract
  function resultsAdded(
    string memory dealId,
    string memory resultsId,
    string memory dataId,
    uint256 instructionCount
  ) public override {

  }
}
