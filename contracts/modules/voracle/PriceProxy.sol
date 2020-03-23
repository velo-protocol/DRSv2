pragma solidity ^0.5.0;

import "../book-room/BaseProxy.sol";

contract PriceProxy is BaseProxy {
    constructor(address gov) public BaseProxy(gov) {}
}
