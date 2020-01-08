pragma solidity ^0.5.0;

import "../../../contracts/modules/core/DigitalReserveSystem.sol";
import "../mocks/MockIHeart.sol";

contract TestDigitalReserveSystem {

    DigitalReserveSystem public drs;
    MockIHeart public mockIHeart;

    function beforeEach() public {
        drs = new DigitalReserveSystem(mockIHeart);
    }

    constructor() public {
    }

    function testSetup_Success() public {
        mockIHeart = new MockIHeart(
            IRM(address(0)),
            IPF(address(0)),
            new ERC20(),
            StableCredit(address(0)),
            StableCredit(address(0))
        );
        drs = new DigitalReserveSystem(address(mockIHeart));

        drs.setup("VELO", "USD", "vUSD", 1);
    }
}
