const FeederFactory = artifacts.require('FeederFactory');
const Feeder = artifacts.require('Feeder');
const MedianizerProxy = artifacts.require('MedianizerProxy');
const Medianizer = artifacts.require('Medianizer');
const Lag = artifacts.require('Lag');
const LagProxy = artifacts.require('LagProxy');
const Price = artifacts.require('Price');
const PriceProxy = artifacts.require('PriceProxy');

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(FeederFactory);

  if (network === 'development' || network === 'local' || network === 'dev') {

    const [defaultDeployer, usdDeployer, thbDeployer, sgdDeployer] = accounts;

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
    const feederUSD = await Feeder.at(usdFeederAddr);

    const feederTHB = await Feeder.at(thbFeederAddr);

    const feederSGD = await Feeder.at(sgdFeederAddr);

    await deployer.deploy(Medianizer);
    const medLogic = await Medianizer.deployed();

    // Medianizer USD
    await deployer.deploy(MedianizerProxy, usdDeployer);
    const medProxyUSD = await MedianizerProxy.deployed();

    const medInstanceUSD = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeUSDCalldata = await medInstanceUSD.methods.initialize(usdDeployer, usdBytes32, veloBytes32).encodeABI();

    await medProxyUSD.initialize(medLogic.address, initializeUSDCalldata);

    const medUSD = await Medianizer.at(medProxyUSD.address);
    await medUSD.addFeeder(feederUSD.address, {from: usdDeployer});

    // Medianizer THB

    await deployer.deploy(MedianizerProxy, thbDeployer);
    const medProxyTHB = await MedianizerProxy.deployed();

    const medInstanceTHB = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeTHBCalldata = await medInstanceTHB.methods.initialize(thbDeployer, thbBytes32, veloBytes32).encodeABI();

    await medProxyTHB.initialize(medLogic.address, initializeTHBCalldata);

    const medTHB = await Medianizer.at(medProxyTHB.address);
    await medTHB.addFeeder(feederTHB.address, {from: thbDeployer});

    // Medianizer SGD

    await deployer.deploy(MedianizerProxy, sgdDeployer);
    const medProxySGD = await MedianizerProxy.deployed();

    const medInstanceSGD = new web3.eth.Contract(Medianizer.abi, medLogic.address);
    const initializeSGDCalldata = await medInstanceSGD.methods.initialize(sgdDeployer, sgdBytes32, veloBytes32).encodeABI();

    await medProxySGD.initialize(medLogic.address, initializeSGDCalldata);

    const medSGD = await Medianizer.at(medProxySGD.address);
    await medSGD.addFeeder(feederSGD.address, {from: sgdDeployer});

    console.log('medProxyUSD', medProxyUSD.address);
    console.log('medProxyTHB', medProxyTHB.address);
    console.log('medProxySGD', medProxySGD.address);

    await deployer.deploy(Lag);
    const lagLogic = await Lag.deployed();

    // Lag USD

    await deployer.deploy(LagProxy, usdDeployer);
    const lagProxyUSD = await LagProxy.deployed();

    const lagInstanceUSD = new web3.eth.Contract(Lag.abi, lagLogic.address);
    const initializeLagUSDCalldata = await lagInstanceUSD.methods.initialize(usdDeployer, medProxyUSD.address).encodeABI();

    await lagProxyUSD.initialize(lagLogic.address, initializeLagUSDCalldata);

    // Lag THB

    await deployer.deploy(LagProxy, thbDeployer);
    const lagProxyTHB = await LagProxy.deployed();

    const lagInstanceTHB = new web3.eth.Contract(Lag.abi, lagLogic.address);
    const initializeLagTHBCalldata = await lagInstanceTHB.methods.initialize(thbDeployer, medProxyTHB.address).encodeABI();

    await lagProxyTHB.initialize(lagLogic.address, initializeLagTHBCalldata);

    // Lag SGD

    await deployer.deploy(LagProxy, sgdDeployer);
    const lagProxySGD = await LagProxy.deployed();

    const lagInstanceSGD = new web3.eth.Contract(Lag.abi, lagLogic.address);
    const initializeLagSGDCalldata = await lagInstanceSGD.methods.initialize(sgdDeployer, medProxySGD.address).encodeABI();

    await lagProxySGD.initialize(lagLogic.address, initializeLagSGDCalldata);

    console.log('lagProxyUSD', lagProxyUSD.address);
    console.log('lagProxyTHB', lagProxyTHB.address);
    console.log('lagProxySGD', lagProxySGD.address);

    await deployer.deploy(Price);
    const priceLogic = await Price.deployed();

    // Price USD

    await deployer.deploy(PriceProxy, usdDeployer);
    const priceProxyUSD = await PriceProxy.deployed();

    const priceInstanceUSD = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceUSDCalldata = await priceInstanceUSD.methods.initialize(usdDeployer, lagProxyUSD.address).encodeABI();

    await priceProxyUSD.initialize(priceLogic.address, initializePriceUSDCalldata);

    // Price THB

    await deployer.deploy(PriceProxy, thbDeployer);
    const priceProxyTHB = await PriceProxy.deployed();

    const priceInstanceTHB = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceTHBCalldata = await priceInstanceTHB.methods.initialize(thbDeployer, lagProxyTHB.address).encodeABI();

    await priceProxyTHB.initialize(priceLogic.address, initializePriceTHBCalldata);

    // Price SGD

    await deployer.deploy(PriceProxy, sgdDeployer);
    const priceProxySGD = await PriceProxy.deployed();

    const priceInstanceSGD = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceSGDCalldata = await priceInstanceSGD.methods.initialize(sgdDeployer, lagProxySGD.address).encodeABI();

    await priceProxySGD.initialize(priceLogic.address, initializePriceSGDCalldata);

    console.log('priceProxyUSD', priceProxyUSD.address);
    console.log('priceProxyTHB', priceProxyTHB.address);
    console.log('priceProxySGD', priceProxySGD.address);

    process.migration = {
      contractAddress: {
        priceProxyUSD: priceProxyUSD.address,
        priceProxyTHB: priceProxyTHB.address,
        priceProxySGD: priceProxySGD.address,
      }
    }

  }

};
