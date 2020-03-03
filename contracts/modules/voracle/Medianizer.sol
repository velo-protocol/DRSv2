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

    uint8 public minFedPrices = 0x1;
    uint256 public validityPeriod = 15 minutes;

    event MedianSet(uint256 price, bool isErr);

    modifier onlyOwner() {
        require(msg.sender == owner, "Med.onlyOwner: caller must be an owner of this contract");
        _;
    }

    function initialize(address _owner, bytes32 _fiatCode, bytes32 _collateralCode) public initializer {
        active = true;
        owner = _owner;
        fiatCode = _fiatCode;
        collateralCode = _collateralCode;

        feeders.init();
    }

    function post() external {
        (uint256 newMedPrice, bool isErr) = compute();
        _set(newMedPrice);
        emit MedianSet(newMedPrice, isErr);
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
    }

    function removeFeeder(address feeder, address prevFeeder) public {
        feeders.remove(feeder, prevFeeder);
    }

    function getFeeders() public view returns (address[] memory) {
        return feeders.getAll();
    }

    function get() external view returns (uint256) {
        require(price > 0, "Med | invalid price");
        return price;
    }

    function getWithError() external view returns (uint256, bool, bool) {
        return (price, active, price <= 0);
    }

    function _set(uint256 newPrice) internal {
        price = newPrice;
    }

    function enable() onlyOwner external {
        active = true;
    }

    function void() onlyOwner external {
        price = 0;
        active = false;
    }

    function getValidityPeriod() external view returns (uint256) {
        return validityPeriod / 1 minutes;
    }

    function setValidityPeriod(uint256 valueInMinute) onlyOwner external {
        validityPeriod = valueInMinute * 1 minutes;
    }
}
