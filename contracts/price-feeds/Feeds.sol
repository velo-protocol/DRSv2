pragma solidity ^0.5.0;

import "../contract-interfaces/IVS.sol";
import "../contract-interfaces/IMED.sol";

contract Feeds is IVS {
    address public owner;
    bytes32 public value;
    bool public active;

    event LogPriceUpdate(bytes32 value);

    constructor(address _owner) public {
        active = true;
        owner = _owner;
    }

    function getWithError() external view returns (bytes32, bool) {
        return (value, !active);
    }

    function get() external view returns (bytes32) {
        require(value > 0, "value not available");
        require(active, "active must be true");
        return value;
    }

    function set(bytes32 newValue) external {
        require(msg.sender == owner, "caller must be owner");
        value = newValue;
    }

    function set(bytes32 newValue, address medAddr) external {
        require(msg.sender == owner, "caller must be owner");
        value = newValue;
        IMED(medAddr).post();

        emit LogPriceUpdate(newValue);
    }

    function disable() external {
        require(msg.sender == owner, "caller must be owner");
        active = false;
    }

}
