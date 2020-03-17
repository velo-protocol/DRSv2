const fs = require("fs");

module.exports = async function (contractName, address) {

  const filePath = './contract-addresses.json';

  let addresses = {};

  try {
    let data = fs.readFileSync(filePath);
    addresses = JSON.parse(data);
  } catch (err) {
    if (!err.toString().includes("no such file or directory")) {
      throw err;
    }
  }
  addresses[contractName] = address;
  const jsonString = JSON.stringify(addresses);
  fs.writeFileSync(filePath, jsonString);


};
