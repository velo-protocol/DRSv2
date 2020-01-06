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
        Assert.equal(heart.getStableCreditCount(), uint(0), "heart.getStableCreditCount() must return 0");
        Assert.equal(address(heart.getRecentStableCredit()), address(1), "heart.getFirstStableCredit() must return StableCredit at address 1");

        heart.addStableCredit(new StableCredit(
                "THB",
                address(10),
                "VELO",
                address(99),
                "vTHB",
                uint(1),
                address(heart)
            ));
        heart.addStableCredit(new StableCredit(
                "USD",
                address(11),
                "VELO",
                address(99),
                "vUSD",
                uint(1),
                address(heart)
            ));

        Assert.equal(heart.getStableCreditCount(), uint(2), "heart.getStableCreditCount() must return 2");

        StableCredit vUSD = heart.getRecentStableCredit();
        Assert.equal(vUSD.name(), "vUSD", "heart.getFirstStableCredit() must return a correct StableCredit");

        StableCredit vTHB = heart.getNextStableCredit(vUSD.getId());
        Assert.equal(vTHB.name(), "vTHB", "heart.getNextStableCredit() must return a correct StableCredit");

    }


}
