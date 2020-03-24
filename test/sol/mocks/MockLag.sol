pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/ILag.sol";

contract MockLag is ILag {
    uint256 public value;

    constructor(uint256 _value) public {
        value = _value;
    }

    function post() external returns (bool){
        return true;
    }

    function getWithError() external view returns (uint256, bool, bool){
        return (value, true, false);
    }

    function void() external {

    }

    function activate() external {

    }

    function setMinimumPeriod(int256) external {

    }
}