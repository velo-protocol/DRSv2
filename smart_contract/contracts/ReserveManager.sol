pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./IRM.sol";
import "./IGOV.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

contract ReserveManager is IRM {

    address public drsAddress;
    IGOV public gov;

    modifier onlyDRSSC() {
        require(drsAddress == msg.sender, "caller is not DRSSC");
        _;
    }

    constructor(address digitalReserveSystemAddr, address governanceAddr) public {
        drsAddress = digitalReserveSystemAddr;
        gov = IGOV(governanceAddr);
    }

    event InjectCollateral(bytes32 indexed assetCode, uint256 amount);

    function injectCollateral(bytes12 assetCode, address to, uint256 amount) external onlyDRSSC {
        gov.getCollateralAsset(assetCode).transfer(to, amount);

        emit InjectCollateral(assetCode, amount);
    }

}
