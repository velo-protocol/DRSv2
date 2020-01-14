pragma solidity ^0.5.0;

import "../interfaces/IHeart.sol";
import "../interfaces/IDRS.sol";
import "../interfaces/IStableCredit.sol";
import "../interfaces/ICollateralAsset.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract DigitalReserveSystem is IDRS {
    using SafeMath for uint256;
    IHeart public heart;

    event Setup(
        string assetCode,
        bytes32 peggedCurrency,
        uint256 peggedValue,
        bytes32 indexed collateralAssetCode,
        address assetAddress
    );

    event Mint(
        string assetCode,
        uint256 mintAmount,
        address indexed assetAddress,
        bytes32 indexed collateralAssetCode,
        uint256 collateralAmount
    );

    modifier onlyTrustedPartner() {
        require(heart.isTrustedPartner(msg.sender), "DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner");
        _;
    }

    constructor(address heartAddr) public {
        heart = IHeart(heartAddr);
    }

    function setup(
        bytes32 collateralAssetCode,
        bytes32 peggedCurrency,
        string calldata assetCode,
        uint256 peggedValue
    ) external onlyTrustedPartner returns (string memory, address) {
        // validate asset code
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.setup: invalid assetCode format");
        bytes32 stableCreditId = Hasher.stableCreditId(assetCode);
        IStableCredit stableCredit = heart.getStableCreditById(stableCreditId);
        require(address(stableCredit) == address(0), "DigitalReserveSystem.setup: assetCode has already been used");

        // validate collateralAssetCode
        ICollateralAsset collateralAsset = heart.getCollateralAsset(collateralAssetCode);
        require(address(collateralAsset) != address(0), "DigitalReserveSystem.setup: collateralAssetCode does not exist");

        // validate collateralAssetCode, peggedCurrency
        bytes32 linkId = Hasher.linkId(collateralAssetCode, peggedCurrency);
        require(heart.isLinkAllowed(linkId), "DigitalReserveSystem.setup: collateralAssetCode - peggedCurrency pair does not exist");

        // validate link price
        require(heart.getPriceFeeders().getMedianPrice(linkId) > 0, "DigitalReserveSystem.setup: price of link must have value more than 0");

        StableCredit newStableCredit = new StableCredit(
            peggedCurrency,
            msg.sender,
            collateralAssetCode,
            address(collateralAsset),
            assetCode,
            peggedValue,
            address(heart)
        );
        heart.addStableCredit(IStableCredit(address(newStableCredit)));
        emit Setup(
            assetCode,
            peggedCurrency,
            peggedValue,
            collateralAssetCode,
            address(newStableCredit)
        );

        return (assetCode, address(newStableCredit));
    }

    function mintFromCollateral(
        uint256 collateralAmount,
        string calldata assetCode
    ) external onlyTrustedPartner payable returns (bool) {

        IStableCredit stableCredit = heart.getStableCreditById(getStableCreditId(assetCode));
        require(address(stableCredit) != address(0), "DigitalReserveSystem.mintFromCollateral: stableCredit not exist");
        bytes32 collateralAssetCode = stableCredit.collateralAssetCode();

        bytes32 linkId = Hasher.linkId(collateralAssetCode, stableCredit.peggedCurrency());

        require(heart.getCollateralAsset(collateralAssetCode) == stableCredit.collateral(), "DigitalReserveSystem.mintFromCollateral: collateralAsset must be the same");
        require(heart.getPriceFeeders().getMedianPrice(linkId) > 0, "DigitalReserveSystem.mintFromCollateral: price of link must have value more than 0");

        (uint256 mintAmount, uint256 fee) = _calMintFromCollateral(stableCredit, linkId, collateralAmount);
        uint256 actualCollateralAmount = _callCollateral(stableCredit, linkId, mintAmount);
        uint256 reserveAmount = collateralAmount.sub(actualCollateralAmount).sub(fee);

        ICollateralAsset collateralAsset = heart.getCollateralAsset(collateralAssetCode);
        collateralAsset.transferFrom(msg.sender, address(heart), fee);
        collateralAsset.transferFrom(msg.sender, address(stableCredit), actualCollateralAmount);
        collateralAsset.transferFrom(msg.sender, address(this), reserveAmount);

        collateralAsset.approve(address(heart.getReserveManager()), reserveAmount);

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

    function mintStableCredit(
        uint256 stableCreditAmount,
        string calldata assetCode
    ) external onlyTrustedPartner payable returns (bool) {
        bytes32 stableCreditId = Hasher.stableCreditId(assetCode);
        IStableCredit stableCredit = heart.getStableCreditById(stableCreditId);
        require(address(stableCredit) != address(0), "DigitalReserveSystem.mintStableCredit: stableCredit not exist");

        bytes32 collateralAssetCode = stableCredit.collateralAssetCode();
        bytes32 linkId = Hasher.linkId(collateralAssetCode, stableCredit.peggedCurrency());

        require(heart.getCollateralAsset(collateralAssetCode) == stableCredit.collateral(), "DigitalReserveSystem.mintStableCredit: collateralAsset must be the same");
        require(heart.getPriceFeeders().getMedianPrice(linkId) > 0, "DigitalReserveSystem.mintStableCredit: price of link must have value more than 0");

        (uint256 mintAmount, uint256 fee) = _calMintStableCredit(stableCredit, linkId, stableCreditAmount);
        uint256 actualCollateralAmount = _callCollateral(stableCredit, linkId, mintAmount);
        uint256 reserveAmount = stableCreditAmount.add(actualCollateralAmount).add(fee);

        ICollateralAsset collateralAsset = heart.getCollateralAsset(collateralAssetCode);
        collateralAsset.transferFrom(msg.sender, address(heart), fee);
        collateralAsset.transferFrom(msg.sender, address(stableCredit), actualCollateralAmount);
        collateralAsset.transferFrom(msg.sender, address(this), reserveAmount);

        collateralAsset.approve(address(heart.getReserveManager()), reserveAmount);

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
    ) external returns (bool) {
        IStableCredit stableCredit = heart.getStableCreditById(getStableCreditId(assetCode));

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
    ) external returns (bool) {
        return _rebalance(creditOwner, assetCode);
    }

    function _rebalance(
        address creditOwner,
        string memory assetCode
    ) private returns (bool) {
        IStableCredit stableCredit = heart.getStableCreditById(getStableCreditId(assetCode));
        bytes32 linkId = keccak256(abi.encodePacked(stableCredit.collateralAssetCode(), stableCredit.peggedCurrency()));

        uint256 collateralAmount = _callCollateral(stableCredit, linkId, stableCredit.totalSupply());

        if (collateralAmount >= stableCredit.collateral().balanceOf(address(stableCredit))) {
            heart.getReserveManager().injectCollateral(stableCredit.collateralAssetCode(), address(stableCredit), collateralAmount.sub(stableCredit.collateral().balanceOf(address(stableCredit))));
        } else {
            heart.getCollateralAsset(stableCredit.collateralAssetCode()).transferFrom(address(stableCredit), address(heart.getReserveManager()), stableCredit.collateral().balanceOf(address(stableCredit)).sub(collateralAmount));
        }

        return true;
    }

    function _calMintFromCollateral(IStableCredit credit, bytes32 linkId, uint256 collateralAmount) private view returns (uint256, uint256) {
        uint256 fee = collateralAmount.mul(heart.getCreditIssuanceFee()).div(10000000);
        uint256 mintAmount = collateralAmount.sub(fee).mul(heart.getPriceFeeders().getMedianPrice(linkId)).mul(10000000).mul(10000000).
        div(heart.getCollateralRatio(credit.collateralAssetCode()).mul(credit.peggedValue()).mul(10000000));
        return (mintAmount, fee);
    }

    function _calMintStableCredit(IStableCredit credit, bytes32 linkId, uint256 stableCreditAmount) private view returns (uint256, uint256) {
        uint256 mintAmount = stableCreditAmount.mul(heart.getCollateralRatio(credit.collateralAssetCode())).mul(credit.peggedValue()).div(heart.getPriceFeeders().getMedianPrice(linkId)).mul(10000000);
        uint256 fee = mintAmount.mul(heart.getCreditIssuanceFee()).div(10000000);

        return (mintAmount, fee);
    }

    function _calCollateralWithFee(IStableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256, uint256) {
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
