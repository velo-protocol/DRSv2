pragma solidity ^0.5.0;

import "./Feeds.sol";

contract FeedsFactory {
    event LogFeedCreated(address indexed owner, address indexed feedAddr);

    function create() public returns (address) {
        Feeds feed = new Feeds(msg.sender);
        emit LogFeedCreated(msg.sender, address(feed));
        return address(feed);
    }
}
