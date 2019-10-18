pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "./StableCredit.sol";
import "./IDRS.sol";

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

    /*
        PriceFeeder related mapping
        collateralAssetFiat map between collateralAssetCode => currency[]
        collateralAssetFiatCounter map between keccak256(collateralAssetCode, currency) => num of fiats
        priceFeeders map between keccak256(collateralAssetCode, currency) => address[] of price feeders
        priceFeederCounter map between keccak256(collateralAssetCode, currency) => num of price feeders
        prices map between keccak256(collateralAssetCode, currency, address of price feeder) => price
    */
    mapping(bytes32 => bytes32[]) public collateralAssetFiat;
    mapping(bytes32 => address[]) public priceFeeders;
    mapping(bytes32 => uint256) public prices;
    mapping(bytes32 => uint256) public medianPrices;

    constructor() public {}

    function setup(
        string calldata assetCode,
        bytes32 collateralAssetCode,
        uint256 peggedValue,
        bytes32 peggedCurrency
    ) external returns(address) {
        bytes32 stableCreditId = keccak256(abi.encodePacked(msg.sender, assetCode));

        require(address(collateralAssets[collateralAssetCode]) != address(0x0), "assetCode has not been whitelisted");
        require(ownedStableCredits[stableCreditId] == false, "trusted partner cannot setup the same asset code");
        require(trustedPartners[msg.sender], "only trusted partner can setup the stable credit");

        int index = getCollateralAssetFiatIndexByFiat(collateralAssetCode, peggedCurrency);
        require(index != -1, "collateral asset had not been linked with the peggedCurrency");

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

        require(medianPrices[linkId] != 0, "median price ref mut not be zero");
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
        uint256 mintAmount = collateralAmountAfterFee.mul(medianPrices[linkId]).div(peggedValue);

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

    function getCollateralAssetFiat(bytes32 assetCode) public view returns (bytes32[] memory){
        return collateralAssetFiat[assetCode];
    }

    function getCollateralAssetFiatIndexByFiat(bytes32 assetCode, bytes32 currency) public view returns (int) {
        int index = -1;
        for (uint i = 0; i < collateralAssetFiat[assetCode].length; i++) {
            if(collateralAssetFiat[assetCode][i] == currency) {
                index = int(i);
            }
        }
        return index;
    }

    function addCollateralAssetFiat(bytes32 assetCode, bytes32 currency) public onlyWhitelistAdmin {
        int index = getCollateralAssetFiatIndexByFiat(assetCode, currency);

        require(address(collateralAssets[assetCode]) != address(0x0), "assetCode has not been whitelisted");
        require(index == -1, "collateral asset has been linked with this fiat");

        collateralAssetFiat[assetCode].push(currency);
    }

    function removeCollateralAssetFiat(bytes32 assetCode, uint index) public onlyWhitelistAdmin returns (bool) {
        if(index >= collateralAssetFiat[assetCode].length) return false;

        for (uint i = index; i<collateralAssetFiat[assetCode].length-1; i++){
            collateralAssetFiat[assetCode][i] = collateralAssetFiat[assetCode][i+1];
        }
        collateralAssetFiat[assetCode].length--;

        return true;
    }

    function setCreditIssuanceFee(uint256 newFee) public onlyWhitelistAdmin {
        creditIssuanceFee = newFee;
    }

    function setTrustedPartner(address addr) public onlyWhitelistAdmin {
        trustedPartners[addr] = true;
    }

    function addPriceFeeder(bytes32 collateralAssetCode, bytes32 currency, address priceFeederAddr) public onlyWhitelistAdmin {
        bytes32 pfId = keccak256(abi.encodePacked(collateralAssetCode, currency));

        int linkIndex = getCollateralAssetFiatIndexByFiat(collateralAssetCode, currency);
        int pfIndex = getPriceFeederIndexByPriceFeederAddress(pfId, priceFeederAddr);

        require(address(collateralAssets[collateralAssetCode]) != address(0x0), "assetCode has not been whitelisted");
        require(linkIndex != -1, "collateral asset has not been linked with this fiat");
        require(pfIndex == -1, "this price feeder has been added to the list");

        priceFeeders[pfId].push(priceFeederAddr);
    }

    function getPriceFeeders(bytes32 pfId) public view returns (address[] memory) {
        return priceFeeders[pfId];
    }

    function removePriceFeeder(bytes32 pfId, uint index) public onlyWhitelistAdmin returns (bool) {
        if(index >= priceFeeders[pfId].length) return false;

        for (uint i = index; i<priceFeeders[pfId].length-1; i++){
            priceFeeders[pfId][i] = priceFeeders[pfId][i+1];
        }
        priceFeeders[pfId].length--;

        return true;
    }

    function getPriceFeederIndexByPriceFeederAddress(bytes32 id, address priceFeeder) public view returns (int) {
        int index = -1;
        for (uint i = 0; i < priceFeeders[id].length; i++) {
            if(priceFeeders[id][i] == priceFeeder) {
                index = int(i);
            }
        }
        return index;
    }

    function setPrice(bytes32 collateralAssetCode, bytes32 currency, uint256 price) public onlyWhitelistAdmin {
        bytes32 linkId = keccak256(abi.encodePacked(collateralAssetCode, currency));
//        int pfIndex = getPriceFeederIndexByPriceFeederAddress(linkId, msg.sender);
//        int linkIndex = getCollateralAssetFiatIndexByFiat(collateralAssetCode, currency);

//        require(address(collateralAssets[collateralAssetCode]) != address(0x0), "assetCode has not been whitelisted");
//        require(linkIndex != -1, "collateral asset has not been linked with this fiat");
//        require(pfIndex != -1, "msg.sender did not have an authority to set the price");

//        bytes32 priceId = keccak256(abi.encodePacked(collateralAssetCode, currency, msg.sender));
//        prices[priceId] = price;
        medianPrices[linkId] = price;
    }
}
