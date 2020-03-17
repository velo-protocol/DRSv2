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
      const isErr = result[1];

      assert.equal(100000000, nextPrice);
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

});