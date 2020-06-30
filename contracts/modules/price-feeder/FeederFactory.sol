pragma solidity ^0.5.0;

import "./Feeder.sol";

contract FeederFactory {
    event FeederCreated(address indexed owner, address indexed feedAddr);

    function create(address f1,address f2 ,address f3,bytes32 fiatCode, bytes32 collateralCode) public returns (address) {
        Feeder feeder = new Feeder(f1,f2,f3,msg.sender, fiatCode, collateralCode);
        emit FeederCreated(msg.sender, address(feeder));
        return address(feeder);
    }
}
