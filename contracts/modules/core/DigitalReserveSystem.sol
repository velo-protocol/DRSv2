pragma solidity ^0.5.0;

import "../interfaces/IHeart.sol";
import "../interfaces/IDRS.sol";
import "../interfaces/IRM.sol";
import "../interfaces/IStableCredit.sol";
import "../interfaces/ICollateralAsset.sol";

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

    event Redeem(
        string assetCode,
        uint256 stableCreditAmount,
        address indexed collateralAssetAddress,
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

    function mintFromCollateralAmount(
        uint256 netCollateralAmount,
        string calldata assetCode
    ) external onlyTrustedPartner payable returns (bool) {
        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);

        (uint256 mintAmount, uint256 actualCollateralAmount, uint256 reserveCollateralAmount, uint256 fee) = _calMintAmountFromCollateral(
            netCollateralAmount,
            heart.getPriceFeeders().getMedianPrice(linkId),
            heart.getCreditIssuanceFee(),
            heart.getCollateralRatio(collateralAssetCode),
            stableCredit.peggedValue(),
            10000000
        );

        _mint(collateralAsset, stableCredit, mintAmount, fee, actualCollateralAmount, reserveCollateralAmount);

        // redeclare collateralAmount, this a workaround for StackTooDeep error
        uint256 _netCollateralAmount = netCollateralAmount;
        emit Mint(
            assetCode,
            mintAmount,
            address(stableCredit),
            collateralAssetCode,
            _netCollateralAmount
        );

        return true;
    }

    function mintFromStableCreditAmount(
        uint256 mintAmount,
        string calldata assetCode
    ) external onlyTrustedPartner payable returns (bool) {
        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);

        (uint256 netCollateralAmount, uint256 actualCollateralAmount, uint256 reserveCollateralAmount, uint256 fee) = _calMintAmountFromStableCredit(
            mintAmount,
            heart.getPriceFeeders().getMedianPrice(linkId),
            heart.getCreditIssuanceFee(),
            heart.getCollateralRatio(collateralAssetCode),
            stableCredit.peggedValue(),
            10000000
        );

        _mint(collateralAsset, stableCredit, mintAmount, fee, actualCollateralAmount, reserveCollateralAmount);
        uint256 _mintAmount = mintAmount;
        emit Mint(
            assetCode,
            _mintAmount,
            address(stableCredit),
            collateralAssetCode,
            netCollateralAmount
        );

        return true;
    }

    function _validateAssetCode(string memory assetCode) private view returns (IStableCredit, ICollateralAsset, bytes32, bytes32) {
        IStableCredit stableCredit = heart.getStableCreditById(Hasher.stableCreditId(assetCode));
        require(address(stableCredit) != address(0), "DigitalReserveSystem._validateAssetCode: stableCredit not exist");

        bytes32 collateralAssetCode = stableCredit.collateralAssetCode();
        ICollateralAsset collateralAsset = heart.getCollateralAsset(collateralAssetCode);
        require(collateralAsset == stableCredit.collateral(), "DigitalReserveSystem._validateAssetCode: collateralAsset must be the same");

        bytes32 linkId = Hasher.linkId(collateralAssetCode, stableCredit.peggedCurrency());
        require(heart.getPriceFeeders().getMedianPrice(linkId) > 0, "DigitalReserveSystem._validateAssetCode: price of link must have value more than 0");

        return (stableCredit, collateralAsset, collateralAssetCode, linkId);
    }

    function _mint(ICollateralAsset collateralAsset, IStableCredit stableCredit, uint256 mintAmount, uint256 fee, uint256 actualCollateralAmount, uint256 reserveCollateralAmount) private returns (bool) {
        bytes32 collateralAssetCode = stableCredit.collateralAssetCode();
        collateralAsset.transferFrom(msg.sender, address(heart), fee);
        collateralAsset.transferFrom(msg.sender, address(stableCredit), actualCollateralAmount);
        collateralAsset.transferFrom(msg.sender, address(this), reserveCollateralAmount);

        IRM resManager = heart.getReserveManager();
        collateralAsset.approve(address(resManager), reserveCollateralAmount);
        resManager.lockReserve(collateralAssetCode, address(this), reserveCollateralAmount);

        stableCredit.mint(msg.sender, mintAmount);
        stableCredit.approveCollateral();

        heart.collectFee(fee, collateralAssetCode);
        return true;
    }

    function redeem(
        uint256 stableCreditAmount,
        string calldata assetCode
    ) external returns (bool) {
        require(stableCreditAmount > 0, "DigitalReserveSystem.redeem: redeem amount must be greater than 0");
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.redeem: invalid assetCode format");

        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);
        require(address(collateralAsset) != address(0), "DigitalReserveSystem.collateralHealthCheck: collateralAssetCode does not exist");

        _rebalance(assetCode);

        uint256 priceInCollateralPerAssetUnit = _calExchangeRate(stableCredit, linkId);

        // collateralAmount = stableCreditAmount * priceInCollateralPerAssetUnit
        uint256 collateralAmount = stableCreditAmount.mul(priceInCollateralPerAssetUnit).div(10000000);

        stableCredit.redeem(msg.sender, stableCreditAmount, collateralAmount);
        stableCredit.approveCollateral();

        emit Redeem(
            assetCode,
            stableCreditAmount,
            address(collateralAsset),
            collateralAssetCode,
            collateralAmount
        );

        return true;
    }

    function rebalance(
        string calldata assetCode
    ) external returns (bool) {
        return _rebalance(assetCode);
    }

    function getExchange(
        string calldata assetCode
    ) external view returns (string memory, bytes32, uint256) {
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.getExchange: invalid assetCode format");

        (IStableCredit stableCredit, , bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);

        (uint256 priceInCollateralPerAssetUnit) = _calExchangeRate(stableCredit, linkId);

        return (assetCode, collateralAssetCode, priceInCollateralPerAssetUnit);
    }

    function collateralHealthCheck(
        string calldata assetCode
    ) external view returns (bytes32, uint256, uint256) {
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.collateralHealthCheck: invalid assetCode format");

        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);
        require(address(collateralAsset) != address(0), "DigitalReserveSystem.collateralHealthCheck: collateralAssetCode does not exist");

        uint256 requiredAmount = _calCollateral(stableCredit, linkId, stableCredit.totalSupply());
        uint256 presentAmount = stableCredit.collateral().balanceOf(address(stableCredit));

        return (collateralAssetCode, requiredAmount, presentAmount);
    }

    function _rebalance(
        string memory assetCode
    ) private returns (bool) {

        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);

        uint256 requireCollateralAmount = _calCollateral(stableCredit, linkId, stableCredit.totalSupply());

        IRM reserveManager = heart.getReserveManager();

        uint256 presentAmount = collateralAsset.balanceOf(address(stableCredit));

        if (requireCollateralAmount >= presentAmount) {
            reserveManager.injectCollateral(collateralAssetCode, address(stableCredit), requireCollateralAmount.sub(presentAmount));
        } else {
            heart.getCollateralAsset(collateralAssetCode).transferFrom(address(stableCredit), address(reserveManager), presentAmount.sub(requireCollateralAmount));
        }

        return true;
    }


    function _calMintAmountFromCollateral(
        uint256 netCollateralAmount,
        uint256 price,
        uint256 issuanceFee,
        uint256 collateralRatio,
        uint256 peggedValue,
        uint256 divider
    ) private pure returns (uint256, uint256, uint256, uint256) {
        // fee = netCollateralAmount * (issuanceFee / divider )
        uint256 fee = netCollateralAmount.mul(issuanceFee).div(divider);

        // collateralAmount = netCollateralAmount - fee
        uint256 collateralAmount = netCollateralAmount.sub(fee);

        // mintAmount = (collateralAmount * priceInCurrencyPerCollateralUnit) / (collateralRatio * peggedValue)
        uint256 mintAmount = collateralAmount.mul(price).mul(divider);
        mintAmount = mintAmount.div(collateralRatio.mul(peggedValue));

        // actualCollateralAmount = collateralAmount / collateralRatio
        uint actualCollateralAmount = collateralAmount.mul(divider).div(collateralRatio);

        // reserveCollateralAmount = collateralAmount - actualCollateralAmount
        uint reserveCollateralAmount = collateralAmount.sub(actualCollateralAmount);

        return (mintAmount, actualCollateralAmount, reserveCollateralAmount, fee);
    }

    function _calMintAmountFromStableCredit(
        uint256 mintAmount,
        uint256 price,
        uint256 issuanceFee,
        uint256 collateralRatio,
        uint256 peggedValue,
        uint256 divider
    ) private pure returns (uint256, uint256, uint256, uint256) {
        // collateralAmount = (mintAmount * collateralRatio * peggedValue) / priceInCurrencyPerCollateralUnit
        uint256 collateralAmount = mintAmount.mul(collateralRatio).mul(peggedValue);
        collateralAmount = collateralAmount.div(price).div(divider);

        // fee = (collateralAmount * issuanceFee) / (divider - issuanceFee)
        uint256 fee = collateralAmount.mul(issuanceFee).div(divider.sub(issuanceFee));

        // netCollateralAmount = collateralAmount + fee
        uint256 netCollateralAmount = collateralAmount.add(fee);

        // actualCollateralAmount = collateralAmount / collateralRatio
        uint actualCollateralAmount = collateralAmount.mul(divider).div(collateralRatio);

        // reserveCollateralAmount = collateralAmount - actualCollateralAmount
        uint reserveCollateralAmount = collateralAmount.sub(actualCollateralAmount);

        return (netCollateralAmount, actualCollateralAmount, reserveCollateralAmount, fee);
    }

    function _calCollateral(IStableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256) {
        return creditAmount.mul(credit.peggedValue()).div(heart.getPriceFeeders().getMedianPrice(linkId));
    }

    function _calExchangeRate(IStableCredit credit, bytes32 linkId) private view returns (uint256) {
        // priceInCollateralPerAssetUnit = (collateralRatio * peggedValue) / priceInCurrencyPerAssetUnit
        uint256 priceInCollateralPerAssetUnit = heart.getCollateralRatio(credit.collateralAssetCode()).mul(credit.peggedValue()).div(heart.getPriceFeeders().getMedianPrice(linkId));
        return (priceInCollateralPerAssetUnit);
    }

    function collateralOf(address creditOwner, string calldata assetCode) external view returns (uint256, address) {
        return heart.getStableCreditById(getStableCreditId(assetCode)).getCollateralDetail();
    }

    function getStableCreditId(string memory assetCode) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(assetCode));
    }
}
