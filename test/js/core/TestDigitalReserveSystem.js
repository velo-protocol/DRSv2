const DigitalReserveSystem = artifacts.require("DigitalReserveSystem");
const Heart = artifacts.require("Heart");
const PriceFeeders = artifacts.require('PriceFeeders');
const StableCredit = artifacts.require("StableCredit");
const MockContract = artifacts.require("MockContract");

const Web3 = require('web3');
const truffleAssert = require('truffle-assertions');
const helper = require('../testhelper');

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
                mock.address
            );
            await mock.givenMethodReturnUint(
                helper.methodABI(pf, "getMedianPrice", [helper.address(0)]),
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
                mock.address
            );
            await mock.givenMethodReturnUint(
                helper.methodABI(pf, "getMedianPrice", [helper.address(0)]),
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

