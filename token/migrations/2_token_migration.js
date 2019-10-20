const Token = artifacts.require('Token');

module.exports = async function (deployer) {
    await deployer.deploy(Token, "TIM", "TIM", 7);
    let timToken = await Token.deployed();
    await timToken.mint('0xdEF0D3660Ed8A0165aBfb21020BC5834565c263C', 10000000000000);
};