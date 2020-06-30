pragma solidity ^0.5.0;

import "../../../contracts/modules/price-feeder/FeederFactory.sol";
import "truffle/Assert.sol";

contract TestFeederFactory {

    FeederFactory public feederFactory;

    constructor() public {
        feederFactory = new FeederFactory();
    }

    function testCreate() public {
        address feederAddr = feederFactory.create(address(this),address(this),address(this),"USD", "VELO");
        Assert.notEqual(feederAddr, address(0), "feeds address should not be empty");
    }
}
