pragma solidity ^0.5.0;

import "../../../contracts/modules/book-room/LL.sol";

contract TestLLrevert {
    LL.List public cats;
    using LL for LL.List;

    address constant CAT1 = 0xb04aD816e86bFA5515c35Ad02081F71D9E848C88;
    address constant CAT2 = 0x633eC0587Aee0cFEC883cAB73cEf822230C54d3a;

    function initLL() public {
        cats = cats.init();
    }

    function llDupAddRevert() public {
        cats = cats.add(CAT1);
        cats = cats.add(CAT1);
    }

    function llRemoveNoExistRevert() public {
        cats = cats.remove(CAT2, CAT1);
    }

    function llRemoveWrongPreRevert() public {
        cats = cats.add(CAT2);
        cats = cats.remove(CAT2, CAT1);
    }
}
