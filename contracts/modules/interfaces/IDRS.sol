pragma solidity ^0.5.0;

interface IDRS {
    function setup(
        bytes32 collateralAssetCode,
        bytes32 peggedCurrency,
        string calldata assetCode,
        uint256 peggedValue
    ) external returns (string memory, address);

    function mintFromCollateralAmount(
        uint256 collateralAmount,
        string calldata assetCode
    ) external payable returns (bool);

    function mintFromStableCreditAmount(
        uint256 stableCreditAmount,
        string calldata assetCode
    ) external payable returns (bool);

    function redeem(
        uint256 amount,
        string calldata assetCode
    ) external returns (bool);

    function rebalance(
        string calldata assetCode
    ) external returns (bool);

    function getExchange(
        string calldata assetCode
    ) external view returns (string memory, bytes32, uint256);
}
