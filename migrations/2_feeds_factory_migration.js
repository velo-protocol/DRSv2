const FeederFactory = artifacts.require('FeederFactory');

module.exports = async function (deployer) {
    await deployer.deploy(FeederFactory);
};
