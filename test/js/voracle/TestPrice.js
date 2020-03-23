const Lag = artifacts.require('Lag');
const Price = artifacts.require('Price');
const MockContract = artifacts.require("MockContract");

const helper = require('../testhelper');
const abi = require('ethereumjs-abi');
const truffleAssert = require('truffle-assertions');

let price, mocks;

contract("Price test", async accounts => {

  before(async () => {
    mocks = {
      lag: await MockContract.new(),
    };
    lag = await Lag.at(mocks.lag.address);
  });

  beforeEach(async () => {
    price = await Price.new();
    await price.initialize(accounts[0], lag.address);
  });

  afterEach(async () => {
    const promises = [];
    for (const key in mocks) {
      promises.push(mocks[key].reset());
    }
    await Promise.all(promises);
  });

  describe("Post", async () => {
    it("should post successfully", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, false]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, currPrice);
      assert.equal(true, isActive);
      assert.equal(false, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = true isErr = true", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, true, true]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = true isErr = false", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, true, false]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = false isErr = true", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, false, true]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = false isErr = false", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, false, false]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 100000000 isActive = true isErr = true", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, true]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, currPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 100000000 isActive = false isErr = true", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, false, true]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, currPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 100000000 isActive = false isErr = false", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, false, false]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, currPrice);
      assert.equal(false, isActive);
      assert.equal(false, isErr);

    });

  });


  describe("GetWithError", async () => {
    it("should get with error successfully", async () => {
      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, false]).toString("hex")
      );
      await price.post();

      const result = await price.getWithError();
      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, currPrice);
      assert.equal(true, isActive);
      assert.equal(false, isErr);
    });

    it("should get with error successfully, when initial Price contract", async () => {
      const result = await price.getWithError();

      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });
  });

  describe("Void", async () => {
    it("should void successfully", async () => {
      const voidResult = await price.void();

      truffleAssert.eventEmitted(voidResult, 'PriceVoid', event => {
        return event.caller === accounts[0]
          && event.price === price.address
          && event.isActive === false
      }, 'contract should emit the event correctly');

    });

    it("should void successfully, price is 100000000", async () => {

      await mocks.lag.givenMethodReturn(
        helper.methodABI(lag, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, false]).toString("hex")
      );
      await price.post();

      const voidResult = await price.void();

      truffleAssert.eventEmitted(voidResult, 'PriceVoid', event => {
        return event.caller === accounts[0]
          && event.price === price.address
          && event.isActive === false
      }, 'contract should emit the event correctly');

    });

    it("should fail, sender is not owner", async () => {
      try {
        await price.void({from: accounts[1]});
      } catch (err) {
        assert.equal(err.reason, "Price.onlyOwner: The message sender is not found or does not have sufficient permission")
      }

    });

  });
});
