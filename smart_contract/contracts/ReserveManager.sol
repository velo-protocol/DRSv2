pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./IRM.sol";
import "./IGOV.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";


contract ReserveManager is IRM {
    /*
        WARNINGS: DO NOT SEND TOKEN TO THIS CONTRACT DIRECTLY
    */

    using SafeMath for uint256;

    struct LockedReserve {
        address owner;
        bytes32 assetCode;
        uint256 amount;
        uint256 blockNumber;
    }

    mapping(bytes32 => LockedReserve) public lockedReserves;

    address public drsAddress;
    IGOV public gov;

    modifier onlyDRSSC() {
        require(drsAddress == msg.sender, "caller is not DRSSC");
        _;
    }

    modifier onlyGov() {
        require(address(gov) == msg.sender, "caller is not Governance");
        _;
    }

    constructor(address digitalReserveSystemAddr, address governanceAddr) public {
        drsAddress = digitalReserveSystemAddr;
        gov = IGOV(governanceAddr);
    }

    event InjectCollateral(bytes32 indexed assetCode, uint256 amount);
    event LockReserve(bytes32 indexed lockedReserveId);

    function lockReserve(bytes32 assetCode, address from, uint256 amount) external {
        gov.getCollateralAsset(assetCode).transferFrom(from, address(this), amount);

        bytes32 lockedReserveId = keccak256(abi.encodePacked(from, assetCode, amount, block.number));
        lockedReserves[lockedReserveId] = LockedReserve(
            from,
            assetCode,
            amount,
            block.number
        );

        emit LockReserve(lockedReserveId);
    }

    function releaseReserve(bytes32 lockedReserveId, bytes32 assetCode, uint256 amount) external {
        require(block.number.sub(lockedReserves[lockedReserveId].blockNumber).mul(3) < gov.getReserveFreeze(assetCode), "release time not reach");
        require(lockedReserves[lockedReserveId].owner == msg.sender, "only owner can release reserve");

        gov.getCollateralAsset(assetCode).transfer(msg.sender, amount);

        lockedReserves[lockedReserveId].amount = lockedReserves[lockedReserveId].amount.sub(amount);
    }

    function injectCollateral(bytes32 assetCode, address to, uint256 amount) external onlyDRSSC {
        gov.getCollateralAsset(assetCode).transfer(to, amount);

        emit InjectCollateral(assetCode, amount);
    }

}
