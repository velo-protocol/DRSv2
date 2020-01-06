pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IRM.sol";

contract MockIRM is IRM {
    constructor() public {
    }

    function lockReserve(bytes32, address, uint256) external {}

    function releaseReserve(bytes32, bytes32, uint256) external {}

    function injectCollateral(bytes32, address, uint256) external {}
}
