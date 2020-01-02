pragma solidity ^0.5.0;

import "../../../contracts/modules/book-room/Hasher.sol";
import "truffle/Assert.sol";

contract TestHasher {
    Hasher public hasher;

    constructor() public {
        hasher = new Hasher();
    }

    function testLinkId() public {
        bytes4 selector = hasher.linkId.selector;
        bytes32 collateralAssetCode = 'VELO';
        bytes32 peggedCurrency = 'THB';
        bytes memory data = abi.encodeWithSelector(selector, collateralAssetCode, peggedCurrency);
        (bool result,) = address(hasher).call(data);

        Assert.isTrue(result, "hasher.linkId() should not throw an error");
    }

    function testLinkIdWithError() public {
        bytes4 selector = hasher.linkId.selector;
        bytes32 collateralAssetCode = bytes32('VELO');
        bytes memory data = abi.encodeWithSelector(selector, collateralAssetCode);
        (bool result,) = address(hasher).call(data);

        Assert.isFalse(result, "hasher.linkId() should throw an error");
    }

    function testHasherCallFunctionLockedReserveId() public {
        bytes4 selector = hasher.lockedReserveId.selector;
        address from = address(1);
        bytes32 collateralAssetCode = "VELO";
        uint256 collateralAmount = 100;
        uint blockNumber = 1;
        bytes memory data = abi.encodeWithSelector(selector, from, collateralAssetCode, collateralAmount, blockNumber);
        (bool result,) = address(hasher).call(data);

        Assert.equal(result, true, "Hasher lockedReserveId() should be true");
    }

    function testHasherLockedReserveIdWithErrorNoneParameters() public {
        bytes4 selector = hasher.lockedReserveId.selector;
        bytes memory data = abi.encodeWithSelector(selector);
        (bool result,) = address(hasher).call(data);

        Assert.equal(result, false, "Hasher lockedReserveId() should throw an error");
    }
}
