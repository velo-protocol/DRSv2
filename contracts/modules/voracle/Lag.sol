pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IMedianizer.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract Lag {
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

    uint16 constant DEFAULT_LAG_TIME = 15 minutes;
    uint16 public lagTime;

    struct MedPrice {
        uint256 price;
        bool isErr;
    }

    MedPrice nextPrice;
    MedPrice currentPrice;

    uint256 timeLastUpdate;

    event LogLaggedPrice(uint256 laggedPrice);
    event VoidLagEvent(address caller, address lagAddress, bool currentStatus);

    constructor(address _owner, address _medianizerAddr) public {
        medianizerAddr = _medianizerAddr;
        owner = _owner;
        lagTime = uint16(DEFAULT_LAG_TIME);
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
        require(lagTime > 0, "Lag | lagTime must more than 0");
        return timestamp.sub(timestamp % lagTime);
    }

    function setLagTime(uint16 newLagTime) external onlyOwner {
        require(newLagTime > 0, "Lag | newLagTime must more than 0");
        lagTime = newLagTime;
    }

    function void() external onlyOwner {
        currentPrice = nextPrice = MedPrice(0, true);
        active = false;
        emit VoidLagEvent(msg.sender, address(this), active);
    }

    function isLagTimePass() public view returns (bool) {
        return getBlockTimestamp() >= timeLastUpdate.add(lagTime);
    }

    function post() external mustBeActive {
        require(isLagTimePass(), "Lag.post: lag time is not pass yet");
        (uint256 medPrice, , bool isErr) = IMedianizer(medianizerAddr).getWithError();
        if (!isErr) {
            currentPrice = nextPrice;
            nextPrice = MedPrice(medPrice, isErr);
            timeLastUpdate = calLastUpdate(getBlockTimestamp());
            emit LogLaggedPrice(currentPrice.price);
        }
    }

    function getWithError() external view returns (uint256, bool) {
        return (currentPrice.price, currentPrice.isErr);
    }

    function getNextWithError() external view returns (uint256, bool) {
        return (nextPrice.price, nextPrice.isErr);
    }

    function get() external view returns (uint256) {
        require(currentPrice.isErr == false, "Lag.get: currentPrice has an error");
        return currentPrice.price;
    }
}
