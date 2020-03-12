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

  describe("Void", async () => {
    it("should void correctly", async () => {
      const result = await medianizer.void();

      assert.equal(false, await medianizer.active());

      truffleAssert.eventEmitted(result, 'MedianVoid', event => {
        return event.medianizer === medianizer.address
          && event.isActive === false
      }, 'contract should emit the event correctly');
    });
  });


});
