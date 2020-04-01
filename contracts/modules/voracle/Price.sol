pragma solidity ^0.5.0;

import "../interfaces/IPrice.sol";
import "../interfaces/ILag.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";


contract Price is Initializable, IPrice {
    using SafeMath for uint256;

    address public owner;
    modifier onlyOwner {
        require(msg.sender == owner, "Price.onlyOwner: The message sender is not found or does not have sufficient permission");
        _;
    }

    uint256 public price;
    address public lagAddr;
    bool isErr;

    bool public active;

    event PriceVoid(address caller, address price, bool isActive);
    event PriceActivate(address caller, address price, bool isActive);

    function setLag(address newLagAddr) external onlyOwner {
        lagAddr = newLagAddr;
    }


    function initialize(address _owner, address _lagAddr) public initializer {
        lagAddr = _lagAddr;
        owner = _owner;
        price = 0;
        active = true;
        isErr = true;
    }

    function post() external {
        (uint256 lagPrice, bool isActive, bool lagIsErr) = ILag(lagAddr).getWithError();
        if (!isActive) {
            active = false;
        }

        isErr = lagIsErr;

        if (lagPrice <= 0) {
            isErr = true;
        }
        price = lagPrice;

    }

    function get() external view returns (uint256){
        require(active == true, "Price.get: price is not active");
        require(price > 0, "Price.get: invalid price");
        require(isErr == false, "Price.get: price has an error");
        return (price);
    }

    function getWithError() external view returns (uint256, bool, bool) {
        return (price, active, isErr);
    }

    function void() external onlyOwner {
        price = 0;
        active = false;
        isErr = true;
        emit PriceVoid(msg.sender, address(this), active);
    }

    function activate() external onlyOwner {
        require(active == false, "Price.activate: price is active");
        require(price > 0, "Price.activate: price is not in a correct state");
        active = true;
        emit PriceActivate(msg.sender, address(this), active);
    }

}