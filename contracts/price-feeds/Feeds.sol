pragma solidity ^0.5.0;

import "../contract-interfaces/IPRS.sol";
import "../contract-interfaces/IMED.sol";

contract Feeds is IPRS {
    address public owner;
    uint256 public value;
    bool public active;

    event LogPriceUpdate(uint256 value);

    constructor(address _owner) public {
        active = true;
        owner = _owner;
    }

    function getWithError() external view returns (uint256, bool) {
        return (value, !active);
    }

    function get() external view returns (uint256) {
        require(value > 0, "Feeds | value not available");
        require(active, "Feeds | active must be true");
        return value;
    }

    function set(uint256 newValue) external {
        require(newValue > 0, "Feeds | newValue must more than 0");
        require(msg.sender == owner, "Feeds | caller must be owner");
        value = newValue;
    }

    function postMed(uint256 newValue, address medAddr) external {
        require(newValue > 0, "Feeds | newValue must more than 0");
        require(msg.sender == owner, "Feeds | caller must be owner");
        value = newValue;
        IMED(medAddr).post();

        emit LogPriceUpdate(newValue);
    }

    function disable() external {
        require(msg.sender == owner, "Feeds | caller must be owner");
        active = false;
    }

}
