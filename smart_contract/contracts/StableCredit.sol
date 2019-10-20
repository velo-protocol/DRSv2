pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/access/roles/WhitelistAdminRole.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

/// @author Velo Team
/// @title A modified ERC20
contract StableCredit is ERC20, ERC20Detailed, WhitelistAdminRole {

    IERC20 public collateral;
    bytes32 public collateralAssetCode;

    uint256 private peggedValue;
    bytes32 private peggedCurrency;

    address public creditOwner;

    constructor (
        address _creditOwner,
        string memory _name,
        string memory _code,
        uint256 _peggedValue,
        bytes32 _peggedCurrency,
        address _collateralAddress,
        bytes32 _collateralAssetCode
    )
    public ERC20Detailed(_name, _code, 7) {
        creditOwner = _creditOwner;
        peggedValue = _peggedValue;
        peggedCurrency = _peggedCurrency;
        collateral = IERC20(_collateralAddress);
        collateralAssetCode = _collateralAssetCode;
    }

    function mint(address recipient, uint256 amount) external onlyWhitelistAdmin {
        _mint(recipient, amount);
    }

    function burn(address tokenOwner, uint256 amount) external onlyWhitelistAdmin {
        _burn(tokenOwner, amount);
    }

    function approveCollateral() external onlyWhitelistAdmin {
        collateral.approve(msg.sender, collateral.balanceOf(address(this)));
    }

    function redeem(address redeemer, uint burnAmount, uint256 returnAmount) external onlyWhitelistAdmin {
        collateral.transfer(redeemer, returnAmount);
        _burn(redeemer, burnAmount);
    }

    function getCollateralDetail() external view returns(uint256, address) {
        return (collateral.balanceOf(address(this)), address(collateral));
    }

    function getCollateralAssetCode() external view returns (bytes32) {
        return collateralAssetCode;
    }

    function getPeggedValue() external view returns (uint256) {
        return peggedValue;
    }

    function getPeggedCurrency() external view returns (bytes32) {
        return peggedCurrency;
    }
}