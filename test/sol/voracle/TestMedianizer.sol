pragma solidity ^0.5.0;

import "../../../contracts/modules/voracle/Medianizer.sol";
import "../mocks/MockIFeeder.sol";
import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";

contract TestMedianizer {
    MockIFeeder public mockIFeeder1;
    MockIFeeder public mockIFeeder2;

    constructor() public {
        mockIFeeder1 = new MockIFeeder(100);
        mockIFeeder2 = new MockIFeeder(120);
    }

    function testInitialize() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.initialize(address(this), "USD", "VELO");
        Assert.equal(med.fiatCode(), "USD", "fiatCode must be USD");
        Assert.equal(med.collateralCode(), "VELO", "fiatCode must be VELO");
    }

    function testGetWhenPriceIsLessThanZero() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        (bool r,) = address(med).call(abi.encodePacked(med.get.selector));

        Assert.equal(r, false, "med.get() must throw error");
    }

    function testGetWithErrorWhenPriceIsLessThanZero() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        (uint256 price, bool isActive, bool isErr) = med.getWithError();

        Assert.equal(price, 0, "med.getWithError() must return price = 0");
        Assert.equal(isActive, true, "med.getWithError() must return isActive = true");
        Assert.equal(isErr, true, "med.getWithError() must return isErr = true");
    }

    function testSetMinFedPrices() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.setMinFedPrices(2);
        Assert.equal(uint(med.minFedPrices()), uint(2), "minFedPrices must be 2");

        med.setMinFedPrices(1);
        Assert.equal(uint(med.minFedPrices()), uint(1), "minFedPrices must be 1");
    }

    function testAddFeeder() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        med.addFeeder(address(mockIFeeder1));
        med.addFeeder(address(mockIFeeder2));

        address[] memory feeders = med.getFeeders();

        Assert.equal(feeders[0], address(mockIFeeder2), "feeder[0] must be address(mockIPRS2)");
        Assert.equal(feeders[1], address(mockIFeeder1), "feeder[1] must be address(mockIPRS1)");
        Assert.equal(feeders.length, 2, "feeders.length must be 2");
    }

    function testRemoveFeeder() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        address[] memory feeders = med.getFeeders();
        Assert.equal(feeders.length, 2, "feeders.length must be 2");

        med.removeFeeder(address(mockIFeeder1));
        feeders = med.getFeeders();
        Assert.equal(feeders.length, 1, "feeders.length must be 1");

        med.removeFeeder(address(mockIFeeder2));
        feeders = med.getFeeders();
        Assert.equal(feeders.length, 0, "feeders.length must be 0");
    }

    function testComputeOddFeeder() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        med.addFeeder(address(mockIFeeder1));
        (uint256 median, bool isErr) = med.compute();

        Assert.equal(median, 100, "median must be 100");
        Assert.equal(isErr, false, "isErr must be false");

        med.removeFeeder(address(mockIFeeder1));
        address[] memory feeders = med.getFeeders();

        Assert.equal(feeders.length, 0, "feeders.length must be 0");
    }

    function testComputeEvenFeeder() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        med.addFeeder(address(mockIFeeder1));
        med.addFeeder(address(mockIFeeder2));

        (uint256 median, bool isErr) = med.compute();

        Assert.equal(median, 110, "median must be 110");
        Assert.equal(isErr, false, "isErr must be false");

        testRemoveFeeder();
    }

    function testPost() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        med.addFeeder(address(mockIFeeder1));
        med.addFeeder(address(mockIFeeder2));

        mockIFeeder1.post(130);

        med.post();

        Assert.equal(med.price(), 125, "med.price() must be 125");

        testRemoveFeeder();
    }

    function testGet() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        uint256 medianPrice = med.get();

        Assert.equal(medianPrice, 125, "med.get() should return 125");
    }

    function testGetWithError() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());

        (uint256 medianPrice, bool isActive, bool isErr) = med.getWithError();

        Assert.equal(medianPrice, 125, "med.getWithError() should return 125");
        Assert.equal(isActive, true, "med.getWithError() should return true");
        Assert.equal(isErr, false, "med.getWithError() should return false");
    }

    function testGetWithErrorWhenActiveIsFalse() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.addFeeder(address(mockIFeeder1));
        med.addFeeder(address(mockIFeeder2));

        mockIFeeder1.post(130);

        med.void();
        med.post();

        (uint256 price, bool isActive, bool isErr) = med.getWithError();

        Assert.equal(price, 125, "med.getWithError() must return price = 125");
        Assert.equal(isActive, false, "med.getWithError() must return isActive = false");
        Assert.equal(isErr, false, "med.getWithError() must return isErr = false");
    }

    function testSetValidityPeriod() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.setValidityPeriod(5 minutes);
        Assert.equal(med.validityPeriod(), 300, "validityPeriod must be 300 seconds");
    }

    function testGetValidityPeriod() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.setValidityPeriod(25);
        Assert.equal(med.getValidityPeriod(), 25, "validityPeriod must be 25");
    }

    function testVoid() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.void();
        Assert.equal(med.active(), false, "active must be false");
        Assert.equal(med.price(), 0, "price must be 0");
    }

    function testActivate() public {
        Medianizer med = Medianizer(DeployedAddresses.Medianizer());
        med.post();
        med.activate();
        Assert.equal(med.active(), true, "active must be true");
    }
}
