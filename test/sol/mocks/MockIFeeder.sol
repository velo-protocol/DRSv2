pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IFeeder.sol";

contract MockIFeeder is IFeeder {

    uint256 public value;

    constructor(uint256 _value) public {
        value = _value;
    }

    function getWithError() external view returns (uint256, uint256, bool) {
        return (value, now, false);
    }

    function get() external view returns (uint256) {
        return value;
    }
    function post(uint256 newValue) external {
        value = newValue;
    }

    function enable() external {
        // do something
    }

    function disable() external {
        // do something
    }
}
