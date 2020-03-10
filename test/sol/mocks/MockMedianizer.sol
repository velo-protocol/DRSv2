pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IMedianizer.sol";

contract MockMedianizer is IMedianizer {

    uint256 public value;

    constructor(uint256 _value) public {
        value = _value;
    }

    function post() external {
        // mock stuffs
    }

    function getWithError() external view returns (uint256, bool, bool){
        return (value, true, false);
    }
}
