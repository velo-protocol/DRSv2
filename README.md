<div style="text-align: center">
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

## Script usage
#### Deployment
##### 1. Deploy locally, suit when developing on local machine
```sh
$ yarn run reset
```

##### 2. Deploy to dev environment (required DEV_SCC_HOST, DEV_SCC_PK configured correctly)
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
```sh
$ yarn run test:cov
```
