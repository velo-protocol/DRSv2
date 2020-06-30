require('dotenv').config();

const FeederFactory = artifacts.require('FeederFactory');
const Feeder = artifacts.require('Feeder');
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
  var  pf1 = 'address';
  var  pf2 = 'address';
  var  pf3 = 'address';
    // 1. Create 3 feeders
    let result = await feederFactory.create(pf1,pf2,pf3,usdBytes32, veloBytes32);
    const usdFeederAddr = result.logs[0].args.feedAddr;
    console.log(`USD Feeder: ${usdFeederAddr}`);

    result = await feederFactory.create(pf1,pf2,pf3,thbBytes32, veloBytes32);
    const thbFeederAddr = result.logs[0].args.feedAddr;
    console.log(`THB Feeder: ${thbFeederAddr}`);

    result = await feederFactory.create(pf1,pf2,pf3,sgdBytes32, veloBytes32);
    const sgdFeederAddr = result.logs[0].args.feedAddr;
    console.log(`SGD Feeder: ${sgdFeederAddr}`);



    await deployer.deploy(Price);
    const priceLogic = await Price.deployed();

    // Price USD

    await deployer.deploy(PriceProxy, process.env.USD_VELO_ADDRESS);
    const priceProxyUSD = await PriceProxy.deployed();

    const priceInstanceUSD = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceUSDCalldata = priceInstanceUSD.methods.initialize(process.env.USD_VELO_ADDRESS, usdFeederAddr).encodeABI();

    await priceProxyUSD.initialize(priceLogic.address, initializePriceUSDCalldata);

    // Price THB

    await deployer.deploy(PriceProxy, process.env.THB_VELO_ADDRESS);
    const priceProxyTHB = await PriceProxy.deployed();

    const priceInstanceTHB = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceTHBCalldata = priceInstanceTHB.methods.initialize(process.env.THB_VELO_ADDRESS, thbFeederAddr).encodeABI();

    await priceProxyTHB.initialize(priceLogic.address, initializePriceTHBCalldata);

    // Price SGD

    await deployer.deploy(PriceProxy, process.env.SGD_VELO_ADDRESS);
    const priceProxySGD = await PriceProxy.deployed();

    const priceInstanceSGD = new web3.eth.Contract(Price.abi, priceLogic.address);
    const initializePriceSGDCalldata = priceInstanceSGD.methods.initialize(process.env.SGD_VELO_ADDRESS, sgdFeederAddr).encodeABI();

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
