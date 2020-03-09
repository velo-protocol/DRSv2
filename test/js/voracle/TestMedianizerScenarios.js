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

    console.log('usdFeeder1', usdFeeder1.address);
    console.log('usdFeeder2', usdFeeder2.address);
    console.log('usdFeeder3', usdFeeder3.address);
    console.log('usdFeeder4', usdFeeder4.address);
    console.log('usdFeeder5', usdFeeder5.address);

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

    const feederAddresses = await medUSD.getFeeders();
    assert.equal(feederAddresses.length, 5);
    assert.ok(feederAddresses.includes(usdFeeder1.address));
    assert.ok(feederAddresses.includes(usdFeeder2.address));
    assert.ok(feederAddresses.includes(usdFeeder3.address));
    assert.ok(feederAddresses.includes(usdFeeder4.address));
    assert.ok(feederAddresses.includes(usdFeeder5.address));


    const feedPrice1 = await usdFeeder1.post(30000000, {from: pf1}); // 3.0 USD/VELO
    assert.equal(feedPrice1.logs[0].args.prevValue, 0);
    assert.equal(feedPrice1.logs[0].args.curPrice, 30000000);


    let feedResult1 = await usdFeeder1.getWithError();
    feedResult1 = {
      price: feedResult1['0'],
      isErr: feedResult1['1']
    };
    assert.equal(feedResult1.price.toString(), 30000000);
    // assert.equal(feedResult.isErr,false);

    feedPrice2 = await usdFeeder2.post(18000000, {from: pf2}); // 1.8 USD/VELO
    assert.equal(feedPrice2.logs[0].args.prevValue, 0);
    assert.equal(feedPrice2.logs[0].args.curPrice, 18000000);

    let feedResult2 = await usdFeeder2.getWithError();
    feedResult2 = {
      price: feedResult2['0'],
      isErr: feedResult2['1']
    };
    assert.equal(feedResult2.price.toString(), 18000000);

    feedPrice3 = await usdFeeder3.post(15000000, {from: pf3}); // 1.5 USD/VELO
    assert.equal(feedPrice3.logs[0].args.prevValue, 0);
    assert.equal(feedPrice3.logs[0].args.curPrice, 15000000);

    let feedResult3 = await usdFeeder3.getWithError();
    feedResult3 = {
      price: feedResult3['0'],
      isErr: feedResult3['1']
    };
    assert.equal(feedResult3.price.toString(), 15000000);


    feedPrice4 = await usdFeeder4.post(45000000, {from: pf4}); // 4.5 USD/VELO
    assert.equal(feedPrice4.logs[0].args.prevValue, 0);
    assert.equal(feedPrice4.logs[0].args.curPrice, 45000000);

    let feedResult4 = await usdFeeder4.getWithError();
    feedResult4 = {
      price: feedResult4['0'],
      isErr: feedResult4['1']
    };
    assert.equal(feedResult4.price.toString(), 45000000);

    feedPrice5 = await usdFeeder5.post(29000000, {from: pf5}); // 2.9 USD/VELO
    assert.equal(feedPrice5.logs[0].args.prevValue, 0);
    assert.equal(feedPrice5.logs[0].args.curPrice, 29000000);

    let feedResult5 = await usdFeeder5.getWithError();
    feedResult5 = {
      price: feedResult5['0'],
      isErr: feedResult5['1']
    };
    assert.equal(feedResult5.price.toString(), 29000000);


    const postResult = await medUSD.post();
    h.assert.equalNumber(postResult.logs[0].args.price.toString(), 29000000);

    h.assert.equalNumber(await medUSD.get(), 29000000);

  });
});