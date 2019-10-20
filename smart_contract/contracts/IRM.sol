pragma solidity ^0.5.0;

interface IRM {
    function injectCollateral(bytes12 assetCode, address to, uint256 amount) external;
}
