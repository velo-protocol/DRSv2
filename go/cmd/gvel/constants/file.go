package constants

import (
	"fmt"
	"os"
	"path"
)

var FsBaseDir = path.Join(os.Getenv("HOME"), "/.gvel")
var FsConfigFileNameFormat = path.Join(FsBaseDir, "%s-config.json")

var FsSharedConfigFile = fmt.Sprintf(FsConfigFileNameFormat, "shared")
var FsTestnetConfigFile = fmt.Sprintf(FsConfigFileNameFormat, EnvMainNet)
var FsMainnetConfigFile = fmt.Sprintf(FsConfigFileNameFormat, EnvTestNet)
