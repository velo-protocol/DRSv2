require('dotenv').config();

const PrivateKeyProvider = require("truffle-privatekey-provider");
const web3 = require('web3');

module.exports = {
  networks: {
    local: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    },
    dev: {
      provider: () => new PrivateKeyProvider(process.env.DEV_SCC_PK, process.env.DEV_SCC_HOST),
      network_id: "*",
      gasPrice: 1000000000,
      gas: 6357193,
      production: false
    },
    evry: {
      provider: () => new PrivateKeyProvider(process.env.EVRY_SCC_PK, process.env.EVRY_SCC_HOST),
      network_id: process.env.EVRY_NETWORK_ID,
      gas: 8000000,
      gasPrice: web3.utils.toWei("1", "gwei"),
      production: false,
      confirmations: 1,
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
      version: '0.5.15',
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        },
        evmVersion: 'constantinople',
      }
    }
  },
  plugins: ["solidity-coverage"]
};
