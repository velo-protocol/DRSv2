pragma solidity ^0.5.0;

import "../interfaces/IFeeder.sol";

contract Feeder is IFeeder {
    address public owner;
    uint256 public value;
    uint256 public valueTimestamp;
    bool public active;

    bytes32 public fiatCode;
    bytes32 public collateralCode;

    event PriceSet(uint256 prevValue, uint prevTimestamp, uint curPrice, uint curTimestamp);

    constructor(address _owner, bytes32 _fiatCode, bytes32 _collateralCode) public {
        active = true;
        owner = _owner;
        fiatCode = _fiatCode;
        collateralCode = _collateralCode;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Feeder.onlyOwner: caller must be an owner of this contract");
        _;
    }

    function getWithError() external view returns (uint256, uint256, bool) {
        return (value, valueTimestamp, !active);
    }

    function get() external view returns (uint256) {
        require(value > 0, "Feeder.get: value not available");
        require(active, "Feeder.get: active must be true");
        return value;
    }

    function post(uint256 newValue) external onlyOwner {
        uint oldValue = value;
        uint oldTimestamp = valueTimestamp;

        require(newValue > 0, "Feeder.set: newValue must be greater than 0");
        value = newValue;
        valueTimestamp = block.timestamp;

        emit PriceSet(
            oldValue,
            oldTimestamp,
            value,
            valueTimestamp
        );
    }

    function enable() external onlyOwner {
        active = true;
    }

    function disable() external onlyOwner {
        active = false;
    }
}
