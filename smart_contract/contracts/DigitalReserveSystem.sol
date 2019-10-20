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

        require(address(collateralAssets[collateralAssetCode]) != address(0x0), "collateralAssetCode has not been whitelisted");
        require(ownedStableCredits[stableCreditId] == false, "trusted partner cannot setup the same asset code");
        require(trustedPartners[msg.sender], "only trusted partner can setup the stable credit");

        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));
        uint256 medPrice = priceFeedersContract.medianPrices(linkId);
        require(medPrice > 0, "collateralAssetCode must has value more than 0");

        StableCredit newStableCredit = new StableCredit(
            assetCode,
            assetCode,
            7,
            peggedValue,
            peggedCurrency
        );

        stableCredits[stableCreditId] = newStableCredit;
        ownedStableCredits[stableCreditId] = true;

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

        (uint256 mintAmount, uint256 fee) = _calMintStableCredit(
                credit.getPeggedValue(), credit.getPeggedCurrency(), collateralAssetCode, collateralAmount);

        require(priceFeedersContract.medianPrices(linkId) != 0, "median price ref mut not be zero");
        require(mintAmount > 0, "mint amount must not be 0");
        require(collateralAssets[collateralAssetCode].balanceOf(msg.sender) > collateralAmount, "not enough collateral balance in Trusted Partner account");
        require(collateralAssets[collateralAssetCode].allowance(msg.sender, address(this)) > collateralAmount, "not enough allowance");

        collateralAssets[collateralAssetCode].transferFrom(msg.sender, address(this), collateralAmount);

        credit.mint(msg.sender, mintAmount);

        _collectFee(fee, collateralAssetCode);

        return true;
    }

    function redeem(
        uint256 amount,
        bytes32 assetCode
    ) external returns(bool) {
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

    function _calMintStableCredit(uint256 peggedValue, bytes32 peggedCurrency, bytes32 collateralAssetCode, uint256 collateralAmount) private view returns (uint256, uint256) {
        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));

        uint256 fee = collateralAmount.mul(creditIssuanceFee).div(10000);
        uint256 collateralAmountAfterFee = collateralAmount.sub(fee);
        uint256 mintAmount = collateralAmountAfterFee.mul(priceFeedersContract.medianPrices(linkId)).div(peggedValue);

        return (mintAmount, fee);
    }

//    function _calCollateralWithFee(uint256 peggedValue, bytes32 peggedCurrency, uint256 amount) private view returns (uint256, uint256) {
//        uint256 collateral = peggedValue.div(prices[peggedCurrency]).mul(amount);
//        uint256 fee = collateral.mul(creditIssuanceFee).div(10000);
//
//        return (collateral, fee);
//    }

    function collateralOf(bytes32 assetCode, address creditOwner) public view returns (uint256, bytes32) {
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
