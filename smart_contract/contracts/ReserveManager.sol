pragma solidity ^0.5.0;

//import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
//import "./DigitalReserveSystem.sol";
//import "./VeloToken.sol";

//contract ReserveManager is WhitelistAdminRole {

//    DigitalReserveSystem private drs;
//    VeloToken private veloToken;
//
//    address public drsAddr;
//
//    constructor(address digitalReserveSystemAddr, address veloTokenAddr) public {
//        drs = DigitalReserveSystem(digitalReserveSystemAddr);
//        drsAddr = digitalReserveSystemAddr;
//
//        veloToken = VeloToken(veloTokenAddr);
//
//        _addWhitelistAdmin(digitalReserveSystemAddr);
//        _removeWhitelistAdmin(msg.sender);
//    }
//
//    event InjectCollateral(address indexed trustedPartnerAddress, bytes32 indexed assetName, uint256 increasedAmount);
//
//    function injectCollateral(address trustedPartnerAddress, bytes32 assetName, uint256 amount) public onlyWhitelistAdmin {
//        veloToken.transfer(drsAddr, amount);
//
//        emit InjectCollateral(trustedPartnerAddress, assetName, amount);
//    }

//}
