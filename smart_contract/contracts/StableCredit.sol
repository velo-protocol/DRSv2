pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";

/// @author Velo Team
/// @title A modified ERC20
contract StableCredit is ERC20, ERC20Detailed, WhitelistAdminRole {

    uint256 public collateralAmount;
    bytes32 public collateralAssetCode;

    uint256 private peggedValue;
    bytes32 private peggedCurrency;

    constructor (
        string memory _name,
        string memory _code,
        uint8 _decimals,
        uint256 _peggedValue,
        bytes32 _peggedCurrency
    )
    public ERC20Detailed(_name, _code, _decimals) {
        peggedValue = _peggedValue;
        peggedCurrency = _peggedCurrency;
    }

    function mint(address recipient, uint256 amount) external onlyWhitelistAdmin {
        _mint(recipient, amount);
    }

    function burn(address tokenOwner, uint256 amount) external onlyWhitelistAdmin {
        _burn(tokenOwner, amount);
    }

    function getCollateralDetail() external view returns(uint256, bytes32) {
        return (collateralAmount, collateralAssetCode);
    }

    function getPeggedValue() external view returns (uint256) {
        return peggedValue;
    }

    function getPeggedCurrency() external view returns (bytes32) {
        return peggedCurrency;
    }
}