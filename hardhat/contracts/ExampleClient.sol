// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ILilypadJobManager.sol";
import "./ILilypadJobClient.sol";

contract ExampleClient is Ownable, Initializable, ILilypadJobClient {

  address private jobManagerAddress;
  ILilypadJobManager private jobManagerContract;

  event JobCompleted(
    uint256 id,
    string dealId,
    string dataId
  );

  function initialize(address _jobManagerAddress) public initializer {
    setJobManagerAddress(_jobManagerAddress);
  }

  function setJobManagerAddress(address _jobManagerAddress) public onlyOwner {
    require(_jobManagerAddress != address(0), "Job manager address");
    jobManagerAddress = _jobManagerAddress;
    jobManagerContract = ILilypadJobManager(jobManagerAddress);
  }

  function runCowsay(
    string memory message
  ) public {
    string[] memory inputs = new string[](1);
    inputs[0] = string(abi.encodePacked("Message=", message));
    jobManagerContract.runJob(
      "cowsay:v0.0.1",
      inputs,
      msg.sender
    );
  }

  function submitResults(
    uint256 id,
    string memory dealId,
    string memory dataId
  ) public override {
    emit JobCompleted(
      id,
      dealId,
      dataId
    );
  }
}
