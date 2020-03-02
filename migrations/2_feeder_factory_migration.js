const FeederFactory = artifacts.require('FeederFactory');
const Feeder = artifacts.require('Feeder');

module.exports = async (deployer, network, accounts) => {
  await deployer.deploy(FeederFactory);

  if (network === 'development' || network === 'local' || network === 'dev') {
    const feederFactory = await FeederFactory.deployed();

    const veloBytes32 = web3.utils.fromAscii("VELO");
    const usdBytes32 = web3.utils.fromAscii("USD");
    const thbBytes32 = web3.utils.fromAscii("THB");
    const sgdBytes32 = web3.utils.fromAscii("SGD");

    // 1. Create 3 feeders
    let result = await feederFactory.create(usdBytes32, veloBytes32);
    const usdFeederAddr = result.logs[0].args.feedAddr;
    console.log(`USD Feeder: ${usdFeederAddr}`);

    result = await feederFactory.create(thbBytes32, veloBytes32);
    const thbFeederAddr = result.logs[0].args.feedAddr;
    console.log(`THB Feeder: ${thbFeederAddr}`);

    result = await feederFactory.create(sgdBytes32, veloBytes32);
    const sgdFeederAddr = result.logs[0].args.feedAddr;
    console.log(`SGD Feeder: ${sgdFeederAddr}`);

    // 2. Feed the price
    const usdFeeder = await Feeder.at(usdFeederAddr);
    await usdFeeder.set(12000000); // 1.2 USD/VELO

    const thbFeeder = await Feeder.at(thbFeederAddr);
    await thbFeeder.set(53000000); // 5.3 THB/VELO

    const sgdFeeder = await Feeder.at(sgdFeederAddr);
    await sgdFeeder.set(27000000); // 2.7 THB/VELO
  }
};
