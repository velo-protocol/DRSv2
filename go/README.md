# Getting Started with Golang Implementation
## Project Structure
```
├── README.md
├── abi              - Go's code generated from `abigen`
├── cmd/gvel         - `gvel` cli stays here
├── libs/vclient     - Go's client, used by anyone who want to integrate with VELO
├── contract_helper  - the script here helps QA team do automation easier, consult with them for usage
├── Makefile         - a bunch of shortcuts (see detail in later section)
├── go.mod
└── go.sum
```

## Prerequisite
* Golang 1.12 or higher

## Go VCLient
VClient is built for anyone who want to integrate with VELO. It contains a bunch of express functionality such `credit setup`, 
`credit mint`, `credit redeem`, etc.

### Basic Usage
```go
import (
	"github.com/velo-protocol/DRSv2/go/libs/vclient"
)
func main() {
	client, _ := vclient.NewClient("<SMART_CONTRACT_NODE_URL>", "<PRIVATE_KEY>", ContractAddress{
		DrsAddress:   "<DRS_CONTRACT_ADDRESS>",   // 0x4Db9c67836A3735f63c0eCe4cFBc486bB80732b0
		HeartAddress: "<HEART_CONTRACT_ADDRESS>", // 0x1F1247eDEa84dC392C857A7887203a5640f3f2Fd
	})
    
	output, _ := client.SetupCredit(context.Background(), &vclient.SetupCreditInput{
		CollateralAssetCode: "VELO",
		PeggedCurrency:      "USD",
		AssetCode:           "vUSD",
		PeggedValue:         "1.0",
	})
}
```
See more example usage in example.go from its directory.

### Development Tips
#### Generating mock
Use the `make` command to generate mock files for `vclient`
```shell script
$ make mockgen/all
```

#### Generating Go's abi
Use the `make` command to generate go's abi files
```shell script
$ make abigen/gen/all
```
The file's will be stored in `./abi`

## Go CLI gvel
### Building gvel
```sh
$ cd cmd/gvel
$ make gvel/build
```
or just
```shell script
$ cd go/cmd/gvel
$ go build . 
```

### Development Tips
#### Software structure
The gvel itself follows `Clean Architecture`. The layers is divided into 
* commands - this is the outermost layer, in other word, this layer is user interface
* logic - this layer control the flow of each command input
* repositories - this layer contains two sublayers
    * database - database repo controls how to interact with database, currently we use leveldb
    * vfactory - vfactory repo controls how the `vclient` will be initiated

#### Generating mock
Use the `make` command to generate mock files for `gvel`
```shell script
$ cd cmd/gvel
$ make mockgen/all
```

