package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
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
	"math/big"
)

// Client struct
type Client struct {
	conn         *ethclient.Client
	privateKey   *ecdsa.PrivateKey
	publicKey    common.Address
	drs          *vabi.DigitalReserveSystem
	heart        *vabi.Heart
	prices       map[string]*vabi.Price
	feeders      map[string]*vabi.Feeder
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
		feeders:      map[string]*vabi.Feeder{},
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

func (i*Client)SetUp(collateralAssetCode string, peggedCurrency string, assetCode string, peggedValue *big.Int) (*types.Transaction,error) {
	drs:=i.drs
	opt := bind.NewKeyedTransactor(i.privateKey)
	opt.GasLimit = constants.GasLimit
	tran,err:= drs.Setup(opt,utils.StringToByte32(collateralAssetCode),utils.StringToByte32(peggedCurrency),assetCode,peggedValue)
	if err!=nil {
		return nil, err
	}
	return tran,nil
}

func (i*Client)MintFromCollateralAmount(netCollateralAmount *big.Int, assetCode string) (*types.Transaction,error) {
	drs:=i.drs
	opt := bind.NewKeyedTransactor(i.privateKey)
	opt.GasLimit = constants.GasLimit
	tran,err:= drs.MintFromCollateralAmount(opt,netCollateralAmount,assetCode)
	if err!=nil {
		return nil, err
	}
	return tran,nil
}

func (i*Client)MintFromStableCreditAmount( mintAmount *big.Int, assetCode string) (*types.Transaction,error) {
	drs:=i.drs
	opt := bind.NewKeyedTransactor(i.privateKey)
	opt.GasLimit = constants.GasLimit
	tran,err:= drs.MintFromStableCreditAmount(opt,mintAmount,assetCode)
	if err!=nil {
		return nil, err
	}
	return tran,nil
}


func (i*Client)Redeem(stableCreditAmount *big.Int, assetCode string) (bool,error) {
	drs:=i.drs
	opt := bind.NewKeyedTransactor(i.privateKey)
	opt.GasLimit = constants.GasLimit
	_,err:= drs.Redeem(opt,stableCreditAmount,assetCode)
	if err!=nil {
		return false, err
	}
	return true,nil
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
		log.Println(err)
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
		log.Println(err)
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

func (i *Client) PostPrice(priceContractName string) (bool,error) {
	opt := bind.NewKeyedTransactor(i.privateKey)
	opt.GasLimit = constants.GasLimit
	_, err := i.prices[priceContractName].Post(opt)
	if err != nil {
		return false,err
	}
	return true,nil
}
func (i *Client) GetPriceWithError(priceContractName string) (string, bool, bool) {
	price, isActive, isErr, err := i.prices[priceContractName].GetWithError(nil)
	if err != nil {
		panic(err)
	}
	return price.String(), isActive, isErr
}

type Addresses struct {
	DrsAddress            string `json:"drsAddress"`
	HeartAddress          string `json:"heartAddress"`
	ReserveManagerAddress string `json:"reserveManagerAddress"`
	CollateralAddress     string `json:"collateralAddress"`
	PriceAddress          string `json:"priceAddress"`
}

func main() {
	client := NewClient(
		"node rpc",
		"xxxxxx",
		"xxxx",
		"xxxxx",
	)
	fmt.Println(client.GetCollateralRatio("VELO"))
	client.SetCollateralRatio("VELO", "1.0", "<privateKey>")
	client.AddCollateral("<collateralName>", "<collateralAddress>")
	fmt.Println(client.GetCollateralBalanceOf("<collateralName>", "<address>"))
	fmt.Println(client.GetCollateralTotalSupply("<collateralName>"))
	// stable credit
	client.AddStableCredit("<stableCreditName>", "<stableCreditAddress>")
	fmt.Println(client.GetStableCreditTotalSupply("<stableCreditName>"))
	// drs
	peggedCurrency,_:=new(big.Int).SetString("<value>", 10)
	client.SetUp("","","",peggedCurrency)
	netCollateralAmount,_:=new(big.Int).SetString("<value>", 10)
	client.MintFromCollateralAmount(netCollateralAmount,"")
	stableCreditAmount,_:=new(big.Int).SetString("<value>", 10)
	client.Redeem(stableCreditAmount,"")


}