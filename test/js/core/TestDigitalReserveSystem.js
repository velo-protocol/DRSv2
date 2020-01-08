const DigitalReserveSystem = artifacts.require("DigitalReserveSystem");
const Heart = artifacts.require("Heart");
const PriceFeeders = artifacts.require('PriceFeeders');
const StableCredit = artifacts.require("StableCredit");
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

            truffleAssert.eventEmitted(result, 'Setup', async event => {
                assert.equal(event.peggedCurrency, web3.utils.padRight(web3.utils.utf8ToHex("USD"), 64));
                assert.equal(event.peggedValue, 1);
                assert.equal(event.collateralAssetCode, web3.utils.padRight(web3.utils.utf8ToHex("VELO"), 64));
                assert.ok(web3.utils.isAddress(event.assetAddress));

                const stableCredit = await StableCredit.at(event.assetAddress);
                const name = await stableCredit.name();
                assert.equal(name, "vUSD");

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
                assert.equal(err.reason, "DigitalReserveSystem.onlyTrustedPartner: caller must be a trusted partner")
            }
        });

        it("should fail, invalid assetCode format", async () => {
            await mock.givenMethodReturnBool(
                heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
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
                heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
                true
            );
            await mock.givenMethodReturnAddress(
                heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
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
                heart.contract.methods.isTrustedPartner(accounts[0]).encodeABI(),
                true
            );
            await mock.givenMethodReturnAddress(
                heart.contract.methods.getStableCreditById(Web3.utils.fromAscii("")).encodeABI(),
                web3.utils.padLeft("0x0", 40)
            );
            await mock.givenMethodReturnAddress(
                heart.contract.methods.getCollateralAsset(Web3.utils.fromAscii("")).encodeABI(),
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
    })
});

