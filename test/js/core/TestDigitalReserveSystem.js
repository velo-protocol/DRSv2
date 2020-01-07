const DigitalReserveSystem = artifacts.require("DigitalReserveSystem");
const Heart = artifacts.require("Heart");
const PriceFeeders = artifacts.require('PriceFeeders');
const StableCredit = artifacts.require("StableCredit");
const Token = artifacts.require('Token');
const ReserveManager = artifacts.require('ReserveManager');
const MockContract = artifacts.require("MockContract");
const CollateralAsset = artifacts.require('ERC20');

const Web3 = require('web3');
const truffleAssert = require('truffle-assertions');
const helper = require('../testhelper');

let drs, heart, mock, priceFeeder, reserveManager, veloCollateralAsset;

contract("DigitalReserveSystem test", async accounts => {

  before(async () => {
    mock = await MockContract.new();
    heart = await Heart.at(mock.address);
    priceFeeder = await PriceFeeders.at(mock.address);
    reserveManager = await ReserveManager.at(mock.address);
    veloCollateralAsset = await Token.at(mock.address);

  });

  beforeEach(async () => {
    drs = await DigitalReserveSystem.new(mock.address);
  });

  afterEach(async () => {
    await mock.reset()
  });

  describe("SetupCredit", async () => {
    it("should setup credit successfully", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        helper.address(0)
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        helper.address(1)
      );
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getPriceFeeders"),
        priceFeeder.address
      );
      await mock.givenMethodReturnUint(
        helper.methodABI(priceFeeder, "getMedianPrice", [helper.address(0)]),
        1
      );

      const result = await drs.setup(
        Web3.utils.fromAscii("VELO"),
        Web3.utils.fromAscii("USD"),
        "vUSD",
        1
      );

      truffleAssert.eventEmitted(result, 'Setup', async event => {
        assert.equal(event.assetCode, "vUSD");
        helper.assertEqualByteString(event.peggedCurrency, "USD");
        assert.equal(event.peggedValue, 1);
        helper.assertEqualByteString(event.collateralAssetCode, "VELO");
        assert.ok(web3.utils.isAddress(event.assetAddress));

        const stableCredit = await StableCredit.at(event.assetAddress);
        const name = await stableCredit.name();
        assert.equal(name, "vUSD");

        return true;
      }, 'contract should emit the event correctly');
    });

    it("should fail, caller must be trusted partner", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        false
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner")
      }
    });

    it("should fail, invalid assetCode format", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSDvUSDvUSDvUSDvUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: invalid assetCode format")
      }
    });

    it("should fail, assetCode has already been used", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: assetCode has already been used")
      }
    });

    it("should fail, collateralAssetCode does not exist", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: collateralAssetCode does not exist")
      }
    });

    it("should fail, collateralAssetCode - peggedCurrency pair does not exist", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        false
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: collateralAssetCode - peggedCurrency pair does not exist")
      }
    });

    it("should fail, price of link must have value more than 0", async () => {
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );
      await mock.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        true
      );
      await mock.givenMethodReturnAddress(
        helper.methodABI(heart, "getPriceFeeders"),
        priceFeeder.address
      );
      await mock.givenMethodReturnUint(
        helper.methodABI(priceFeeder, "getMedianPrice", [helper.address(0)]),
        0
      );

      try {
        await drs.setup(
          Web3.utils.fromAscii("VELO"),
          Web3.utils.fromAscii("USD"),
          "vUSD",
          1
        );
      } catch (err) {
        assert.equal(err.reason, "DigitalReserveSystem.setup: price of link must have value more than 0")
      }
    });
  });

  describe("MintFromCollateral", async () => {
    it("should mint from collateral correctly", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("THB"),
        accounts[1],
        Web3.utils.fromAscii('VELO'),
        veloCollateralAsset.address,
        'vTHB',
        1000000,
        heart.address
      );

      await mock.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );

      await mock.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("stableCreditId")).encodeABI(),
        stableCredit.address
      );

      await mock.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );

      await mock.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        10000000
      );

      await mock.givenMethodReturnUint(
        heart.contract.methods.getCreditIssuanceFee().encodeABI(),
        100000
      );

      await mock.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii('VELO')).encodeABI(),
        10000000
      );

      await mock.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("VELO")).encodeABI(),
        veloCollateralAsset.address
      );
      await mock.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], heart.address, 1).encodeABI(),
        true
      );

      await mock.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], veloCollateralAsset.address, 1).encodeABI(),
        true
      );

      await mock.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], drs.address, 8).encodeABI(),
        true
      );

      await mock.givenMethodReturnAddress(
        heart.contract.methods.getDrsAddress().encodeABI(),
        drs.address
      );

      await mock.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.approve(accounts[0], 1).encodeABI(),
        true
      );

      await mock.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      await mock.givenMethodReturnBool(
        reserveManager.contract.methods.lockReserve(Web3.utils.fromAscii('VELO'), drs.address, 1).encodeABI(),
        true
      );

      const result = await drs.mintFromCollateral(100000000, 'vTHB');

      truffleAssert.eventEmitted(result, 'Mint', event => {
        return event.assetCode === "vTHB"
          && new web3.utils.BN(event.mintAmount).toNumber() === 990000000
          && new web3.utils.BN(event.collateralAmount).toNumber() === 99000000
          && event.assetAddress === stableCredit.address
          && event.collateralAssetCode === web3.utils.padRight(web3.utils.utf8ToHex("VELO"), 64);
      }, 'contract should emit the event correctly');

    });

    it("should fail, caller must be trusted partner", async () => {
      await mock.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        false
      );
      try {
        await drs.mintFromCollateral(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    });

    it("should fail, caller must be setup credit already", async () => {
      await mock.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x0", 40)
      );
      try {
        await drs.mintFromCollateral(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintFromCollateral: stableCredit not exist", err.reason)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      stableCredit = await StableCredit.new(
          Web3.utils.fromAscii("USD"),
          accounts[1],
          Web3.utils.fromAscii("VELO"),
          veloCollateralAsset.address,
          'vUSD',
          1000000,
          heart.address
      );
      const badCollateralAsset = await CollateralAsset.new();
      heart.addStableCredit(stableCredit.address);

      await mock.givenMethodReturnBool(
          heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
          true
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
          stableCredit.address
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
          badCollateralAsset.address
      );
      try {
        await drs.mintFromCollateral(
            1,
            "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintFromCollateral: collateralAsset must be the same", err.reason)
      }
    });

    it("should fail, median price ref mut not be zero", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[2],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );
      await heart.addStableCredit(stableCredit.address);

      await mock.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        priceFeeder.address
      );
      await mock.givenMethodReturnUint(
        priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        0
      );
      try {
        await drs.mintFromCollateral(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintFromCollateral: price of link must have value more than 0", err.reason)
      }
    });
  });

  describe("MintStableCredit", async () => {
    it("should mint from stable credit correctly", async () => {
      stableCredit = await StableCredit.new(
          Web3.utils.fromAscii('USD'),
          accounts[1],
          Web3.utils.fromAscii('VELO'),
          veloCollateralAsset.address,
          'vUSD',
          1000000,
          heart.address
      );

      heart.addStableCredit(stableCredit.address);

      await mock.givenMethodReturnBool(
          heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
          true
      );

      await mock.givenMethodReturnAddress(
          heart.contract.methods.getStableCreditById(Web3.utils.fromAscii('stableCreditId')).encodeABI(),
          stableCredit.address
      );

      await mock.givenMethodReturnAddress(
          heart.contract.methods.getPriceFeeders().encodeABI(),
          mock.address
      );

      await mock.givenMethodReturnUint(
          priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii('')).encodeABI(),
          10000000
      );

      await mock.givenMethodReturnUint(
          heart.contract.methods.getCreditIssuanceFee().encodeABI(),
          100000
      );

      await mock.givenMethodReturnAddress(
          heart.contract.methods.getPriceFeeders().encodeABI(),
          mock.address
      );

      await mock.givenMethodReturnUint(
          priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii('')).encodeABI(),
          10000000
      );

      await mock.givenMethodReturnUint(
          heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii('VELO')).encodeABI(),
          1
      );

      await mock.givenMethodReturnAddress(
          heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii('VELO')).encodeABI(),
          veloCollateralAsset.address
      );
      await mock.givenMethodReturnBool(
          veloCollateralAsset.contract.methods.transferFrom(accounts[0], heart.address, 1).encodeABI(),
          true
      );

      await mock.givenMethodReturnBool(
          veloCollateralAsset.contract.methods.transferFrom(accounts[0], veloCollateralAsset.address, 1).encodeABI(),
          true
      );

      await mock.givenMethodReturnBool(
          veloCollateralAsset.contract.methods.transferFrom(drs.address, accounts[0], 10).encodeABI(),
          true
      );

      await mock.givenMethodReturnAddress(
          heart.contract.methods.getDrsAddress().encodeABI(),
          drs.address
      );

      await mock.givenMethodReturnBool(
          veloCollateralAsset.contract.methods.approve(accounts[0], 10).encodeABI(),
          true
      );

      await mock.givenMethodReturnAddress(
          heart.contract.methods.getReserveManager().encodeABI(),
          reserveManager.address
      );

      await mock.givenMethodReturnBool(
          reserveManager.contract.methods.lockReserve(Web3.utils.fromAscii('VELO'), drs.address, 1).encodeABI(),
          true
      );

      const result = await drs.mintStableCredit(10, Web3.utils.fromAscii("vUSD"));
      const stableCreditName = await stableCredit.name();
      truffleAssert.eventEmitted(result, 'Mint', event => {
        let BN = web3.utils.BN;
        const eventMintAmount = new BN(event.mintAmount).toNumber();
        const eventCollateralAmount = new BN(event.collateralAmount).toNumber();
        return eventMintAmount === 10000000
            && web3.utils.isAddress(event.assetAddress) === true
            && web3.utils.hexToUtf8(event.assetCode) === 'vUSD'
            && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
            && eventCollateralAmount === 1000000
            && web3.utils.hexToUtf8(event.assetCode) === stableCreditName
      }, 'contract should emit the event correctly');
    });

    it("should fail, caller must be trusted partner", async () => {
      await mock.givenMethodReturnBool(
          heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
          false
      );
      try {
        await drs.mintStableCredit(
            1,
            "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    });

    it("should fail, caller must be setup credit already", async () => {
      await mock.givenMethodReturnBool(
          heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
          true
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
          web3.utils.padLeft("0x0", 40)
      );
      try {
        await drs.mintStableCredit(
            1,
            "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintStableCredit: stableCredit not exist", err.reason)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      const collateralAsset = await Token.new('Velo', 'VELO', 7);
      stableCredit = await StableCredit.new(
          Web3.utils.fromAscii("USD"),
          accounts[1],
          Web3.utils.fromAscii("VELO"),
          collateralAsset.address,
          'vUSD',
          1000000,
          heart.address
      );
      const badCollateralAsset = await CollateralAsset.new();
      heart.addStableCredit(stableCredit.address);

      await mock.givenMethodReturnBool(
          heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
          true
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
          stableCredit.address
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
          badCollateralAsset.address
      );
      try {
        await drs.mintStableCredit(
            1,
            "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintStableCredit: collateralAsset must be the same", err.reason)
      }
    });

    it("should fail, median price ref mut not be zero", async () => {
      const collateralAsset = await Token.new('Velo', 'VELO', 7);
      stableCredit = await StableCredit.new(
          Web3.utils.fromAscii("USD"),
          accounts[2],
          Web3.utils.fromAscii("VELO"),
          collateralAsset.address,
          'vUSD',
          1000000,
          heart.address
      );
      await heart.addStableCredit(stableCredit.address);

      await mock.givenMethodReturnBool(
          heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
          true
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
          stableCredit.address
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
          collateralAsset.address
      );
      await mock.givenMethodReturnAddress(
          heart.contract.methods.getPriceFeeders().encodeABI(),
          mock.address
      );
      await mock.givenMethodReturnUint(
          priceFeeder.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
          0
      );
      try {
        await drs.mintStableCredit(
            1,
            "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintStableCredit: price of link must have value more than 0", err.reason)
      }
    });
  });
});

