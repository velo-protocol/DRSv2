pragma solidity ^0.5.0;

import "../../contracts/voracle/Lag.sol";
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
        lag.halt();

        Assert.isTrue(lag.halted(), "lag.halted() are should be true");
    }

    function testHaltWhenNotGov() public {
        lagNoneGov = new Lag(address(1), address(2));
        (bool r,) = address(lagNoneGov).call(abi.encodePacked(lagNoneGov.halt.selector));

        Assert.equal(r, false, "lag.halted() are should be false");
    }

    function testResume() public {
        lag.halt();
        lag.resume();

        Assert.equal(lag.halted(), false, "lag.resume() are should be false");
    }

    function testSetPriceRefStorage() public {
        address expectedPriceRefStorage = address(mockIPRS);
        lag.setPriceRefStorage(address(mockIPRS));

        Assert.equal(lag.priceRefStorage(), expectedPriceRefStorage, "lag.setPriceRefStorage() are should be 0xb04aD816e86bFA5515c35Ad02081F71D9E848C88");
    }

    function testSetLagTime() public {
        bytes4 selector = lag.setLagTime.selector;
        uint16 newLagTime = 60 * 60 * 16;
        bytes memory data = abi.encodeWithSelector(selector, newLagTime);
        (bool r,) = address(lag).call(data);

        Assert.isTrue(r, "lag.setLagTime() are should be true");
    }

    function testSetLagTimeWithError() public {
        bytes4 selector = lag.setLagTime.selector;
        uint16 newLagTime = 0;
        bytes memory data = abi.encodeWithSelector(selector, newLagTime);
        (bool r,) = address(lag).call(data);

        Assert.isFalse(r, "lag.setLagTime() should throw the error");
    }

    function testVoid() public {
        lag.void();

        Assert.equal(lag.halted(), true, "lag.void() variable halted should be true");
    }

    function testIsLagTimePass() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.isLagTimePass.selector));

        Assert.isTrue(r, "lag.isLagTimePass() are should be true");
    }

    function testPost() public {
        lag.setPriceRefStorage(address(mockIPRS));
        lag.addConsumer(address(this));

        lag.post();

        (uint256 curr, bool isErr) = lag.getWithError();

        Assert.equal(curr, 0, "curr are should be 0");
        Assert.equal(isErr, false, "isErr are should be false");
    }

    function testPostWhenPriceRefNotSet() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.post.selector));

        Assert.equal(r, false, "lag.post() should throw the error");
    }

    function testPostWhenHalted() public {
        lag.halt();
        (bool r,) = address(lag).call(abi.encodePacked(lag.post.selector));

        Assert.equal(r, false, "lag.post() must throw error Lag | Lag has been halted");
    }

    function testGetWithError() public {
        lag.addConsumer(address(this));
        (uint256 curr, bool isErr) = lag.getWithError();

        Assert.equal(curr, 0, "lag.getWithError() curr are should be 0");
        Assert.isFalse(isErr, "lag.getWithError() isErr are should be false");
    }

    function testGetWithErrorWithOutConsumerList() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.getWithError.selector));

        Assert.isFalse(r, "lag.getWithError() are should be false because not consumer list");
    }

    function testGetNextWithError() public {
        lag.addConsumer(address(this));
        (uint256 curr, bool isErr) = lag.getNextWithError();

        Assert.equal(curr, 0, "lag.getNextWithError() curr are should be 0");
        Assert.isFalse(isErr, "lag.getNextWithError() isErr are should be false");
    }

    function testGetNextWithErrorWithOutConsumerList() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.getNextWithError.selector));

        Assert.isFalse(r, "lag.getNextWithError() should throw the error");
    }

    function testGet() public {
        lag.addConsumer(address(this));
        uint256 currPrice = lag.get();

        Assert.equal(currPrice, 0, "lag.get() are should be 0");
    }

    function testGetWithErrorWhenAddressNotInConsumerList() public {
        (bool r,) = address(lag).call(abi.encodePacked(lag.get.selector));

        Assert.isFalse(r, "lag.get() should throw error");
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

        Assert.equal(r, false, "lag.addConsumer() should throw the error");
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

        Assert.equal(r, false, "lag.removeConsumer() should throw the error");
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
