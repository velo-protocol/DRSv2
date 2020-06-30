const Feeder = artifacts.require("Feeder");

contract("Feeder test", async accounts => {
  let feeder;

  beforeEach(async () => {
    feeder = await Feeder.new(accounts[0],accounts[0],accounts[0],accounts[0], web3.utils.fromAscii("USD"), web3.utils.fromAscii("VELO"));
  });

  it("should startOracle  getLastPrice correctly", async () => {
    const expected = 120;
    await feeder.startOracle(expected);
    const value = await feeder.getLastPrice();
    assert.equal(value[0], expected, "feeds.value should be 120");
  });

  it("should commitPrice setValue correctly", async () => {
    const expected = 120;
    await feeder.startOracle(expected);
    await feeder.setValue(3,0);
    await feeder.commitPrice(200);
    const value = await feeder.getLastPrice();
    assert.equal(value[0], expected, "feeds.value should be 120");
  });

  it("should commitPrice setValue correctly", async () => {
    const expected = 200;
    await feeder.startOracle(120);
    await feeder.setValue(3,0);
    await feeder.commitPrice(expected);
    await sleep(3000)
    await feeder.commitPrice(expected);
    const value = await feeder.getLastPrice();
    assert.equal(value[0], expected, "feeds.value should be 200");
  });


});


function sleep (time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}
