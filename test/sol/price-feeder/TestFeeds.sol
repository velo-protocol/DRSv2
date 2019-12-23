pragma solidity ^0.5.0;

import "../../../contracts/modules/price-feeder/Feeds.sol";
import "../mocks/MockMed.sol";
import "truffle/Assert.sol";

contract TestFeeds {
    Feeds public feeds;
    MockMed public mockMed;

    constructor() public {
        feeds = new Feeds(address(this));
        mockMed = new MockMed();
    }

    function testInit() public {
        Assert.equal(feeds.value(), 0, "feeds.value should be 0x0");
        Assert.equal(feeds.active(), true, "feeds.active should be true");
        Assert.equal(feeds.owner(), address(this), "owner should be address(this)");
    }

    function testSet() public {
        uint256 expected = 120;
        feeds.set(expected);
        Assert.equal(feeds.value(), expected, "feeds.value should be 120");
    }

    function testPostMed() public {
        uint256 expected = 130;
        feeds.postMed(expected, address(mockMed));
        Assert.equal(feeds.value(), expected, "feeds.value should be 130");
    }

    function testGetWithError() public {
        (uint256 value, bool isErr) = feeds.getWithError();
        Assert.equal(value, feeds.value(), "value should be the same as feeds.value");
        Assert.equal(isErr, false, "isErr should be false");
    }

    function testGet() public {
        uint256 value = feeds.get();
        Assert.equal(value, feeds.value(), "value should be the same as feeds.value");
    }

    function testDisable() public {
        feeds.disable();
        Assert.equal(feeds.active(), false, "feeds.active should be false");
    }
}
