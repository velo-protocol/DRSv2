pragma solidity ^0.5.0;

import "./MockTargetProxyV1.sol";

contract MockTargetProxyV2 is MockTargetProxyV1 {
    function post(uint256 newValue) external {
        value = newValue*2;
    }
}
