const Medianizer = artifacts.require('Medianizer');

const truffleAssert = require('truffle-assertions');

contract("Medianizer test", async accounts => {
  let medianizer;

  const veloBytes32 = web3.utils.fromAscii("VELO");
  const usdBytes32 = web3.utils.fromAscii("USD");


  beforeEach(async () => {
    medianizer = await Medianizer.new();
    await medianizer.initialize(accounts[0], usdBytes32, veloBytes32)
  });

  describe("Add Feeder", async () => {
    it("should add feeder correctly", async () => {
      const result = await medianizer.addFeeder(accounts[1]);

      truffleAssert.eventEmitted(result, 'FeederAdd', event => {
        return event.caller === accounts[0]
          && event.medianizer === medianizer.address
          && event.feeder === accounts[1]
      }, 'contract should emit the event correctly');
    });
  });

  describe("Remove Feeder", async () => {
    it("should remove feeder correctly", async () => {
      await medianizer.addFeeder(accounts[1]);
      const result = await medianizer.removeFeeder(accounts[1]);

      truffleAssert.eventEmitted(result, 'FeederRemove', event => {
        return event.caller === accounts[0]
          && event.medianizer === medianizer.address
          && event.feeder === accounts[1]
      }, 'contract should emit the event correctly');
    });

    it("should fail, feeder address does not exist", async () => {
      try {
        await medianizer.removeFeeder(accounts[1]);
      } catch (err) {
        assert.equal(err.reason, "Medianizer.removeFeeder: address does not exist")
      }
    });
  });

  describe("Void", async () => {
    it("should void correctly", async () => {
      const result = await medianizer.void();

      assert.equal(false, await medianizer.active());

      truffleAssert.eventEmitted(result, 'MedianVoid', event => {
        return event.caller === accounts[0]
          && event.medianizer === medianizer.address
          && event.isActive === false
      }, 'contract should emit the event correctly');
    });
  });

  describe("SetValidityPeriod", async () => {
    it("should setValidityPeriod correctly", async () => {
      const validityPeriod = 300;
      await medianizer.setValidityPeriod(validityPeriod);

      assert.equal(validityPeriod, await medianizer.getValidityPeriod(), "medianizer.validityPeriod validityPeriod be 300");

    });

    it("should fail, validityPeriod must be greater than 0", async () => {
      try {
        await medianizer.setValidityPeriod(-1);
      } catch (err) {
        assert.equal(err.reason, "Medianizer.setValidityPeriod: validityPeriod must be greater than 0")
      }
    });

    it("should fail, validityPeriod must be greater than 0", async () => {
      try {
        await medianizer.setValidityPeriod(-1999999393999999);
      } catch (err) {
        assert.equal(err.reason, "Medianizer.setValidityPeriod: validityPeriod must be greater than 0")
      }
    });
    it("should fail, validityPeriod must be greater than 0", async () => {
      try {
        await medianizer.setValidityPeriod(1999999393999999);
      } catch (err) {
        assert.equal(err.reason, "Medianizer.setValidityPeriod: validityPeriod must not be greater than maxValidityPeriod")
      }
    });
  });

  describe("GetValidityPeriod", async () => {
    it("should getValidityPeriod correctly", async () => {
      const validityPeriod = 300;
      await medianizer.setValidityPeriod(validityPeriod);

      assert.equal(validityPeriod, await medianizer.getValidityPeriod(), "medianizer.validityPeriod validityPeriod be 300");

    });
  });

});
