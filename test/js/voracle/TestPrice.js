
const Price = artifacts.require('Price');
const Feeder = artifacts.require('Feeder');
const MockContract = artifacts.require("MockContract");

const helper = require('../testhelper');
const abi = require('ethereumjs-abi');
const truffleAssert = require('truffle-assertions');

let price, mocks;

contract("Price test", async accounts => {
    before(async () => {
        mocks = {
            feeder: await MockContract.new(),
        };
        feeder = await Feeder.at(mocks.feeder.address);
    });
    beforeEach(async () => {
    price = await Price.new();
    await price.initialize(accounts[0],feeder.address);
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
      await mocks.feeder.givenMethodReturn(
        helper.methodABI(feeder, "getLastPrice"),
        '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
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


  });

  describe("GetWithError", async () => {
    it("should get with error successfully", async () => {
        await mocks.feeder.givenMethodReturn(
            helper.methodABI(feeder, "getLastPrice"),
            '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
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

  describe("Get", async () => {
    it("should get successfully", async () => {
        await mocks.feeder.givenMethodReturn(
            helper.methodABI(feeder, "getLastPrice"),
            '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
        );
      await price.post();
      const BN = web3.utils.BN;
      const result = await price.get();
      const currPrice = new BN(result).toNumber();

      assert.equal(100000000, currPrice);
    });

    it("should fail, when initial Price contract", async () => {
      try {
        await price.get();
      } catch (err) {
        assert.equal(err, "Error: Returned error: VM Exception while processing transaction: revert Price.get: invalid price")
      }
    });

    it("should fail, active = false", async () => {
        await mocks.feeder.givenMethodReturn(
            helper.methodABI(feeder, "getLastPrice"),
            '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
        );
      await price.post();

      try {
        await price.get();
      } catch (err) {
        assert.equal(err, "Error: Returned error: VM Exception while processing transaction: revert Price.get: price is not active")
      }
    });

    it("should fail, isErr = true", async () => {
        await mocks.feeder.givenMethodReturn(
            helper.methodABI(feeder, "getLastPrice"),
            '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
        );
      await price.post();

      try {
        await price.get();
      } catch (err) {
        assert.equal(err, "Error: Returned error: VM Exception while processing transaction: revert Price.get: price has an error")
      }
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

      const result = await price.getWithError();

      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);


    });

    it("should void successfully, price is 100000000", async () => {

        await mocks.feeder.givenMethodReturn(
            helper.methodABI(feeder, "getLastPrice"),
            '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
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

  describe("Activate", async () => {
    it("should activate successfully", async () => {
        await mocks.feeder.givenMethodReturn(
            helper.methodABI(feeder, "getLastPrice"),
            '0x' + abi.rawEncode(['uint','uint' ], [100000000, 0]).toString("hex")
        );

      await price.void();
      await price.post();

      const activateResult = await price.activate();

      truffleAssert.eventEmitted(activateResult, 'PriceActivate', event => {
        return event.caller === accounts[0]
          && event.price === price.address
          && event.isActive === true
      }, 'contract should emit the event correctly');

    });

    it("should fail, price is active", async () => {

      try {
        await price.activate();

      } catch (err) {
        assert.equal(err.reason, "Price.activate: price is active")
      }

    });

    it("should fail, price is not in a correct state", async () => {

      try {
        await price.void();
        await price.activate();

      } catch (err) {
        assert.equal(err.reason, "Price.activate: price is not in a correct state")
      }

    });

  });
});
