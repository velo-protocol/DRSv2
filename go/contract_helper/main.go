package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
)

// Client struct
type Client struct {
	privateKey  ecdsa.PrivateKey
	publicKey   common.Address
	drs         *vabi.DigitalReserveSystem
	heart       *vabi.Heart
	priceFeeder *vabi.PriceFeeders
	collateral  *vabi.Token
}

func NewClient(contractUrl, drsAddress, heartAddress, priceFeederAddress string) Client {
	conn, err := ethclient.Dial(contractUrl)
	if err != nil {
		panic(err)
	}

	drsContract, err := vabi.NewDigitalReserveSystem(common.HexToAddress(drsAddress), conn)
	if err != nil {
		panic(err)
	}

	heartContract, err := vabi.NewHeart(common.HexToAddress(heartAddress), conn)
	if err != nil {
		panic(err)
	}

	priceFeederContract, err := vabi.NewPriceFeeders(common.HexToAddress(priceFeederAddress), conn)
	if err != nil {
		panic(err)
	}

	return Client{

		drs:         drsContract,
		heart:       heartContract,
		priceFeeder: priceFeederContract}
}

func HexToPrivateKey(hex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic("invalid private key")
	}
	return privKey
}

func (i *Client) SetCollateralRatio(assetCode, ratio, privateKey string) string {

	privKey := HexToPrivateKey(privateKey)
	intRatio, _ := utils.StringToAmount(ratio)

	opt := bind.NewKeyedTransactor(privKey)
	opt.GasLimit = constants.GasLimit
	result, err := i.heart.SetCollateralRatio(opt, utils.StringToByte32(assetCode), intRatio)
	if err != nil {
		panic(err)
	}

	return result.Hash().String()

}

func (i *Client) GetCollateralRatio(assetCode string) string {
	result, err := i.heart.GetCollateralRatio(nil, utils.StringToByte32(assetCode))
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(result)
}

func (i *Client) GetCollateralBalanceOf(address string) string {
	result, _ := i.collateral.BalanceOf(nil, common.HexToAddress(address))
	return utils.AmountToString(result)
}

func (i *Client) SetPrice(assetCode, currency, price string) string {

	intRatio, _ := utils.StringToAmount(price)

	opt := bind.NewKeyedTransactor(&i.privateKey)
	opt.GasLimit = constants.GasLimit
	result, err := i.priceFeeder.SetPrice(opt, utils.StringToByte32(assetCode), utils.StringToByte32(currency), intRatio)
	if err != nil {
		panic(err)
	}
	return result.Hash().String()

}

func (i *Client) GetPrice(assetCode, currency string) string {
	bytes32Ty, _ := abi.NewType("bytes32", "", nil)

	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
	}
	bytes, _ := arguments.Pack(
		utils.StringToByte32(assetCode),
		utils.StringToByte32(currency),
	)

	hash := crypto.Keccak256(bytes)

	linkId := utils.BytesToBytes32(hash)
	price, err := i.priceFeeder.GetMedianPrice(nil, linkId)
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(price)
}

func main() {
	client := NewClient("http://127.0.0.1:7545", "0xb06601682f9c32A16C9F3aBE70aACe03676C09C0", "0xf5A513a8CD2ba17836954eC7f3868181302fEfc5", "0x584B126780ff0d68bFCB920dE0F56F4eF1D9aFe3")
	client.GetPrice("VELO", "USD")
}
