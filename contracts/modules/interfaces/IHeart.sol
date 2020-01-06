pragma solidity ^0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IPF.sol";
import "./IRM.sol";
import "../core/StableCredit.sol";

interface IHeart {
    function setReserveManager(address newReserveManager) external;
    function getReserveManager() external view returns (IRM);
    function setReserveFreeze(bytes32 assetCode, uint32 newSeconds) external;
    function getReserveFreeze(bytes32 assetCode) external view returns (uint32);
    function setDrsAddress(address newDrsAddress) external;
    function getDrsAddress() external view returns (address);
    function setCollateralAsset(bytes32 assetCode, address addr, uint ratio) external;
    function getCollateralAsset(bytes32 assetCode) external view returns (IERC20);
    function setCollateralRatio(bytes32 assetCode, uint ratio) external;
    function getCollateralRatio(bytes32 assetCode) external view returns (uint);
    function setCreditIssuanceFee(uint256 newFee) external;
    function getCreditIssuanceFee() external view returns (uint256);
    function setTrustedPartner(address addr) external;
    function isTrustedPartner(address addr) external view returns (bool);
    function setPriceFeeders(address newPriceFeeders) external;
    function getPriceFeeders() external view returns (IPF);
    function collectFee(uint256 fee, bytes32 collateralAssetCode) external;
    function getCollectedFee(bytes32 collateralAssetCode) external view returns (uint256);
    function withdrawFee(bytes32 collateralAssetCode, uint256 amount) external;
    function addStableCredit(StableCredit stableCredit) external;
    function getStableCredit(bytes32 stableCreditId) external view returns (StableCredit);
}
