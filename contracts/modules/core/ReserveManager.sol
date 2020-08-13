pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/access/roles/WhitelistAdminRole.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "../interfaces/IRM.sol";
import "../interfaces/IHeart.sol";


contract ReserveManager is IRM,ReentrancyGuard {
    IHeart public heart;
    /*
        WARNINGS: DO NOT SEND TOKEN TO THIS CONTRACT DIRECTLY
    */

    using SafeMath for uint256;

    struct LockedReserve {
        address owner;
        bytes32 assetCode;
        uint256 amount;
        uint256 blockNumber;
        uint256 time;
    }

    /*
        lockedReserveId = keccak256(abi.encodePacked(from, assetCode, amount, block.number))
        lockedReserveId => LockedReserve
    */
    mapping(bytes32 => LockedReserve) public lockedReserves;

    modifier onlyDRSSC() {
        require(heart.getDrsAddress() == msg.sender, "caller is not DRSSC");
        _;
    }

    modifier onlyHeart() {
        require(address(heart) == msg.sender, "caller is not Heart");
        _;
    }

    constructor(address heartAddr) public {
        heart = IHeart(heartAddr);
    }

    event InjectCollateral(bytes32 indexed assetCode, uint256 amount);
    event LockReserve(bytes32 indexed lockedReserveId);

    function lockReserve(bytes32 assetCode, address from, uint256 amount) external nonReentrant {

        uint256 balance= heart.getCollateralAsset(assetCode).balanceOf(address(this));
        heart.getCollateralAsset(assetCode).transferFrom(from, address(this), amount);
        uint256 balance1= heart.getCollateralAsset(assetCode).balanceOf(address(this));
        require(balance1==balance.add(amount),"");
        bytes32 lockedReserveId = keccak256(abi.encodePacked(from, assetCode, amount, block.number));
        require(lockedReserves[lockedReserveId].owner==address(0));
        lockedReserves[lockedReserveId] = LockedReserve(
            from,
            assetCode,
            amount,
            block.number,
            now
        );

        emit LockReserve(lockedReserveId);
    }

    function releaseReserve(bytes32 lockedReserveId, bytes32 assetCode, uint256 amount) external nonReentrant {
        require(now.sub(lockedReserves[lockedReserveId].time) > heart.getReserveFreeze(assetCode), "release time not reach");
        require(lockedReserves[lockedReserveId].owner == msg.sender, "only owner can release reserve");
        heart.getCollateralAsset(assetCode).transfer(msg.sender, amount);

        lockedReserves[lockedReserveId].amount = lockedReserves[lockedReserveId].amount.sub(amount);
    }

    function injectCollateral(bytes32 assetCode, address to, uint256 amount) external onlyDRSSC {
        heart.getCollateralAsset(assetCode).transfer(to, amount);

        emit InjectCollateral(assetCode, amount);
    }

}
