pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./IPF.sol";
import "./IGOV.sol";
import "./IRM.sol";

contract Governance is WhitelistAdminRole, IGOV {
    using SafeMath for uint256;

    address public drsAddr;

    /*
        collateralRatios is 1/100
        collateralRatios 1000 = require collateral 10x
        collateralRatios 125 = require collateral 1.25x
        collateralRatios 100 = require collateral 1x
    */
    mapping(bytes32 => IERC20) public collateralAssets;
    mapping(bytes32 => uint) public collateralRatios;

    /*
        creditIssuanceFee is  1/10,000
        creditIssuanceFee 185 = 1.85% = 0.0185x
        creditIssuanceFee 100 = 1% = 0.01x
    */
    uint256 public creditIssuanceFee;
    mapping(bytes32 => uint256) public collectedFee;

    mapping(address => bool) public trustedPartners;

    IPF public priceFeeders;
    IRM public reserveManager;

    constructor() public { }

    function setReserveManager(address newReserveManager) external onlyWhitelistAdmin {
        reserveManager = IRM(newReserveManager);
    }

    function getReserveManager() external view returns (IRM) {
        return reserveManager;
    }

    function setDrsAddress(address newDrsAddress) external onlyWhitelistAdmin {
        drsAddr = newDrsAddress;
    }

    function setCollateralAsset(bytes32 assetCode, address addr, uint ratio) external onlyWhitelistAdmin {
        collateralAssets[assetCode] = IERC20(addr);
        collateralRatios[assetCode] = ratio;
    }

    function getCollateralAsset(bytes32 assetCode) external view returns (IERC20) {
        return collateralAssets[assetCode];
    }

    function setCollateralRatio(bytes32 assetCode, uint ratio) external onlyWhitelistAdmin {
        require(address(collateralAssets[assetCode]) != address(0x0), "assetCode has not been added");
        collateralRatios[assetCode] = ratio;
    }

    function getCollateralRatio(bytes32 assetCode) external view returns (uint) {
        return collateralRatios[assetCode];
    }

    function setCreditIssuanceFee(uint256 newFee) external onlyWhitelistAdmin {
        creditIssuanceFee = newFee;
    }

    function getCreditIssuanceFee() external view returns (uint256) {
        return creditIssuanceFee;
    }

    function setTrustedPartner(address addr) external onlyWhitelistAdmin {
        trustedPartners[addr] = true;
    }

    function isTrustedPartner(address addr) external view returns (bool) {
        return trustedPartners[addr];
    }

    function setPriceFeeders(address newPriceFeeders) external onlyWhitelistAdmin {
        priceFeeders = IPF(newPriceFeeders);
    }

    function getPriceFeeders() external view returns (IPF) {
        return priceFeeders;
    }

    function collectFee(uint256 fee, bytes32 collateralAssetCode) external {
        require(msg.sender == drsAddr, "only DRSSC can update the collected fee");
        collectedFee[collateralAssetCode] = collectedFee[collateralAssetCode].add(fee);
    }

    function getCollectedFee(bytes32 collateralAssetCode) external view returns (uint256) {
        return collectedFee[collateralAssetCode];
    }

    function withdrawFee(uint256 amount) public onlyWhitelistAdmin {
        //        require(amount <= collectedFee, "amount must <= to collectedFee");
        //        veloToken.transfer(msg.sender, amount);
        //        collectedFee = collectedFee.sub(amount);
    }
}
