const Web3 = require('web3');

const MedProxy = artifacts.require('MedProxy');
const Med = artifacts.require('Med');

module.exports = async function (deployer, network, accounts) {
    // await deployer.deploy(Med);
    // let med = await Med.deployed();
    //
    // await deployer.deploy(MedProxy, accounts[0]);
    // let medProxy = await MedProxy.deployed();
    //
    // let initCalldata = med.methods.initialize(accounts[0], Web3.utils.fromAscii("VELOUSD"));
    //
    // console.log("======= initialize Med in MedProxy context =======");
    // await medProxy.initialize(med.address, initCalldata);
};
