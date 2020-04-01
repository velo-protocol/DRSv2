const Heart = artifacts.require('Heart');
const ReserveManager = artifacts.require('ReserveManager');
const Price = artifacts.require('Price');
const DRS = artifacts.require('DigitalReserveSystem');
const Token = artifacts.require('Token');
const Hasher = artifacts.require('Hasher');
const StableCredit = artifacts.require('StableCredit');
const h = require("../testhelper");
const MockContract = artifacts.require("MockContract");

let drs, heart, price, reserveManager, veloCollateralAsset, otherCollateralAsset, stableCreditVUSD, mocks;

const velo = "VELO";
const veloBytes32 = web3.utils.padRight(web3.utils.fromAscii(velo), 64);
const usdBytes32 = web3.utils.padRight(web3.utils.fromAscii("USD"), 64);

contract("DigitalReserveSystem Scenario Test", async accounts => {

  before(async () => {
    mocks = {
      price: await MockContract.new(),
    }
    price = await Price.at(mocks.price.address);
  });

  it("should work!", async () => {
    const calInputs = {
      issuanceFeeRate: h.decimal7(0.05), // 0.05 (5%)
      price: h.decimal7(10.00), // 10.00 USD/VELO
      collateralRatio: h.decimal7(1.3), // 1.30
      peggedValue: h.decimal7(1.00) // 1.00
    };

    const [deployer, bob, alice, pf] = accounts;

    // Setup the whole ecosystem
    heart = await Heart.new();
    drs = await DRS.new(heart.address);
    reserveManager = await ReserveManager.new(heart.address);
    veloCollateralAsset = await Token.new(velo, velo, 7);

    await mocks.price.givenMethodReturnUint(
      price.contract.methods.get().encodeABI(),
      calInputs.price
    );

    // Setup heart
    heart.addPrice(web3.utils.soliditySha3(veloBytes32, usdBytes32), price.address);
    heart.setReserveManager(reserveManager.address);
    heart.setDrsAddress(drs.address);

    // Mint VELO for Bob, so he can spend VELO afterward
    await veloCollateralAsset.mint(bob, 10000000000000);

    // Setup Heart
    await heart.setCollateralAsset(veloBytes32, veloCollateralAsset.address, calInputs.collateralRatio);
    await heart.setTrustedPartner(bob);
    await heart.setCreditIssuanceFee(calInputs.issuanceFeeRate);
    await heart.setAllowedLink(web3.utils.soliditySha3(veloBytes32, usdBytes32), true);

    // Approve DRS to spend VELO
    await veloCollateralAsset.approve(drs.address, 10000000000, {from: bob});

    // 1. Test drs.setup
    const setupResult = await drs.setup(veloBytes32, usdBytes32, "vUSD", calInputs.peggedValue, {from: bob}); // pegged value 1.00
    const setupEvent = setupResult.logs[0].args;
    assert.equal(setupEvent.assetCode, "vUSD");
    h.assert.equalByteString(setupEvent.peggedCurrency, "USD");
    h.assert.equalByteString(setupEvent.collateralAssetCode, velo);
    h.assert.equalNumber(setupEvent.peggedValue, calInputs.peggedValue);
    assert.ok(setupEvent.assetAddress);

    // Assert the credit that has been setup
    const vUSDStableCredit = await StableCredit.at(setupEvent.assetAddress);
    assert.equal(await vUSDStableCredit.assetCode(), "vUSD");
    h.assert.equalNumber(await vUSDStableCredit.peggedValue(), calInputs.peggedValue);
    h.assert.equalNumber(await vUSDStableCredit.balanceOf(bob), 0);

    // 2. Test drs.mintFromCollateralAmount
    const mintCResult = await drs.mintFromCollateralAmount(22000000, "vUSD", {from: bob});
    const mintCEvent = mintCResult.logs[0].args;
    assert.equal(mintCEvent.assetCode, "vUSD");
    h.assert.equalNumber(mintCEvent.mintAmount, 160769230);
    assert.equal(mintCEvent.assetAddress, setupEvent.assetAddress);
    h.assert.equalByteString(mintCEvent.collateralAssetCode, velo);
    h.assert.equalNumber(mintCEvent.collateralAmount, 22000000);

    // Assert the credit that has been minted
    h.assert.equalNumber(await vUSDStableCredit.balanceOf(bob), 160769230); // Bob's vUSD balance
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(bob), 10000000000000 - 22000000); // Bob's VELO balance
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(heart.address), 1100000); // fee
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(vUSDStableCredit.address), 16076923 + 4823077); // actualCollateralAmount + reserveCollateralAmount

    // 3. Test drs.mintFromStableCreditAmount
    const mintSResult = await drs.mintFromStableCreditAmount(160769230, "vUSD", {from: bob});
    const mintSEvent = mintSResult.logs[0].args;
    assert.equal(mintSEvent.assetCode, "vUSD");
    h.assert.equalNumber(mintSEvent.mintAmount, 160769230);
    assert.equal(mintSEvent.assetAddress, setupEvent.assetAddress);
    h.assert.equalByteString(mintSEvent.collateralAssetCode, velo);
    h.assert.equalNumber(mintSEvent.collateralAmount, 21999998);

    // Assert the credit that has been minted
    // Noted that there is little calculation error
    h.assert.equalNumber(await vUSDStableCredit.balanceOf(bob), 160769230 + 160769230); // Bob's vUSD balance
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(bob), 10000000000000 - 22000000 - 21999998); // Bob's VELO balance
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(heart.address), 1100000 + 1099999); // fee
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(vUSDStableCredit.address), 16076923 + 16076922 + 4823077 + 4823077); // actualCollateralAmount + reserveCollateralAmount

    // 4. Test drs.getExchange
    let getExchangeResult = await drs.getExchange("vUSD");
    getExchangeResult = {
      assetCode: getExchangeResult['0'],
      collateralAssetCode: getExchangeResult['1'],
      priceInCollateralPerAssetUnit: getExchangeResult['2']
    };
    assert.equal(getExchangeResult.assetCode, "vUSD");
    h.assert.equalByteString(getExchangeResult.collateralAssetCode, velo);
    h.assert.equalNumber(getExchangeResult.priceInCollateralPerAssetUnit.toString(), 1300000);

    // 5. Test drs.redeem
    const redeemResult = await drs.redeem(160769230, "vUSD", {from: bob});
    const redeemEvent = redeemResult.logs[0].args;
    assert.equal(redeemEvent.assetCode, "vUSD");
    h.assert.equalNumber(redeemEvent.stableCreditAmount, 160769230);
    assert.equal(redeemEvent.collateralAsset, veloCollateralAsset.assetAddress);
    h.assert.equalByteString(redeemEvent.collateralAssetCode, velo);
    h.assert.equalNumber(redeemEvent.collateralAmount, 20899999);

    // Assert the credit that has been redeemed
    // Noted that there is little calculation error
    h.assert.equalNumber(await vUSDStableCredit.balanceOf(bob), 160769230 + 160769230 - 160769230); // Bob's vUSD balance
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(bob), 10000000000000 - 22000000 - 21999998 + 20899999); // Bob's VELO balance
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(vUSDStableCredit.address), 20899999 + 1); // actualCollateralAmount '1' from rebalance error
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(reserveManager.address), 0); // reserveCollateralAmount

    // 6. Test drs.collateralHealthCheck
    let healthCheckResult = await drs.collateralHealthCheck("vUSD");
    healthCheckResult = {
      collateralAssetAddress: healthCheckResult['0'],
      collateralAssetCode: healthCheckResult['1'],
      requiredAmount: healthCheckResult['2'],
      presentAmount: healthCheckResult['3']
    };
    h.assert.equalByteString(healthCheckResult.collateralAssetCode, velo);
    h.assert.equalNumber(healthCheckResult.requiredAmount.toString(), 20899999); // vUSD balance * peggedValue * collateralRatio / price: 160769230 * 1 * 1.3 / 10
    h.assert.equalNumber(healthCheckResult.presentAmount.toString(), 20900000);
    assert.equal(veloCollateralAsset.address, healthCheckResult.collateralAssetAddress);

    // 7. Test drs.rebalance

    // change collateral ratio
    const newRatio = h.decimal7(1.2);
    await heart.setCollateralRatio(veloBytes32, newRatio);

    const rebalanceResult = await drs.rebalance("vUSD");
    const rebalanceEvent = rebalanceResult.logs[0].args;
    h.assert.equalString(rebalanceEvent.assetCode, "vUSD");
    h.assert.equalByteString(rebalanceEvent.collateralAssetCode, velo);
    h.assert.equalNumber(rebalanceEvent.requiredAmount.toString(), 19292307); // vUSD balance * peggedValue * collateralRatio / price: 160769230 * 1 * 1.2 / 10
    h.assert.equalNumber(rebalanceEvent.presentAmount.toString(), 20900000);

  });
});
