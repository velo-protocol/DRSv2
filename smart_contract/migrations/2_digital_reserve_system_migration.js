const DRS = artifacts.require('DigitalReserveSystem');

module.exports = async function (deployer) {
    await deployer.deploy(DRS);
};