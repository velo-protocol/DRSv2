const PriceFeeders = artifacts.require('PriceFeeders');
const DRS = artifacts.require('DigitalReserveSystem');

module.exports = async function (deployer) {
    await deployer.deploy(PriceFeeders);
    let priceFeedersInstance = await PriceFeeders.deployed();
    await deployer.deploy(DRS, priceFeedersInstance.address);
};