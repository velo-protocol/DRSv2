pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../interfaces/IFeeder.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";

contract Med is Initializable, IFeeder {
    using SafeMath for uint256;

    address public gov;
    bool public active;

    LL.List public feeders;
    using LL for LL.List;

    uint256 public price;
    bytes32 public pair;

    uint8 public minFedPrices = 0x1;

    event LogMedPrice(uint256 price, bool isErr);

    modifier onlyGov() {
        require(msg.sender == gov, "Med | caller must be GOV");
        _;
    }

    function initialize(address _gov, bytes32 _pair) public initializer {
        active = true;
        gov = _gov;
        pair = _pair;

        feeders.init();
    }

    function post() external {
        (uint256 newMedPrice, bool isErr) = compute();
        _set(newMedPrice);
        emit LogMedPrice(newMedPrice, isErr);
    }

    function compute() public view returns (uint256, bool) {
        address curr = feeders.next[address(1)];
        uint256[] memory fedPrices = new uint256[](feeders.llSize);
        uint8 controller = 0;

        for(uint8 i = 0; curr != address(1); i++) {
            (uint256 fedPrice, bool isErr) = IFeeder(curr).getWithError();
            if(!isErr) {
                if(controller == 0 || fedPrice >= fedPrices[controller - 1]) {
                    fedPrices[controller] = fedPrice;
                } else {
                    uint8 j = 0;
                    while(fedPrice >= fedPrices[j]) {
                        j = j+1;
                    }
                    for(uint8 k = controller; k > j; k--) {
                        fedPrices[k] = fedPrices[k-1];
                    }
                    fedPrices[j] = fedPrice;
                }
                controller++;
            }
            curr = feeders.next[curr];
        }

        if(controller < minFedPrices) {
            return (price, true);
        }

        uint256 medPrice;
        if (controller % 2 == 0) {
            uint256 val1 = fedPrices[(controller/2) - 1];
            uint256 val2 = fedPrices[(controller/2)];
            medPrice = val1.add(val2).div(2);
        } else {
            medPrice = fedPrices[(controller-1)/2];
        }

        return (medPrice, false);
    }

    function setMinFedPrices(uint8 newMinFedPrices) onlyGov public {
        require(newMinFedPrices != 0, "Med | minFedPrices must more than 0");
        minFedPrices = newMinFedPrices;
    }

    function addFeeder(address feeder) onlyGov public {
        feeders.add(feeder);
    }

    function removeFeeder(address feeder, address prevFeeder) public  {
        feeders.remove(feeder, prevFeeder);
    }

    function getFeeders() public view returns (address[] memory) {
        return feeders.getAll();
    }

    function get() external view returns (uint256) {
        require(price > 0, "Med | invalid price");
        return price;
    }

    function getWithError() external view returns (uint256, bool) {
        return (price, price <= 0 || !active);
    }

    function set(uint256 newPrice) external {
        require(msg.sender == address(this), "Med | caller must be Med");
        _set(newPrice);
    }

    function _set(uint256 newPrice) internal {
        price = newPrice;
    }

    function disable() onlyGov external {
        active = false;
    }
}
