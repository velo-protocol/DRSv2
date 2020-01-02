pragma solidity ^0.5.0;

contract Hasher {

    function linkId(bytes32 collateralAssetCode, bytes32 peggedCurrency) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(collateralAssetCode, peggedCurrency));
    }

    function lockedReserveId(address from, bytes32 collateralAssetCode, uint256 collateralAmount, uint blockNumber) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(from, collateralAssetCode, collateralAmount, blockNumber));
    }

    function stableCredit(address creditOwner, string memory assetCode) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(creditOwner, assetCode));
    }
}
