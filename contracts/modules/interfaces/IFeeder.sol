pragma solidity ^0.5.0;

interface IFeeder {
    function getWithError() external view returns (uint256, bool);
    function get() external view returns (uint256);
    function set(uint256 newValue) external;
    function disable() external;

    event LogSetValue(uint256 oldValue, uint256 newValue);
}
