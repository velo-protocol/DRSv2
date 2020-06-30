pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IFeeder.sol";

contract MockIFeeder is IFeeder {
    uint  public value;
    function commitPrice(uint priceInWei) external  returns (bool){
        value=priceInWei;
        return true;
    }
    function startOracle(uint priceInWei) external  returns (bool){
        value=priceInWei;
        return true;
    }
    function getLastPrice() external view returns(uint, uint){
        return (value,1);
    }
}