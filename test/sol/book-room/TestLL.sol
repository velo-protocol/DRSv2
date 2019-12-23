pragma solidity ^0.5.0;

import "../../../contracts/modules/book-room/LL.sol";
import "truffle/Assert.sol";

contract TestLL {
    LL.List public cats;
    using LL for LL.List;

    address constant CAT1 = 0xb04aD816e86bFA5515c35Ad02081F71D9E848C88;
    address constant CAT2 = 0x633eC0587Aee0cFEC883cAB73cEf822230C54d3a;

    function testInit() public {
        cats = cats.init();
        address expected = address(1);
        Assert.equal(cats.next[address(1)], expected,
            "next address should be address(1)");
    }

    function testAdd() public {
        cats = cats.add(CAT1);
        Assert.equal(cats.next[address(1)], CAT1,
            "next address should be CAT1");

        cats = cats.add(CAT2);
        Assert.equal(cats.next[address(1)], CAT2,
            "next address should be CAT2");
        Assert.equal(cats.next[CAT1], address(1),
            "next address should be 0x1");
    }

    function testHas() public {
        bool isHas = cats.has(CAT1);
        Assert.equal(isHas, true, "should has CAT1");

        isHas = cats.has(CAT2);
        Assert.equal(isHas, true, "should has CAT2");

        isHas = cats.has(address(3));
        Assert.equal(isHas, false, "should not has address(3)");
    }

    function testRemove() public {
        cats = cats.remove(CAT2, address(1));
        Assert.equal(cats.next[address(1)], CAT1,
            "next address should be CAT1");
    }

    function testGetAll() public {
        cats = cats.add(CAT2);
        Assert.equal(cats.next[address(1)], CAT2,
            "next address should be CAT2");
        Assert.equal(cats.next[CAT1], address(1),
            "next address should be 0x1");

        address[] memory allCats = cats.getAll();
        Assert.equal(allCats[0], CAT2, "index[0] must be CAT2");
        Assert.equal(allCats[1], CAT1, "index[1] must be CAT1");
    }
}
