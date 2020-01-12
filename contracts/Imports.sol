pragma solidity ^0.5.0;

// We import the contract so truffle compiles it, and we have the ABI
// available when working from truffle console.
import "@gnosis.pm/mock-contract/contracts/MockContract.sol";

// Declaring imports as library to avoid crash when running 'truffle test' on solidity file
library Imports {}
