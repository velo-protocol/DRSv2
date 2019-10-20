pragma solidity ^0.5.0;

interface IPF {
    function setAsset(bytes32 assetCode, address assetAddress) external;
    function addAssetFiat(bytes32 assetCode, bytes32 currency) external;
    function getAssetFiat(bytes32 assetCode) external view returns (bytes32[] memory);
    function removeAssetFiat(bytes32 assetCode, uint index) external returns (bool);
    function addPriceFeeder(bytes32 assetCode, bytes32 currency, address priceFeederAddr) external;
    function getPriceFeeders(bytes32 linkId) external view returns (address[] memory);
    function removePriceFeeder(bytes32 pfId, uint index) external returns (bool);
    function setPrice(bytes32 assetCode, bytes32 currency, uint256 price) external;
}
