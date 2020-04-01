const Web3 = require('web3');

const Heart = artifacts.require('Heart');
const ReserveManager = artifacts.require('ReserveManager');
const Price = artifacts.require('Price');
const DRS = artifacts.require('DigitalReserveSystem');
const Token = artifacts.require('Token');
const Hasher = artifacts.require('Hasher');
const StableCredit = artifacts.require('StableCredit');

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(Hasher);
  await deployer.link(Hasher, [DRS, Heart, StableCredit]);

  await deployer.deploy(Heart);
  const heartInstance = await Heart.deployed();
  await deployer.deploy(Price);
  const priceInstance = await Price.deployed();
  await deployer.deploy(DRS, heartInstance.address);
  const drsInstance = await DRS.deployed();
  await deployer.deploy(ReserveManager, heartInstance.address);
  const reserveManagerInstance = await ReserveManager.deployed();


  if (network === 'development' || network === 'local' || network === 'dev') {
    const hasher = await Hasher.deployed();

    const veloBytes32 = Web3.utils.fromAscii("VELO");
    const usdBytes32 = Web3.utils.fromAscii("USD");

    heartInstance.addPrice(await hasher.linkId(veloBytes32, usdBytes32), priceInstance.address);
    heartInstance.setReserveManager(reserveManagerInstance.address);
    heartInstance.setDrsAddress(drsInstance.address);

    const adminAddress = accounts[0];

    await deployer.deploy(Token, "VELO", "VELO", 7);
    const veloToken = await Token.deployed();
    await veloToken.mint(adminAddress, 10000000000000);

    console.log("Set Collateral assets");
    await heartInstance.setCollateralAsset(veloBytes32, veloToken.address, 13000000); // 1.3
    await heartInstance.setTrustedPartner(adminAddress);
    await heartInstance.setCreditIssuanceFee(500000);  // 0.05 (5%)
    await heartInstance.setAllowedLink(await hasher.linkId(veloBytes32, usdBytes32), true);

    console.log("Approve DRS to spend VELO");
    await veloToken.approve(drsInstance.address, 10000000000);
  }
};

