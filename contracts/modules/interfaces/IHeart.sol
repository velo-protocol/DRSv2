pragma solidity ^0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IPrice.sol";
import "./IRM.sol";
import "./IStableCredit.sol";
import "../core/StableCredit.sol";

interface IHeart {
    function setReserveManager(address newReserveManager) external;

    function getReserveManager() external view returns (IRM);

    function setReserveFreeze(bytes32 assetCode, uint32 newSeconds) external;

    function getReserveFreeze(bytes32 assetCode) external view returns (uint32);

    function setDrsAddress(address newDrsAddress) external;

    function getDrsAddress() external view returns (address);

    function setCollateralAsset(bytes32 assetCode, address addr, uint ratio) external;

    function getCollateralAsset(bytes32 assetCode) external view returns (ICollateralAsset);

    function setCollateralRatio(bytes32 assetCode, uint ratio) external;

    function getCollateralRatio(bytes32 assetCode) external view returns (uint);

    function setCreditIssuanceFee(uint256 newFee) external;

    function getCreditIssuanceFee() external view returns (uint256);

    function setTrustedPartner(address addr) external;

    function isTrustedPartner(address addr) external view returns (bool);

    function setGovernor(address addr) external;

    function isGovernor(address addr) external view returns (bool);

    function addPrice(bytes32 linkId, IPrice newPrice) external;

    function getPriceContract(bytes32 linkId) external view returns (IPrice);

    function collectFee(uint256 fee, bytes32 collateralAssetCode) external;

    function getCollectedFee(bytes32 collateralAssetCode) external view returns (uint256);

    function withdrawFee(bytes32 collateralAssetCode, uint256 amount) external;

    function addStableCredit(IStableCredit stableCredit) external;

    function getStableCreditById(bytes32 stableCreditId) external view returns (IStableCredit);

    function getRecentStableCredit() external view returns (IStableCredit);

    function getNextStableCredit(bytes32 stableCreditId) external view returns (IStableCredit);

    function getStableCreditCount() external view returns (uint8);

    function setAllowedLink(bytes32 linkId, bool enable) external;

    function isLinkAllowed(bytes32 linkId) external view returns (bool);
}
