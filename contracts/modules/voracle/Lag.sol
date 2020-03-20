pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IMedianizer.sol";
import "../interfaces/ILag.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";

contract Lag is Initializable, ILag {
    using SafeMath for uint256;
    using SafeMath for uint16;

    // Governance
    address public owner;
    modifier onlyOwner {
        require(msg.sender == owner, "Lag.onlyOwner: The message sender is not found or does not have sufficient permission");
        _;
    }

    // active
    bool public active;
    modifier mustBeActive {
        require(active == true, "Lag.mustBeActive: lag is inactive");
        _;
    }

    address public medianizerAddr;

    uint256 public minimumPeriod;
    uint256 public minimumPeriodLimit;

    struct MedPrice {
        uint256 price;
        bool isErr;
    }

    MedPrice nextPrice;
    MedPrice currentPrice;

    uint256 public timeLastUpdate;

    event LagVoid(address caller, address lag, bool isActive);
    event LagActivate(address caller, address lag, bool isActive);

    function initialize(address _owner, address _medianizerAddr) public initializer {
        medianizerAddr = _medianizerAddr;
        owner = _owner;
        minimumPeriod = 15 minutes;
        minimumPeriodLimit = 30 days;
        active = true;
        currentPrice = MedPrice(0, true);
        nextPrice = MedPrice(0, true);
    }

    function deactivate() external onlyOwner {
        active = false;
    }

    function activate() external onlyOwner {
        require(active == false, "Lag.activate: lag is active");
        require(currentPrice.price > 0 && nextPrice.price > 0, "Lag.activate: price is not in a correct state");
        active = true;
        emit LagActivate(msg.sender, address(this), active);
    }

    function setMedianizer(address newMedianizerAddr) external onlyOwner {
        medianizerAddr = newMedianizerAddr;
    }

    function getBlockTimestamp() internal view returns (uint) {
        return block.timestamp;
    }

    function setMinimumPeriod(int256 newMinimumPeriod) external onlyOwner {
        require(newMinimumPeriod >= 0, "Lag.setMinimumPeriod: minimumPeriod value must not be less than 0");
        require(uint256(newMinimumPeriod) <= minimumPeriodLimit, "Lag.setMinimumPeriod: minimumPeriod value must not be greater than 2592000");
        minimumPeriod = uint256(newMinimumPeriod);
    }

    function void() external onlyOwner {
        currentPrice = MedPrice(0, true);
        active = false;
        emit LagVoid(msg.sender, address(this), active);
    }

    function isMinimumPeriodPass() public view returns (bool) {
        return getBlockTimestamp() >= timeLastUpdate.add(minimumPeriod);
    }

    function post() external returns (bool) {
        require(isMinimumPeriodPass(), "Lag.post: minimum period for the post function has not passed");
        (uint256 medPrice, bool isActive, bool isErr) = IMedianizer(medianizerAddr).getWithError();
        if (!isActive) {
            active = false;
            return true;
        }

        if (!isErr && medPrice > 0) {
            currentPrice = nextPrice;
            nextPrice = MedPrice(medPrice, isErr);
            timeLastUpdate = getBlockTimestamp();
        }
        return true;
    }

    function getWithError() external view returns (uint256, bool, bool) {
        return (currentPrice.price, active, currentPrice.isErr);
    }

    function getNextWithError() external view returns (uint256, bool, bool) {
        return (nextPrice.price, active, nextPrice.isErr);
    }

    function get() external view returns (uint256) {
        require(currentPrice.isErr == false, "Lag.get: currentPrice has an error");
        require(active, "Lag.get: lag is inactive");
        return currentPrice.price;
    }
}
