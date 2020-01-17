const TestLLrevert = artifacts.require("TestLLrevert");

contract("LL", async accounts => {
    let testLLrevert;
    let item2 = accounts[0];
    let catchRevert = require("../testhelper/exceptions.js").catchRevert;
  
    beforeEach(async () => {
        testLLrevert = await TestLLrevert.new(accounts);
    });

    describe("TestLL and LL contract, assertion revert test:", function() {
        before(async function() {
            testLLrevert = await TestLLrevert.new();
        });

        it("should abort with an revert", async function() {
            // should fail with revert messsage "addr is already in the list"
            await catchRevert(testLLrevert.testRevert());
        });
    });
});
