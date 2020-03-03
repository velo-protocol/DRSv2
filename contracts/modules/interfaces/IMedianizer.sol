pragma solidity ^0.5.0;

contract IMedianizer {
    function post() external;
    function getWithError() external view returns (uint256, bool, bool);
}
