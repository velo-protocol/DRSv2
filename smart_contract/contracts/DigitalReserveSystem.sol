pragma solidity ^0.5.0;

import "./IGOV.sol";
import "./StableCredit.sol";
import "./IDRS.sol";

contract DigitalReserveSystem is IDRS {
    using SafeMath for uint256;

    /*
        StableCredit related mapping
        stableCredits map between keccak256(stableCreditOwnerAddress, stableCreditCode) => StableCredit
    */
    mapping(bytes32 => StableCredit) public stableCredits;

    IGOV gov;

    event Setup(
        string indexed assetCode,
        address indexed assetAddress
    );

    event Mint(
        string indexed assetCode,
        uint256 mintAmount,
        address indexed assetAddress,
        bytes12 indexed collateralAssetCode,
        uint256 collateralAmount
    );

    constructor(address governanceAddress) public {
        gov = IGOV(governanceAddress);
    }

    function setup(
        string calldata assetCode,
        bytes12 collateralAssetCode,
        bytes12 peggedCurrency,
        uint256 peggedValue
    ) external returns(address) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(msg.sender, assetCode));
        address collateralAddress = address(gov.getCollateralAsset(collateralAssetCode));

        require(collateralAddress != address(0x0), "collateralAssetCode has not been whitelisted");
        require(address(stableCredits[stableCreditId]) == address(0x0), "trusted partner cannot setup the same asset code");
        require(gov.isTrustedPartner(msg.sender), "only trusted partner can setup the stable credit");

        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));

        require(gov.getPriceFeeders().getMedianPrice(linkId) > 0, "collateralAssetCode must has value more than 0");

        StableCredit newStableCredit = new StableCredit(
            msg.sender,
            assetCode,
            peggedValue,
            peggedCurrency,
            collateralAddress,
            collateralAssetCode
        );

        stableCredits[stableCreditId] = newStableCredit;

        emit Setup(assetCode, address(newStableCredit));

        return address(newStableCredit);
    }

    /*
        TODO:
            - mint(uint256 mintAmount, string calldata assetCode, bytes12 collateralAssetCode)
    */
    function mint(
        bytes12 collateralAssetCode,
        uint256 collateralAmount,
        string calldata assetCode
    ) external payable returns(bool) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(msg.sender, assetCode));

        require(gov.isTrustedPartner(msg.sender), "only trusted partner can mint the stable credit");
        require(address(stableCredits[stableCreditId]) == address(0x0), "stableCredit not exist");

        StableCredit credit = stableCredits[stableCreditId];
        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, credit.peggedCurrency()));

        (uint256 mintAmount, uint256 fee) = _calMintStableCredit(credit, linkId, collateralAmount);

        require(gov.getCollateralAsset(collateralAssetCode) == credit.collateral(), "collateralAsset must be the same");
        require(gov.getPriceFeeders().getMedianPrice(linkId) != 0, "median price ref mut not be zero");

        gov.getCollateralAsset(collateralAssetCode).transferFrom(msg.sender, address(gov), fee);
        gov.getCollateralAsset(collateralAssetCode).transferFrom(msg.sender, address(credit), collateralAmount.sub(fee));

        credit.mint(msg.sender, mintAmount);
        credit.approveCollateral();

        gov.collectFee(fee, collateralAssetCode);

        emit Mint(
            assetCode,
            mintAmount,
            address(credit),
            collateralAssetCode,
            collateralAmount
        );

        return true;
    }

    function redeem(
        address creditOwner,
        uint256 amount,
        string calldata assetCode
    ) external returns(bool) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(creditOwner, assetCode));

        require(address(stableCredits[stableCreditId]) != address(0x0), "stableCredit not existed");

        StableCredit credit = stableCredits[stableCreditId];
        bytes32 linkId = keccak256(abi.encodePacked(credit.collateralAssetCode(), credit.peggedCurrency()));

        _rebalance(creditOwner, assetCode);
        uint256 returnAmount = _callCollateral(credit, linkId, amount);

        credit.redeem(msg.sender, amount, returnAmount);
        credit.approveCollateral();

        return true;
    }

    function rebalance(
        address creditOwner,
        string calldata assetCode
    ) external returns(bool) {
        return _rebalance(creditOwner, assetCode);
    }

    function _rebalance(
        address creditOwner,
        string memory assetCode
    ) private returns(bool) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(creditOwner, assetCode));

        StableCredit credit = stableCredits[stableCreditId];
        bytes32 linkId = keccak256(abi.encodePacked(credit.collateralAssetCode(), credit.peggedCurrency()));

        uint256 collateralAmount = _callCollateral(credit, linkId, credit.totalSupply());

        if (collateralAmount >= credit.collateral().balanceOf(address(credit))) {
            gov.getReserveManager().injectCollateral(credit.collateralAssetCode(), address(credit), collateralAmount.sub(credit.collateral().balanceOf(address(credit))));
        } else {
            gov.getCollateralAsset(credit.collateralAssetCode()).transferFrom(address(credit), address(gov.getReserveManager()), credit.collateral().balanceOf(address(credit)).sub(collateralAmount));
        }

        return true;
    }

    function _calMintStableCredit(StableCredit credit, bytes32 linkId, uint256 collateralAmount) private view returns (uint256, uint256) {
        uint256 fee = collateralAmount.mul(gov.getCreditIssuanceFee()).div(10000);
        uint256 mintAmount = collateralAmount.sub(fee).mul(100).div(gov.getCollateralRatio(credit.collateralAssetCode())).mul(gov.getPriceFeeders().getMedianPrice(linkId)).div(credit.peggedValue());

        return (mintAmount, fee);
    }

    function _calCollateralWithFee(StableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256, uint256) {
        uint256 collateral = _callCollateral(credit, linkId, creditAmount);
        uint256 fee = collateral.mul(gov.getCreditIssuanceFee()).div(10000);

        return (collateral, fee);
    }

    function _callCollateral(StableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256) {
        uint256 collateralAmountAfterRatio = creditAmount.mul(credit.peggedValue()).div(gov.getPriceFeeders().getMedianPrice(linkId)).mul(gov.getCollateralRatio(credit.collateralAssetCode())).div(100);
        return collateralAmountAfterRatio;
    }

    function collateralOf(address creditOwner, string calldata assetCode) external view returns (uint256, address) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(creditOwner, assetCode));
        return stableCredits[stableCreditId].getCollateralDetail();
    }
}
