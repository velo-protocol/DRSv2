const TestLLrevert = artifacts.require("TestLLrevert");

contract("LL", async accounts => {
    let testLLrevert;
    let catchRevert = require("../testhelper/exceptions.js").catchRevert;

    beforeEach(async () => {
        testLLrevert = await TestLLrevert.new(accounts);
        testLLrevert.initLL();
    });

    describe("TestLL and LL contract, assertion revert test:", function() {
        it("should fail with revert", async function() {
            // expected error messsage "addr is already in the list"
            await catchRevert(testLLrevert.llDupAddRevert());
        });

        it("should fail with revert", async function() {
            // expected error messsage "addr not whitelisted yet."
            await catchRevert(testLLrevert.llRemoveNoExistRevert());
        });

        it("should fail with revert", async function() {
            // expected error messsage "wrong prevConsumer."
            await catchRevert(testLLrevert.llRemoveWrongPreRevert());
        });
    });
});
