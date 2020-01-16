pragma solidity ^0.5.0;

import "@openzeppelin/contracts/access/roles/WhitelistAdminRole.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";
import "../interfaces/ICollateralAsset.sol";

/// @author Velo Team
/// @title A modified ERC20
contract Token is ERC20, ERC20Detailed, WhitelistAdminRole {
    constructor (
        string memory _name,
        string memory _code,
        uint8 _decimals
    )
    public ERC20Detailed(_name, _code, _decimals) { }

    function mint(address recipient, uint256 amount) external onlyWhitelistAdmin {
        _mint(recipient, amount);
    }

    function burn(address tokenOwner, uint256 amount) external onlyWhitelistAdmin {
        _burn(tokenOwner, amount);
    }

}
