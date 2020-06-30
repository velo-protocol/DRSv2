pragma solidity ^0.5.0;

import "../../../contracts/modules/voracle/Price.sol";
import "../mocks/MockIFeeder.sol";
import "truffle/Assert.sol";

contract TestPrice {
    Price public priceContract;
    MockIFeeder public mockFeeder;


    function beforeEach() public {
        mockFeeder = new MockIFeeder();
        priceContract = new Price();
        priceContract.initialize(address(this), address(mockFeeder));
        mockFeeder.commitPrice(100);
    }

    function testPost() public {
         priceContract.post();
        (uint256 price, bool isActive, bool isErr) = priceContract.getWithError();
        Assert.equal(price, 100, "price should be 100");
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
        priceContract.post();
        priceContract.activate();
        (uint256 price, bool isActive, bool isErr) = priceContract.getWithError();
        Assert.equal(price, 100, "price should be 0");
        Assert.equal(isActive, true, "isActive are should be false");
        Assert.equal(isErr, false, "isErr are should be false");
    }

}