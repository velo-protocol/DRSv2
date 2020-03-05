pragma solidity ^0.5.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";
import "../book-room/Hasher.sol";
import "../interfaces/IHeart.sol";
import "../interfaces/ICollateralAsset.sol";

/// @author Velo Team
/// @title A modified ERC20
contract StableCredit is ERC20, ERC20Detailed {
    IHeart public heart;

    ICollateralAsset public collateral;
    bytes32 public collateralAssetCode;

    uint256 public peggedValue;
    bytes32 public peggedCurrency;

    address public creditOwner;
    address public drsAddress;

    modifier onlyDRSSC() {
        require(heart.getDrsAddress() == msg.sender, "caller is not DRSSC");
        _;
    }

    constructor (
        bytes32 _peggedCurrency,
        address _creditOwner,
        bytes32 _collateralAssetCode,
        address _collateralAddress,
        string memory _code,
        uint256 _peggedValue,
        address heartAddr
    )
    public ERC20Detailed(_code, _code, 7) {
        creditOwner = _creditOwner;
        peggedValue = _peggedValue;
        peggedCurrency = _peggedCurrency;
        collateral = ICollateralAsset(_collateralAddress);
        collateralAssetCode = _collateralAssetCode;
        heart = IHeart(heartAddr);
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

    function getCollateralDetail() external view returns (uint256, address) {
        return (collateral.balanceOf(address(this)), address(collateral));
    }

    function getId() external view returns (bytes32) {
        return Hasher.stableCreditId(this.name());
    }

    function assetCode() external view returns (string memory) {
        return this.name();
    }

    function transferCollateralToReserve(uint256 amount) external onlyDRSSC returns (bool) {
        ICollateralAsset(collateral).transfer(address(heart.getReserveManager()), amount);
        return true;
    }

}
