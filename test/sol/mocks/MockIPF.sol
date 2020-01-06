pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IPF.sol";

contract MockIPF is IPF {
    function setAsset(bytes32, address) external {}
    function addAssetFiat(bytes32, bytes32) external {}

    function getAssetFiat(bytes32) external view returns (bytes32[] memory) {
        bytes32[] memory assetFiat;
        assetFiat[0] = "vTHB";
        return assetFiat;
    }

    function removeAssetFiat(bytes32, uint) external returns (bool) {
        return true;
    }

    function addPriceFeeder(bytes32, bytes32, address) external {}

    function getPriceFeeders(bytes32) external view returns (address[] memory) {
        address[] memory priceFeeders;
        priceFeeders[0] = address(1);
        return priceFeeders;
    }

    function removePriceFeeder(bytes32, uint) external returns (bool) {
        return true;
    }

    function setPrice(bytes32, bytes32, uint256) external {}

    function getMedianPrice(bytes32) external view returns (uint256) {
        return 1;
    }
}
