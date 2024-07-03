// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;
//import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract LilypadPow is Initializable, OwnableUpgradeable {
    struct POWSubmission {
        address walletAddress;
        string nodeId;
        uint256 nonce;
        uint256 start_timestamp;
        uint256 complete_timestamp; //used to estimate hashrate of this submission
        bytes32 challenge; //record this to provent user never change challenge
        uint256 difficulty;
    }

    struct Challenge {
        bytes32 challenge;
        uint256 difficulty;
        string nodeId;
        uint256 timestamp;
    }

    uint256 public targetDifficulty; // =
    //555460709263765739036470010701196062214039696708679004195670928130048;
    mapping(address => POWSubmission[]) public powSubmissions;
    address[] public miners;

    mapping(address => Challenge) public lastChallenges;
    uint256 public validProofs;
    uint256 public startTime;

    uint256 public window_start;
    uint256 public window_end;
    /**
     * Init
     */

    // https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
    function initialize() public initializer {
        __Ownable_init();
        // this difficulty was calculate with this tool https://github.com/hunjixin/pow-tool/tree/main/difficulty
        // Theoretically  A machine with a hash rate of 40M has a probability of no more than 0.00001 of not finding a nonce that meets the difficulty within 30 blocks.
        targetDifficulty = 92576780592126171815437600338300430792573009392238517278497593884672;
    }

    function getMinerCount() public view returns (uint256) {
        return miners.length;
    }

    function getMiners() public view returns (address[] memory) {
        return miners;
    }

    function getMinerPowSubmissionCount(
        address addr
    ) public view returns (uint256) {
        return powSubmissions[addr].length;
    }

    function getMinerPowSubmissions(
        address addr
    ) public view returns (POWSubmission[] memory) {
        return powSubmissions[addr];
    }

    // generateChallenge gen a byte32 value as challenge value, Sc store this one for verify
    function generateChallenge(string calldata nodeId) external {
        checkTimeWindow();

        bytes32 challenge = keccak256(
            abi.encodePacked(block.timestamp, window_start, msg.sender, nodeId)
        );

        uint256 difficulty = calculate_difficulty();
        lastChallenges[msg.sender] = Challenge(
            challenge,
            difficulty,
            nodeId,
            block.timestamp
        );
        emit GenerateChallenge(challenge, difficulty);
    }

    function change_difficulty(uint256 difficulty) public onlyOwner {
        targetDifficulty = difficulty;
    }

    function calculate_difficulty() public view returns (uint256) {
        uint256 percentChange = 90 + (block.prevrandao % 21);
        return (targetDifficulty * percentChange) / 100;
    }

    // submitWork miner submint a nonce value, sc check the difficulty and emit a valid pow event when success
    function submitWork(uint256 nonce, string calldata nodeId) external {
        checkTimeWindow();

        Challenge memory lastChallenge = lastChallenges[msg.sender];
        bytes32 challenge = keccak256(
            abi.encodePacked(
                lastChallenge.timestamp,
                window_start,
                msg.sender,
                nodeId
            )
        );

        require(
            lastChallenge.challenge == challenge,
            "Work submit not compatable with challenge"
        );

        bytes32 hash = keccak256(abi.encodePacked(challenge, nonce));
        require(
            uint256(hash) < lastChallenge.difficulty,
            "Work does not meet difficulty target"
        );

        validProofs++;

        POWSubmission[] storage onwMinerPowSubmissions = powSubmissions[
            msg.sender
        ];
        if (onwMinerPowSubmissions.length == 0) {
            miners.push(msg.sender);
        }
        onwMinerPowSubmissions.push(
            POWSubmission(
                msg.sender,
                nodeId,
                nonce,
                lastChallenge.timestamp,
                block.timestamp,
                lastChallenge.challenge,
                lastChallenge.difficulty
            )
        );

        //clean last challenge to submit the same proof
        lastChallenges[msg.sender] = Challenge(0, 0, "", 0);
        emit ValidPOWSubmitted(
            msg.sender,
            nodeId,
            nonce,
            lastChallenge.timestamp,
            block.timestamp,
            lastChallenge.challenge,
            lastChallenge.difficulty
        );
    }

    function triggerNewPowRound() external onlyOwner {
        window_start = block.number;
        window_end = block.number + 36; // give 6 block time to confirm challenge, 30 block time used to calculate proof
        emit NewPowRound();
    }

    function checkTimeWindow() public view {
        require(block.number < window_end, "proof windows has closed");
    }

    event ValidPOWSubmitted(
        address walletAddress,
        string nodeId,
        uint256 nonce,
        uint256 start_timestamp,
        uint256 complete_timestamp,
        bytes32 challenge,
        uint256 difficulty
    );
    event GenerateChallenge(bytes32 challenge, uint256 difficulty);
    event NewPowRound();
}
