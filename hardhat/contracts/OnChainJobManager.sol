// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ILilypadJobManager.sol";
import "./ILilypadJobClient.sol";
import "./ControllerOwnable.sol";
import "./ILilypadToken.sol";

contract JobManager is ILilypadJobManager, ControllerOwnable, Initializable {

  // keep track of which contract has called which job
  // mapping(string => SharedStructs.Result) private results;

  address private tokenAddress;
  ILilypadToken private tokenContract;

  // simplistic way of the solver updating how much things currently cost
  // this is based on market pricing - the solver will post the cheapest
  // TODO: allow more detail here - for example, per module pricing
  uint256 public defaultModuleCost;

  // auto increment job id
  uint256 public nextJobID;

  // map of job id onto offer
  mapping(uint256 => SharedStructs.JobOffer) private jobOffers;

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
    nextJobID = 0;
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
  ) public override returns (uint256) {
    nextJobID = nextJobID + 1;
    jobOffers[nextJobID] = SharedStructs.JobOffer(
      nextJobID,
      msg.sender,
      module,
      inputs
    );
    // move X tokens to the solver (i.e. the controller address)
    // emit the event so the solver can kick in and run the job on behalf
    // of the on chain client
    tokenContract.payOnChainManager(getControllerAddress(), defaultModuleCost);

    return nextJobID;
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
