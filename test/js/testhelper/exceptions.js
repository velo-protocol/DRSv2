const ALREADY_INTHELIST_ERROR = 
    "Returned error: VM Exception while processing transaction: revert addr is already in the list -- Reason given: addr is already in the list.";

async function tryCatch(promise, message) {
    try {
        await promise;
        throw null;
    }
    catch (error) {
        assert(error, "Expected an error but did not get one");
        assert(
            error.message === ALREADY_INTHELIST_ERROR, 
            "Expected an error message: '" + ALREADY_INTHELIST_ERROR + "' but got '" + error.message + "' instead",
            );
    }
};

module.exports = {
    catchRevert            : async function(promise) {await tryCatch(promise, "revert"             );},
    catchOutOfGas          : async function(promise) {await tryCatch(promise, "out of gas"         );},
    catchInvalidJump       : async function(promise) {await tryCatch(promise, "invalid JUMP"       );},
    catchInvalidOpcode     : async function(promise) {await tryCatch(promise, "invalid opcode"     );},
    catchStackOverflow     : async function(promise) {await tryCatch(promise, "stack overflow"     );},
    catchStackUnderflow    : async function(promise) {await tryCatch(promise, "stack underflow"    );},
    catchStaticStateChange : async function(promise) {await tryCatch(promise, "static state change");},
};
