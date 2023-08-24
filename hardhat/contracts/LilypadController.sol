// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "hardhat/console.sol";
import "./SharedStructs.sol";
import "./ILilypadStorage.sol";

contract LilypadController is Ownable, Initializable {
  bool private initialized;
  address private storageAddress;
  address private tokenAddress;
  ILilypadStorage private storageContract;
  IERC20 private tokenContract;

  event DealAdded(uint256 indexed dealId);
  event ResourceProviderAgreed(uint256 indexed dealId);
  event JobCreatorAgreed(uint256 indexed dealId);
  event DealAgreed(uint256 indexed dealId);
  event ResultAdded(uint256 indexed dealId);
  
  // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
  function initialize(address _storageAddress, address _tokenAddress) public initializer {
    setStorageAddress(_storageAddress);
    setTokenAddress(_tokenAddress);
  }

  function setStorageAddress(address _storageAddress) public onlyOwner {
    require(_storageAddress != address(0), "Storage address must be defined");
    storageAddress = _storageAddress;
    storageContract = ILilypadStorage(storageAddress);
  }

  function setTokenAddress(address _tokenAddress) public onlyOwner {
    require(_tokenAddress != address(0), "Token address must be defined");
    tokenAddress = _tokenAddress;
    tokenContract = IERC20(storageAddress);
  }

  function agree(
    uint256 dealId,
    address resourceProvider,
    address jobCreator,
    uint256 instructionPrice,
    uint256 timeout,
    uint256 timeoutCollateral,
    uint256 jobCollateral,
    uint256 resultsCollateral
  ) public returns (SharedStructs.Agreement memory) {
    require(!storageContract.hasAgreement(dealId), "Deal has already been agreed");

    // we already have this deal and so the values must line up
    if(storageContract.hasDeal(dealId)) {
      SharedStructs.Deal memory existingDeal = storageContract.getDeal(dealId);
      require(existingDeal.resourceProvider == resourceProvider, "Resource provider does not match");
      require(existingDeal.jobCreator == jobCreator, "Job creator does not match");
      require(existingDeal.instructionPrice == instructionPrice, "Instruction price does not match");
      require(existingDeal.timeout == timeout, "Timeout does not match");
      require(existingDeal.timeoutCollateral == timeoutCollateral, "Timeout collateral does not match");
      require(existingDeal.jobCollateral == jobCollateral, "Job collateral does not match");
      require(existingDeal.resultsCollateral == resultsCollateral, "Results collateral does not match");
    }
    else {
      // we don't have this deal yet so add it
      storageContract.addDeal(
        dealId,
        resourceProvider,
        jobCreator,
        instructionPrice,
        timeout,
        timeoutCollateral,
        jobCollateral,
        resultsCollateral
      );
    }

    SharedStructs.Deal memory deal = storageContract.getDeal(dealId);
    SharedStructs.Agreement memory agreement = storageContract.getAgreement(dealId);

    bool isResourceProvider = tx.origin == deal.resourceProvider;
    bool isJobCreator = tx.origin == deal.jobCreator;
    require(isResourceProvider || isJobCreator, "Only RP or JC can agree to deal");
    if(isResourceProvider) {
      storageContract.agreeResourceProvider(dealId);
      emit ResourceProviderAgreed(dealId);
    }
    else if(isJobCreator) {
      storageContract.agreeJobCreator(dealId);
      emit JobCreatorAgreed(dealId);
    }
    if(storageContract.hasAgreement(dealId)) {
      emit DealAgreed(dealId);
    }
    return agreement;
  }
}





// function addResult(
//     uint256 dealId,
//     uint256 resultsId,
//     uint256 instructionCount
//   ) public returns (SharedStructs.Result memory) {
//     require(hasDeal(dealId), "Deal does not exist");
//     require(hasAgreement(dealId), "Agreement does not exist");
//     SharedStructs.Result memory newResult = SharedStructs.Result(
//       dealId,
//       resultsId,
//       instructionCount
//     );
//     results[dealId] = newResult;
//     return newResult;
//   }