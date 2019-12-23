const StableCredit = artifacts.require("StableCredit");
const Token = artifacts.require("Token");

const Web3 = require('web3');

let stableCredit;
let veloToken;

contract("StableCredit test", async accounts => {

  beforeEach(async () => {
    veloToken = await Token.new('Velo', 'VELO', 7);
    stableCredit = await StableCredit.new(Web3.utils.fromAscii("USD"), accounts[1],
        Web3.utils.fromAscii(veloToken.symbol()), veloToken.address, 'testCredit', '1');
  });

  it("should mint correctly", async () => {
    stableCredit.mint(accounts[0], 10);

    const balance = await stableCredit.balanceOf(accounts[0]);

    assert.equal(balance, 10, "stableCredit.balanceOf account 0 should be 10");
  });

  it("should burn correctly", async () => {
    stableCredit.mint(accounts[0], 100);
    stableCredit.burn(accounts[0], 10);

    const balance = await stableCredit.balanceOf(accounts[0]);

    assert.equal(balance, 90, "stableCredit.balanceOf account 0 should be 90");
  });

});
