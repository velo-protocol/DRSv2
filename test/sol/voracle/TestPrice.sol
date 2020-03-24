pragma solidity ^0.5.0;

import "../../../contracts/modules/voracle/Price.sol";
import "../mocks/MockLag.sol";
import "truffle/Assert.sol";

contract TestPrice {
    Price public priceContract;
    MockLag public mockLag;

    function beforeEach() public {
        priceContract = new Price();
        priceContract.initialize(address(this), address(0));
    }

    function testPost() public {
        mockLag = new MockLag(100000000);
        priceContract.setLag(address(mockLag));
        priceContract.post();

        (uint256 price, bool isActive, bool isErr) = priceContract.getWithError();
        Assert.equal(price, 100000000, "price should be 100000000");
        Assert.equal(isActive, true, "isActive are should be false");
        Assert.equal(isErr, false, "isErr are should be false");
    }

    function testGetWithError() public {
        mockLag = new MockLag(100000000);
        priceContract.setLag(address(mockLag));
        priceContract.post();

        (uint256 price, bool isActive, bool isErr) = priceContract.getWithError();
        Assert.equal(price, 100000000, "price should be 100000000");
        Assert.equal(isActive, true, "isActive are should be false");
        Assert.equal(isErr, false, "isErr are should be false");
    }

    function testVoid() public {
        priceContract.void();
        (uint256 price, bool isActive, bool isErr) = priceContract.getWithError();
        Assert.equal(price, 0, "price should be 0");
        Assert.equal(isActive, false, "isActive are should be false");
        Assert.equal(isErr, true, "isErr are should be false");
    }

    function testActivate() public {
        priceContract.void();
        mockLag = new MockLag(100000000);
        priceContract.setLag(address(mockLag));
        priceContract.post();
        priceContract.activate();
        (uint256 price, bool isActive, bool isErr) = priceContract.getWithError();
        Assert.equal(price, 100000000, "price should be 0");
        Assert.equal(isActive, true, "isActive are should be false");
        Assert.equal(isErr, false, "isErr are should be false");
    }

}