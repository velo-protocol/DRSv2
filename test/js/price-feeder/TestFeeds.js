const Feeder = artifacts.require("Feeder");

contract("Feeder test", async accounts => {
  let feeds;

  beforeEach(async () => {
    feeds = await Feeder.new(accounts[0], web3.utils.fromAscii("USD"), web3.utils.fromAscii("VELO"));
  });

  it("should set correctly", async () => {
    const expected = 120;
    await feeds.set(expected);

    const value = await feeds.value();

    assert.equal(value, expected, "feeds.value should be 120");
  });

  it("should get correctly", async () => {
    const expected = 120;
    await feeds.set(expected);

    const value = await feeds.get();

    assert.equal(value, expected, "feeds.value should be 120");
  });

  it("should getWithError correctly", async () => {
    const expected = 120;
    await feeds.set(expected);

    const result = await feeds.getWithError();
    const value = result['0'];
    const active = result['1'];

    assert.equal(value, expected, "value should be 120");
    assert.equal(active, false, "active should be false");
  });

  it("should disable correctly", async () => {
    await feeds.disable();
    const active = await feeds.active();

    assert.equal(active, false, "active should be false");
  });
});
