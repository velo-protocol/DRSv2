const LL = artifacts.require("LL");

contract("LL", async accounts => {
    let ll;
    let item1 = accounts[0];
    let item2 = accounts[1];
    let catchRevert = require("../testhelper/exceptions.js").catchRevert;
  
    beforeEach(async () => {
      ll = await LL.new(accounts);
    });

    describe("LL add assertion:", function() {
        before(async function() {
            ll = await LL.new();
        });
        it("should complete successfully", async function() {
            await ll.add(item1); // should add successfully
            await ll.add(item2); // should add successfully
            assert.equal(ll[1], item1, "item1 is the 1st item");
        });
        it("should abort with an revert", async function() {
            await catchRevert(ll.add(item1)); // should fail with revert messsage "addr is already in the list"
        });
    });
});
