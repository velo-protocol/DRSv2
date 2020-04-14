<div style="text-align: center">
<a href="https://velo.org"><img alt="Stellar" src="https://raw.githubusercontent.com/velo-protocol/assets/master/images/logo.png" width="368" /></a>
</div>
<br>

# Introduction
This repository contains the implementation of Velo Protocol version 2, which includes
* Smart contract
* Golang implementations
  * Go SDK (vclient)
  * Go CLI (gvel)

# Getting Started with Smart Contract

## Project Structure
```
├── contracts          - solidity smart contract code stays here
├── migrations         - migration script
├── test               - smart contract test
├── truffle-config.js  - config truffle env and behaviour here
├── package.json
├── yarn.lock
...
└── go (Golang implementations)
```
See [Golang implementations](./go/README.md) here

## Installation

To install DRSv2, you need to install 
NodeJs 10.16.3,
[Truffle](https://www.trufflesuite.com/docs/truffle/getting-started/installation) and 
[Ganache](https://www.trufflesuite.com/ganache) first.

1. You can use the command below to install [Truffle](https://www.trufflesuite.com/docs/truffle/getting-started/installation).

```sh
$ yarn global add truffle
```

2. Clone this repository in your machine, go to folder DRSv2, then use `yarn install` to install all dependencies:

```sh
$ yarn install
```

3. Run Ganache

4. After that you can run the command below to deploy Velo Protocol smart contracts to your local network
```sh
$ yarn run reset
```
5. Now you can interact with the smart contract via `truffle console` 
6. Modify the migration script in `./migrations` to control the how the smart contract should be setup.

## Script usage
#### Deployment
##### 1. Deploy locally, suit when developing on local machine
```sh
$ yarn run reset
```

##### 2. Deploy to dev environment (required variables from .env configured properly, Env File section)
```sh
$ yarn run reset:dev
```


#### Running test
##### 1. Test a single file, suit when focusing on a single file 
```sh
$ yarn run test test/path/to/file.js
or
$ yarn run test test/path/to/file.sol
```

##### 2. Test only js files
```sh
$ yarn run test:js
```

##### 3. Test only sol files
```sh
$ yarn run test:sol
```

##### 4. Test all files
```sh
$ yarn run test:all
```

##### 5. Test all files and generate a coverage report
```shell script
$ yarn run test:cov
```

#### Utils

##### Create Smart Contract .bin and .abi
After smart contract compilation, it yields ABI in .json format. Golang's `abigen` tool requires
ABI in .bin and .abi to generate Go smart contract function. (see more in Go abigen section)
```shell script
# compile your smart contract before proceed (yarn run reset)
$ yarn run abi:extract
```

## Env File
Please copy `.env.example` and rename it to `.env`, config its value properly
```
DEV_SCC_HOST - the rpc url to the Ethereum based chain
DEV_SCC_PK - the private key of the smart contract deployer account (the account must own some ETH balance)
USD_VELO_PK, THB_VELO_PK, SGD_VELO_PK - the private key of the VOracle Module deployer account of each currency pair, this
key can be the same as DEV_SCC_PK for testing purpose
```






