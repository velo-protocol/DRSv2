module.exports = {
    address(addressNumber) {
        return web3.utils.padLeft("0x" + addressNumber, 40)
    },
    methodABI(contractInstance, methodName, methodArgs = []) {
        return contractInstance.contract.methods[methodName](...methodArgs).encodeABI();
    },
    assertEqualByteString(actualBytes, expectedString) {
        const bytesLength = actualBytes.length - 2;
        assert.equal(actualBytes, web3.utils.padRight(web3.utils.utf8ToHex(expectedString), bytesLength));
    }
};
