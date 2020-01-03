pragma solidity ^0.5.0;

import "../../../contracts/modules/core/Heart.sol";
import "../mocks/MockIRM.sol";
import "../mocks/MockIERC20.sol";
import "truffle/Assert.sol";

contract TestHeart {

    Heart public heart;
    MockIRM public irm;
    MockIERC20 public ierc20;

    function beforeEach() public {
        heart = new Heart();
    }

    constructor() public {
        irm = new MockIRM();
        ierc20 = new MockIERC20();
    }

    function testSetReserveManager() public {
        address expected = address(irm);
        heart.setReserveManager(expected);

        Assert.equal(address(heart.getReserveManager()), expected, "heart.setReserveManager() should set reserve manager properly");
    }

    function testGetReserveManager() public {
        address expected = address(irm);
        heart.setReserveManager(expected);

        Assert.equal(address(heart.getReserveManager()), expected, "heart.getReserveManager() should get reserve manager properly");
    }

    function testSetReserveFreeze() public {
        bytes32 assetCode = "VELO";
        uint32 newSeconds = 15;
        heart.setReserveFreeze(assetCode, newSeconds);

        uint32 reserveFreezeSecond = heart.getReserveFreeze(assetCode);

        Assert.equal(uint(reserveFreezeSecond), uint(newSeconds), "heart.setReserveFreeze() should set reserve freeze seconds properly");
    }

    function testGetReserveFreeze() public {
        bytes32 assetCode = "VELO";
        bytes32 assetCode2 = "BTC";
        uint32 newSeconds = 15;
        heart.setReserveFreeze(assetCode, newSeconds);
        heart.setReserveFreeze(assetCode2, newSeconds);

        uint32 reserveFreezeSecond = heart.getReserveFreeze(assetCode);

        Assert.equal(uint(reserveFreezeSecond), uint(newSeconds), "heart.getReserveFreeze() should get reserve freeze seconds properly");
    }

    function testSetDrsAddress() public {
        address expected = 0xBe40F54aDc30b5e2d2b7CEAc2Fc600e786c49B81;
        heart.setDrsAddress(expected);

        Assert.equal(heart.getDrsAddress(), expected, "heart.setDrsAddress() should set drs address properly");
    }

    function testGetDrsAddress() public {
        address firstAddress = 0xBe40F54aDc30b5e2d2b7CEAc2Fc600e786c49B81;
        address expected = 0x4b8c3fa65CE99A0067c5716ae886328e0A0E86cf;
        heart.setDrsAddress(firstAddress);
        heart.setDrsAddress(expected);

        Assert.equal(heart.getDrsAddress(), expected, "heart.getDrsAddress() should get drs address properly");
    }

    function testSetCollateralAsset() public {
        bytes32 assetCode = "VELO";
        address ierc20Address = 0xBe40F54aDc30b5e2d2b7CEAc2Fc600e786c49B81;
        uint ratio = 1;
        heart.setCollateralAsset(assetCode, ierc20Address, ratio);

        address collateralAssetAddress = address(heart.getCollateralAsset(assetCode));
        uint collateralRatios = heart.getCollateralRatio(assetCode);

        Assert.equal(collateralAssetAddress, ierc20Address, "heart.setCollateralAsset() should set with ierc20 address properly");
        Assert.equal(collateralRatios, ratio, "heart.setCollateralAsset() should set with ratio properly");
    }

    function testGetCollateralAsset() public {
        bytes32 assetCode = "VELO";
        address ierc20Address = 0xBe40F54aDc30b5e2d2b7CEAc2Fc600e786c49B81;
        uint ratio = 1;
        heart.setCollateralAsset(assetCode, ierc20Address, ratio);

        address collateralAssetAddress = address(heart.getCollateralAsset(assetCode));

        Assert.equal(collateralAssetAddress, ierc20Address, "heart.getCollateralAsset() should get with collateral asset properly");
    }

    function testSetCollateralRatio() public {
        bytes32 assetCode = "VELO";
        address ierc20Address = 0xBe40F54aDc30b5e2d2b7CEAc2Fc600e786c49B81;
        uint ratio = 2;
        heart.setCollateralAsset(assetCode, ierc20Address, 1);
        heart.setCollateralRatio(assetCode, ratio);

        uint collateralRatios = heart.getCollateralRatio(assetCode);

        Assert.equal(collateralRatios, ratio, "heart.setCollateralRatio() should set properly");
    }

    function testSetCollateralRatioWithErrorAssetCodeHasNotBeenAdded() public {
        bytes4 selector = heart.setCollateralRatio.selector;
        bytes32 assetCode = "VELO";
        uint ratio = 2;
        bytes memory data = abi.encodeWithSelector(selector, assetCode, ratio);

        (bool result,) = address(heart).call(data);

        Assert.isFalse(result, "heart.setCollateralRatio() should throw an error assetCode has not been added");

    }

    function testGetCollateralRatio() public {
        bytes32 assetCode = "VELO";
        address ierc20Address = 0xBe40F54aDc30b5e2d2b7CEAc2Fc600e786c49B81;
        uint ratio = 2;
        heart.setCollateralAsset(assetCode, ierc20Address, ratio);

        uint collateralRatios = heart.getCollateralRatio(assetCode);

        Assert.equal(collateralRatios, ratio, "heart.getCollateralRatio() should get properly");
    }
}
