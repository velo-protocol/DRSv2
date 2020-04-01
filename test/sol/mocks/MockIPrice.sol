pragma solidity ^0.5.0;

import "../../../contracts/modules/interfaces/IPrice.sol";

contract MockIPrice is IPrice {

    function post() external {}

    function get() external view returns (uint256){
        return 1;
    }

    function getWithError() external view returns (uint256, bool, bool){
        return (1, true, false);
    }

    function void() external {}

    function activate() external {}
}