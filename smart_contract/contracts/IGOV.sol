pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "./IPF.sol";
import "./IRM.sol";

interface IGOV {
    function setReserveManager(address newReserveManager) external;
    function getReserveManager() external view returns (IRM);
    function setDrsAddress(address newDrsAddress) external;
    function setCollateralAsset(bytes12 assetCode, address addr, uint ratio) external;
    function getCollateralAsset(bytes12 assetCode) external view returns (IERC20);
    function setCollateralRatio(bytes12 assetCode, uint ratio) external;
    function getCollateralRatio(bytes12 assetCode) external view returns (uint);
    function setCreditIssuanceFee(uint256 newFee) external;
    function getCreditIssuanceFee() external view returns (uint256);
    function setTrustedPartner(address addr) external;
    function isTrustedPartner(address addr) external view returns (bool);
    function setPriceFeeders(address newPriceFeeders) external;
    function getPriceFeeders() external view returns (IPF);
    function collectFee(uint256 fee, bytes12 collateralAssetCode) external;
    function getCollectedFee(bytes12 collateralAssetCode) external view returns (uint256);
}
