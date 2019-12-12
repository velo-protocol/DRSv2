const PrivateKeyProvider = require("truffle-privatekey-provider");

module.exports = {
  networks: {
    development: {
      provider: () => new PrivateKeyProvider(process.env.SCC_PK, process.env.SCC_HOST),
      network_id: "*"
    },
    testnet: {
      provider: () => new PrivateKeyProvider(process.env.SCC_PK, process.env.SCC_HOST),
      network_id: 15,
      gasPrice: 1000000000,
      production: false
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
  }
}
