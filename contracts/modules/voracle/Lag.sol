pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IMedianizer.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "../interfaces/ILag.sol";

contract Lag is ILag {
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

    uint16 constant DEFAULT_MINIMUM_PERIOD = 15 minutes;
    uint256 public minimumPeriod;

    struct MedPrice {
        uint256 price;
        bool isErr;
    }

    MedPrice nextPrice;
    MedPrice currentPrice;

    uint256 timeLastUpdate;

    event LagVoid(address caller, address lag, bool isActive);

    constructor(address _owner, address _medianizerAddr) public {
        medianizerAddr = _medianizerAddr;
        owner = _owner;
        minimumPeriod = uint16(DEFAULT_MINIMUM_PERIOD);
        active = true;
    }

    function deactivate() external onlyOwner {
        active = false;
    }

    function activate() external onlyOwner {
        active = true;
    }

    function setMedianizer(address newMedianizerAddr) external onlyOwner {
        medianizerAddr = newMedianizerAddr;
    }

    function getBlockTimestamp() internal view returns (uint) {
        return block.timestamp;
    }

    function calLastUpdate(uint timestamp) internal view returns (uint256) {
        require(minimumPeriod > 0, "Lag | lagTime must more than 0");
        return timestamp.sub(timestamp % minimumPeriod);
    }

    function setLagTime(uint256 newLagTime) external onlyOwner {
        require(newLagTime > 0, "Lag | newLagTime must more than 0");
        minimumPeriod = newLagTime;
    }

    function void() external onlyOwner {
        currentPrice = MedPrice(0, true);
        active = false;
        emit LagVoid(msg.sender, address(this), active);
    }

    function isMinimumPeriodPass() public view returns (bool) {
        return getBlockTimestamp() >= timeLastUpdate.add(minimumPeriod);
    }

    function post() external mustBeActive returns (bool) {
        require(isMinimumPeriodPass(), "Lag.post: minimum period for the post function has not passed");
        (uint256 medPrice, bool isActive, bool isErr) = IMedianizer(medianizerAddr).getWithError();
        if (!isActive) {
            active = false;
            nextPrice = MedPrice(nextPrice.price, isErr);
            return true;
        }

        if (!isErr && medPrice > 0) {
            currentPrice = nextPrice;
            nextPrice = MedPrice(medPrice, isErr);
            timeLastUpdate = calLastUpdate(getBlockTimestamp());
        } else {
            nextPrice = MedPrice(nextPrice.price, true);
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
