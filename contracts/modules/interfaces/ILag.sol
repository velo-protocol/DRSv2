pragma solidity ^0.5.0;

contract ILag {
    function post() external returns (bool);

    function getWithError() external view returns (uint256, bool, bool);

    function void() external;

    function activate() external;

    function setMinimumPeriod(int256) external;

}
