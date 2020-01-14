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
    const heartInstance = await Heart.deployed();
    await deployer.deploy(PriceFeeders);
    const priceFeedersInstance = await PriceFeeders.deployed();
    await deployer.deploy(DRS, heartInstance.address);
    const drsInstance = await DRS.deployed();
    await deployer.deploy(ReserveManager, heartInstance.address);
    const reserveManagerInstance = await ReserveManager.deployed();


    heartInstance.setPriceFeeders(priceFeedersInstance.address);
    heartInstance.setReserveManager(reserveManagerInstance.address);
    heartInstance.setDrsAddress(drsInstance.address);

    if (network === 'development' || network === 'local' || network === 'evrynet') {
        const hasher = await Hasher.deployed();

        const adminAddress = accounts[0];

        await deployer.deploy(Token, "VELO", "VELO", 7);
        const veloToken = await Token.deployed();
        await veloToken.mint(adminAddress, 10000000000000);

        const veloTokenAscii = Web3.utils.fromAscii("VELO");
        const usdAscii = Web3.utils.fromAscii("USD");

        console.log("Setting up Price Feeder for VELO");
        await priceFeedersInstance.setAsset(veloTokenAscii, veloToken.address);
        await priceFeedersInstance.addAssetFiat(veloTokenAscii, usdAscii);
        await priceFeedersInstance.addPriceFeeder(veloTokenAscii, usdAscii, adminAddress);
        await priceFeedersInstance.setPrice(veloTokenAscii, usdAscii, 10000000);

        console.log("Set Collateral assets");
        await heartInstance.setCollateralAsset(veloTokenAscii, veloToken.address, 100);
        await heartInstance.setTrustedPartner(adminAddress);
        await heartInstance.setAllowedLink(await hasher.linkId(veloTokenAscii, usdAscii), true);

        console.log("Approve DRS to spend VELO");
        await veloToken.approve(drsInstance.address, 10000000000);
    }
};
