pragma solidity ^0.5.0;

import "../interfaces/IHeart.sol";
import "../interfaces/IDRS.sol";
import "./StableCredit.sol";

contract DigitalReserveSystem is IDRS {
    using SafeMath for uint256;
    IHeart public heart;

    event Setup(
        string indexed assetCode,
        address indexed assetAddress
    );

    event Mint(
        string indexed assetCode,
        uint256 mintAmount,
        address indexed assetAddress,
        bytes32 indexed collateralAssetCode,
        uint256 collateralAmount
    );

    constructor(address heartAddr) public {
        heart = IHeart(heartAddr);
    }

    function setup(
        bytes32 collateralAssetCode,
        bytes32 peggedCurrency,
        string calldata assetCode,
        uint256 peggedValue
    ) external returns(address) {
        bytes32 stableCreditId = getStableCreditId(assetCode);
        address stableCreditAddr = address(heart.getStableCreditById(stableCreditId));
        address collateralAddr = address(heart.getCollateralAsset(collateralAssetCode));

        require(collateralAddr != address(0x0), "collateralAssetCode has not been whitelisted");
        require(stableCreditAddr == address(0x0), "trusted partner cannot setup the same asset code");
        require(heart.isTrustedPartner(msg.sender), "only trusted partner can setup the stable credit");

        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));
        require(heart.getPriceFeeders().getMedianPrice(linkId) > 0, "collateralAssetCode must has value more than 0");

        StableCredit newStableCredit = new StableCredit(
            peggedCurrency,
            msg.sender,
            collateralAssetCode,
            collateralAddr,
            assetCode,
            peggedValue,
            address(heart)
        );
        heart.addStableCredit(newStableCredit);

        emit Setup(assetCode, address(newStableCredit));

        return address(newStableCredit);
    }

    /*
        TODO:
            - mint(bytes32 collateralAssetCode, uint256 mintAmount, string calldata assetCode)
    */
    function mint(
        bytes32 collateralAssetCode,
        uint256 collateralAmount,
        string calldata assetCode
    ) external payable returns(bool) {
        StableCredit stableCredit = heart.getStableCreditById(getStableCreditId(assetCode));

        require(heart.isTrustedPartner(msg.sender), "only trusted partner can mint the stable credit");
        require(address(stableCredit) != address(0x0), "stableCredit not exist");

        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, stableCredit.peggedCurrency()));

        require(heart.getCollateralAsset(collateralAssetCode) == stableCredit.collateral(), "collateralAsset must be the same");
        require(heart.getPriceFeeders().getMedianPrice(linkId) != 0, "median price ref mut not be zero");

        (uint256 mintAmount, uint256 fee) = _calMintStableCredit(stableCredit, linkId, collateralAmount);
        uint256 actualCollateralAmount = _callCollateral(stableCredit, linkId, mintAmount);
        uint256 reserveAmount = collateralAmount.sub(actualCollateralAmount).sub(fee);

        heart.getCollateralAsset(collateralAssetCode).transferFrom(msg.sender, address(heart), fee);
        heart.getCollateralAsset(collateralAssetCode).transferFrom(msg.sender, address(stableCredit), actualCollateralAmount);
        heart.getCollateralAsset(collateralAssetCode).transferFrom(msg.sender, address(this), reserveAmount);

        heart.getCollateralAsset(collateralAssetCode).approve(address(heart.getReserveManager()), reserveAmount);

        heart.getReserveManager().lockReserve(collateralAssetCode, address(this), reserveAmount);

        stableCredit.mint(msg.sender, mintAmount);
        stableCredit.approveCollateral();

        heart.collectFee(fee, collateralAssetCode);

        emit Mint(
            assetCode,
            mintAmount,
            address(stableCredit),
            collateralAssetCode,
            actualCollateralAmount
        );

        return true;
    }

    function redeem(
        address creditOwner,
        uint256 amount,
        string calldata assetCode
    ) external returns(bool) {
        StableCredit stableCredit = heart.getStableCreditById(getStableCreditId(assetCode));

        require(address(stableCredit) != address(0x0), "stableCredit not existed");

        bytes32 linkId = keccak256(abi.encodePacked(stableCredit.collateralAssetCode(), stableCredit.peggedCurrency()));

        _rebalance(creditOwner, assetCode);
        uint256 returnAmount = _callCollateral(stableCredit, linkId, amount);

        stableCredit.redeem(msg.sender, amount, returnAmount);
        stableCredit.approveCollateral();

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
        StableCredit stableCredit = heart.getStableCreditById(getStableCreditId(assetCode));
        bytes32 linkId = keccak256(abi.encodePacked(stableCredit.collateralAssetCode(), stableCredit.peggedCurrency()));

        uint256 collateralAmount = _callCollateral(stableCredit, linkId, stableCredit.totalSupply());

        if (collateralAmount >= stableCredit.collateral().balanceOf(address(stableCredit))) {
            heart.getReserveManager().injectCollateral(stableCredit.collateralAssetCode(), address(stableCredit), collateralAmount.sub(stableCredit.collateral().balanceOf(address(stableCredit))));
        } else {
            heart.getCollateralAsset(stableCredit.collateralAssetCode()).transferFrom(address(stableCredit), address(heart.getReserveManager()), stableCredit.collateral().balanceOf(address(stableCredit)).sub(collateralAmount));
        }

        return true;
    }

    function _calMintStableCredit(StableCredit credit, bytes32 linkId, uint256 collateralAmount) private view returns (uint256, uint256) {
        uint256 fee = collateralAmount.mul(heart.getCreditIssuanceFee()).div(10000);
        uint256 mintAmount = collateralAmount.sub(fee).mul(100).div(heart.getCollateralRatio(credit.collateralAssetCode())).mul(heart.getPriceFeeders().getMedianPrice(linkId)).div(credit.peggedValue());

        return (mintAmount, fee);
    }

    function _calCollateralWithFee(StableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256, uint256) {
        uint256 collateral = _callCollateral(credit, linkId, creditAmount);
        uint256 fee = collateral.mul(heart.getCreditIssuanceFee()).div(10000);

        return (collateral, fee);
    }

    function _callCollateral(StableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256) {
        return creditAmount.mul(credit.peggedValue()).div(heart.getPriceFeeders().getMedianPrice(linkId));
    }

    function collateralOf(address creditOwner, string calldata assetCode) external view returns (uint256, address) {
        return heart.getStableCreditById(getStableCreditId(assetCode)).getCollateralDetail();
    }

    function getStableCreditId(string memory assetCode) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(assetCode));
    }
}
