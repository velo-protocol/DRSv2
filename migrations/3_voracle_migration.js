const Web3 = require('web3');

const MedianizerProxy = artifacts.require('MedianizerProxy');
const Medianizer = artifacts.require('Medianizer');
const Lag = artifacts.require('Lag');

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(Medianizer);
  const medLogic = await Medianizer.deployed();

  await deployer.deploy(MedianizerProxy, accounts[0]);
  const medProxy = await MedianizerProxy.deployed();

  const medInstance = new web3.eth.Contract(Medianizer.abi, medLogic.address);
  const initializeCalldata = medInstance.methods.initialize(accounts[0], Web3.utils.fromAscii("USD"), Web3.utils.fromAscii("VELO")).encodeABI();

  await medProxy.initialize(medLogic.address, initializeCalldata);

  // TODO: await deployer.deploy(Lag, gov.address, medProxy.address)
  await deployer.deploy(Lag, medProxy.address, medProxy.address);
};
