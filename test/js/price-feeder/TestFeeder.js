const Feeder = artifacts.require("Feeder");

contract("Feeder test", async accounts => {
  let feeder;

  beforeEach(async () => {
    feeder = await Feeder.new(accounts[0], web3.utils.fromAscii("USD"), web3.utils.fromAscii("VELO"));
  });

  it("should post correctly", async () => {
    const expected = 120;
    await feeder.post(expected);

    const value = await feeder.value();

    assert.equal(value, expected, "feeds.value should be 120");
  });

  it("should get correctly", async () => {
    const expected = 120;
    await feeder.post(expected);

    const value = await feeder.get();

    assert.equal(value, expected, "feeds.value should be 120");
  });

  it("should getWithError correctly", async () => {
    const expected = 120;

    const now1 = Math.floor((new Date()).valueOf() / 1000);
    await feeder.post(expected);
    const now2 = Math.floor((new Date()).valueOf() / 1000);

    const result = await feeder.getWithError();
    const value = result['0'];
    const timestamp = result['1'];
    const active = result['2'];


    assert.equal(value, expected, "value should be 120");
    assert.equal(now1<= timestamp && timestamp <= now2, true, "value should be 120");
    assert.equal(active, false, "active should be false");
  });

  it("should disable correctly", async () => {
    await feeder.disable();
    const active = await feeder.active();

    assert.equal(active, false, "active should be false");
  });
});
