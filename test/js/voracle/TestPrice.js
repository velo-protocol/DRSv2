const Lag = artifacts.require('Lag');
const Price = artifacts.require('Price');
const MockContract = artifacts.require("MockContract");

const helper = require('../testhelper');
const abi = require('ethereumjs-abi');

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

    });
  });


});