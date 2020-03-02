pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IFeeder.sol";
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

    address public priceRefStorage;

    uint16 constant DEFAULT_LAG_TIME = 60*60*15;
    uint16 public lagTime;

    struct MedPrice {
        uint256 price;
        bool isErr;
    }

    MedPrice next;
    MedPrice curr;

    uint256 lastUpdate;

    // Consumer contracts
    LL.List public consumers;
    using LL for LL.List;
    modifier onlyConsumer { require(consumers.has(msg.sender), "Lag | caller must be in consumer list"); _; }

    event LogLaggedPrice(uint256 laggedPrice);

    constructor(address _gov, address _priceRefStorage) public {
        consumers.init();
        priceRefStorage = _priceRefStorage;
        gov = _gov;
        lagTime = uint16(DEFAULT_LAG_TIME);
    }

    function halt() external onlyGov {
        halted = true;
    }

    function resume() external onlyGov {
        halted = false;
    }

    function setPriceRefStorage(address newPriceRefStorage) external onlyGov {
        priceRefStorage = newPriceRefStorage;
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
        curr = next = MedPrice(0, true);
        halted = true;
    }

    function isLagTimePass() public view returns (bool) {
        return getBlockTimestamp() >= lastUpdate.add(lagTime);
    }

    function post() external haltable {
        require(isLagTimePass(), "Lag | lagTime isn't pass yet");
        (uint256 medPrice, bool isErr) = IFeeder(priceRefStorage).getWithError();
        if (!isErr) {
            curr = next;
            next = MedPrice(medPrice, isErr);
            lastUpdate = calLastUpdate(getBlockTimestamp());
            emit LogLaggedPrice(curr.price);
        }
    }

    function getWithError() external view onlyConsumer returns (uint256, bool) {
        return (curr.price, curr.isErr);
    }

    function getNextWithError() external view onlyConsumer returns (uint256, bool) {
        return (next.price, next.isErr);
    }

    function get() external view onlyConsumer returns (uint256) {
        require(curr.isErr == false, "Lag | curr must not error");
        return curr.price;
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
