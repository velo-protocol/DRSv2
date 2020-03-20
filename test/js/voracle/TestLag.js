const Lag = artifacts.require('Lag');
const Medianizer = artifacts.require('Medianizer');
const MockContract = artifacts.require("MockContract");

const helper = require('../testhelper');
const abi = require('ethereumjs-abi');
const truffleAssert = require('truffle-assertions');

let lag, mocks;

contract("Lag test", async accounts => {

  before(async () => {
    mocks = {
      medianizer: await MockContract.new(),
    };
    medianizer = await Medianizer.at(mocks.medianizer.address);
  });

  beforeEach(async () => {
    lag = await Lag.new();
    await lag.initialize(accounts[0], medianizer.address);
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
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, false]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, nextPrice);
      assert.equal(true, isActive);
      assert.equal(false, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = true isErr = true", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, true, true]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = true isErr = false", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, true, false]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = false isErr = true", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, false, true]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 0 isActive = false isErr = false", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [0, false, false]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 10 isActive = true isErr = true", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [10, true, true]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 10 isActive = false isErr = true", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [10, false, true]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should post successfully, getWithError return price = 10 isActive = false isErr = false", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [10, false, false]).toString("hex")
      );
      await lag.post();

      const result = await lag.getNextWithError();
      const BN = web3.utils.BN;
      const nextPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, nextPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

    it("should fail, minimum period for the post function has not passed", async () => {

      try {
        await mocks.medianizer.givenMethodReturn(
          helper.methodABI(medianizer, "getWithError"),
          '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, false]).toString("hex")
        );
        await lag.post();
        await lag.setMinimumPeriod(2592000);
        await lag.post();
      } catch (err) {
        assert.equal(err.reason, "Lag.post: minimum period for the post function has not passed")
      }

    });
  });

  describe("GetWithError", async () => {
    it("should get with error successfully", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], ['100000000', true, false]).toString("hex")
      );
      await lag.post();
      await lag.setMinimumPeriod(0);
      await lag.post();

      const result = await lag.getWithError();

      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(100000000, currPrice);
      assert.equal(true, isActive);
      assert.equal(false, isErr);

    });

    it("should get with error successfully, when initial Lag contract", async () => {
      const result = await lag.getWithError();

      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(true, isActive);
      assert.equal(true, isErr);

    });

    it("should get with error successfully with inactive flag", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], ['100000000', false, false]).toString("hex")
      );
      await lag.post();
      await lag.setMinimumPeriod(1);
      await lag.post();

      const result = await lag.getWithError();

      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(false, isActive);
      assert.equal(true, isErr);

    });

  });

  describe("Void", async () => {
    it("should void successfully", async () => {
      const voidResult = await lag.void();

      truffleAssert.eventEmitted(voidResult, 'LagVoid', event => {
        return event.caller === accounts[0]
          && event.lag === lag.address
          && event.isActive === false
      }, 'contract should emit the event correctly');

    });

    it("should fail, sender is not owner", async () => {
      try {
        await lag.void({from: accounts[1]});
      } catch (err) {
        assert.equal(err.reason, "Lag.onlyOwner: The message sender is not found or does not have sufficient permission")
      }

    });

  });

  describe("Activate", async () => {
    it("should activate successfully", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], [100000000, true, false]).toString("hex")
      );
      await lag.post();
      await lag.setMinimumPeriod(0);
      await lag.void();
      await lag.post();

      const activateResult = await lag.activate();

      truffleAssert.eventEmitted(activateResult, 'LagActivate', event => {
        return event.caller === accounts[0]
          && event.lag === lag.address
          && event.isActive === true
      }, 'contract should emit the event correctly');

    });

    it("should fail, lag is active", async () => {

      try {
        await lag.activate();

      } catch (err) {
        assert.equal(err.reason, "Lag.activate: lag is active")
      }

    });

    it("should fail, price is not in a correct state", async () => {

      try {
        await lag.void();
        await lag.activate();

      } catch (err) {
        assert.equal(err.reason, "Lag.activate: price is not in a correct state")
      }

    });
  });

  describe("SetMinimumPeriod", async () => {
    it("should set minimum period successfully, with minimum period = 0", async () => {

      await lag.setMinimumPeriod(0);

      minimumPeriod = (await lag.minimumPeriod()).toString();

      assert.equal(0, minimumPeriod);
    });

    it("should set minimum period successfully, with minimum period = 2592000", async () => {

      await lag.setMinimumPeriod(2592000);

      minimumPeriod = (await lag.minimumPeriod()).toString();

      assert.equal(2592000, minimumPeriod);
    });

    it("should fail, minimum period < 0", async () => {
      try {
        await lag.setMinimumPeriod(-1);
      } catch (err) {
        assert.equal(err.reason, "Lag.setMinimumPeriod: minimumPeriod value must not be less than 0")
      }
    });

    it("should fail, minimum period > 2592000", async () => {
      try {
        await lag.setMinimumPeriod(2592001);
      } catch (err) {
        assert.equal(err.reason, "Lag.setMinimumPeriod: minimumPeriod value must not be greater than 2592000")
      }
    });

  });
});