pragma solidity ^0.5.0;

import "../../../contracts/modules/core/DigitalReserveSystem.sol";

contract TestDigitalReserveSystem {

    DigitalReserveSystem public drs;
    MockIHeart public mockIHeart;

    function beforeEach() public {
        drs = new DigitalReserveSystem();
    }

    constructor() public {
        mockHeart = new Heart();
    }
}
