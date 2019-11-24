pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./IPF.sol";

contract PriceFeeders is IPF, WhitelistAdminRole {

    mapping(bytes32 => address) private assetList;

    /*
        PriceFeeder related mapping
        linkId = keccak256(assetCode, currency)
        pfId = keccak256(assetCode, currency, address of price feeder)
        collateralAssetFiat map between collateralAssetCode => currency[]
        priceFeeders map between linkId => address[] of price feeders
        prices map between pfId => price
        medianPrices map between linkId => medianPrice
    */
    mapping(bytes32 => bytes32[]) public assetFiat;
    mapping(bytes32 => address[]) public priceFeeders;
    mapping(bytes32 => uint256) public prices;
    mapping(bytes32 => uint256) public medianPrices;

    function setAsset(bytes32 assetCode, address assetAddress) external onlyWhitelistAdmin {
        assetList[assetCode] = assetAddress;
    }

    function addAssetFiat(bytes32 assetCode, bytes32 currency) external onlyWhitelistAdmin {
        int index = _getLinkIndex(assetCode, currency);

        require(address(assetList[assetCode]) != address(0x0), "assetCode has not been added to the assetList");
        require(index == -1, "asset has been linked with this fiat");

        assetFiat[assetCode].push(currency);
    }

    function getAssetFiat(bytes32 assetCode) external view returns (bytes32[] memory){
        return assetFiat[assetCode];
    }

    function _getLinkIndex(bytes32 assetCode, bytes32 currency) private view returns (int) {
        int index = -1;
        for (uint i = 0; i < assetFiat[assetCode].length; i++) {
            if(assetFiat[assetCode][i] == currency) {
                index = int(i);
            }
        }
        return index;
    }

    function removeAssetFiat(bytes32 assetCode, uint index) external onlyWhitelistAdmin returns (bool) {
        if(index >= assetFiat[assetCode].length) return false;

        for (uint i = index; i<assetFiat[assetCode].length-1; i++){
            assetFiat[assetCode][i] = assetFiat[assetCode][i+1];
        }
        assetFiat[assetCode].length--;

        return true;
    }

    function addPriceFeeder(bytes32 assetCode, bytes32 currency, address priceFeederAddr) external onlyWhitelistAdmin {
        bytes32 linkId = keccak256(abi.encodePacked(assetCode, currency));

        int linkIndex = _getLinkIndex(assetCode, currency);
        int pfIndex = _getPriceFeederIndex(linkId, priceFeederAddr);

        require(address(assetList[assetCode]) != address(0x0), "assetCode has not been added to assetList");
        require(linkIndex != -1, "asset has not been linked with this fiat");
        require(pfIndex == -1, "this price feeder has been added to the list");

        priceFeeders[linkId].push(priceFeederAddr);
    }

    function getPriceFeeders(bytes32 linkId) external view returns (address[] memory) {
        return priceFeeders[linkId];
    }

    function _getPriceFeederIndex(bytes32 id, address priceFeeder) private view returns (int) {
        int index = -1;
        for (uint i = 0; i < priceFeeders[id].length; i++) {
            if(priceFeeders[id][i] == priceFeeder) {
                index = int(i);
            }
        }
        return index;
    }

    function removePriceFeeder(bytes32 pfId, uint index) external onlyWhitelistAdmin returns (bool) {
        if(index >= priceFeeders[pfId].length) return false;

        for (uint i = index; i<priceFeeders[pfId].length-1; i++){
            priceFeeders[pfId][i] = priceFeeders[pfId][i+1];
        }
        priceFeeders[pfId].length--;

        return true;
    }

    /*
    TODO:
        - Calculate median everytime the price is updated
    */
    function setPrice(bytes32 assetCode, bytes32 currency, uint256 price) external {
        bytes32 linkId = keccak256(abi.encodePacked(assetCode, currency));

        int pfIndex = _getPriceFeederIndex(linkId, msg.sender);
        int linkIndex = _getLinkIndex(assetCode, currency);

        require(address(assetList[assetCode]) != address(0x0), "assetCode has not been added to assetList");
        require(linkIndex != -1, "asset has not been linked with this fiat");
        require(pfIndex != -1, "msg.sender did not have an authority to set the price");

        bytes32 pfId = keccak256(abi.encodePacked(assetCode, currency, msg.sender));
        prices[pfId] = price;
        medianPrices[linkId] = price;
    }

    function getMedianPrice(bytes32 linkId) external view returns (uint256) {
        return medianPrices[linkId];
    }
}
