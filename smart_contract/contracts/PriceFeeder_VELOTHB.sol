pragma solidity ^0.5.0;

import "./oraclize/usingOraclize.sol";
import "./DigitalReserveSystem.sol";

contract PriceFeeder is usingOraclize {

    address public digitalReserveSystemAddr;

    event LogConstructorInitiated(string nextStep);
    event LogPriceUpdated(bytes32 id, string price);
    event LogNewOraclizeQuery(string description);

    constructor(address _digitalReserveSystemAddr) public payable {
        OAR = OraclizeAddrResolverI(0x845aBB4B697Dc2049Db1fA0FD71aDe29f268010E);
        digitalReserveSystemAddr = _digitalReserveSystemAddr;

        emit LogConstructorInitiated("`PriceFeeder` contract was initiated.");
    }

    function __callback(bytes32 id, string memory result) public {
        if (msg.sender != oraclize_cbAddress()) revert();

        DigitalReserveSystem(digitalReserveSystemAddr).setPrice("THB", parseInt(result, 7));

        emit LogPriceUpdated(id, result);
    }

    function updatePrice(uint delay) public payable {
        emit LogNewOraclizeQuery("oraclize query was sent, standing by for the answer..");
        oraclize_query(delay, "URL", "json(https://dev-internal-system.lightnetapis.io/velo-fx-rates?convert=THB).data.\"1\".quote.THB.price");
    }

}
