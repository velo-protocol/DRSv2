pragma solidity ^0.5.0;

import "../../contracts/price-feeds/FeedsFactory.sol";
import "truffle/Assert.sol";

contract TestFeedsFactory {

    FeedsFactory public feedFactory;

    constructor() public {
        feedFactory = new FeedsFactory();
    }

    function testCreate() public {
        address feedsAddr = feedFactory.create();
        Assert.notEqual(feedsAddr, address(0), "feeds address should not be empty");
    }
}
