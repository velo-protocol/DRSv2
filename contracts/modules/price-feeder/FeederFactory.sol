pragma solidity ^0.5.0;

import "./Feeder.sol";

contract FeederFactory {
    event FeederCreated(address indexed owner, address indexed feedAddr);

    function create(bytes32 fiatCode, bytes32 collateralCode) public returns (address) {
        Feeder feeder = new Feeder(msg.sender, fiatCode, collateralCode);
        emit FeederCreated(msg.sender, address(feeder));
        return address(feeder);
    }
}
