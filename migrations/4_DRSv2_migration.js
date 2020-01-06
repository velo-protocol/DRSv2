const Web3 = require('web3');

const Heart = artifacts.require('Heart');
const ReserveManager = artifacts.require('ReserveManager');
const PriceFeeders = artifacts.require('PriceFeeders');
const DRS = artifacts.require('DigitalReserveSystem');
const Token = artifacts.require('Token');
const Hasher = artifacts.require('Hasher');
const StableCredit = artifacts.require('StableCredit');

module.exports = async function (deployer, network, accounts) {
    await deployer.deploy(Hasher);
    await deployer.link(Hasher, [DRS, Heart, StableCredit]);

    await deployer.deploy(Heart);
    let heartInstance = await Heart.deployed();
    await deployer.deploy(PriceFeeders);
    let priceFeedersInstance = await PriceFeeders.deployed();
    await deployer.deploy(DRS, heartInstance.address);
    let drsInstance = await DRS.deployed();
    await deployer.deploy(ReserveManager, heartInstance.address);
    let reserveManagerInstance = await ReserveManager.deployed();

    heartInstance.setPriceFeeders(priceFeedersInstance.address);
    heartInstance.setReserveManager(reserveManagerInstance.address);
    heartInstance.setDrsAddress(drsInstance.address);

    if (network === 'development' || network === 'evrynet') {
        let adminAddress = accounts[0];

        await deployer.deploy(Token, "VELO", "VELO", 7);
        let veloToken = await Token.deployed();
        await veloToken.mint(adminAddress, 10000000000000);

        let veloTokenAscii = Web3.utils.fromAscii("VELO");
        let usdAscii = Web3.utils.fromAscii("USD");

        console.log("Setting up Price Feeder for VELO");
        await priceFeedersInstance.setAsset(veloTokenAscii, veloToken.address);
        await priceFeedersInstance.addAssetFiat(veloTokenAscii, usdAscii);
        await priceFeedersInstance.addPriceFeeder(veloTokenAscii, usdAscii, adminAddress);
        await priceFeedersInstance.setPrice(veloTokenAscii, usdAscii, 10000000);

        console.log("Set Collateral assets");
        await heartInstance.setCollateralAsset(veloTokenAscii, veloToken.address, 100);
        await heartInstance.setTrustedPartner(adminAddress);

        console.log("Approve DRS to spend VELO");
        await veloToken.approve(drsInstance.address, 10000000000);
    }
};
