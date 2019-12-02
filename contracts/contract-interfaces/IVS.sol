pragma solidity ^0.5.0;

interface IVS {
    function getWithError() external view returns (bytes32, bool);
    function get() external view returns (bytes32);
    function set(bytes32 newValue) external;
    function disable() external;

    event LogSetValue(bytes32 oldValue, bytes32 newValue);
}
