pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IPRS.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";

contract MockTargetProxyV1 is Initializable, IPRS {

    uint256 public value;

    function initialize() public initializer {
        value = 10;
    }

    function getWithError() external view returns (uint256, bool) {
        return (value, false);
    }

    function get() external view returns (uint256) {
        return value;
    }
    function set(uint256 newValue) external {
        value = newValue;
    }

    function disable() external {
        // do something
    }
}
