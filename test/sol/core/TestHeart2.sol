pragma solidity ^0.5.0;

import "../../../contracts/modules/core/Heart.sol";
import "../../../contracts/modules/book-room/LL.sol";
import "../../../contracts/modules/book-room/Hasher.sol";
import "truffle/Assert.sol";

contract TestHeart2 {
    Heart public heart;

    function helper_addVTHB() public {
        StableCredit vTHB = new StableCredit(
            "THB",
            address(10),
            "VELO",
            address(99),
            "vTHB",
            uint(1),
            address(heart)
        );
        heart.addStableCredit(IStableCredit(address(vTHB)));
    }

    function helper_addVUSD() public {
        StableCredit vUSD = new StableCredit(
            "USD",
            address(10),
            "VELO",
            address(99),
            "vUSD",
            uint(1),
            address(heart)
        );
        heart.addStableCredit(IStableCredit(address(vUSD)));
    }

    function beforeEach() public {
        heart = new Heart();
        heart.setDrsAddress(address(this));
    }

    function testAddStableCredit_Success() public {
        bytes memory data = abi.encodeWithSelector(
            heart.addStableCredit.selector,
            new StableCredit(
                "THB",
                address(10),
                "VELO",
                address(99),
                "vTHB",
                uint(1),
                address(heart)
            )
        );

        (bool result,) = address(heart).call(data);

        Assert.isTrue(result, "heart.addStableCredit must not throw");
        IStableCredit vTHB = heart.getStableCreditById(Hasher.stableCreditId("vTHB"));
        Assert.equal(vTHB.assetCode(), "vTHB", "heart.addStableCredit must add StableCredit correctly");
    }

    function testAddStableCredit_Fail_StableCreditAddressMustNotBeZero() public {
        bytes memory data = abi.encodeWithSelector(
            heart.addStableCredit.selector,
            IStableCredit(address(0))
        );

        (bool result,) = address(heart).call(data);
        Assert.isFalse(result, "heart.addStableCredit must throw an error");
    }

    function testAddStableCredit_Fail_StableCreditHasAlreadyExist() public {
        helper_addVTHB();

        bytes memory data = abi.encodeWithSelector(
            heart.addStableCredit.selector,
            new StableCredit(
                "THB",
                address(10),
                "VELO",
                address(99),
                "vTHB",
                uint(1),
                address(heart)
            )
        );

        (bool result,) = address(heart).call(data);
        Assert.isFalse(result, "heart.addStableCredit must throw an error");
    }

    function testGetStableCreditById_Success() public {
        helper_addVTHB();
        IStableCredit vTHB = heart.getStableCreditById(Hasher.stableCreditId("vTHB"));
        Assert.equal(vTHB.assetCode(), "vTHB", "heart.getStableCreditById must return a correct StableCredit");
    }

    function testGetRecentStableCredit_Success() public {
        helper_addVUSD();
        helper_addVTHB();
        IStableCredit vTHB = heart.getRecentStableCredit();
        Assert.equal(vTHB.assetCode(), "vTHB", "heart.getRecentStableCredit must return a correct StableCredit");
    }

    function testGetNextStableCredit_Success() public {
        helper_addVUSD();
        helper_addVTHB();
        IStableCredit vUSD = heart.getNextStableCredit(heart.getRecentStableCredit().getId());
        Assert.equal(vUSD.assetCode(), "vUSD", "heart.getNextStableCredit must return a correct StableCredit");
    }

    function testGetStableCreditCount_Success() public {
        Assert.equal(heart.getStableCreditCount(), uint(0), "StableCredit count must be 0");
        helper_addVUSD();
        Assert.equal(heart.getStableCreditCount(), uint(1), "StableCredit count must be 1");
        helper_addVTHB();
        Assert.equal(heart.getStableCreditCount(), uint(2), "StableCredit count must be 2");
    }
}
