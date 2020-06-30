const DigitalReserveSystem = artifacts.require("DigitalReserveSystem");
const Heart = artifacts.require("Heart");
const Price = artifacts.require('Price');
const StableCredit = artifacts.require("StableCredit");
const Token = artifacts.require('Token');
const ReserveManager = artifacts.require('ReserveManager');
const MockContract = artifacts.require("MockContract");
const Web3 = require('web3');
const truffleAssert = require('truffle-assertions');
const helper = require('../testhelper');
const abi = require('ethereumjs-abi');

let drs, heart, price, reserveManager, veloCollateralAsset, otherCollateralAsset, stableCreditVTHB,
  stableCreditVUSD, mocks;

contract("DigitalReserveSystem test", async accounts => {

  before(async () => {
    mocks = {
      heart: await MockContract.new(),
      price: await MockContract.new(),
      reserveManager: await MockContract.new(),
      veloCollateralAsset: await MockContract.new(),
      otherCollateralAsset: await MockContract.new(),
      stableCreditVUSD: await MockContract.new(),
      stableCreditVTHB: await MockContract.new(),
    };

    heart = await Heart.at(mocks.heart.address);
    price = await Price.at(mocks.price.address);
    reserveManager = await ReserveManager.at(mocks.reserveManager.address);
    veloCollateralAsset = await Token.at(mocks.veloCollateralAsset.address);
    otherCollateralAsset = await Token.at(mocks.otherCollateralAsset.address);
    stableCreditVUSD = await StableCredit.at(mocks.stableCreditVUSD.address);
    stableCreditVTHB = await StableCredit.at(mocks.stableCreditVTHB.address);
  });

  beforeEach(async () => {
    drs = await DigitalReserveSystem.new(heart.address);
  });

  afterEach(async () => {
    const promises = [];
    for (const key in mocks) {
      promises.push(mocks[key].reset());
    }
    await Promise.all(promises);
  });

  describe("SetupCredit", async () => {
    it("should setup credit successfully", async () => {
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        helper.address(0)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        helper.address(1)
      );
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        true
      );

      const result = await drs.setup(
        Web3.utils.fromAscii("VELO"),
        Web3.utils.fromAscii("USD"),
        "vUSD",
        1
      );

      truffleAssert.eventEmitted(result, 'Setup', event => {
        return event.assetCode === "vUSD"
          && web3.utils.hexToUtf8(event.peggedCurrency) === 'USD'
          && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
          && new web3.utils.BN(event.peggedValue).toNumber() === 1
          && web3.utils.isAddress(event.assetAddress);
      }, 'contract should emit the event correctly');
    });

    it("should fail, caller must be trusted partner", async () => {
      await mocks.heart.givenMethodReturnBool(
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
      await mocks.heart.givenMethodReturnBool(
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
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
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
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mocks.heart.givenMethodReturnAddress(
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
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );
      await mocks.heart.givenMethodReturnBool(
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
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isTrustedPartner", [accounts[0]]),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getStableCreditById", [helper.address(0)]),
        web3.utils.padLeft("0x0", 40)
      );
      await mocks.heart.givenMethodReturnAddress(
        helper.methodABI(heart, "getCollateralAsset", [helper.address(0)]),
        web3.utils.padLeft("0x1", 40)
      );
      await mocks.heart.givenMethodReturnBool(
        helper.methodABI(heart, "isLinkAllowed", [helper.address(0)]),
        true
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

  describe("MintFromCollateralAmount", async () => {
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

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("stableCreditId")).encodeABI(),
        stableCreditVTHB.address
      );

      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.creditOwner().encodeABI(),
        accounts[0]
      );

      await mocks.stableCreditVTHB.givenMethodReturn(
        stableCreditVTHB.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii("VELO")]).toString("hex")
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("VELO")).encodeABI(),
        veloCollateralAsset.address
      );

      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        100000000000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCreditIssuanceFee().encodeABI(),
        1000000000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii('VELO')).encodeABI(),
        100000000000
      );

      await mocks.stableCreditVTHB.givenMethodReturnUint(
        stableCreditVTHB.contract.methods.peggedValue().encodeABI(),
        15000000
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], heart.address, 1).encodeABI(),
        true
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], veloCollateralAsset.address, 1).encodeABI(),
        true
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], drs.address, 8).encodeABI(),
        true
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.approve(accounts[0], 1).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      await mocks.reserveManager.givenMethodReturnBool(
        reserveManager.contract.methods.lockReserve(Web3.utils.fromAscii('VELO'), drs.address, 1).encodeABI(),
        true
      );

      const result = await drs.mintFromCollateralAmount(100000000, 'vTHB');

      truffleAssert.eventEmitted(result, 'Mint', event => {
        return event.assetCode === "vTHB"
          && new web3.utils.BN(event.mintAmount).toNumber() === 660000000000
          && new web3.utils.BN(event.collateralAmount).toNumber() === 100000000
          && event.assetAddress === stableCreditVTHB.address
          && event.collateralAssetCode === web3.utils.padRight(web3.utils.utf8ToHex("VELO"), 64);
      }, 'contract should emit the event correctly');

    });

    it("should fail, caller must be trusted partner", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        false
      );
      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    });

    it("should fail, caller must be setup credit already", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x0", 40)
      );
      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: stableCredit not exist", err.reason)
      }
    });

    it("should fail, the stable credit does not belong to you", async () => {

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVTHB.address
      );
      await mocks.stableCreditVTHB.givenMethodReturn(
        stableCreditVTHB.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii("VELO")]).toString("hex")
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("VELO")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        10000000
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.creditOwner().encodeABI(),
        accounts[1]
      );

      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.mintFromCollateralAmount: the stable credit does not belong to you", err.reason)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVTHB.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.creditOwner().encodeABI(),
        accounts[0]
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err.reason)
      }
    });

    it("should fail, median price ref mut not be zero", async () => {

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVTHB.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.creditOwner().encodeABI(),
        accounts[0]
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        0
      );
      try {
        await drs.mintFromCollateralAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: valid price not found", err.reason)
      }
    });
  });

  describe("MintFromStableCreditAmount", async () => {
    it("should mint from stable credit correctly", async () => {

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii('stableCreditId')).encodeABI(),
        stableCreditVUSD.address
      );

      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.creditOwner().encodeABI(),
        accounts[0]
      );

      await mocks.stableCreditVUSD.givenMethodReturn(
        stableCreditVUSD.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii("VELO")]).toString("hex")
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("VELO")).encodeABI(),
        veloCollateralAsset.address
      );

      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        100000000000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCreditIssuanceFee().encodeABI(),
        1000000000
      );

      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii('VELO')).encodeABI(),
        100000000000
      );

      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.peggedValue().encodeABI(),
        10000000
      );

      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(accounts[0], heart.address, 1).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getDrsAddress().encodeABI(),
        drs.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.approve(accounts[0], 10).encodeABI(),
        true
      );

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      await mocks.reserveManager.givenMethodReturnBool(
        reserveManager.contract.methods.lockReserve(Web3.utils.fromAscii('VELO'), drs.address, 1).encodeABI(),
        true
      );

      const result = await drs.mintFromStableCreditAmount(100000000, Web3.utils.fromAscii("vUSD"));
      truffleAssert.eventEmitted(result, 'Mint', event => {
        const BN = web3.utils.BN;
        const eventMintAmount = new BN(event.mintAmount).toNumber();
        const eventCollateralAmount = new BN(event.collateralAmount).toNumber();
        return eventMintAmount === 100000000
          && web3.utils.isAddress(event.assetAddress) === true
          && web3.utils.hexToUtf8(event.assetCode) === 'vUSD'
          && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
          && eventCollateralAmount === 10101
      }, 'contract should emit the event correctly');
    });

    it("should fail, caller must be trusted partner", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        false
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    });

    it("should fail, caller must be setup credit already", async () => {
      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x0", 40)
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: stableCredit not exist", err.reason)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {

        assert.equal("DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err.reason)
      }
    });

    it("should fail, median price ref mut not be zero", async () => {
      const stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii("VELO"),
        veloCollateralAsset.address,
        'vUSD',
        1000000,
        heart.address
      );

      await mocks.heart.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCredit.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        0
      );
      try {
        await drs.mintFromStableCreditAmount(
          1,
          "vUSD"
        );
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: valid price not found", err.reason)
      }
    });
  });

  describe("RedeemCredit", async () => {
    it("should redeem correctly", async () => {

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVTHB.address
      );
      await mocks.stableCreditVTHB.givenMethodReturn(
        stableCreditVTHB.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii('VELO')]).toString("hex")
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );
      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        180000000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnUint(
        veloCollateralAsset.contract.methods.balanceOf(stableCreditVTHB.address).encodeABI(),
        100000000
      );
      await mocks.veloCollateralAsset.givenMethodReturnBool(
        veloCollateralAsset.contract.methods.transferFrom(stableCreditVTHB.address, reserveManager.address, 90000000).encodeABI(),
        true
      );
      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii("")).encodeABI(),
        130000000000
      );
      await mocks.stableCreditVTHB.givenMethodReturnUint(
        stableCreditVTHB.contract.methods.peggedValue().encodeABI(),
        1
      );
      await mocks.stableCreditVTHB.givenMethodReturnBool(
        stableCreditVTHB.contract.methods.redeem(accounts[0], 0, 0).encodeABI(),
        true
      );
      await mocks.stableCreditVTHB.givenMethodReturnBool(
        stableCreditVTHB.contract.methods.approveCollateral().encodeABI(),
        true
      );

      const result = await drs.redeem(20000000000000, "vTHB");
      truffleAssert.eventEmitted(result, 'Redeem', event => {
        const BN = web3.utils.BN;
        const eventMintAmount = new BN(event.stableCreditAmount).toNumber();
        const eventCollateralAmount = new BN(event.collateralAmount).toNumber();

        return event.assetCode === "vTHB"
          && eventMintAmount === 20000000000000
          && event.collateralAssetAddress === veloCollateralAsset.address
          && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
          && eventCollateralAmount === 144;
      }, 'contract should emit the event correctly');

    });
    it("should fail, redeem amount is zero", async () => {
      try {
        await drs.redeem(0, "vUSD");
      } catch (err) {
        assert.equal("DigitalReserveSystem.redeem: redeem amount must be greater than 0", err.reason)
      }

    });
    it("should fail, invalid assetCode format", async () => {
      try {
        await drs.redeem(10000000, "");
      } catch (err) {
        assert.equal("DigitalReserveSystem.redeem: invalid assetCode format", err.reason)
      }
    });
    it("should fail, stableCredit not exist", async () => {
      try {
        await drs.redeem(10000000, "vTHB");
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: stableCredit not exist", err.reason)
      }
    });
    it("should fail, collateralAsset must be the same", async () => {

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVTHB.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.collateral().encodeABI(),
        otherCollateralAsset.address
      );

      try {
        await drs.redeem(10000000, "vTHB");
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err.reason)
      }
    });
    it("should fail, collateralAssetCode does not exist", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        10000000
      );

      await helper.assert.throwsWithReason(async () => {
        await drs.redeem(10000000, "vTHB")
      }, 'DigitalReserveSystem.redeem: collateralAssetCode does not exist');
    });
    it("should fail, valid price not found", async () => {

      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVTHB.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVTHB.givenMethodReturnAddress(
        stableCreditVTHB.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        0
      );

      try {
        await drs.redeem(10000000, "vTHB");
      } catch (err) {
        assert.equal("DigitalReserveSystem._validateAssetCode: valid price not found", err.reason)
      }
    });
  });

  describe("GetExchangeRate", async () => {
    it("should get exchange rate correctly", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        10000000
      );

      const result = await drs.getExchange("vUSD");
      const BN = web3.utils.BN;
      const pricePerAssetUnit = new BN(result[2]).toNumber();

      assert.equal("vUSD", result[0]);
      assert.equal(web3.utils.padRight(0x0, 64), result[1]);
      assert.equal(0, pricePerAssetUnit);
    });

    it("should fail, invalid assetCode format", async () => {
      const expectedErr = "Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem.getExchange: invalid assetCode format";
      try {
        await drs.getExchange("");
      } catch (err) {
        assert.equal(expectedErr, err)
      }

      try {
        await drs.getExchange("vUSDTHBKHRSGDJPYEURBDHUSSkbph");
      } catch (err) {
        assert.equal(expectedErr, err)
      }
    });

    it("should fail, stableCredit not exist", async () => {
      try {
        await drs.getExchange("vUSD");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: stableCredit not exist", err)
      }
    });

    it("should fail, collateralAsset must be the same", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      try {
        await drs.getExchange("vUSD");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: collateralAsset must be the same", err)
      }
    });

    it("should fail, valid price not found", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        0
      );

      try {
        await drs.getExchange("vUSD");
      } catch (err) {
        assert.equal("Error: Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: valid price not found", err)
      }
    });
  });

  describe("GetCollateralHealthCheck", async () => {
    it("should collateral health check correctly", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditVUSD.givenMethodReturn(
        stableCreditVUSD.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii('VELO')]).toString("hex")
      );
      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.totalSupply().encodeABI(),
        1000000000000
      );
      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii("")).encodeABI(),
        1000000000000
      );

      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.peggedValue().encodeABI(),
        15000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnUint(
        veloCollateralAsset.contract.methods.balanceOf(stableCreditVUSD.address).encodeABI(),
        100000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        100000000000
      );

      const result = await drs.collateralHealthCheck("vUSD");
      const BN = web3.utils.BN;
      const collateralAddress = result[0];
      const requiredAmount = new BN(result[2]).toNumber();
      const presentAmount = new BN(result[3]).toNumber();
      assert.equal(web3.utils.hexToUtf8(result[1]), "VELO");

      assert.equal(1500000000, requiredAmount); //100000000 * 15000000 * 100000000 / 10000000
      assert.equal(100000000, presentAmount);
      assert.equal(veloCollateralAsset.address, collateralAddress);
    });

    it("should fail, invalid assetCode format", async () => {
      const expectedErr = "Returned error: VM Exception while processing transaction: revert DigitalReserveSystem.collateralHealthCheck: invalid assetCode format";

      await helper.assert.throwsWithMessage(async () => {
        await drs.collateralHealthCheck("");
      }, expectedErr);

      await helper.assert.throwsWithMessage(async () => {
        await drs.collateralHealthCheck("vTHBBBBBBBBBB");
      }, expectedErr);
    });

    it("should fail, stableCredit not exist", async () => {
      await helper.assert.throwsWithMessage(async () => {
        await drs.collateralHealthCheck("vUSD");
      }, 'Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: stableCredit not exist');
    });

    it("should fail, collateralAsset must be the same", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      await helper.assert.throwsWithMessage(async () => {
        await drs.collateralHealthCheck("vUSD");
      }, 'Returned error: VM Exception while processing transaction: revert DigitalReserveSystem._validateAssetCode: collateralAsset must be the same');
    });

    it("should fail, collateralAssetCode does not exist", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        10000000
      );

      await helper.assert.throwsWithMessage(async () => {
        await drs.collateralHealthCheck("vUSD");
      }, 'Returned error: VM Exception while processing transaction: revert DigitalReserveSystem.collateralHealthCheck: collateralAssetCode does not exist');
    });
  });

  describe("Rebalance", async () => {
    it("should rebalance correctly requiredAmount > presentAmount", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditVUSD.givenMethodReturn(
        stableCreditVUSD.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii('VELO')]).toString("hex")
      );
      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.totalSupply().encodeABI(),
        1000000000000
      );
      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii("")).encodeABI(),
        100000000000
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.peggedValue().encodeABI(),
        15000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnUint(
        veloCollateralAsset.contract.methods.balanceOf(stableCreditVUSD.address).encodeABI(),
        100000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        100000000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      const result = await drs.rebalance("vUSD");
      truffleAssert.eventEmitted(result, 'Rebalance', event => {
        const BN = web3.utils.BN;
        const eventRequiredAmount = new BN(event.requiredAmount).toNumber();
        const eventPresentAmount = new BN(event.presentAmount).toNumber();
        return event.assetCode === "vUSD"
          && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
          && eventRequiredAmount === 150000000
          && eventPresentAmount === 100000000
      }, 'contract should emit the event correctly');
    });

    it("should rebalance correctly requiredAmount < presentAmount", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditVUSD.givenMethodReturn(
        stableCreditVUSD.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii('VELO')]).toString("hex")
      );
      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.totalSupply().encodeABI(),
        1000000000000
      );
      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii("")).encodeABI(),
        100000000000
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.peggedValue().encodeABI(),
        15000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnUint(
        veloCollateralAsset.contract.methods.balanceOf(stableCreditVUSD.address).encodeABI(),
        200000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        100000000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnBool(
        stableCreditVUSD.contract.methods.transferCollateralToReserve(10000000).encodeABI(),
        true
      );

      const result = await drs.rebalance("vUSD");
      truffleAssert.eventEmitted(result, 'Rebalance', event => {
        const BN = web3.utils.BN;
        const eventRequiredAmount = new BN(event.requiredAmount).toNumber();
        const eventPresentAmount = new BN(event.presentAmount).toNumber();
        return event.assetCode === "vUSD"
          && web3.utils.hexToUtf8(event.collateralAssetCode) === 'VELO'
          && eventRequiredAmount === 150000000
          && eventPresentAmount === 200000000
      }, 'contract should emit the event correctly');
    });

    it("should fail, invalid assetCode format", async () => {
      const expectedErr = "DigitalReserveSystem.rebalance: invalid assetCode format";

      await helper.assert.throwsWithReason(async () => {
        await drs.rebalance("");
      }, expectedErr);

      await helper.assert.throwsWithReason(async () => {
        await drs.rebalance("vTHBBBBBBBBBB");
      }, expectedErr);
    });

    it("should fail, stableCredit not exist", async () => {
      await helper.assert.throwsWithReason(async () => {
        await drs.rebalance("vUSD");
      }, 'DigitalReserveSystem._validateAssetCode: stableCredit not exist');
    });

    it("should fail, collateralAsset must be the same", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        otherCollateralAsset.address
      );

      await helper.assert.throwsWithReason(async () => {
        await drs.rebalance("vUSD");
      }, 'DigitalReserveSystem._validateAssetCode: collateralAsset must be the same');
    });

    it("should fail, collateralAssetCode does not exist", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        10000000
      );

      await helper.assert.throwsWithReason(async () => {
        await drs.rebalance("vUSD");
      }, 'DigitalReserveSystem.rebalance: collateralAssetCode does not exist');
    });

    it("should not rebalance, rebalance of collateral have been equal none require a rebalance", async () => {
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        stableCreditVUSD.address
      );
      await mocks.stableCreditVUSD.givenMethodReturn(
        stableCreditVUSD.contract.methods.collateralAssetCode().encodeABI(),
        '0x' + abi.rawEncode(['bytes32'], [Web3.utils.fromAscii('VELO')]).toString("hex")
      );
      await mocks.stableCreditVUSD.givenMethodReturnAddress(
        stableCreditVUSD.contract.methods.collateral().encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.totalSupply().encodeABI(),
        1000000000000
      );
      await mocks.heart.givenMethodReturnUint(
        heart.contract.methods.getCollateralRatio(Web3.utils.fromAscii("")).encodeABI(),
        100000000000
      );
      await mocks.stableCreditVUSD.givenMethodReturnUint(
        stableCreditVUSD.contract.methods.peggedValue().encodeABI(),
        15000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        veloCollateralAsset.address
      );
      await mocks.veloCollateralAsset.givenMethodReturnUint(
        veloCollateralAsset.contract.methods.balanceOf(stableCreditVUSD.address).encodeABI(),
        150000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getPriceContract(Web3.utils.fromAscii("")).encodeABI(),
        price.address
      );

      await mocks.price.givenMethodReturnUint(
        price.contract.methods.get().encodeABI(),
        100000000000
      );
      await mocks.heart.givenMethodReturnAddress(
        heart.contract.methods.getReserveManager().encodeABI(),
        reserveManager.address
      );

      const result = await drs.rebalance("vUSD");
      // this implies that rebalance returns false (no event emitted)
      assert.equal(result.logs.length, 0);

    });
  });
});

