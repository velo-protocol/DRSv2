pragma solidity ^0.5.0;

import "../mocks/MockTargetProxyV1.sol";
import "../mocks/MockTargetProxyV2.sol";
import "../../../contracts/modules/book-room/BaseProxy.sol";
import "truffle/Assert.sol";

contract TestBaseProxy {
    MockTargetProxyV1 public mockTargetProxyV1;
    MockTargetProxyV2 public mockTargetProxyV2;
    BaseProxy public baseProxy;

    constructor() public {
        mockTargetProxyV1 = new MockTargetProxyV1();
        mockTargetProxyV2 = new MockTargetProxyV2();
        baseProxy = new BaseProxy(address(this));
    }

    function testProxyCallFuncWhenLogicNotSet() public {
        (bool r, ) = address(baseProxy).call(abi.encodePacked(mockTargetProxyV1.get.selector));
        Assert.equal(r, false, "baseProxy should throw the error when the logic contract not set but client try to forward the func");
    }

    function testInitializeWithoutCalldata() public {
        baseProxy.initialize(address(mockTargetProxyV1), "");
        Assert.equal(baseProxy.getLogic(), address(mockTargetProxyV1), "baseProxy.getLogic() must be the same as address(mockTargetProxy)");
        Assert.equal(MockTargetProxyV1(address(baseProxy)).value(), 0, "value as baseProxy context must be 0");

        baseProxy.upgradeTo(address(0));
        Assert.equal(baseProxy.getLogic(), address(0), "baseProxy.getLogic() must be address(0)");
    }

    function testUpgradeToSameLogic() public {
        bytes4 selector = baseProxy.upgradeTo.selector;
        bytes memory data = abi.encodeWithSelector(selector, address(0));
        (bool r, ) = address(baseProxy).call(data);

        Assert.equal(r, false, "baseProxy should throw the error when owner try to upgradeTo the same Logic");
    }

    function testInitializeWithMalformatCalldata() public {
        bytes4 selector = baseProxy.initialize.selector;
        bytes memory data = abi.encodeWithSelector(selector, address(mockTargetProxyV2), "velogogo");
        (bool r, ) = address(baseProxy).call(data);

        Assert.equal(r, false, "baseProxy should throw the error when the calldata is not right");
    }

    function testInitialize() public {
        baseProxy.initialize(address(mockTargetProxyV1), abi.encodePacked(mockTargetProxyV1.initialize.selector));
        Assert.equal(baseProxy.getLogic(), address(mockTargetProxyV1), "baseProxy.getLogic() must be the same as address(mockTargetProxy)");
        Assert.equal(MockTargetProxyV1(address(baseProxy)).value(), 10, "value as baseProxy context must be 10");
    }

    function testDoubleInitialize() public {
        bytes4 selector = baseProxy.initialize.selector;
        bytes memory data = abi.encodeWithSelector(selector, address(1), "velogogo");
        (bool r, ) = address(baseProxy).call(data);

        Assert.equal(r, false, "baseProxy should throw the error when owner try to double initialize");
    }

    function testGetLogic() public {
        Assert.equal(baseProxy.getLogic(), address(mockTargetProxyV1), "baseProxy.getLogic() must be the same as address(mockTargetProxy)");
    }

    function testCallLogicV1Function() public {
        MockTargetProxyV1(address(baseProxy)).post(100);
        Assert.equal(MockTargetProxyV1(address(baseProxy)).value(), 100, "value as baseProxy context must be 100");
    }

    function testUpgradeTo() public {
        baseProxy.upgradeTo(address(mockTargetProxyV2));
        Assert.equal(baseProxy.getLogic(), address(mockTargetProxyV2), "baseProxy.getLogic() must be the same as address(mockTargetProxy)");
        Assert.equal(MockTargetProxyV1(address(baseProxy)).value(), 100, "value as baseProxy context must be 10");
    }

    function testCallLogicV2Function() public {
        MockTargetProxyV2(address(baseProxy)).post(5);
        Assert.equal(MockTargetProxyV1(address(baseProxy)).value(), 10, "value as baseProxy context must be 100");
    }

    function testTransferOwnershipToAddressZero() public {
        bytes4 selector = baseProxy.transferOwnership.selector;
        bytes memory data = abi.encodeWithSelector(selector, address(0));
        (bool r, ) = address(baseProxy).call(data);

        Assert.equal(r, false, "baseProxy should throw the error when owner try to transferOwnership(address(0))");
    }

    function testTransferOwnership() public {
        baseProxy.transferOwnership(address(1));
        Assert.equal(baseProxy.getProxyOwner(), address(1), "baseProxy.getProxyOwner() must be address(0)");
    }

    function testNonProxyOwnerTryToUpgrade() public {
        bytes4 selector = baseProxy.upgradeTo.selector;
        bytes memory data = abi.encodeWithSelector(selector, address(3));
        (bool r, ) = address(baseProxy).call(data);

        Assert.equal(r, false, "baseProxy should throw the error");
    }
}
