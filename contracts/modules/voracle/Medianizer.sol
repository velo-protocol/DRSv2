pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IFeeder.sol";
import "../interfaces/IMedianizer.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";

contract Medianizer is Initializable, IMedianizer {
    using SafeMath for uint256;

    LL.List public feeders;
    using LL for LL.List;

    address public owner;
    bool public active;
    uint256 public price;
    bytes32 public fiatCode;
    bytes32 public collateralCode;

    uint8 public minFedPrices;
    uint256 public validityPeriod;
    uint256 private maxValidityPeriod;

    event FeederAdd(address caller, address medianizer, address feeder);
    event FeederRemove(address caller, address medianizer, address feeder);
    event MedianActivate(address caller, address medianizer, bool isActive);
    event MedianVoid(address caller, address medianizer, bool isActive);

    modifier onlyOwner() {
        require(msg.sender == owner, "Medianizer.onlyOwner: The message sender is not found or does not have sufficient permission");
        _;
    }

    function initialize(address _owner, bytes32 _fiatCode, bytes32 _collateralCode) public initializer {
        active = true;
        owner = _owner;
        fiatCode = _fiatCode;
        collateralCode = _collateralCode;
        minFedPrices = 0x1;
        validityPeriod = 15 minutes;
        maxValidityPeriod = 30 days;

        feeders.init();
    }

    function post() external {
        (uint256 newMedPrice,) = compute();
        _set(newMedPrice);
    }

    function compute() public view returns (uint256, bool) {
        address curr = feeders.next[address(1)];
        uint256[] memory fedPrices = new uint256[](feeders.llSize);
        uint8 controller = 0;

        for (uint8 i = 0; curr != address(1); i++) {
            (uint256 fedPrice, uint256 timestamp, bool isErr) = IFeeder(curr).getWithError();

            if (timestamp < now - validityPeriod) {
                isErr = true;
            }

            if (!isErr) {
                if (controller == 0 || fedPrice >= fedPrices[controller - 1]) {
                    fedPrices[controller] = fedPrice;
                } else {
                    uint8 j = 0;
                    while (fedPrice >= fedPrices[j]) {
                        j = j + 1;
                    }
                    for (uint8 k = controller; k > j; k--) {
                        fedPrices[k] = fedPrices[k - 1];
                    }
                    fedPrices[j] = fedPrice;
                }
                controller++;
            }
            curr = feeders.next[curr];
        }

        if (controller < minFedPrices) {
            return (0, true);
        }

        uint256 medPrice;
        if (controller % 2 == 0) {
            uint256 val1 = fedPrices[(controller / 2) - 1];
            uint256 val2 = fedPrices[(controller / 2)];
            medPrice = val1.add(val2).div(2);
        } else {
            medPrice = fedPrices[(controller - 1) / 2];
        }

        return (medPrice, false);
    }

    function setMinFedPrices(uint8 newMinFedPrices) onlyOwner public {
        require(newMinFedPrices != 0, "Med | minFedPrices must more than 0");
        minFedPrices = newMinFedPrices;
    }

    function addFeeder(address feeder) onlyOwner public {
        feeders.add(feeder);
        emit FeederAdd(msg.sender, address(this), feeder);
    }

    function removeFeeder(address feeder) onlyOwner public {
        address[] memory feedersArray = getFeeders();
        address prevFeeder = address(1);
        bool found = false;
        for (uint i = 0; i < feedersArray.length; i++) {
            if (feeder == feedersArray[i]) {
                found = true;
                break;
            }

            prevFeeder = feedersArray[i];
        }

        require(found, "Medianizer.removeFeeder: address does not exist");

        feeders.remove(feeder, prevFeeder);
        emit FeederRemove(msg.sender, address(this), feeder);
    }

    function getFeeders() public view returns (address[] memory) {
        return feeders.getAll();
    }

    function get() external view returns (uint256) {
        require(price > 0, "Medianizer.get: invalid price");
        require(active, "Medianizer.get: medianizer is inactive");
        return price;
    }

    function getWithError() external view returns (uint256, bool, bool) {
        return (price, active, price <= 0);
    }

    function _set(uint256 newPrice) internal {
        price = newPrice;
    }

    function activate() onlyOwner external {
        require(!active, "Medianizer.activate: the medianizer is active");
        require(price > 0, "Medianizer.activate: the medianizer does not have a valid price");
        active = true;
        emit MedianActivate(msg.sender, address(this), active);
    }

    function void() onlyOwner external {
        price = 0;
        active = false;
        emit MedianVoid(msg.sender, address(this), active);
    }

    function getValidityPeriod() external view returns (uint256) {
        return validityPeriod;
    }

    function setValidityPeriod(int256 newValidityPeriod) onlyOwner external {
        require(newValidityPeriod > 0, "Medianizer.setValidityPeriod: validityPeriod must be greater than 0");
        require(uint256(newValidityPeriod) <= maxValidityPeriod, "Medianizer.setValidityPeriod: validityPeriod must not be greater than maxValidityPeriod");
        validityPeriod = uint256(newValidityPeriod);
    }
}
