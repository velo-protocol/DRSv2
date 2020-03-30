require('dotenv').config();

const HDWalletProvider = require("@truffle/hdwallet-provider");

const privateKeys = [
  process.env.DEV_SCC_PK,
  process.env.USD_VELO_PK,
  process.env.THB_VELO_PK,
  process.env.SGD_VELO_PK,
];

module.exports = {
  networks: {
    local: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    },
    dev: {
      provider: () => new HDWalletProvider(privateKeys, process.env.DEV_SCC_HOST,0,4),
      network_id: "*",
      gasPrice: 1000000000,
      gas: 6357193,
      production: false
    },
    coverage: {
      host: "localhost",
      network_id: "*",
      port: 8555,
      gas: 0xfffffffffff,
      gasPrice: 0x01
    }
  },
  // Set default mocha options here, use special reporters etc.
  mocha: {
    // timeout: 100000
  },
  compilers: {
    solc: {
      settings: {
       optimizer: {
         enabled: true,
         runs: 200
       },
      }
    }
  },
  plugins: ["solidity-coverage"]
};
