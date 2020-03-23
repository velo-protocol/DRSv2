pragma solidity ^0.5.0;

import "../interfaces/IPrice.sol";
import "../interfaces/ILag.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/upgrades/contracts/Initializable.sol";


contract Price is Initializable, IPrice {
    using SafeMath for uint256;

    address public owner;
    modifier onlyOwner {
        require(msg.sender == owner, "Lag.onlyOwner: The message sender is not found or does not have sufficient permission");
        _;
    }

    uint256 public price;
    address public lagAddr;
    bool isErr;

    bool public active;

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
}