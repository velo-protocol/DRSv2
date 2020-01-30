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

    if (network === 'development' || network === 'local' || network === 'dev') {
        const hasher = await Hasher.deployed();

        const adminAddress = accounts[0];

        await deployer.deploy(Token, "VELO", "VELO", 7);
        const veloToken = await Token.deployed();
        await veloToken.mint(adminAddress, 10000000000000);

        const veloBytes32 = Web3.utils.fromAscii("VELO");
        const usdBytes32 = Web3.utils.fromAscii("USD");

        console.log("Setting up Price Feeder for VELO");
        await priceFeedersInstance.setAsset(veloBytes32, veloToken.address);
        await priceFeedersInstance.addAssetFiat(veloBytes32, usdBytes32);
        await priceFeedersInstance.addPriceFeeder(veloBytes32, usdBytes32, adminAddress);
        await priceFeedersInstance.setPrice(veloBytes32, usdBytes32, 100000000); // 10 USD/VELO

        console.log("Set Collateral assets");
        await heartInstance.setCollateralAsset(veloBytes32, veloToken.address, 13000000); // 1.3
        await heartInstance.setTrustedPartner(adminAddress);
        await heartInstance.setCreditIssuanceFee(500000);  // 0.05 (5%)
        await heartInstance.setAllowedLink(await hasher.linkId(veloBytes32, usdBytes32), true);

        console.log("Approve DRS to spend VELO");
        await veloToken.approve(drsInstance.address, 10000000000);
    }
};

