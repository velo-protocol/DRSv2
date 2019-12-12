const Web3 = require('web3');

const FeedsFactory = artifacts.require('FeedsFactory');

module.exports = async function (deployer) {
    await deployer.deploy(FeedsFactory);
};
