const FeederFactory = artifacts.require('FeederFactory');
const Feeder = artifacts.require('Feeder');
const MedianizerProxy = artifacts.require('MedianizerProxy');
const Medianizer = artifacts.require('Medianizer');
const Lag = artifacts.require('Lag');
const LagProxy = artifacts.require('LagProxy');
const Price = artifacts.require('Price');
const PriceProxy = artifacts.require('PriceProxy');

module.exports = async (deployer, network, accounts) => {
  await deployer.deploy(FeederFactory);

  if (network === 'development' || network === 'local' || network === 'dev') {
    const feederFactory = await FeederFactory.deployed();

    const veloBytes32 = web3.utils.fromAscii("VELO");
    const usdBytes32 = web3.utils.fromAscii("USD");
    const thbBytes32 = web3.utils.fromAscii("THB");
    const sgdBytes32 = web3.utils.fromAscii("SGD");

    // 1. Create 3 feeders
    let result = await feederFactory.create(usdBytes32, veloBytes32);
    const usdFeederAddr = result.logs[0].args.feedAddr;
    console.log(`USD Feeder: ${usdFeederAddr}`);

    result = await feederFactory.create(thbBytes32, veloBytes32);
    const thbFeederAddr = result.logs[0].args.feedAddr;
    console.log(`THB Feeder: ${thbFeederAddr}`);

    result = await feederFactory.create(sgdBytes32, veloBytes32);
    const sgdFeederAddr = result.logs[0].args.feedAddr;
    console.log(`SGD Feeder: ${sgdFeederAddr}`);

    // 2. Feed the price
    const usdFeeder = await Feeder.at(usdFeederAddr);
    await usdFeeder.post(12000000); // 1.2 USD/VELO

    const thbFeeder = await Feeder.at(thbFeederAddr);
    await thbFeeder.post(53000000); // 5.3 THB/VELO

    const sgdFeeder = await Feeder.at(sgdFeederAddr);
    await sgdFeeder.post(27000000); // 2.7 THB/VELO

    await deployer.deploy(Medianizer);
    const medLogic = await Medianizer.deployed();

    // Medianizer USD
    await deployer.deploy(MedianizerProxy, accounts[0]);
    const medProxyUSD = await MedianizerProxy.deployed();

    const medInstanceUSD = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeUSDCalldata = medInstanceUSD.methods.initialize(accounts[0], usdBytes32, veloBytes32).encodeABI();

    await medProxyUSD.initialize(medLogic.address, initializeUSDCalldata);

    const medUSD = await Medianizer.at(medProxyUSD.address);
    await medUSD.addFeeder(usdFeeder.address);

    // Medianizer THB

    await deployer.deploy(MedianizerProxy, accounts[0]);
    const medProxyTHB = await MedianizerProxy.deployed();

    const medInstanceTHB = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeTHBCalldata = medInstanceTHB.methods.initialize(accounts[0], thbBytes32, veloBytes32).encodeABI();

    await medProxyTHB.initialize(medLogic.address, initializeTHBCalldata);

    const medTHB = await Medianizer.at(medProxyTHB.address);
    await medTHB.addFeeder(thbFeeder.address);

    // Medianizer SGD

    await deployer.deploy(MedianizerProxy, accounts[0]);
    const medProxySGD = await MedianizerProxy.deployed();

    const medInstanceSGD = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeSGDCalldata = medInstanceSGD.methods.initialize(accounts[0], sgdBytes32, veloBytes32).encodeABI();

    await medProxySGD.initialize(medLogic.address, initializeSGDCalldata);

    const medSGD = await Medianizer.at(medProxySGD.address);
    await medSGD.addFeeder(sgdFeeder.address);

    console.log('medProxyUSD', medProxyUSD.address);
    console.log('medProxyTHB', medProxyTHB.address);
    console.log('medProxySGD', medProxySGD.address);

    await deployer.deploy(Lag);
    const lagLogic = await Lag.deployed();

    // Lag USD

    await deployer.deploy(LagProxy, accounts[0]);
    const lagProxyUSD = await LagProxy.deployed();

    const lagInstanceUSD = new web3.eth.Contract(Lag.abi, lagLogic.address);
    const initializeLagUSDCalldata = lagInstanceUSD.methods.initialize(accounts[0], medProxyUSD.address).encodeABI();

    await lagProxyUSD.initialize(lagLogic.address, initializeLagUSDCalldata);

    // Lag THB

    await deployer.deploy(LagProxy, accounts[0]);
    const lagProxyTHB = await LagProxy.deployed();

    const lagInstanceTHB = new web3.eth.Contract(Lag.abi, lagLogic.address);
    const initializeLagTHBCalldata = lagInstanceTHB.methods.initialize(accounts[0], medProxyTHB.address).encodeABI();

    await lagProxyTHB.initialize(lagLogic.address, initializeLagTHBCalldata);
    // Lag SGD

    await deployer.deploy(LagProxy, accounts[0]);
    const lagProxySGD = await LagProxy.deployed();

    const lagInstanceSGD = new web3.eth.Contract(Lag.abi, lagLogic.address);
    const initializeLagSGDCalldata = lagInstanceSGD.methods.initialize(accounts[0], medProxySGD.address).encodeABI();

    await lagProxySGD.initialize(lagLogic.address, initializeLagSGDCalldata);

    console.log('lagProxyUSD', lagProxyUSD.address);
    console.log('lagProxyTHB', lagProxyTHB.address);
    console.log('lagProxySGD', lagProxySGD.address);

    await deployer.deploy(Price);
    const priceLogic = await Price.deployed();

    // Price USD

    await deployer.deploy(PriceProxy, accounts[0]);
    const priceProxyUSD = await PriceProxy.deployed();

    const priceInstanceUSD = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceUSDCalldata = priceInstanceUSD.methods.initialize(accounts[0], lagProxyUSD.address).encodeABI();

    await priceProxyUSD.initialize(priceLogic.address, initializePriceUSDCalldata);

    // Price THB

    await deployer.deploy(PriceProxy, accounts[0]);
    const priceProxyTHB = await PriceProxy.deployed();

    const priceInstanceTHB = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceTHBCalldata = priceInstanceTHB.methods.initialize(accounts[0], lagProxyTHB.address).encodeABI();

    await priceProxyTHB.initialize(priceLogic.address, initializePriceTHBCalldata);

    // Price SGD

    await deployer.deploy(PriceProxy, accounts[0]);
    const priceProxySGD = await PriceProxy.deployed();

    const priceInstanceSGD = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceSGDCalldata = priceInstanceSGD.methods.initialize(accounts[0], lagProxySGD.address).encodeABI();

    await priceProxySGD.initialize(priceLogic.address, initializePriceSGDCalldata);

    console.log('priceProxyUSD', priceProxyUSD.address);
    console.log('priceProxyTHB', priceProxyTHB.address);
    console.log('priceProxySGD', priceProxySGD.address);
  }
};
