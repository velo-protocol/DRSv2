pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./VeloToken.sol";

contract DigitalReserveSystem is WhitelistAdminRole {
    using SafeMath for uint256;

    VeloToken private veloToken;

    uint256 public creditIssuanceFee;
    uint256 public collectedFee;

    mapping(address => bool) public trustedPartners;

    struct DigitalCredit {
        uint256 collateral;
        uint256 assetAmount;
        uint256 peggedValue;
        bytes32 peggedCurrency;
        bytes32 stellarIssuerAddress;
    }

    mapping(bytes32 => DigitalCredit) public digitalCredits;

    mapping(bytes32 => address) public priceFeeders;
    mapping(bytes32 => uint) public prices;

    event Mint(
        bytes32 assetName,
        uint256 amount,
        uint256 collateral,
        address indexed issuer,
        string indexed stellarAddr,
        bytes32 hashSecret
    );

    constructor(address veloTokenAddr) public {
        veloToken = VeloToken(veloTokenAddr);
    }

    function mint(
        uint256 peggedValue,
        bytes32 peggedCurrency,
        bytes32 assetName,
        uint256 amount,
        string memory stellarAddr,
        bytes32 hashSecret
    ) public payable {
        require(trustedPartners[msg.sender] == true, "only trusted partner can mint the digital credits");
        require(priceFeeders[peggedCurrency] != address(0), "price feeder for given currency hasn't been set");

        bytes32 key = keccak256(abi.encodePacked(msg.sender, assetName));

        if (digitalCredits[key].peggedValue != 0) {
            require(peggedValue == 0, "peggedValue can only be set at the 1st time digital credit minted");
            peggedValue = digitalCredits[key].peggedValue;
            peggedCurrency = digitalCredits[key].peggedCurrency;
        }

        (uint256 requiredCollateral, uint256 fee) = _calCollateralWithFee(peggedValue, peggedCurrency, amount);

        require(requiredCollateral > 0, "every digital credits must have collateral");
        require(veloToken.balanceOf(msg.sender) > requiredCollateral, "not enough VELO");
        require(veloToken.allowance(msg.sender, address(this)) > requiredCollateral.add(fee), "not enough allowance");

        veloToken.transferFrom(msg.sender, address(this), requiredCollateral.add(fee));

        digitalCredits[key] = DigitalCredit(
            digitalCredits[key].collateral.add(requiredCollateral),
            digitalCredits[key].assetAmount.add(amount),
            peggedValue,
            peggedCurrency,
            digitalCredits[key].stellarIssuerAddress
        );

        collectedFee = collectedFee.add(fee);

        emit Mint(assetName, amount, requiredCollateral, msg.sender, stellarAddr, hashSecret);
    }

    function __callbackConfirmIssueCredit(address issuer, bytes32 assetName, string memory stellarIssuerAddress) public onlyWhitelistAdmin {
        bytes32 key = keccak256(abi.encodePacked(msg.sender, assetName));
        digitalCredits[key].stellarIssuerAddress = keccak256(abi.encodePacked(stellarIssuerAddress));
    }

    function redeem(
        bytes32 assetName,
        uint256 amount
    ) public {

    }

    function checkRequiredCollateralWithFee(
        uint256 peggedValue,
        bytes32 peggedCurrency,
        uint256 amount
    ) public view returns (uint256, uint256) {
        (uint256 requiredCollateral, uint256 fee) = _calCollateralWithFee(peggedValue, peggedCurrency, amount);

        return (requiredCollateral, fee);
    }

    function _calCollateralWithFee(uint256 peggedValue, bytes32 peggedCurrency, uint256 amount) private view returns (uint256, uint256) {
        uint256 collateral = peggedValue.div(prices[peggedCurrency]).mul(amount);
        uint256 fee = collateral.mul(creditIssuanceFee).div(10000);

        return (collateral, fee);
    }

    function collateralOf(bytes32 assetName, address issuer) public view returns (uint256) {
        bytes32 key = keccak256(abi.encodePacked(issuer, assetName));
        return digitalCredits[key].collateral;
    }

    function getCollectedFee() public view returns (uint256) {
        return collectedFee;
    }

    function withdrawFee(uint256 amount) public onlyWhitelistAdmin {
        require(amount <= collectedFee, "amount must <= to collectedFee");
        veloToken.transfer(msg.sender, amount);
        collectedFee = collectedFee.sub(amount);
    }

    function setCreditIssuanceFee(uint256 newFee) public onlyWhitelistAdmin {
        creditIssuanceFee = newFee;
    }

    function setTrustedPartner(address addr) public onlyWhitelistAdmin {
        trustedPartners[addr] = true;
    }

    function setPriceFeeder(bytes32 currency, address priceFeederAddr) public onlyWhitelistAdmin {
        priceFeeders[currency] = priceFeederAddr;
    }

    function setPrice(bytes32 currency, uint newPrice) public {
        require(msg.sender == priceFeeders[currency], "only price feeder can update the price");
        prices[currency] = newPrice;
    }

    function getPrice(bytes32 currency) public view returns (uint) {
        return prices[currency];
    }
}
