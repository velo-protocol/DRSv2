pragma solidity ^0.5.0;

import "../../contracts/voracle/Med.sol";
import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";

contract TestMed {
    function testInitialize() public {
        Med med = Med(DeployedAddresses.Med());
        med.initialize(address(this), "VELOUSD");

        Assert.equal(med.pair(), "VELOUSD", "pair must be VELOUSD");
    }
}
