// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract PoWToken is Ownable, Initializable {
    struct POWSubmission {
      address walletAddress;
      string nodeId;
      uint256 nonce;
      uint256 timestamp;
      bytes32 challenge; //record this to provent user never change challenge
    }

    mapping(address => POWSubmission[]) public powSubmissions;
    mapping(address => bytes32) public lastChallenges;
    uint256 public validProofs;
    uint256 public startTime;
    
    /**
     * Init
     */

    // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
    function initialize() public initializer {
        
    }

    // generateChallenge gen a byte32 value as challenge value, Sc store this one for verify
    function generateChallenge(string calldata nodeId) external {
       bytes32 challenge =  keccak256(abi.encodePacked(block.timestamp, msg.sender, nodeId));
       lastChallenges[msg.sender] = challenge;
       emit GenerateChallenge(challenge);
    }

    // submitWork miner submint a nonce value, sc check the difficulty and emit a valid pow event when success
    function submitWork(uint256 nonce, string calldata  nodeId) external {
        POWSubmission[] storage posSubmissions = powSubmissions[msg.sender];

        bytes32 lastChallenge = lastChallenges[msg.sender];
        bytes32 challengeString =  keccak256(abi.encodePacked(lastChallenge, msg.sender, nodeId));

        bytes32 hash = keccak256(abi.encodePacked(challengeString, nonce));
        uint256 difficulty = type(uint256).max / (block.prevrandao + 1);  //this difficulty need to check and test
        require(uint256(hash) < difficulty, "Work does not meet difficulty target");

        validProofs++;

        emit ValidPOWSubmitted(msg.sender, nodeId, nonce, block.timestamp, lastChallenge);
        posSubmissions.push(POWSubmission(msg.sender, nodeId, nonce, block.timestamp, lastChallenge));
    }


    function triggerNewPowRound() external onlyOwner {
        emit NewPostRound();
    }
    event NewPostRound();

    function getHashrate() external view returns (uint256) {
        uint256 elapsedTime = block.timestamp - startTime;
        // Assume each valid proof corresponds to an average of 2^256 / difficulty hashes
        uint256 averageDifficulty = type(uint256).max / validProofs;  // Simplified average difficulty
        uint256 estimatedHashes = validProofs * averageDifficulty;
        return estimatedHashes / elapsedTime;
    }

    event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge);
    event GenerateChallenge(bytes32 challenge);

}