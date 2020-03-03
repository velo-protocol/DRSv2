pragma solidity ^0.5.0;

import "../../../contracts/modules/price-feeder/Feeder.sol";
import "truffle/Assert.sol";

contract TestFeeder {
    Feeder public feeder;

    constructor() public {
        feeder = new Feeder(address(this), "USD", "VELO");
    }

    function testInit() public {
        Assert.equal(feeder.value(), 0, "feeds.value should be 0x0");
        Assert.equal(feeder.active(), true, "feeds.active should be true");
        Assert.equal(feeder.owner(), address(this), "owner should be address(this)");
    }

    function testPost() public {
        uint256 expected = 120;
        feeder.post(expected);
        Assert.equal(feeder.value(), expected, "feeds.value should be 120");
        Assert.equal(feeder.valueTimestamp(), block.timestamp, "feeder.valueTimestamp must be set to now timestamp");
    }

    function testGetWithError() public {
        (uint256 value, uint256 valueTimestamp, bool isErr) = feeder.getWithError();
        Assert.equal(value, feeder.value(), "value should be the same as feeds.value");
        Assert.equal(valueTimestamp, feeder.valueTimestamp(), "valueTimestamp should be the same as feeds.valueTimestamp");
        Assert.equal(isErr, false, "isErr should be false");
    }

    function testGet() public {
        uint256 value = feeder.get();
        Assert.equal(value, feeder.value(), "value should be the same as feeds.value");
    }

    function testDisable() public {
        feeder.disable();
        Assert.isFalse(feeder.active(), "feeds.active should be false");
    }

    function testEnable() public {
        feeder.enable();
        Assert.isTrue(feeder.active(), "feeds.active should be true");
    }
}
