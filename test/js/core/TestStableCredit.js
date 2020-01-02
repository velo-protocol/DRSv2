const StableCredit = artifacts.require("StableCredit");
const Token = artifacts.require("Token");
const Heart = artifacts.require("Heart");

const Web3 = require('web3');

contract("StableCredit test", async accounts => {
  let stableCredit, veloToken, heart;

  beforeEach(async () => {
    veloToken = await Token.new('Velo', 'VELO', 7);
    heart = await Heart.new();
    await heart.setDrsAddress(accounts[0]);

    stableCredit = await StableCredit.new(
        Web3.utils.fromAscii("USD"),
        accounts[1],
        Web3.utils.fromAscii(veloToken.symbol()),
        veloToken.address,
        'testCredit',
        '1',
        heart.address
    );
  });

  it("should mint correctly", async () => {
    await stableCredit.mint(accounts[0], 10);

    const balance = await stableCredit.balanceOf(accounts[0]);

    assert.equal(balance, 10, "stableCredit.balanceOf account 0 should be 10");
  });

  it("should burn correctly", async () => {
    await stableCredit.mint(accounts[0], 100);
    await stableCredit.burn(accounts[0], 10);

    const balance = await stableCredit.balanceOf(accounts[0]);

    assert.equal(balance, 90, "stableCredit.balanceOf account 0 should be 90");
  });

});
