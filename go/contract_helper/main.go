package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"io/ioutil"
)

// Client struct
type Client struct {
	conn         *ethclient.Client
	privateKey   *ecdsa.PrivateKey
	publicKey    common.Address
	drs          *vabi.DigitalReserveSystem
	heart        *vabi.Heart
	prices       map[string]*vabi.Price
	collateral   map[string]*vabi.Token
	stableCredit map[string]*vabi.StableCredit
}

func NewClient(contractUrl, drsAddress, heartAddress, privateKey string) Client {
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

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic(err)
	}

	return Client{
		drs:          drsContract,
		heart:        heartContract,
		conn:         conn,
		privateKey:   privKey,
		publicKey:    crypto.PubkeyToAddress(*pubKeyECDSA),
		prices:       map[string]*vabi.Price{},
		collateral:   map[string]*vabi.Token{},
		stableCredit: map[string]*vabi.StableCredit{},
	}
}

func ConfirmTx(ctx context.Context, client *Client, tx *types.Transaction, from common.Address) error {
	receipt, err := bind.WaitMined(ctx, client.conn, tx)
	if err != nil {
		return errors.Wrap(err, "fail to confirm transaction")
	}
	if receipt.Status == 0 {
		untrimmedBytes, err := client.conn.CallContract(ctx, ethereum.CallMsg{
			From:     from,
			To:       tx.To(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		}, receipt.BlockNumber)
		if err != nil {
			return errors.Wrap(err, "fail to get revert message")
		}
		return errors.New(utils.ParseRevertMessage(untrimmedBytes))
	}
	return nil
}

func HexToPrivateKey(hex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic("invalid private key")
	}
	return privKey
}

func (i *Client) AddPriceContract(name string, address string) {
	pricesContract, err := vabi.NewPrice(common.HexToAddress(address), i.conn)
	if err != nil {
		panic(err)
	}
	i.prices[name] = pricesContract
}

func (i *Client) AddStableCredit(name string, address string) {
	stableCreditContract, err := vabi.NewStableCredit(common.HexToAddress(address), i.conn)
	if err != nil {
		panic(err)
	}
	i.stableCredit[name] = stableCreditContract
}

func (i *Client) AddCollateral(name string, address string) {
	collateralContract, err := vabi.NewToken(common.HexToAddress(address), i.conn)
	if err != nil {
		panic(err)
	}
	i.collateral[name] = collateralContract
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

	err = ConfirmTx(context.Background(), i, result, i.publicKey)
	if err != nil {
		fmt.Println(err)
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

func (i *Client) GetCollateralBalanceOf(collateralName, holderAddress string) string {
	result, err := i.collateral[collateralName].BalanceOf(nil, common.HexToAddress(holderAddress))
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(result)
}

func (i *Client) ApproveCollateral(collateralName, spender, allowAmount string) string {
	amount, _ := utils.StringToAmount(allowAmount)
	spenderAddress := common.HexToAddress(spender)
	opt := bind.NewKeyedTransactor(i.privateKey)
	opt.GasLimit = constants.GasLimit
	result, err := i.collateral[collateralName].Approve(opt, spenderAddress, amount)
	if err != nil {
		panic(err)
	}

	err = ConfirmTx(context.Background(), i, result, i.publicKey)
	if err != nil {
		fmt.Println(err)
	}
	return result.Hash().String()

}

func (i *Client) GetCollateralTotalSupply(collateralAddress string) string {
	result, err := i.collateral[collateralAddress].TotalSupply(nil)
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(result)
}

func (i *Client) GetStableCreditTotalSupply(stableCreditName string) string {
	result, err := i.stableCredit[stableCreditName].TotalSupply(nil)
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(result)
}

func (i *Client) GetPrice(priceContractName string) string {
	price, err := i.prices[priceContractName].Get(nil)
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(price)
}

func (i *Client) GetPriceWithError(priceContractName string) (string, bool, bool) {
	price, isActive, isErr, err := i.prices[priceContractName].GetWithError(nil)
	if err != nil {
		panic(err)
	}
	return utils.AmountToString(price), isActive, isErr
}

type Addresses struct {
	DrsAddress            string `json:"drsAddress"`
	HeartAddress          string `json:"heartAddress"`
	ReserveManagerAddress string `json:"reserveManagerAddress"`
	CollateralAddress     string `json:"collateralAddress"`
	PriceAddress          string `json:"priceAddress"`
}

func GetContractAddresses() *Addresses {

	byteValue, err := ioutil.ReadFile("../contract-addresses.json")
	if err != nil {
		panic(err)
	}
	addresses := new(Addresses)
	err = json.Unmarshal(byteValue, addresses)
	if err != nil {
		panic(err)
	}

	return addresses
}

func main() {
	client := NewClient(
		"http://127.0.0.1:7545",
		"0xf93BF6d3bE161793F1688dFf38E51341552f5aA9",
		"0xaC06374fc95955fb495cE3De37f18eCF241e62F7",
		"91c351a1080a4eb4e63ff2e376f3360ddc469f032fdd6d2b136357a6849758dc",
	)
	// price
	client.AddPriceContract("<priceName>", "<priceAddress>")
	fmt.Println(client.GetPrice("<priceName>"))
	price, isActive, isErr := client.GetPriceWithError("<priceName>")
	fmt.Println(price)
	fmt.Println(isActive)
	fmt.Println(isErr)

	// collateral
	fmt.Println(client.GetCollateralRatio("VELO"))
	client.SetCollateralRatio("VELO", "1.0", "<privateKey>")

	client.AddCollateral("<collateralName>", "<collateralAddress>")
	fmt.Println(client.GetCollateralBalanceOf("<collateralName>", "<address>"))
	fmt.Println(client.GetCollateralTotalSupply("<collateralName>"))

	// stable credit
	client.AddStableCredit("<stableCreditName>", "<stableCreditAddress>")
	fmt.Println(client.GetStableCreditTotalSupply("<stableCreditName>"))
}
