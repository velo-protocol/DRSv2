pragma solidity ^0.5.0;

interface IRM {
    function lockReserve(bytes32 assetCode, address from, uint256 amount) external;
    function releaseReserve(bytes32 lockedReserveId, bytes32 assetCode, uint256 amount) external;
    function injectCollateral(bytes32 assetCode, address to, uint256 amount) external;
}
