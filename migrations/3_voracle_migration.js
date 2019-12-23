const Web3 = require('web3');

const MedProxy = artifacts.require('MedProxy');
const Med = artifacts.require('Med');
const Lag = artifacts.require('Lag');

module.exports = async function (deployer, network, accounts) {
    await deployer.deploy(Med);
    let med = await Med.deployed();

    await deployer.deploy(MedProxy, accounts[0]);
    let medProxy = await MedProxy.deployed();

    let medInstance = new web3.eth.Contract(Med.abi, med.address);
    let initializeCalldata = medInstance.methods.initialize(accounts[0], Web3.utils.fromAscii("VELOUSD")).encodeABI();

    await medProxy.initialize(med.address, initializeCalldata);

    // TODO: await deployer.deploy(Lag, gov.address, medProxy.address)
    await deployer.deploy(Lag, medProxy.address, medProxy.address);
};
