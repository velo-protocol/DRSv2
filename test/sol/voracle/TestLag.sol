pragma solidity ^0.5.0;

import "../../../contracts/modules/voracle/Lag.sol";
import "../mocks/MockMedianizer.sol";
import "truffle/Assert.sol";

contract TestLag {
    Lag public lag;
    Lag public lagNoneGov;
    MockMedianizer public mockMedianizer;
    MockMedianizer public mockMedianizer2;

    function beforeEach() public {
        mockMedianizer = new MockMedianizer(10);
        mockMedianizer2 = new MockMedianizer(11);
        lag = new Lag(address(this), address(0));
    }

    function testDeactivate() public {
        lag.activate();
        Assert.isTrue(lag.active(), "lag.active() should be true");
        lag.deactivate();
        Assert.isFalse(lag.active(), "lag.active() should be false");
    }

    function testDeactivateWhenNotGov() public {
        lagNoneGov = new Lag(address(1), address(2));
        (bool r,) = address(lagNoneGov).call(abi.encodePacked(lagNoneGov.deactivate.selector));

        Assert.equal(r, false, "lag.active() should be false");
    }

    function testActivate() public {
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

    function testSetLagTime() public {
        bytes4 selector = lag.setLagTime.selector;
        uint16 newLagTime = 60 * 60 * 16;
        bytes memory data = abi.encodeWithSelector(selector, newLagTime);
        (bool r,) = address(lag).call(data);

        Assert.isTrue(r, "lag.setLagTime() should not throw an error");
    }

    function testSetLagTimeWithError() public {
        bytes4 selector = lag.setLagTime.selector;
        uint16 newLagTime = 0;
        bytes memory data = abi.encodeWithSelector(selector, newLagTime);
        (bool r,) = address(lag).call(data);

        Assert.isFalse(r, "lag.setLagTime() should throw an error");
    }

    function testVoid() public {
        lag.void();
        Assert.equal(lag.active(), false, "lag.active() should be false after lag.void() has been called");

        (uint256 price, bool isErr) = lag.getWithError();
        Assert.equal(price, 0, "lag.getWithError() should return price equal to 0");
        Assert.equal(isErr, true, "lag.getWithError() should return isErr equal to true");

    }

    function testIsLagTimePass() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.isLagTimePass.selector));

        Assert.isTrue(r, "lag.isLagTimePass() should not throw an error");
    }

    function testPost() public {
        lag.activate();
        lag.setMedianizer(address(mockMedianizer));

        lag.post();

        (uint256 curr, bool isErr) = lag.getWithError();

        Assert.equal(curr, 0, "curr should be 0");
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
        (uint256 curr, bool isErr) = lag.getWithError();

        Assert.equal(curr, 0, "lag.getWithError() curr should be 0");
        Assert.isFalse(isErr, "lag.getWithError() isErr should be false");
    }

    function testGetNextWithError() public {
        (uint256 curr, bool isErr) = lag.getNextWithError();

        Assert.equal(curr, 0, "lag.getNextWithError() curr should be 0");
        Assert.isFalse(isErr, "lag.getNextWithError() isErr should be false");
    }

    function testGet() public {
        uint256 currPrice = lag.get();

        Assert.equal(currPrice, 0, "lag.get() should be 0");
    }

    function testGetAfterPost() public {
        lag.activate();
        lag.setMedianizer(address(mockMedianizer));
        lag.post();

        uint256 currPrice = lag.get();
        Assert.equal(currPrice, 0, "lag.get() should be 0");
    }
}
