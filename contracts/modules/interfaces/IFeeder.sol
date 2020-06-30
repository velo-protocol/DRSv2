pragma solidity ^0.5.0;

interface IFeeder {
    function commitPrice(uint priceInWei) external  returns (bool);
    function startOracle(uint priceInWei) external  returns (bool);
    function getLastPrice() external view returns(uint, uint);
    event CommitPrice(uint indexed priceInWei, uint indexed timeInSecond, address sender, uint index);
    event AcceptPrice(uint indexed priceInWei, uint indexed timeInSecond, address sender);
}
