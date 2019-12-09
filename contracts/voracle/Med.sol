pragma solidity ^0.5.0;

import "../book-room/LL.sol";
import "../contract-interfaces/IVS.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";

contract Med is Initializable, IVS {
    using SafeMath for uint256;

    address public gov;
    bool public active;

    LL.List public feeders;
    using LL for LL.List;

    bytes32 public price;
    bytes32 public pair;

    uint8 public minFedPrices = 0x1;

    event LogMedPrice(bytes32 price, bool isErr);

    modifier onlyGov() {
        require(msg.sender == gov, "caller must be GOV");
        _;
    }

    function initialize(address _gov, bytes32 _pair) public initializer {
        active = true;
        gov = _gov;
        pair = _pair;

        feeders.init();
    }

    function post() external {
        (bytes32 newMedPrice, bool isErr) = compute();
        set(newMedPrice);
        emit LogMedPrice(newMedPrice, isErr);
    }

    function compute() public view returns (bytes32, bool) {
        address curr = feeders.next[address(1)];
        bytes32[] memory fedPrices = new bytes32[](feeders.llSize);
        uint8 controller = 0;

        for(uint8 i = 0; curr != address(1); i++) {
            (bytes32 fedPrice, bool isErr) = IVS(curr).getWithError();
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

        bytes32 medPrice;
        if (controller % 2 == 0) {
            uint256 val1 = uint256(fedPrices[(controller/2) - 1]);
            uint256 val2 = uint256(fedPrices[(controller/2)]);
            medPrice = bytes32(val1.add(val2).div(2));
        } else {
            medPrice = fedPrices[(controller-1)/2];
        }

        return (medPrice, false);
    }

    function set(uint8 newMinFedPrices) onlyGov public {
        require(newMinFedPrices != 0, "minFedPrices must more than 0");
        minFedPrices = newMinFedPrices;
    }

    function set(address feeder) onlyGov public {
        feeders.add(feeder);
    }

    function removeFeeder(address feeder, address prevFeeder) public  {
        feeders.remove(feeder, prevFeeder);
    }

    function getFeeders() public view returns (address[] memory) {
        return feeders.getAll();
    }

    function get() external view returns (bytes32) {
        require(price > 0, "invalid price");
        return price;
    }

    function getWithError() external view returns (bytes32, bool) {
        return (price, price < 0 && active);
    }

    function set(bytes32 newPrice) public {
        require(msg.sender == address(this), "caller must be Med");
        price = newPrice;
    }

    function disable() onlyGov external {
        active = false;
    }
}
