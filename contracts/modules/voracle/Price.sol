pragma solidity ^0.5.0;

import "../interfaces/IPrice.sol";
import "../interfaces/IFeeder.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";


contract Price is Initializable, IPrice {

    using SafeMath for uint256;

    address public owner;
    modifier onlyOwner {
        require(msg.sender == owner, "Price.onlyOwner: The message sender is not found or does not have sufficient permission");
        _;
    }

    uint256 public price=0;
    address public lagAddr;
    address  public  feederAddr;
    bool isErr;

    bool public active;

    event PriceVoid(address caller, address price, bool isActive);
    event PriceActivate(address caller, address price, bool isActive);


    function initialize(address _owner, address _feederAddr) public initializer {
        feederAddr = _feederAddr;
        owner = _owner;
        price = 0;
        active = true;
        isErr = true;
    }

    function post() external {
        (uint lagPrice,) = IFeeder(feederAddr).getLastPrice();
        require(lagPrice>0,"price need >0");
        isErr=false;
        price = lagPrice;
    }


    function getWithError() external view returns (uint256, bool, bool) {
        return (price, active, isErr);
    }
    function get() external view returns (uint256){
        require(active == true, "Price.get: price is not active");
        require(price > 0, "Price.get: invalid price");
        require(isErr == false, "Price.get: price has an error");
        return (price);
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