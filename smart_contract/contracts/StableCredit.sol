pragma solidity ^0.5.0;

import "./openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "./openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

/// @author Velo Team
/// @title A modified ERC20
contract StableCredit is ERC20, ERC20Detailed {

    IERC20 public collateral;
    bytes12 public collateralAssetCode;

    uint256 public peggedValue;
    bytes12 public peggedCurrency;

    address public creditOwner;
    address public drsAddress;

    modifier onlyDRSSC() {
        require(drsAddress == msg.sender, "caller is not DRSSC");
        _;
    }

    constructor (
        address _creditOwner,
        string memory _code,
        uint256 _peggedValue,
        bytes12 _peggedCurrency,
        address _collateralAddress,
        bytes12 _collateralAssetCode
    )
    public ERC20Detailed(_code, _code, 7) {
        drsAddress = msg.sender;
        creditOwner = _creditOwner;
        peggedValue = _peggedValue;
        peggedCurrency = _peggedCurrency;
        collateral = IERC20(_collateralAddress);
        collateralAssetCode = _collateralAssetCode;
    }

    function mint(address recipient, uint256 amount) external onlyDRSSC {
        _mint(recipient, amount);
    }

    function burn(address tokenOwner, uint256 amount) external onlyDRSSC {
        _burn(tokenOwner, amount);
    }

    function approveCollateral() external onlyDRSSC {
        collateral.approve(msg.sender, collateral.balanceOf(address(this)));
    }

    function redeem(address redeemer, uint burnAmount, uint256 returnAmount) external onlyDRSSC {
        collateral.transfer(redeemer, returnAmount);
        _burn(redeemer, burnAmount);
    }

    function getCollateralDetail() external view returns(uint256, address) {
        return (collateral.balanceOf(address(this)), address(collateral));
    }
}