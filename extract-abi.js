const fs = require("fs");
const dirName = "./build/contracts/";
const abiDirName = dirName + 'abi/';
const binDirName = dirName + 'bin/';

fs.readdir(dirName, (err, fileNames) => {
  if (err) throw err;

  try {
    fs.mkdirSync(abiDirName);
  } catch (e) {

  }
  try {
    fs.mkdirSync(binDirName);
  } catch (e) {
  }

  for (const fileName of fileNames) {
    if (fileName.split('.')[1] === "json") {
      fs.readFile(dirName + fileName, 'utf-8', (err, content) => {
        if (err) throw err;

        // parse contract
        const contract = JSON.parse(content);

        const abiFileName = fileName.split(".")[0] + "." + "abi";
        const binFileName = fileName.split(".")[0] + "." + "bin";
        // stringify abi and write to file
        fs.writeFileSync(abiDirName + abiFileName, JSON.stringify(contract.abi));
        fs.writeFileSync(binDirName + binFileName, contract.bytecode);
      });
    }
  }
});
