pragma solidity ^0.5.0;

interface IDRS {
    function setup(
        string calldata assetCode,
        bytes12 collateralAssetCode,
        bytes12 peggedCurrency,
        uint256 peggedValue
    ) external returns(address);

    function mint(
        bytes12 collateralAssetCode,
        uint256 collateralAmount,
        string calldata assetCode
    ) external payable returns(bool);

    function redeem(
        address creditOwner,
        uint256 amount,
        string calldata assetCode
    ) external returns(bool);

    function rebalance(
        address creditOwner,
        string calldata assetCode
    ) external returns(bool);
}
