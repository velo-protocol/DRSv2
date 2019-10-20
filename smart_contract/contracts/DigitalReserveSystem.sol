pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "./StableCredit.sol";
import "./IDRS.sol";
import "./PriceFeeders.sol";

contract DigitalReserveSystem is IDRS, WhitelistAdminRole {
    using SafeMath for uint256;

    mapping(bytes32 => IERC20) public collateralAssets;

    uint256 public creditIssuanceFee;
    mapping(bytes32 => uint256) public collectedFee;

    mapping(address => bool) public trustedPartners;

    /*
        StableCredit related mapping
        ownedStableCredits map between keccak256(stableCreditOwnerAddress, stableCreditCode) => bool
        stableCredits map between keccak256(stableCreditOwnerAddress, stableCreditCode) => StableCredit
    */
    mapping(bytes32 => bool) public ownedStableCredits;
    mapping(bytes32 => StableCredit) public stableCredits;

    PriceFeeders priceFeedersContract;

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

    constructor(address priceFeedersAddress) public {
        priceFeedersContract = PriceFeeders(priceFeedersAddress);
    }

    function setup(
        string calldata assetCode,
        bytes32 collateralAssetCode,
        uint256 peggedValue,
        bytes32 peggedCurrency
    ) external returns(address) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(msg.sender, assetCode));
        address collateralAddress = address(collateralAssets[collateralAssetCode]);

        require(collateralAddress != address(0x0), "collateralAssetCode has not been whitelisted");
        require(ownedStableCredits[stableCreditId] == false, "trusted partner cannot setup the same asset code");
        require(trustedPartners[msg.sender], "only trusted partner can setup the stable credit");

        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));

        require(priceFeedersContract.medianPrices(linkId) > 0, "collateralAssetCode must has value more than 0");

        StableCredit newStableCredit = new StableCredit(
            msg.sender,
            assetCode,
            assetCode,
            peggedValue,
            peggedCurrency,
            collateralAddress,
            collateralAssetCode
        );

        stableCredits[stableCreditId] = newStableCredit;
        ownedStableCredits[stableCreditId] = true;

        emit Setup(assetCode, address(newStableCredit));

        return address(newStableCredit);
    }

    function mint(
        bytes32 collateralAssetCode,
        uint256 collateralAmount,
        string calldata assetCode
    ) external payable returns(bool) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(msg.sender, assetCode));

        require(trustedPartners[msg.sender] == true, "only trusted partner can mint the stable credit");
        require(ownedStableCredits[stableCreditId] == true, "msg.sender must be the owner of the stable credit");

        StableCredit credit = stableCredits[stableCreditId];
        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, credit.getPeggedCurrency()));

        (uint256 mintAmount, uint256 fee) = _calMintStableCredit(credit, linkId, collateralAmount);

        require(collateralAssets[collateralAssetCode] == credit.collateral(), "collateralAsset must be the same");
        require(priceFeedersContract.medianPrices(linkId) != 0, "median price ref mut not be zero");
        require(mintAmount > 0, "mint amount must not be 0");
        require(collateralAssets[collateralAssetCode].balanceOf(msg.sender) > collateralAmount, "not enough collateral balance in Trusted Partner account");
        require(collateralAssets[collateralAssetCode].allowance(msg.sender, address(this)) > collateralAmount, "not enough allowance");

        collateralAssets[collateralAssetCode].transferFrom(msg.sender, address(this), fee);
        collateralAssets[collateralAssetCode].transferFrom(msg.sender, address(credit), collateralAmount.sub(fee));

        credit.mint(msg.sender, mintAmount);
        credit.approveCollateral();

        _collectFee(fee, collateralAssetCode);

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
        bytes32 linkId = keccak256(abi.encodePacked(credit.getCollateralAssetCode(), credit.getPeggedCurrency()));

        uint256 returnAmount = _callCollateral(credit, linkId, amount);

        credit.redeem(msg.sender, amount, returnAmount);
        credit.approveCollateral();

        return true;
    }

    function rebalance(
        address creditOwner,
        string calldata assetCode
    ) external returns(bool) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(creditOwner, assetCode));

        StableCredit credit = stableCredits[stableCreditId];
        bytes32 linkId = keccak256(abi.encodePacked(credit.getCollateralAssetCode(), credit.getPeggedCurrency()));

        uint256 collateralAmount = _callCollateral(credit, linkId, credit.totalSupply());

        if (collateralAmount >= credit.collateral().balanceOf(address(credit))) {
            // TODO: Inject collateral from reserve manager
        } else {
            // TODO: Transfer collateral from StableCredit contract to Reserve Manager
//            collateralAssets[credit.collateralAssetCode].transferFrom(address(credit), address(reserveManager), amount)
        }

        return true;
    }

//    function checkRequiredCollateralWithFee(
//        uint256 peggedValue,
//        bytes32 peggedCurrency,
//        uint256 amount
//    ) public view returns (uint256, uint256) {
//        (uint256 requiredCollateral, uint256 fee) = _calCollateralWithFee(peggedValue, peggedCurrency, amount);
//
//        return (requiredCollateral, fee);
//    }

    function _collectFee(uint256 fee, bytes32 collateralAssetCode) private {
        collectedFee[collateralAssetCode] = collectedFee[collateralAssetCode].add(fee);
    }

    function _calMintStableCredit(StableCredit credit, bytes32 linkId, uint256 collateralAmount) private view returns (uint256, uint256) {
        uint256 fee = collateralAmount.mul(creditIssuanceFee).div(10000);
        uint256 collateralAmountAfterFee = collateralAmount.sub(fee);
        uint256 mintAmount = collateralAmountAfterFee.mul(priceFeedersContract.medianPrices(linkId)).div(credit.getPeggedValue());

        return (mintAmount, fee);
    }

    function _callCollateral(StableCredit credit, bytes32 linkId, uint256 creditAmount) private view returns (uint256) {
        uint256 collateralAmount = creditAmount.mul(credit.getPeggedValue()).div(priceFeedersContract.medianPrices(linkId));

        return collateralAmount;
    }

//    function _calCollateralWithFee(uint256 peggedValue, bytes32 peggedCurrency, uint256 amount) private view returns (uint256, uint256) {
//        uint256 collateral = peggedValue.div(prices[peggedCurrency]).mul(amount);
//        uint256 fee = collateral.mul(creditIssuanceFee).div(10000);
//
//        return (collateral, fee);
//    }

    function collateralOf(bytes32 assetCode, address creditOwner) public view returns (uint256, address) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(creditOwner, assetCode));
        return stableCredits[stableCreditId].getCollateralDetail();
    }

    function getCollectedFee(bytes32 collateralAssetCode) public view returns (uint256) {
        return collectedFee[collateralAssetCode];
    }

    function withdrawFee(uint256 amount) public onlyWhitelistAdmin {
//        require(amount <= collectedFee, "amount must <= to collectedFee");
//        veloToken.transfer(msg.sender, amount);
//        collectedFee = collectedFee.sub(amount);
    }

    function setCollateralAsset(bytes32 assetCode, address addr) public onlyWhitelistAdmin {
        collateralAssets[assetCode] = IERC20(addr);
    }

    function setCreditIssuanceFee(uint256 newFee) public onlyWhitelistAdmin {
        creditIssuanceFee = newFee;
    }

    function setTrustedPartner(address addr) public onlyWhitelistAdmin {
        trustedPartners[addr] = true;
    }
}
