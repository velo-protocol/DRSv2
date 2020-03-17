const Lag = artifacts.require('Lag');
const Medianizer = artifacts.require('Medianizer');
const MockContract = artifacts.require("MockContract");

const helper = require('../testhelper');
const abi = require('ethereumjs-abi');

let lag, mocks;

contract("Lag test", async accounts => {

  before(async () => {
    mocks = {
      medianizer: await MockContract.new(),
    };
    medianizer = await Medianizer.at(mocks.medianizer.address);
  });

  beforeEach(async () => {
    lag = await Lag.new(accounts[0], medianizer.address);
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
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], ['100000000', true, false]).toString("hex")
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

    it("should fail, minimum period for the post function has not passed", async () => {

      try {
        await lag.setLagTime(25920000000);
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

      const result = await lag.getWithError();

      const BN = web3.utils.BN;
      const currPrice = new BN(result[0]).toNumber();
      const isActive = result[1];
      const isErr = result[2];

      assert.equal(0, currPrice);
      assert.equal(true, isActive);
      assert.equal(false, isErr);

    });

    it("should get with error successfully with in active flag", async () => {
      await mocks.medianizer.givenMethodReturn(
        helper.methodABI(medianizer, "getWithError"),
        '0x' + abi.rawEncode(['uint256', 'bool', 'bool'], ['100000000', false, false]).toString("hex")
      );
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
});