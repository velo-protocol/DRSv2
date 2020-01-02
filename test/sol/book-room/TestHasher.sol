pragma solidity ^0.5.0;

import "../../../contracts/modules/book-room/Hasher.sol";
import "truffle/Assert.sol";

contract TestHasher {
    Hasher public hasher;

    constructor() public {
        hasher = new Hasher();
    }

    function testLinkId() public {
        bytes32 expected = 0x4cb01440ae0bc757f71beb1faa542cd6a1aba2477a2b38aa35a7eee3283535e9;
        bytes32 collateralAssetCode = "VELO";
        bytes32 peggedCurrency = "vTHB";
        bytes32 result = hasher.linkId(collateralAssetCode, peggedCurrency);

        Assert.equal(result,expected, "hasher.linkId() result should be 0x4cb01440ae0bc757f71beb1faa542cd6a1aba2477a2b38aa35a7eee3283535e9");
    }

    function testLockedReserveId() public {
        bytes32 expected = 0x4afd1c8d71908c0a77df9b5973780be5f71f53c6a2ceb7f102746ef30eb5c9df;
        address from = 0x51800B4E5acD84BcA3b1a2be1d07158E35851Cb3;
        bytes32 collateralAssetCode = "VELO";
        uint256 collateralAmount = 100;
        uint blockNumber = 1;
        bytes32 result = hasher.lockedReserveId(from, collateralAssetCode, collateralAmount, blockNumber);

        Assert.equal(result,expected, "Hasher lockedReserveId() result should be 0x4afd1c8d71908c0a77df9b5973780be5f71f53c6a2ceb7f102746ef30eb5c9df");
    }

    function testStableCreditId() public {
        bytes32 expected = 0x82d8c2282adba9ac0ab3f1a555a37692290b0598488f7fc4dfc2e2f8991ecbaf;
        address creditOwner = 0x51800B4E5acD84BcA3b1a2be1d07158E35851Cb3;
        string memory assetCode = "vTHB";
        bytes32 result = hasher.stableCredit(creditOwner, assetCode);

        Assert.equal(result, expected, "Hasher stableCredit() result should be 0x82d8c2282adba9ac0ab3f1a555a37692290b0598488f7fc4dfc2e2f8991ecbaf");
    }

}
