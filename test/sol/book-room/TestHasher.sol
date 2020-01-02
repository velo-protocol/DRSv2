pragma solidity ^0.5.0;

import "../../../contracts/modules/book-room/Hasher.sol";
import "truffle/Assert.sol";

contract TestHasher {
    Hasher public hasher;

    constructor() public {
        hasher = new Hasher();
    }

    function testGetLinkId() public {
        bytes4 selector = hasher.getLinkId.selector;
        bytes32 collateralAssetCode = 'VELO';
        bytes32 peggedCurrency = 'THB';
        bytes memory data = abi.encodeWithSelector(selector, collateralAssetCode, peggedCurrency);
        (bool result,) = address(hasher).call(data);

        Assert.isTrue(result, "hasher.getLinkId() should not throw an error");
    }

    function testGetLinkIdWithError() public {
        bytes4 selector = hasher.getLinkId.selector;
        bytes32 collateralAssetCode = bytes32('VELO');
        bytes memory data = abi.encodeWithSelector(selector, collateralAssetCode);
        (bool result,) = address(hasher).call(data);

        Assert.isFalse(result, "hasher.getLinkId() should throw an error");
    }

}
