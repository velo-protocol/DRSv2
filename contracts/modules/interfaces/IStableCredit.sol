pragma solidity ^0.5.0;

import "../interfaces/ICollateralAsset.sol";

interface IStableCredit {
    function mint(address recipient, uint256 amount) external;

    function burn(address tokenOwner, uint256 amount) external;

    function redeem(address redeemer, uint burnAmount, uint256 returnAmount) external;

    function approveCollateral() external;

    function getCollateralDetail() external view returns (uint256, address);

    function getId() external view returns (bytes32);

    function transferCollateralToReserve(uint256 amount) external returns (bool);

    // Getter functions
    function assetCode() external view returns (string memory);

    function peggedValue() external view returns (uint256);

    function peggedCurrency() external view returns (bytes32);

    function creditOwner() external view returns (address);

    function collateral() external view returns (ICollateralAsset);

    function collateralAssetCode() external view returns (bytes32);

    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);
}
