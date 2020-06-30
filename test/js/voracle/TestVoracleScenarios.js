const FeederFactory = artifacts.require('FeederFactory');
const Feeder = artifacts.require('Feeder');
const Price = artifacts.require('Price');
const PriceProxy = artifacts.require('PriceProxy');
const h = require("../testhelper");

const veloBytes32 = web3.utils.fromAscii("VELO");
const usdBytes32 = web3.utils.fromAscii("USD");

contract("Voracle Scenario Test", async accounts => {
  it("should work!", async () => {

    const [deployer, pf1, pf2, pf3, pf4, pf5] = accounts;

    const feederFactory = await FeederFactory.new();

    let result = await feederFactory.create( pf1, pf2, pf3,usdBytes32, veloBytes32, {from: pf1});
    const usdFeederAddr1 = result.logs[0].args.feedAddr;

    result = await feederFactory.create( pf1, pf2, pf3,usdBytes32, veloBytes32, {from: pf2});
    const usdFeederAddr2 = result.logs[0].args.feedAddr;

    result = await feederFactory.create( pf1, pf2, pf3,usdBytes32, veloBytes32, {from: pf3});
    const usdFeederAddr3 = result.logs[0].args.feedAddr;

    const usdFeeder1 = await Feeder.at(usdFeederAddr1);
    const usdFeeder2 = await Feeder.at(usdFeederAddr2);
    const usdFeeder3 = await Feeder.at(usdFeederAddr3);
    const feedPrice1 = await usdFeeder1.startOracle(30000000, {from: pf1}); // 3.0 USD/VELO
    h.assert.equalNumber(feedPrice1.logs[0].args[0], 30000000);
    let feedResult1 = await usdFeeder1.getLastPrice();
    feedResult1 = {
      price: feedResult1['0']
    };
    h.assert.equalNumber(feedResult1.price, 30000000);
    const feedPrice2 = await usdFeeder2.startOracle(18000000, {from: pf2}); // 1.8 USD/VELO
    h.assert.equalNumber(feedPrice2.logs[0].args[0], 18000000);
    //
    let feedResult2 = await usdFeeder2.getLastPrice();
    feedResult2 = {
      price: feedResult2['0']
    };
    h.assert.equalNumber(feedResult2.price, 18000000);
    const feedPrice3 = await usdFeeder3.startOracle(15000000, {from: pf3}); // 1.5 USD/VELO
    h.assert.equalNumber(feedPrice3.logs[0].args[0], 15000000);
    let feedResult3 = await usdFeeder3.getLastPrice();
    feedResult3 = {
      price: feedResult3['0']
    };
    h.assert.equalNumber(feedResult3.price, 15000000);

  });
});
