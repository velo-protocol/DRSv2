const DigitalReserveSystem = artifacts.require("DigitalReserveSystem");
const Heart = artifacts.require("Heart");
const PriceFeeders = artifacts.require('PriceFeeders');
const MockContract = artifacts.require("MockContract");
const Web3 = require('web3');
const truffleAssert = require('truffle-assertions');
let drs, heart, mock, pf;
contract("DigitalReserveSystem test", async accounts => {
  before(async () => {
    heart = await Heart.new();
    pf = await PriceFeeders.new();
    mock = await MockContract.new();
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
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
        true
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x0", 40)
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
        web3.utils.padLeft("0x1", 40)
      );
      await mock.givenMethodReturnBool(
        heart.contract.methods.isLinkAllowed(Web3.utils.fromAscii("")).encodeABI(),
        true
      );
      await mock.givenMethodReturnAddress(
        heart.contract.methods.getPriceFeeders().encodeABI(),
        mock.address
      );
      await mock.givenMethodReturnUint(
        pf.contract.methods.getMedianPrice(Web3.utils.fromAscii("")).encodeABI(),
        1
      );
      const result = await drs.setup(
        Web3.utils.fromAscii("VELO"),
        Web3.utils.fromAscii("USD"),
        "vUSD",
        1
      );
      truffleAssert.eventEmitted(result, 'Setup', event => {
        assert.equal(event.peggedCurrency, web3.utils.padRight(web3.utils.utf8ToHex("USD"), 64));
        assert.equal(event.peggedValue, 1);
        assert.equal(event.collateralAssetCode, web3.utils.padRight(web3.utils.utf8ToHex("VELO"), 64));
        assert.ok(web3.utils.isAddress(event.assetAddress));
        return true;
      }, 'contract should emit the event correctly');
    });
    it("should fail, caller must be trusted partner", async () => {
      await mock.givenMethodReturnBool(
        heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
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
        assert.equal("DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner", err.reason)
      }
    })
  })
});