pragma solidity ^0.5.0;

import "../../../contracts/modules/core/Heart.sol";
import "../mocks/MockIRM.sol";
import "../mocks/MockIERC20.sol";
import "../mocks/MockIPrice.sol";
import "truffle/Assert.sol";
import "../../../contracts/modules/book-room/Hasher.sol";

contract TestHeart1 {

    Heart public heart;
    MockIRM public mockIRM;
    MockIERC20 public mockIERC20;
    MockIPrice public mockIPrice;

    function beforeEach() public {
        heart = new Heart();
    }

    constructor() public {
        mockIPrice = new MockIPrice();
        mockIRM = new MockIRM();
        mockIERC20 = new MockIERC20();
    }

    function testConstructor_Success() public {
        Assert.equal(heart.getStableCreditCount(), uint(0), "StableCredit count must be 0");
        Assert.equal(address(heart.getRecentStableCredit()), address(1), "getRecentStableCredit() must return address 0");
    }

    function testSetReserveManager_Success() public {
        address expected = address(mockIRM);
        heart.setReserveManager(expected);

        Assert.equal(address(heart.getReserveManager()), expected, "heart.setReserveManager() should set reserve manager properly");
    }

    function testGetReserveManager_Success() public {
        Assert.equal(address(heart.getReserveManager()), address(0), "heart.getReserveManager() should return address(0)");

        address expected = address(mockIRM);
        heart.setReserveManager(expected);

        Assert.equal(address(heart.getReserveManager()), expected, "heart.getReserveManager() should return reserveManager correctly");
    }

    function testSetReserveFreeze_Success() public {
        bytes32 assetCode = "VELO";
        uint32 newSeconds = 15;
        heart.setReserveFreeze(assetCode, newSeconds);

        uint32 reserveFreezeSecond = heart.getReserveFreeze(assetCode);

        Assert.equal(uint(reserveFreezeSecond), uint(newSeconds), "heart.setReserveFreeze() should set reserve freeze seconds properly");
    }

    function testGetReserveFreeze_Success() public {
        bytes32 assetCode = "VELO";
        uint32 newSeconds = 15;

        Assert.equal(heart.getReserveFreeze(assetCode), uint(0), "heart.getReserveFreeze() should return reserve freeze second");

        heart.setReserveFreeze(assetCode, newSeconds);
        uint32 reserveFreezeSecond = heart.getReserveFreeze(assetCode);

        Assert.equal(uint(reserveFreezeSecond), uint(newSeconds), "heart.getReserveFreeze() should return reserve freeze second");
    }

    function testSetDrsAddress_Success() public {
        address expected = address(10);
        heart.setDrsAddress(expected);

        Assert.equal(heart.getDrsAddress(), expected, "heart.setDrsAddress() should set drs address properly");
    }

    function testGetDrsAddress_Success() public {
        Assert.equal(heart.getDrsAddress(), address(0), "heart.getDrsAddress() should return address 0");

        address expected = address(11);
        heart.setDrsAddress(expected);

        Assert.equal(heart.getDrsAddress(), expected, "heart.getDrsAddress() should return drs address correctly");
    }

    function testSetCollateralAsset_Success() public {
        bytes32 assetCode = "VELO";
        address mockIERC20Address = address(10);
        uint ratio = 1;
        heart.setCollateralAsset(assetCode, mockIERC20Address, ratio);

        address collateralAssetAddress = address(heart.getCollateralAsset(assetCode));
        uint collateralRatios = heart.getCollateralRatio(assetCode);

        Assert.equal(collateralAssetAddress, mockIERC20Address, "heart.setCollateralAsset() should set collateral address properly");
        Assert.equal(collateralRatios, ratio, "heart.setCollateralAsset() should set collateral ratio properly");
    }

    function testGetCollateralAsset_Success() public {
        bytes32 assetCode = "VELO";
        address mockIERC20Address = address(10);
        uint ratio = 1;
        heart.setCollateralAsset(assetCode, mockIERC20Address, ratio);

        address collateralAssetAddress = address(heart.getCollateralAsset(assetCode));

        Assert.equal(collateralAssetAddress, mockIERC20Address, "heart.getCollateralAsset() should return collateral asset correctly");
    }

    function testSetCollateralRatio_Success() public {
        bytes32 assetCode = "VELO";
        address mockIERC20Address = address(10);
        uint ratio = 2;
        heart.setCollateralAsset(assetCode, mockIERC20Address, 1);
        heart.setCollateralRatio(assetCode, ratio);
        uint collateralRatios = heart.getCollateralRatio(assetCode);

        Assert.equal(collateralRatios, ratio, "heart.setCollateralRatio() should set collateral ratio properly");
    }

    function testSetCollateralRatio_Fail_WithErrorAssetCodeHasNotBeenAdded() public {
        bytes4 selector = heart.setCollateralRatio.selector;
        bytes32 assetCode = "VELO";
        uint ratio = 2;
        bytes memory data = abi.encodeWithSelector(selector, assetCode, ratio);
        (bool result,) = address(heart).call(data);

        Assert.isFalse(result, "heart.setCollateralRatio() should throw an error assetCode has not been added");
    }

    function testGetCollateralRatio_Success() public {
        bytes32 assetCode = "VELO";
        address mockIERC20Address = address(10);
        uint ratio = 2;
        heart.setCollateralAsset(assetCode, mockIERC20Address, ratio);
        uint collateralRatios = heart.getCollateralRatio(assetCode);

        Assert.equal(collateralRatios, ratio, "heart.getCollateralRatio() should return collateral ratio correctly");
    }

    function testSetCreditIssuanceFee_Success() public {
        uint256 expectedNewFee = 1;
        heart.setCreditIssuanceFee(expectedNewFee);

        Assert.equal(heart.getCreditIssuanceFee(), expectedNewFee, "heart.setCreditIssuanceFee() should set credit issuance fee properly");
    }

    function testGetCreditIssuanceFee_Success() public {
        uint256 fee = 1;
        heart.setCreditIssuanceFee(fee);

        uint256 expectedCreditIssuanceFee = 1;
        heart.getCreditIssuanceFee();

        Assert.equal(heart.getCreditIssuanceFee(), expectedCreditIssuanceFee, "heart.getCreditIssuanceFee() should return credit issuance fee correctly");
    }

    function testSetTrustedPartner_Success() public {
        Assert.isFalse(heart.isTrustedPartner(address(1)), "heart.setTrustedPartner() from address(1) should be false");
        heart.setTrustedPartner(address(1));

        Assert.isTrue(heart.isTrustedPartner(address(1)), "heart.setTrustedPartner() from address(1) should be true");
    }

    function testIsTrustedPartner_Success() public {
        heart.setTrustedPartner(address(1));

        Assert.isTrue(heart.isTrustedPartner(address(1)), "heart.isTrustedPartner() from address(1) should be true");
    }

    function testAddPrice_Success() public {
        bytes32 linkId = Hasher.linkId("VELO", "USD");
        heart.addPrice(linkId, mockIPrice);

        Assert.equal(address(heart.getPriceContract(linkId)), address(mockIPrice), "heart.addPrice() should add price properly");
    }

    function testGetPrice_Success() public {
        bytes32 linkId = Hasher.linkId("VELO", "USD");
        heart.addPrice(linkId, mockIPrice);

        Assert.equal(address(heart.getPriceContract(linkId)), address(mockIPrice), "heart.getPriceContract() should return price correctly");
    }

    function testCollectFee_Success() public {
        heart.setDrsAddress(address(this));

        uint256 expectedFee = 1;
        bytes32 expectedCollateralAsset = "VELO";

        bytes4 selector = heart.collectFee.selector;
        bytes memory data = abi.encodeWithSelector(selector, expectedFee, expectedCollateralAsset);
        (bool result,) = address(heart).call(data);

        Assert.isTrue(result, "heart.collectFee() should not throw any error");
        Assert.equal(heart.getCollectedFee(expectedCollateralAsset), uint(1), "heart.collectFee() should calculate collected fee correctly");
    }

    function testCollectFee_Fail_WhenNotDrsAddress() public {
        uint256 expectedFee = 1;
        bytes32 expectedCollateralAsset = "VELO";

        bytes4 selector = heart.collectFee.selector;
        bytes memory data = abi.encodeWithSelector(selector, expectedFee, expectedCollateralAsset);
        (bool result,) = address(heart).call(data);

        Assert.isFalse(result, "heart.collectFee() should throw error only DRSSC can update the collected fee");
    }

    function testGetCollectedFee_Success() public {
        uint256 expectedFee = 1;
        bytes32 expectedCollateralAsset = "VELO";

        heart.setDrsAddress(address(this));
        heart.collectFee(expectedFee, expectedCollateralAsset);
        heart.getCollectedFee(expectedCollateralAsset);

        Assert.equal(heart.getCollectedFee(expectedCollateralAsset), expectedFee, "heart.getCollectedFee() should return collected fee correctly");
    }

    function testWithdrawFee_Success() public {
        uint256 expectedCollectedFee = 20;
        uint256 expectedAmount = 1;
        uint256 expectedBalanceCollectedFee = 19;
        bytes32 expectedCollateralAsset = "VELO";

        address mockIERC20Address = address(mockIERC20);
        uint ratio = 1;

        heart.setDrsAddress(address(this));
        heart.setCollateralAsset(expectedCollateralAsset, mockIERC20Address, ratio);
        heart.setCollateralRatio(expectedCollateralAsset, ratio);
        heart.collectFee(expectedCollectedFee, expectedCollateralAsset);
        Assert.equal(heart.getCollectedFee(expectedCollateralAsset), expectedCollectedFee, "heart.getCollectedFee() should be collectedFee 20");


        bytes4 selector = heart.withdrawFee.selector;
        bytes memory data = abi.encodeWithSelector(selector, expectedCollateralAsset, expectedAmount);
        (bool result,) = address(heart).call(data);
        uint256 balanceCollectedFee = heart.getCollectedFee(expectedCollateralAsset);

        Assert.isTrue(result, "heart.withdrawFee() should not throw any error");
        Assert.equal(balanceCollectedFee, expectedBalanceCollectedFee, "heart.withdrawFee() should withdraw fee correctly");
    }

    function testWithdrawFee_Fail_WithAmountMoreThanCollectedFee() public {
        uint256 expectedFee = 1;
        uint256 expectedAmount = 2;
        bytes32 expectedCollateralAsset = "VELO";

        heart.setDrsAddress(address(this));
        address mockIERC20Address = address(mockIERC20);
        uint ratio = 1;

        heart.setDrsAddress(address(this));
        heart.setCollateralAsset(expectedCollateralAsset, mockIERC20Address, ratio);
        heart.setCollateralRatio(expectedCollateralAsset, ratio);
        heart.collectFee(expectedFee, expectedCollateralAsset);

        bytes4 selector = heart.withdrawFee.selector;
        bytes memory data = abi.encodeWithSelector(selector, expectedCollateralAsset, expectedAmount);
        (bool result,) = address(heart).call(data);

        Assert.isFalse(result, "heart.withdrawFee() should throw error when amount more than collectedFee");
    }
}
