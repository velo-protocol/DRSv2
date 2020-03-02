pragma solidity ^0.5.0;

interface IFeeder {
    function getWithError() external view returns (uint256, bool);
    function get() external view returns (uint256);
    function set(uint256 newValue) external;
    function enable() external;
    function disable() external;

    event PriceSet(uint256 prevValue, uint prevTimestamp, uint curPrice, uint curTimestamp);
}
