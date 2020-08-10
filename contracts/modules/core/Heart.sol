pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IPrice.sol";
import "../interfaces/IHeart.sol";
import "../interfaces/IRM.sol";
import "../interfaces/IStableCredit.sol";
import "../interfaces/ICollateralAsset.sol";
import "../book-room/LL.sol";
import "../book-room/Hasher.sol";

contract Heart is IHeart {
    using SafeMath for uint256;

    address public drsAddr;
    mapping(bytes32 => IPrice) public prices;
    IRM public reserveManager;

    /*
        collateralRatios is 1/100
        collateralRatios 1000 = require collateral 10x
        collateralRatios 125 = require collateral 1.25x
        collateralRatios 100 = require collateral 1x

        ex. You need 1.6x of Eth as collateral to issue DAI

        collateralAssetCode => ratio, ERC20 token
    */
    mapping(bytes32 => ICollateralAsset) public collateralAssets;
    mapping(bytes32 => uint) public collateralRatios;

    /*
        StableCredit related mapping
        stableCredits map between keccak256(stableCreditOwnerAddress, stableCreditCode) => StableCredit
    */
    mapping(bytes32 => IStableCredit) public stableCredits;
    LL.List public stableCreditsLL;
    using LL for LL.List;

    /*
        creditIssuanceFee is  1/10,000
        creditIssuanceFee 185 = 1.85% = 0.0185x
        creditIssuanceFee 100 = 1% = 0.01x
    */
    uint256 public creditIssuanceFee;
    /*
        collateralAssetCode => collectedFee
    */
    mapping(bytes32 => uint256) public collectedFee;

    /*
        trusted partner address => bool
        governor address => bool
    */
    mapping(address => bool) public trustedPartners;
    mapping(address => bool) public governor;

    /*
        Allowed peggedCurrency - collateralAsset pair
        linkId => bool
    */
    mapping(bytes32 => bool) allowedLinks;
    bool public active;


    modifier onlyGovernor() {
        require(isGovernor(msg.sender), "Heart.onlyGovernor: The message sender is not found or does not have sufficient permission");
        _;
    }

    modifier onlyTrustedPartner() {
        require(isTrustedPartner(msg.sender), "Heart.onlyTrustedPartner: The message sender is not found or does not have sufficient permission");
        _;
    }

    modifier onlyDRS() {
        require(msg.sender == drsAddr, "Heart.onlyDRS: caller must be DRS");
        _;
    }

    /*
        reserveFreeze collateralAssetCode => seconds
    */
    mapping(bytes32 => uint256) public reserveFreeze;


    constructor() public {
        governor[msg.sender] = true;
        stableCreditsLL.init();
        active=false;
    }

    function setReserveManager(address newReserveManager) external onlyGovernor {
        reserveManager = IRM(newReserveManager);
    }

    function getReserveManager() external view returns (IRM) {
        return reserveManager;
    }

    function setReserveFreeze(bytes32 assetCode, uint256 newSeconds) external onlyGovernor{
        reserveFreeze[assetCode] = newSeconds;
    }

    function getReserveFreeze(bytes32 assetCode) external view returns (uint256) {
        return reserveFreeze[assetCode];
    }

    function setDrsAddress(address newDrsAddress) external onlyGovernor {
      if(active==false){
          drsAddr = newDrsAddress;
          active=true;
      }

    }

    function getDrsAddress() external view returns (address) {
        return drsAddr;
    }

    function setCollateralAsset(bytes32 assetCode, address addr, uint ratio) external onlyGovernor {
        collateralAssets[assetCode] = ICollateralAsset(addr);
        collateralRatios[assetCode] = ratio;
    }

    function getCollateralAsset(bytes32 assetCode) external view returns (ICollateralAsset) {
        return collateralAssets[assetCode];
    }

    function setCollateralRatio(bytes32 assetCode, uint ratio) external onlyGovernor {
        require(address(collateralAssets[assetCode]) != address(0x0), "assetCode has not been added");
        collateralRatios[assetCode] = ratio;
    }

    function getCollateralRatio(bytes32 assetCode) external view returns (uint) {
        return collateralRatios[assetCode];
    }

    function setCreditIssuanceFee(uint256 newFee) external onlyGovernor {
        creditIssuanceFee = newFee;
    }

    function getCreditIssuanceFee() external view returns (uint256) {
        return creditIssuanceFee;
    }

    function setTrustedPartner(address addr) external onlyGovernor {
        trustedPartners[addr] = true;
    }

    function setGovernor(address addr) external onlyGovernor {
        governor[addr] = true;
    }

    function isTrustedPartner(address addr) public view returns (bool) {
        return trustedPartners[addr];
    }

    function isGovernor(address addr) public view returns (bool) {
        return governor[addr];
    }

    function addPrice(bytes32 linkId, IPrice newPrice) external onlyGovernor {
        require(address(newPrice) != address(0), "newPrice address must not be 0");
        require(address(prices[linkId]) == address(0), "Price has already existed");
        prices[linkId] = IPrice(newPrice);
    }

    function getPriceContract(bytes32 linkId) external view returns (IPrice) {
        return prices[linkId];
    }

    function collectFee(uint256 fee, bytes32 collateralAssetCode) external {
        require(msg.sender == drsAddr, "only DRSSC can update the collected fee");
        collectedFee[collateralAssetCode] = collectedFee[collateralAssetCode].add(fee);
    }

    function getCollectedFee(bytes32 collateralAssetCode) external view returns (uint256) {
        return collectedFee[collateralAssetCode];
    }

    function withdrawFee(bytes32 collateralAssetCode, uint256 amount) external onlyGovernor {
        require(amount <= collectedFee[collateralAssetCode], "amount must <= to collectedFee");

        collateralAssets[collateralAssetCode].transfer(msg.sender, amount);

        collectedFee[collateralAssetCode] = collectedFee[collateralAssetCode].sub(amount);
    }

    function addStableCredit(IStableCredit newStableCredit) external onlyDRS {
        require(address(newStableCredit) != address(0), "newStableCredit address must not be 0");
        bytes32 stableCreditId = Hasher.stableCreditId(newStableCredit.assetCode());
        require(address(stableCredits[stableCreditId]) == address(0), "stableCredit has already existed");

        stableCredits[stableCreditId] = newStableCredit;
        stableCreditsLL = stableCreditsLL.add(address(newStableCredit));
    }

    function getStableCreditById(bytes32 stableCreditId) external view returns (IStableCredit) {
        return stableCredits[stableCreditId];
    }

    function getRecentStableCredit() external view returns (IStableCredit) {
        address addr = stableCreditsLL.getNextOf(address(1));
        return IStableCredit(addr);
    }

    function getNextStableCredit(bytes32 stableCreditId) external view returns (IStableCredit) {
        address currentAddr = address(stableCredits[stableCreditId]);
        address nextAddr = stableCreditsLL.getNextOf(currentAddr);
        return IStableCredit(nextAddr);
    }

    function getStableCreditCount() external view returns (uint8) {
        return stableCreditsLL.llSize;
    }

    function setAllowedLink(bytes32 linkId, bool enable) external onlyGovernor {
        allowedLinks[linkId] = enable;
    }

    function isLinkAllowed(bytes32 linkId) external view returns (bool) {
        return allowedLinks[linkId];
    }
}
