const Web3 = require('web3');

const PriceFeeders = artifacts.require('PriceFeeders');
const DRS = artifacts.require('DigitalReserveSystem');
const Token = artifacts.require('Token');

module.exports = async function (deployer, network) {
    await deployer.deploy(PriceFeeders);
    let priceFeedersInstance = await PriceFeeders.deployed();
    await deployer.deploy(DRS, priceFeedersInstance.address);
    let drsInstance = await DRS.deployed();

    if (network === 'development') {
        let adminAddress = '0xdEF0D3660Ed8A0165aBfb21020BC5834565c263C';

        await deployer.deploy(Token, "TIM", "TIM", 7);
        let timToken = await Token.deployed();
        await timToken.mint(adminAddress, 10000000000000);

        let timTokenAscii = Web3.utils.fromAscii("TIM");
        let usdAscii = Web3.utils.fromAscii("USD");

        await priceFeedersInstance.setAsset(timTokenAscii, timToken.address);
        await priceFeedersInstance.addAssetFiat(timTokenAscii, usdAscii);
        await priceFeedersInstance.addPriceFeeder(timTokenAscii, usdAscii, adminAddress);
        await priceFeedersInstance.setPrice(timTokenAscii, usdAscii, 10000000);

        await drsInstance.setCollateralAsset(timTokenAscii, timToken.address);
        await drsInstance.setTrustedPartner(adminAddress);

        await timToken.approve(drsInstance.address, 100000000);
    }
};