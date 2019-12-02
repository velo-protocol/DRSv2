pragma solidity ^0.5.0;

import "../../contracts/price-feeds/Feeds.sol";
import "../../contracts/contract-interfaces/IMED.sol";
import "truffle/Assert.sol";

contract MockMed is IMED {
    function post() external {
        // mock stuffs
    }
}

contract TestFeeds {
    Feeds public feeds;
    MockMed public mockMed;

    constructor() public {
        feeds = new Feeds(address(this));
        mockMed = new MockMed();
    }

    function testInit() public {
        Assert.equal(feeds.value(), "", "feeds.value should be 0x0");
        Assert.equal(feeds.active(), true, "feeds.active should be true");
        Assert.equal(feeds.owner(), address(this), "owner should be address(this)");
    }

    function testSet() public {
        bytes32 expected = "120";
        feeds.set(expected);
        Assert.equal(feeds.value(), expected, "feeds.value should be 120");
    }

    function testSetMed() public {
        bytes32 expected = "130";
        feeds.set(expected, address(mockMed));
        Assert.equal(feeds.value(), expected, "feeds.value should be 130");
    }

    function testGetWithError() public {
        (bytes32 value, bool isErr) = feeds.getWithError();
        Assert.equal(value, feeds.value(), "value should be the same as feeds.value");
        Assert.equal(isErr, false, "isErr should be false");
    }

    function testGet() public {
        bytes32 value = feeds.get();
        Assert.equal(value, feeds.value(), "value should be the same as feeds.value");
    }

    function testDisable() public {
        feeds.disable();
        Assert.equal(feeds.active(), false, "feeds.active should be false");
    }
}
