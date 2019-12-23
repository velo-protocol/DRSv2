pragma solidity ^0.5.0;

import "../../../contracts/modules/voracle/Lag.sol";
import "../mocks/MockIPRS.sol";
import "truffle/Assert.sol";

contract TestLag {
    Lag public lag;
    Lag public lagNoneGov;
    MockIPRS public mockIPRS;
    MockIPRS public mockIPRS2;

    function beforeEach() public {
        mockIPRS = new MockIPRS(10);
        mockIPRS2 = new MockIPRS(11);
        lag = new Lag(address(this), address(0));
    }

    function testHalt() public {
        Assert.isFalse(lag.halted(), "lag.halted() should be false");
        lag.halt();
        Assert.isTrue(lag.halted(), "lag.halted() should be true");
    }

    function testHaltWhenNotGov() public {
        lagNoneGov = new Lag(address(1), address(2));
        (bool r,) = address(lagNoneGov).call(abi.encodePacked(lagNoneGov.halt.selector));

        Assert.equal(r, false, "lag.halted() should be false");
    }

    function testResume() public {
        Assert.isFalse(lag.halted(), "lag.halted() should be false");
        lag.halt();
        Assert.isTrue(lag.halted(), "lag.halted() should be true");
        lag.resume();
        Assert.isFalse(lag.halted(), "lag.halted() should be false");
    }

    function testSetPriceRefStorage() public {
        address expectedPriceRefStorage = address(mockIPRS);
        lag.setPriceRefStorage(address(mockIPRS));

        Assert.equal(lag.priceRefStorage(), expectedPriceRefStorage, "lag.setPriceRefStorage() should be 0xb04aD816e86bFA5515c35Ad02081F71D9E848C88");
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

        Assert.equal(lag.halted(), true, "lag.halted() should be true after lag.void() has been called");
    }

    function testIsLagTimePass() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.isLagTimePass.selector));

        Assert.isTrue(r, "lag.isLagTimePass() should not throw an error");
    }

    function testPost() public {
        lag.setPriceRefStorage(address(mockIPRS));
        lag.addConsumer(address(this));

        lag.post();

        (uint256 curr, bool isErr) = lag.getWithError();

        Assert.equal(curr, 0, "curr should be 0");
        Assert.equal(isErr, false, "isErr are should be false");
    }

    function testPostWhenPriceRefNotSet() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.post.selector));

        Assert.equal(r, false, "lag.post() should throw an error");
    }

    function testPostWhenHalted() public {
        lag.halt();
        (bool r,) = address(lag).call(abi.encodePacked(lag.post.selector));

        Assert.equal(r, false, "lag.post() should throw an error");
    }

    function testGetWithError() public {
        lag.addConsumer(address(this));
        (uint256 curr, bool isErr) = lag.getWithError();

        Assert.equal(curr, 0, "lag.getWithError() curr should be 0");
        Assert.isFalse(isErr, "lag.getWithError() isErr should be false");
    }

    function testGetWithErrorWithOutConsumerList() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.getWithError.selector));

        Assert.isFalse(r, "lag.getWithError() should throw an error");
    }

    function testGetNextWithError() public {
        lag.addConsumer(address(this));
        (uint256 curr, bool isErr) = lag.getNextWithError();

        Assert.equal(curr, 0, "lag.getNextWithError() curr should be 0");
        Assert.isFalse(isErr, "lag.getNextWithError() isErr should be false");
    }

    function testGetNextWithErrorWithOutConsumerList() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.getNextWithError.selector));

        Assert.isFalse(r, "lag.getNextWithError() should throw an error");
    }

    function testGet() public {
        lag.addConsumer(address(this));
        uint256 currPrice = lag.get();

        Assert.equal(currPrice, 0, "lag.get() should be 0");
    }

    function testGetAfterPost() public {
        lag.setPriceRefStorage(address(mockIPRS));
        lag.addConsumer(address(this));
        lag.post();

        uint256 currPrice = lag.get();
        Assert.equal(currPrice, 0, "lag.get() should be 0");
    }

    function testGetWithErrorWhenAddressNotInConsumerList() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.get.selector));

        Assert.isFalse(r, "lag.get() should throw an error");
    }

    function testAddConsumer() public {
        lag.addConsumer(address(mockIPRS));

        address[] memory consumers = lag.getConsumers();

        Assert.equal(consumers.length, 1, "consumers.length must be 1");
        Assert.equal(consumers[0], address(mockIPRS), "consumers[0] must be address(this)");
    }

    function testAddConsumerWithError() public {
        bytes4 selector = lag.addConsumer.selector;
        bytes memory data = abi.encodeWithSelector(selector);
        (bool r,) = address(lag).call(data);

        Assert.equal(r, false, "lag.addConsumer() should throw an error");
    }

    function testRemoveConsumer() public {
        lag.addConsumer(address(mockIPRS));
        lag.addConsumer(address(mockIPRS2));
        address[] memory consumers = lag.getConsumers();

        Assert.equal(consumers[0], address(mockIPRS2), "consumers[0] must be address(this)");
        Assert.equal(consumers[1], address(mockIPRS), "consumers[1] must be address(this)");

        lag.removeConsumer(address(mockIPRS), address(mockIPRS2));
        Assert.equal(consumers.length, 2, "consumers.length must be 2");
    }

    function testRemoveConsumerWithErrorWhenNonePrevConsumer() public {
        bytes4 selector = lag.removeConsumer.selector;
        bytes memory data = abi.encodeWithSelector(selector, address(mockIPRS));
        (bool r,) = address(lag).call(data);

        Assert.equal(r, false, "lag.removeConsumer() should throw an error");
    }

    function testGetConsumers() public {
        address[] memory consumers = lag.getConsumers();

        Assert.equal(consumers.length, 0, "consumers.length must be 0");
    }

    function testAddConsumers() public {
        address[] memory newConsumer = new address[](2);
        newConsumer[0] = address(2);
        newConsumer[1] = address(3);

        lag.addConsumers(newConsumer);

        Assert.equal(newConsumer.length, 2, "consumers.length must be 2");
    }
}
