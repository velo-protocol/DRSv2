pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IHeart.sol";
import "../../../contracts/modules/interfaces/IRM.sol";
import "../../../contracts/modules/interfaces/IPF.sol";
import "../../../contracts/modules/interfaces/IERC20.sol";
import "../../../contracts/modules/core/StableCredit.sol";

contract MockIHeart is IHeart {

    IRM public rm;
    IPF public pf;
    IERC20 public collateralAsset;
    StableCredit public vTHB;
    StableCredit public vUSD;

    constructor (IRM _rm, IPF _pf, IERC20 _collateralAsset, StableCredit _vTHB, StableCredit _vUSD) {
        rm = _rm;
        pf = _pf;
        collateralAsset = _collateralAsset;
        vTHB = _vTHB;
        vUSD = _vUSD;
    }

    function setReserveManager(address) external {}

    function getReserveManager() external view returns (IRM) {
        return rm;
    }

    function setReserveFreeze(bytes32, uint32) external {}

    function getReserveFreeze(bytes32) external view returns (uint32) {
        return uint(10);
    }

    function setDrsAddress(address) external {}

    function getDrsAddress() external view returns (address) {
        return address(1);
    }

    function setCollateralAsset(bytes32, address, uint) external {}

    function getCollateralAsset(bytes32) external view returns (IERC20) {
        return collateralAsset;
    }

    function setCollateralRatio(bytes32, uint) external {}

    function getCollateralRatio(bytes32) external view returns (uint) {
        return uint(10);
    }

    function setCreditIssuanceFee(uint256) external {}

    function getCreditIssuanceFee() external view returns (uint256);

    function setTrustedPartner(address) external {}

    function isTrustedPartner(address) external view returns (bool) {
        return true;
    }

    function setGovernor(address) external {}

    function isGovernor(address) external view returns (bool) {
        return true;
    }

    function setPriceFeeders(address) external {}

    function getPriceFeeders() external view returns (IPF) {
        return pf;
    }

    function collectFee(uint256 fee, bytes32 collateralAssetCode) external {}

    function getCollectedFee(bytes32 collateralAssetCode) external view returns (uint256) {
        return uint(10);
    }

    function withdrawFee(bytes32, uint256) external {}

    function addStableCredit(StableCredit) external {}

    function getStableCreditById(bytes32) external view returns (StableCredit) {
        return vTHB;
    }

    function getRecentStableCredit() external view returns (StableCredit) {
        return vUSD;
    }

    function getNextStableCredit(bytes32 stableCreditId) external view returns (StableCredit) {
        return vTHB;
    }

    function getStableCreditCount() external view returns (uint8) {
        return uint(2);
    }

    function setAllowedLink(bytes32 linkId, bool enable) external {}

    function isLinkAllowed(bytes32 linkId) external view returns (bool) {
        return true;
    }
}
