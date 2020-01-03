pragma solidity ^0.5.0;

import "../../../contracts/modules/core/Heart.sol";
import "../../../contracts/modules/book-room/LL.sol";
import "../../../contracts/modules/book-room/Hasher.sol";
import "truffle/Assert.sol";

contract TestHeart2 {
    event debug(bytes32 x);
    Heart public heart;

    function beforeEach() public {
        heart = new Heart();
    }

    function testGetAddStableCredit() public {
        uint count = heart.getStableCreditCount();
        Assert.equal(count, uint(0), "heart.getStableCreditCount() must return 0");
        Assert.equal(address(heart.getFirstStableCredit()), address(1), "heart.getFirstStableCredit() must return StableCredit at address 1");

        heart.addStableCredit(Hasher.stableCreditId("vTHB"), StableCredit(10));
        heart.addStableCredit(Hasher.stableCreditId("vUSD"), StableCredit(10));
    }
}
