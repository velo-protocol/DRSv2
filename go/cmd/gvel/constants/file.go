package constants

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var FsBaseDir = path.Join(os.Getenv("HOME"), "/.gvel")
var FsConfigFileNameFormat = path.Join(FsBaseDir, "%s-config.json")

var FsSharedConfigFile = fmt.Sprintf(FsConfigFileNameFormat, "shared")
var FsTestnetConfigFile = fmt.Sprintf(FsConfigFileNameFormat, strings.ToLower(EnvTestNet))
var FsMainnetConfigFile = fmt.Sprintf(FsConfigFileNameFormat, strings.ToLower(EnvMainNet))
