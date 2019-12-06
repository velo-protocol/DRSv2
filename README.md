<div align="center">
<a href="https://velo.org"><img alt="Stellar" src="https://raw.githubusercontent.com/velo-protocol/assets/master/images/logo.png" width="368" /></a>
</div>
<br>

DRSv2 is a VELO Protocol Version Smart Contract.

# Get Started
## Installation

To install DRSv2, you need to install 
[Truffle](https://www.trufflesuite.com/docs/truffle/getting-started/installation), 
[Ganache](https://www.trufflesuite.com/ganache) and 
[Metamask](https://metamask.io/) first.

1. The first need [Truffle](https://www.trufflesuite.com/docs/truffle/getting-started/installation) installed (**NodeJS v8.9.4 or later required**), then you can use the below npm command to install truffle.

```sh
$ npm install -g truffle
```

2. Clone this Repository it in your machine and then go to folder DRSv2:

```node
$ yarn install
```

3. After download ganache and installed you need run ganache and run this command
```bash
$ truffle build
$ truffle deploy
```

4. Used Metamask extension from google chrome
 - use custom network to link in your ganache server (http://localhost:7545) 
 - import your ether account from private key to metamask
 - you can interact with metamask for send asset and see your transaction history from metamask or ganache
 
5. DRSv2 Environment Migration run this command
 
```bash
$ truffle migrations/
```