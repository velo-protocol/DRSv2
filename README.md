<div align="center">
<a href="https://velo.org"><img alt="Stellar" src="https://raw.githubusercontent.com/velo-protocol/assets/master/images/logo.png" width="368" /></a>
</div>
<br>

# Introduction
This repository contains the implementation of Velo Protocol in Solidity.

# Get Started
## Installation

To install DRSv2, you need to install 
[Truffle](https://www.trufflesuite.com/docs/truffle/getting-started/installation) and 
[Ganache](https://www.trufflesuite.com/ganache) first.

1. You can use the command below to install [Truffle](https://www.trufflesuite.com/docs/truffle/getting-started/installation). Please note that Truffle requires NodeJS v8.9.4 or later.

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
$ truffle migrate --reset --network development
```

## Run unit test with solidity-coverage
You can use the command below to run all unit test with full coverage report
```sh
$ yarn solidity-coverage
```