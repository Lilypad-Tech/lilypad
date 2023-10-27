// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SharedStructs.sol";
import "./ILilypadJobManager.sol";
import "./ILilypadJobClient.sol";
import "./ControllerOwnable.sol";
import "./ILilypadToken.sol";

contract LilypadOnChainJobCreator is ILilypadJobManager, ControllerOwnable, Initializable {

  // the token contract
  // we check to see what allowance has been granted to be spent on behalf
  // of the customer of a job
  address private tokenAddress;
  ILilypadToken private tokenContract;

  // the minimum amount that must be "approved" on the smart contract for the solver to spend
  // for it to consider running a job for a client
  // the solver will update this as the market pricing changes
  uint256 public requiredDeposit;

  // auto increment job id
  uint256 public nextJobID;

  // map of job id onto offer
  mapping(uint256 => SharedStructs.JobOffer) private jobOffers;

  event JobAdded(
    uint256 id,
    address calling_contract,
    address payee,
    string module,
    string[] inputs
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

  function getTokenAddress() public view returns (address) {
    return tokenAddress;
  }

  function setRequiredDeposit(uint256 cost) public onlyController {
    require(cost > 0, "Min deposit");
    requiredDeposit = cost;
  }

  function getRequiredDeposit() public view returns (uint256) {
    return requiredDeposit;
  }

  // called by on-chain clients to make an offer for a job
  // this will return a ticketID which is a unique onchain identifier for the job
  function runJob(
    // what is the module name we are making an offer for
    string memory module,
    // an array of key=value pairs that will be the inputs to the job
    string[] memory inputs,
    // the address of the client who is paying for the job
    // they must have called the increaseAllowance function
    // giving the controller (i.e. solver) permission to spend their tokens
    address payee
  ) public override returns (uint256) {
    // this makes sure that the person paying for the job has
    // already called "approve" on the token contract so the solver can
    // work on it's behalf
    require(tokenContract.allowance(payee, getControllerAddress()) >= requiredDeposit, "Token allowance not enough");

    nextJobID = nextJobID + 1;
    jobOffers[nextJobID] = SharedStructs.JobOffer(
      nextJobID,
      msg.sender,
      payee,
      module,
      inputs
    );
    emit JobAdded(
      nextJobID,
      msg.sender,
      payee,
      module,
      inputs
    );

    return nextJobID;
  }

  // this is called by the solver once we've got results out of the controller
  // it will call the "resultsAdded" function on the original client contract
  function submitResults(
    uint256 id,
    string memory dealId,
    string memory dataId
  ) public onlyController override {
    SharedStructs.JobOffer storage offer = jobOffers[id];
    require(offer.id != 0, "Job not found");
    ILilypadJobClient(offer.calling_contract).submitResults(
      id,
      dealId,
      dataId
    );
  }
}
