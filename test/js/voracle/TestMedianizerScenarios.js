const FeederFactory = artifacts.require('FeederFactory');
const Feeder = artifacts.require('Feeder');
const MedianizerProxy = artifacts.require('MedianizerProxy');
const Medianizer = artifacts.require('Medianizer');
const h = require("../testhelper");

const veloBytes32 = web3.utils.fromAscii("VELO");
const usdBytes32 = web3.utils.fromAscii("USD");

contract("Medianizer Scenario Test", async accounts => {
  it("should work!", async () => {

    const [deployer, pf1, pf2, pf3, pf4, pf5] = accounts;

    const feederFactory = await FeederFactory.new();

    let result = await feederFactory.create(usdBytes32, veloBytes32, {from: pf1});
    const usdFeederAddr1 = result.logs[0].args.feedAddr;

    result = await feederFactory.create(usdBytes32, veloBytes32, {from: pf2});
    const usdFeederAddr2 = result.logs[0].args.feedAddr;

    result = await feederFactory.create(usdBytes32, veloBytes32, {from: pf3});
    const usdFeederAddr3 = result.logs[0].args.feedAddr;

    result = await feederFactory.create(usdBytes32, veloBytes32, {from: pf4});
    const usdFeederAddr4 = result.logs[0].args.feedAddr;

    result = await feederFactory.create(usdBytes32, veloBytes32, {from: pf5});
    const usdFeederAddr5 = result.logs[0].args.feedAddr;

    const usdFeeder1 = await Feeder.at(usdFeederAddr1);
    const usdFeeder2 = await Feeder.at(usdFeederAddr2);
    const usdFeeder3 = await Feeder.at(usdFeederAddr3);
    const usdFeeder4 = await Feeder.at(usdFeederAddr4);
    const usdFeeder5 = await Feeder.at(usdFeederAddr5);

    // Medianizer
    const medLogic = await Medianizer.new();

    const medProxyUSD = await MedianizerProxy.new(medLogic.address);

    const medInstanceUSD = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeUSDCalldata = await medInstanceUSD.methods.initialize(deployer, usdBytes32, veloBytes32).encodeABI();
    await medProxyUSD.initialize(medLogic.address, initializeUSDCalldata);

    const medUSD = await Medianizer.at(medProxyUSD.address);
    await medUSD.addFeeder(usdFeeder1.address);
    await medUSD.addFeeder(usdFeeder2.address);
    await medUSD.addFeeder(usdFeeder3.address);
    await medUSD.addFeeder(usdFeeder4.address);
    await medUSD.addFeeder(usdFeeder5.address);

    await usdFeeder1.post(25000000, {from: pf1}); // 2.5 USD/VELO
    await usdFeeder2.post(26000000, {from: pf2}); // 2.6 USD/VELO
    await usdFeeder3.post(27000000, {from: pf3}); // 2.7 USD/VELO
    await usdFeeder4.post(28000000, {from: pf4}); // 2.8 USD/VELO
    await usdFeeder5.post(29000000, {from: pf5}); // 2.9 USD/VELO

    await medUSD.post();

    h.assert.equalNumber(await medUSD.get(), 27000000);
  });
});