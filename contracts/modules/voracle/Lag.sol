pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IMedianizer.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract Lag {
    using SafeMath for uint256;
    using SafeMath for uint16;

    // Governance
    address public gov;
    modifier onlyGov {
        require(msg.sender == gov, "Lag | caller must be Gov");
        _;
    }

    // Halt
    bool public halted;
    modifier haltable { require(halted == false, "Lag | Lag has been halted"); _; }

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

    // Consumer contracts
    LL.List public consumers;
    using LL for LL.List;
    modifier onlyConsumer { require(consumers.has(msg.sender), "Lag | caller must be in consumer list"); _; }

    event LogLaggedPrice(uint256 laggedPrice);

    constructor(address _gov, address _priceRefStorage) public {
        consumers.init();
        medianizerAddr = _priceRefStorage;
        gov = _gov;
        lagTime = uint16(DEFAULT_LAG_TIME);
    }

    function halt() external onlyGov {
        halted = true;
    }

    function resume() external onlyGov {
        halted = false;
    }

    function setMedianizer(address newMedianizerAddr) external onlyGov {
        medianizerAddr = newMedianizerAddr;
    }

    function getBlockTimestamp() internal view returns (uint) {
        return block.timestamp;
    }

    function calLastUpdate(uint timestamp) internal view returns (uint256) {
        require(lagTime > 0, "Lag | lagTime must more than 0");
        return timestamp.sub(timestamp % lagTime);
    }

    function setLagTime(uint16 newLagTime) external onlyGov {
        require(newLagTime > 0, "Lag | newLagTime must more than 0");
        lagTime = newLagTime;
    }

    function void() external onlyGov {
        currentPrice = nextPrice = MedPrice(0, true);
        halted = true;
    }

    function isLagTimePass() public view returns (bool) {
        return getBlockTimestamp() >= timeLastUpdate.add(lagTime);
    }

    function post() external haltable {
        require(isLagTimePass(), "Lag.post: lag time is not pass yet");
        (uint256 medPrice, , bool isErr) = IMedianizer(medianizerAddr).getWithError();
        if (!isErr) {
            currentPrice = nextPrice;
            nextPrice = MedPrice(medPrice, isErr);
            timeLastUpdate = calLastUpdate(getBlockTimestamp());
            emit LogLaggedPrice(currentPrice.price);
        }
    }

    function getWithError() external view onlyConsumer returns (uint256, bool) {
        return (currentPrice.price, currentPrice.isErr);
    }

    function getNextWithError() external view onlyConsumer returns (uint256, bool) {
        return (nextPrice.price, nextPrice.isErr);
    }

    function get() external view onlyConsumer returns (uint256) {
        require(currentPrice.isErr == false, "Lag | curr must not error");
        return currentPrice.price;
    }

    function addConsumer(address newConsumer) external onlyGov {
        require(newConsumer != address(0), "Lag | newConsumer must no be address(0)");
        consumers.add(newConsumer);
    }

    function removeConsumer(address consumer, address prevConsumer) external onlyGov {
        consumers.remove(consumer, prevConsumer);
    }

    function getConsumers() public view returns (address[] memory) {
        return consumers.getAll();
    }

    function addConsumers(address[] calldata newConsumers) external onlyGov {
        for(uint i = 0; i < newConsumers.length; i++) {
            require(newConsumers[i] != address(0), "Lag | newConsumers[x] must not be address(0)");
            consumers.add(newConsumers[i]);
        }
    }

}
