pragma solidity ^0.5.0;

interface IDRS {
    function setup(
        string calldata assetCode,
        bytes32 collateralAssetCode,
        uint256 peggedValue,
        bytes32 peggedCurrency
    ) external returns(address);

    function mint(
        bytes32 collateralAssetCode,
        uint256 collateralAmount,
        string calldata assetCode
    ) external payable returns(bool);

    function redeem(
        address creditOwner,
        uint256 amount,
        string calldata assetCode
    ) external returns(bool);
}
