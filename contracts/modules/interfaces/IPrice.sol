pragma solidity ^0.5.0;

contract IPrice {
    function post() external;

    function get() external view returns (uint256);

    function getWithError() external view returns (uint256, bool, bool);

    function void() external;

    function activate() external;
}
