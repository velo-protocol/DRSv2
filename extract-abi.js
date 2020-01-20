const fs = require("fs");
const dirName = "./build/contracts/";
const abiDirName = dirName;

fs.readdir(dirName, (err, fileNames) => {
  if (err) throw err;

  for (const fileName of fileNames) {
    fs.readFile(dirName + fileName, 'utf-8',  (err, content) => {
      if (err) throw err;

      // parse contract
      const contract = JSON.parse(content);

      const newFileName = fileName.split(".")[0] + "." + "abi";
      // stringify abi and write to file
      fs.writeFileSync(abiDirName + newFileName, JSON.stringify(contract.abi));
    });
  }
});
