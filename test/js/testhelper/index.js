module.exports = {
  address(addressNumber) {
    return web3.utils.padLeft("0x" + addressNumber, 40);
  },
  decimal7(float) {
    return Math.trunc(float * Math.pow(10, 7));
  },
  methodABI(contractInstance, methodName, methodArgs = []) {
    return contractInstance.contract.methods[methodName](...methodArgs).encodeABI();
  },
  assertEqualByteString(actualBytes, expectedString) {
    const bytesLength = actualBytes.length - 2;
    assert.equal(actualBytes, web3.utils.padRight(web3.utils.utf8ToHex(expectedString), bytesLength));
  },
  assert: {
    equalByteString(actualBytes, expectedString) {
      const bytesLength = actualBytes.length - 2;
      assert.equal(actualBytes, web3.utils.padRight(web3.utils.utf8ToHex(expectedString), bytesLength));
    },
    equalNumber(actualNumber, expectedNumber) {
      assert.equal(actualNumber.toString(), expectedNumber.toString());
    },
    async throwsWithReason(asyncFn, reason) {
      let result;
      try {
        result = await asyncFn();
      } catch (err) {
        assert.equal(err.reason, reason);
      } finally {
        assert.ok(typeof result === 'undefined');
      }
    },
    async throwsWithMessage(asyncFn, message) {
      let result;
      try {
        result = await asyncFn();
      } catch (err) {
        assert.equal(err.message, message);
      } finally {
        assert.ok(typeof result === 'undefined');
      }
    },
  }
};
