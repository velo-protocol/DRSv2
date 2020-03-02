const FeederFactory = artifacts.require('FeederFactory');

module.exports = async (deployer, network, accounts) => {
  await deployer.deploy(FeederFactory);

  if (network === 'development' || network === 'local' || network === 'dev') {
    const feederFactory = await FeederFactory.deployed();

    const veloBytes32 = web3.utils.fromAscii("VELO");
    const usdBytes32 = web3.utils.fromAscii("USD");
    const thbBytes32 = web3.utils.fromAscii("THB");
    const sgdBytes32 = web3.utils.fromAscii("SGD");

    let result = await feederFactory.create(usdBytes32, veloBytes32);
    console.log(`USD Feeder: ${result.logs[0].args.feedAddr}`);

    result = await feederFactory.create(thbBytes32, veloBytes32);
    console.log(`THB Feeder: ${result.logs[0].args.feedAddr}`);

    result = await feederFactory.create(sgdBytes32, veloBytes32);
    console.log(`SGD Feeder: ${result.logs[0].args.feedAddr}`);
  }
};
