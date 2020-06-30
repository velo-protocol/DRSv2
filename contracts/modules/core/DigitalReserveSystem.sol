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

    event Rebalance(
        string assetCode,
        bytes32 indexed collateralAssetCode,
        uint256 requiredAmount,
        uint256 presentAmount
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

        // validate stable credit belong to the message sender
        require(msg.sender == stableCredit.creditOwner(), "DigitalReserveSystem.mintFromCollateralAmount: the stable credit does not belong to you");

        (uint256 mintAmount, uint256 actualCollateralAmount, uint256 reserveCollateralAmount, uint256 fee) = _calMintAmountFromCollateral(
            netCollateralAmount,
            heart.getPriceContract(linkId).get(),
            heart.getCreditIssuanceFee(),
            heart.getCollateralRatio(collateralAssetCode),
            stableCredit.peggedValue(),
            100000000000
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

        // validate stable credit belong to the message sender
        require(msg.sender == stableCredit.creditOwner(), "DigitalReserveSystem.mintFromStableCreditAmount: the stable credit does not belong to you");

        (uint256 netCollateralAmount, uint256 actualCollateralAmount, uint256 reserveCollateralAmount, uint256 fee) = _calMintAmountFromStableCredit(
            mintAmount,
            heart.getPriceContract(linkId).get(),
            heart.getCreditIssuanceFee(),
            heart.getCollateralRatio(collateralAssetCode),
            stableCredit.peggedValue(),
            100000000000
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
        require(heart.getPriceContract(linkId).get() > 0, "DigitalReserveSystem._validateAssetCode: valid price not found");

        return (stableCredit, collateralAsset, collateralAssetCode, linkId);
    }

    function _mint(ICollateralAsset collateralAsset, IStableCredit stableCredit, uint256 mintAmount, uint256 fee, uint256 actualCollateralAmount, uint256 reserveCollateralAmount) private returns (bool) {
        bytes32 collateralAssetCode = stableCredit.collateralAssetCode();
        collateralAsset.transferFrom(msg.sender, address(heart), fee);
        collateralAsset.transferFrom(msg.sender, address(stableCredit), actualCollateralAmount.add(reserveCollateralAmount));
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
        require(address(collateralAsset) != address(0), "DigitalReserveSystem.redeem: collateralAssetCode does not exist");

        _rebalance(assetCode);

        uint256 collateralAmount = _calExchangeRate(stableCredit, linkId, stableCreditAmount);

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
    ) external payable returns (bool) {
        return _rebalance(assetCode);
    }

    function getExchange(
        string calldata assetCode
    ) external view returns (string memory, bytes32, uint256) {
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.getExchange: invalid assetCode format");

        (IStableCredit stableCredit, , bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);

        (uint256 priceInCollateralPerAssetUnit) = _calExchangeRate(stableCredit, linkId, 100000000000);

        return (assetCode, collateralAssetCode, priceInCollateralPerAssetUnit);
    }

    function collateralHealthCheck(
        string calldata assetCode
    ) external view returns (address, bytes32, uint256, uint256) {
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.collateralHealthCheck: invalid assetCode format");
        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);
        require(address(collateralAsset) != address(0), "DigitalReserveSystem.collateralHealthCheck: collateralAssetCode does not exist");
        uint256 requiredAmount = _calCollateral(stableCredit, linkId, stableCredit.totalSupply(), heart.getCollateralRatio(collateralAssetCode)).div(100000000000);
        uint256 presentAmount = stableCredit.collateral().balanceOf(address(stableCredit));
        return (address(collateralAsset), collateralAssetCode, requiredAmount, presentAmount);
    }
    function getStableCreditAmount(
        string calldata assetCode
    ) external view returns ( uint256) {
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.collateralHealthCheck: invalid assetCode format");

        (IStableCredit stableCredit,,, ) = _validateAssetCode(assetCode);
        return stableCredit.totalSupply();

    }

    function _rebalance(
        string memory assetCode
    ) private returns (bool) {
        require(bytes(assetCode).length > 0 && bytes(assetCode).length <= 12, "DigitalReserveSystem.rebalance: invalid assetCode format");

        (IStableCredit stableCredit, ICollateralAsset collateralAsset, bytes32 collateralAssetCode, bytes32 linkId) = _validateAssetCode(assetCode);
        require(address(collateralAsset) != address(0), "DigitalReserveSystem.rebalance: collateralAssetCode does not exist");
        uint256 requiredAmount = _calCollateral(stableCredit, linkId, stableCredit.totalSupply(), heart.getCollateralRatio(collateralAssetCode)).div(100000000000);
        uint256 presentAmount = stableCredit.collateral().balanceOf(address(stableCredit));

        if (requiredAmount == presentAmount) {return false;}

        IRM reserveManager = heart.getReserveManager();

        if (requiredAmount > presentAmount) {
            reserveManager.injectCollateral(collateralAssetCode, address(stableCredit), requiredAmount.sub(presentAmount));
        } else {
            stableCredit.transferCollateralToReserve(presentAmount.sub(requiredAmount));
        }

        emit Rebalance(
            assetCode,
            collateralAssetCode,
            requiredAmount,
            presentAmount
        );

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

    function _calCollateral(IStableCredit credit, bytes32 linkId, uint256 creditAmount, uint256 collateralRatio) private view returns (uint256) {
        // collateral = (creditAmount * peggedValue * collateralRatio) / priceInCurrencyPerAssetUnit
        return creditAmount.mul(credit.peggedValue().mul(collateralRatio)).div(heart.getPriceContract(linkId).get());
    }

    function _calExchangeRate(IStableCredit credit, bytes32 linkId, uint256 stableCreditAmount) private view returns (uint256) {
        // priceInCollateral = (collateralRatio * peggedValue * stableCreditAmount) / priceInCurrencyPerAssetUnit
        uint256 priceInCollateralPerAssetUnit = heart.getCollateralRatio(credit.collateralAssetCode()).mul(credit.peggedValue()).mul(stableCreditAmount).div(heart.getPriceContract(linkId).get()).div(100000000000);
        return (priceInCollateralPerAssetUnit);
    }
}


