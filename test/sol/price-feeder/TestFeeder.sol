pragma solidity ^0.5.0;

import "../../../contracts/modules/price-feeder/Feeder.sol";
import "truffle/Assert.sol";

contract TestFeeder {
    Feeder public feeder;

    address  pf1 = address(this);
    address  pf2 = address(this);
    address  pf3 = address(this);

    constructor() public {
        feeder = new Feeder(pf1,pf1,pf1,address(this), "USD", "VELO");
    }

    function testInit() public {
        Assert.equal(feeder.owner(), address(this), "owner should be address(this)");
    }

    function testStartOracle() public {
        uint  expected = 120;
        feeder.startOracle(expected);
        (uint price,)=feeder.getLastPrice();
        Assert.equal(price, expected, "feeds.value should be 120");
    }

    function testCommitPrice() public {
        uint  expected = 200;
        feeder.setValue(3,0);
        feeder.commitPrice(expected);
        (uint price,)=feeder.getLastPrice();
        Assert.equal(price, 120, "feeds.value should be 120");
    }

}
