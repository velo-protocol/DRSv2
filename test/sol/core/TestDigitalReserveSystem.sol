pragma solidity ^0.5.0;

import "../../../contracts/modules/core/DigitalReserveSystem.sol";

contract TestDigitalReserveSystem {

    DigitalReserveSystem public drs;

    function beforeEach() public {
        drs = new DigitalReserveSystem();
    }

    constructor() public {

        mockHeart = new MockHeart();
    }

    function testMintFromCollateral_Success() public {
        uint256 collateralAmount = 100;
        string assetCode = "VELO";
        drs.mintFromCollateral(collateralAmount,assetCode);
    }
}
