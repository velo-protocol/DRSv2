pragma solidity ^0.5.0;

import "../../../contracts/modules/voracle/Lag.sol";
import "../mocks/MockMedianizer.sol";
import "truffle/Assert.sol";

contract TestLag {
    Lag public lag;
    Lag public lagNoneGov;
    MockMedianizer public mockMedianizer;

    function beforeEach() public {
        mockMedianizer = new MockMedianizer(10);
        lag = new Lag();
        lag.initialize(address(this), address(0));
    }

    function testDeactivate() public {
        lag.deactivate();
        Assert.isFalse(lag.active(), "lag.active() should be false");
    }

    function testDeactivateWhenNotGov() public {
        lagNoneGov = new Lag();
        lagNoneGov.initialize(address(1), address(2));

        (bool r,) = address(lagNoneGov).call(abi.encodePacked(lagNoneGov.deactivate.selector));

        Assert.equal(r, false, "lag.active() should be false");
    }

    function testActivate() public {
        lag.setMedianizer(address(mockMedianizer));
        lag.setMinimumPeriod(0);
        lag.post();
        lag.post();
        lag.deactivate();
        Assert.isFalse(lag.active(), "lag.active() should be false");
        lag.activate();
        Assert.isTrue(lag.active(), "lag.active() should be true");
    }

    function testSetPriceRefStorage() public {
        address expectedMedianizer = address(mockMedianizer);
        lag.setMedianizer(address(mockMedianizer));

        Assert.equal(lag.medianizerAddr(), expectedMedianizer, "lag.setMedianizer() should be 0xb04aD816e86bFA5515c35Ad02081F71D9E848C88");
    }

    function testSetMinimumPeriod() public {
        bytes4 selector = lag.setMinimumPeriod.selector;
        uint16 newMinPeriod = 60 * 60 * 16;
        bytes memory data = abi.encodeWithSelector(selector, newMinPeriod);
        (bool r,) = address(lag).call(data);

        Assert.isTrue(r, "lag.setMinimumPeriod() should not throw an error");
    }

    function testSetMinimumPeriodWithError() public {
        bytes4 selector = lag.setMinimumPeriod.selector;
        uint256 newLagTime = 31 days;
        bytes memory data = abi.encodeWithSelector(selector, newLagTime);
        (bool r,) = address(lag).call(data);

        Assert.isFalse(r, "lag.setLagTime() should throw an error");
    }

    function testVoid() public {
        lag.void();
        Assert.equal(lag.active(), false, "lag.active() should be false after lag.void() has been called");

        (uint256 price, bool isActive, bool isErr) = lag.getWithError();
        Assert.equal(price, 0, "lag.getWithError() should return price equal to 0");
        Assert.equal(isActive, false, "lag.getWithError() should return isActive equal to true");
        Assert.equal(isErr, true, "lag.getWithError() should return isErr equal to true");
    }

    function testIsMinimumPeriodPass() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.isMinimumPeriodPass.selector));

        Assert.isTrue(r, "lag.isMinimumPeriodPass() should not throw an error");
    }

    function testPost() public {
        lag.setMedianizer(address(mockMedianizer));
        lag.setMinimumPeriod(0);
        lag.post();

        (uint256 price, bool isActive, bool isErr) = lag.getNextWithError();

        Assert.equal(price, 10, "price should be 10");
        Assert.equal(isActive, true, "isActive are should be false");
        Assert.equal(isErr, false, "isErr are should be false");
    }

    function testPostWhenPriceRefNotSet() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.post.selector));

        Assert.equal(r, false, "lag.post() should throw an error");
    }

    function testPostWhenDeactivated() public {
        lag.deactivate();
        (bool r,) = address(lag).call(abi.encodePacked(lag.post.selector));

        Assert.equal(r, false, "lag.post() should throw an error");
    }

    function testGetWithError() public {
        lag.setMedianizer(address(mockMedianizer));
        lag.setMinimumPeriod(0);
        lag.post();
        lag.post();

        (uint256 curr, bool isActive, bool isErr) = lag.getWithError();

        Assert.equal(curr, 10, "lag.getWithError() curr should be 10");
        Assert.isTrue(isActive, "lag.getWithError() isActive should be true");
        Assert.isFalse(isErr, "lag.getWithError() isErr should be false");
    }

    function testGetNextWithError() public {
        lag.setMedianizer(address(mockMedianizer));
        lag.setMinimumPeriod(0);
        lag.post();

        (uint256 curr, bool isActive, bool isErr) = lag.getNextWithError();

        Assert.equal(curr, 10, "lag.getNextWithError() curr should be 10");
        Assert.isTrue(isActive, "lag.getNextWithError() isActive should be true");
        Assert.isFalse(isErr, "lag.getNextWithError() isErr should be false");
    }

    function testGet() public {
        lag.setMedianizer(address(mockMedianizer));
        lag.setMinimumPeriod(0);
        lag.post();
        lag.post();

        uint256 currPrice = lag.get();

        Assert.equal(currPrice, 10, "lag.get() should be 10");
    }
}
