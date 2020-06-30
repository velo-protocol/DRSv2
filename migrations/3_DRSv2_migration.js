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
  await deployer.deploy(DRS, heartInstance.address);
  const drsInstance = await DRS.deployed();
  await deployer.deploy(ReserveManager, heartInstance.address);
  const reserveManagerInstance = await ReserveManager.deployed();

  if (network === 'development' || network === 'local' || network === 'dev') {
    const hasher = await Hasher.deployed();

    const veloBytes32 = Web3.utils.fromAscii("VELO");
    const usdBytes32 = Web3.utils.fromAscii("USD");
    const thbBytes32 = Web3.utils.fromAscii("THB");
    const sgdBytes32 = Web3.utils.fromAscii("SGD");

    const veloMintAmount = 10000000000000;

    const priceUSD = await Price.at(process.migration.contractAddress.priceProxyUSD);
    const priceTHB = await Price.at(process.migration.contractAddress.priceProxyTHB);
    const priceSGD = await Price.at(process.migration.contractAddress.priceProxySGD);

    await heartInstance.addPrice(await hasher.linkId(veloBytes32, usdBytes32), priceUSD.address);
    await heartInstance.addPrice(await hasher.linkId(veloBytes32, thbBytes32), priceTHB.address);
    await heartInstance.addPrice(await hasher.linkId(veloBytes32, sgdBytes32), priceSGD.address);
    await heartInstance.setReserveManager(reserveManagerInstance.address);
    await heartInstance.setDrsAddress(drsInstance.address);

    const adminAddress = accounts[0];

    await deployer.deploy(Token, "VELO", "VELO", 7);
    const veloToken = await Token.deployed();
    await veloToken.mint(adminAddress, veloMintAmount);

    console.log("Set Collateral assets");
    await heartInstance.setCollateralAsset(veloBytes32, veloToken.address, 130000000000); // 1.3
    await heartInstance.setTrustedPartner(adminAddress);
    await heartInstance.setCreditIssuanceFee(5000000000);  // 0.05 (5%)
    await heartInstance.setAllowedLink(await hasher.linkId(veloBytes32, usdBytes32), true);
    await heartInstance.setAllowedLink(await hasher.linkId(veloBytes32, thbBytes32), true);
    await heartInstance.setAllowedLink(await hasher.linkId(veloBytes32, sgdBytes32), true);

    console.log("Approve DRS to spend VELO");
    await veloToken.approve(drsInstance.address, veloMintAmount);
  }
};
