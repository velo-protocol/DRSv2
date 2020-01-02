pragma solidity ^0.5.0;

contract Hasher {

    function getLinkId(bytes32 collateralAssetCode, bytes32 peggedCurrency) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));
    }
}
