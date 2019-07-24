package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	drsAbi "gitlab.com/velo-lab/core/smart_contract/abi/digital_reserve_system"
	priceFeederTHBAbi "gitlab.com/velo-lab/core/smart_contract/abi/price_feeder_velo_thb"
	priceFeederUSDAbi "gitlab.com/velo-lab/core/smart_contract/abi/price_feeder_velo_usd"
	veloTokenAbi "gitlab.com/velo-lab/core/smart_contract/abi/velo_token"
	warpAbi "gitlab.com/velo-lab/core/smart_contract/abi/warp"
	whitelistAssetsAbi "gitlab.com/velo-lab/core/smart_contract/abi/whitelist_assets"
	"log"
	"math/big"
)

func main() {
	ethEndpoint := "http://ec2-13-250-122-92.ap-southeast-1.compute.amazonaws.com:8501"
	deployerPrvKey := "CDAAC34D74F4CC371A1A23A754E749D3CF60D0B916F7CC3350BBAA99BC663B95"

	ethNode, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		log.Fatalf("unable to connect to ethereum node: %v \n", err)
	}

	prvKey, err := crypto.HexToECDSA(deployerPrvKey)
	txOpts := bind.NewKeyedTransactor(prvKey)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v \n", err)
	}
	txOpts.GasLimit = 4000000

	whitelistAssetsAddress := deployWhitelistAssetsContract(txOpts, ethNode)
	veloTokenContractAddress := deployVeloTokenContract(txOpts, ethNode)
	warpContractAddress := deployWarpContract(txOpts, ethNode, whitelistAssetsAddress)
	addWarpAsAdmin(txOpts, ethNode, veloTokenContractAddress.Hex(), warpContractAddress.Hex())
	addAsset(txOpts, ethNode, whitelistAssetsAddress.Hex(), "VELO", veloTokenContractAddress.Hex())

	//getAssetContractAddress(nil, ethNode, "0x94b6e47a8f459304B041E5D8d5eaB163679A1689", "VELO")
	//unlockAsset(txOpts, ethNode, "0xBfc15f814295C50B22f5F896cD7B7D04822ff462", "VELO", "0x212C2a2891227f39B48D655C5ecA0b1377daFF90", big.NewInt(1110000000))

	//drsAddress := deployDigitalReserveSystemContract(txOpts, ethNode, "0x61a5d56c9214377a777ba410f2a796b8d953c792")
	//_ = deployPriceFeederTHBContract(txOpts, ethNode, drsAddress)
	//_ = deployPriceFeederUSDContract(txOpts, ethNode, common.HexToAddress("0x2329abE57D45c9798Dd267f83ec70aD8284ea8bE"))
	//setPriceFeeder(txOpts, ethNode, "0x2329abE57D45c9798Dd267f83ec70aD8284ea8bE", "THB", "0x54454F06ac60d35b815AdBD81E222Cf8aa4AD4C5")
	//setPriceFeeder(txOpts, ethNode, "0x2329abE57D45c9798Dd267f83ec70aD8284ea8bE", "USD", "0x75851853cDb6C2f6E9d79c4425DBd7A4C2E3B5a4")
	//
	//updateTHBPrice(txOpts, ethNode, "0x54454F06ac60d35b815AdBD81E222Cf8aa4AD4C5")
	//updateUSDPrice(txOpts, ethNode, "0x75851853cDb6C2f6E9d79c4425DBd7A4C2E3B5a4")
	//getPrice(nil, ethNode, "0x6e9c25bBDE26111c7daD0AF8db023B6B39c3B76F", "THB")

	//approve(txOpts, ethNode, "0x61a5d56c9214377a777ba410f2a796b8d953c792", "0xdA1Dabd0465A59cf2899B02F00faf43361f06Fb7", big.NewInt(8860000000))
	//getAllowance(nil, ethNode, "0x61a5d56c9214377a777ba410f2a796b8d953c792", "0x212C2a2891227f39B48D655C5ecA0b1377daFF90", "0xdA1Dabd0465A59cf2899B02F00faf43361f06Fb7")
	//
	//updateTrustedPartner(txOpts, ethNode, "0xdA1Dabd0465A59cf2899B02F00faf43361f06Fb7","0x212C2a2891227f39B48D655C5ecA0b1377daFF90")
	//mint(
	//	txOpts,
	//	ethNode,
	//	"0x1e31E59da8BC065DB817db8A9cB066A76FafEE9D",
	//	big.NewInt(3500000),
	//	"THB",
	//	"fVELO",
	//	big.NewInt(10000000),
	//	"GDB5JLOVDB53YQ3UTAN6PMFFIXL6NAS3ET67SRQPVQYYQQXRJ7JU2AHA",
	//	"685db6a78d5af37aae9cb7531ffc034444a562c774e54a73201cc17d7388fcbd",
	//)
	//getCollateralOf(nil, ethNode, "0xd30782C76f245aC20b3eC0a86BB5E6348902be42", "fVELO", "0x212C2a2891227f39B48D655C5ecA0b1377daFF90")

	//checkRequiredCollateralWithFee(
	//	nil,
	//	ethNode,
	//	"0xd30782C76f245aC20b3eC0a86BB5E6348902be42",
	//	big.NewInt(3500000),
	//	"THB",
	//	big.NewInt(10000000),
	//)
	//checkAllowance(nil, ethNode, "0xa9359751e9446E3b3e1f3c673A4b1d62dD5Cd51C", "0x212C2a2891227f39B48D655C5ecA0b1377daFF90")
}

func deployVeloTokenContract(opts *bind.TransactOpts, client *ethclient.Client) common.Address {
	veloTokenAddress, _, _, _ := veloTokenAbi.DeployVeloToken(opts, client,
		"GD5ZEDPHHSBYW4SS6YDSLZH7TTUXYKKEUTFL2GS5MYJRGQWIKB7XGKIL", "Velo", "VELO", 7)
	fmt.Printf("VeloToken's contract Address -> %s \n", veloTokenAddress.String())

	return veloTokenAddress
}

func deployWhitelistAssetsContract(opts *bind.TransactOpts, client *ethclient.Client) common.Address {
	whitelistAssetsAddress, _, _, _ := whitelistAssetsAbi.DeployWhitelistAssets(opts, client)
	fmt.Printf("Whitelist's contract Address -> %s \n", whitelistAssetsAddress.String())

	return whitelistAssetsAddress
}

func deployWarpContract(opts *bind.TransactOpts, client *ethclient.Client, whitelistAssetAddr common.Address) common.Address {
	warpContractAddress, _, _, _ := warpAbi.DeployWarp(opts, client, whitelistAssetAddr, 0)
	fmt.Printf("Warp's contract Address -> %s \n", warpContractAddress.String())

	return warpContractAddress
}

func deployPriceFeederTHBContract(opts *bind.TransactOpts, client *ethclient.Client, drsAddr common.Address) common.Address {
	priceFeederAddress, _, _, err := priceFeederTHBAbi.DeployPriceFeeder(opts, client, drsAddr)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("PriceFeeder_THB's contract Address -> %s \n", priceFeederAddress.String())

	return priceFeederAddress
}

func deployPriceFeederUSDContract(opts *bind.TransactOpts, client *ethclient.Client, drsAddr common.Address) common.Address {
	priceFeederAddress, _, _, err := priceFeederUSDAbi.DeployPriceFeeder(opts, client, drsAddr)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("PriceFeeder_USD's contract Address -> %s \n", priceFeederAddress.String())

	return priceFeederAddress
}

func deployDigitalReserveSystemContract(opts *bind.TransactOpts, client *ethclient.Client, veloTokenAddr string) common.Address {
	digitalReserveSystemAddress, _, _, _ := drsAbi.DeployDigitalReserveSystem(opts, client, common.HexToAddress(veloTokenAddr))
	fmt.Printf("DRS's contract Address -> %s \n", digitalReserveSystemAddress.String())

	return digitalReserveSystemAddress
}

func addAsset(opts *bind.TransactOpts, client *ethclient.Client, whitelistAssetsContractAddress string, assetName string, assetContractAddr string) {
	contractInstance, err := whitelistAssetsAbi.NewWhitelistAssets(common.HexToAddress(whitelistAssetsContractAddress), client)
	if err != nil {
		log.Fatalf("can't init WhitelistAssets instance: %v \n", err)
	}

	assetNameB32 := [32]byte{}
	copy(assetNameB32[:], []byte(assetName))

	addAsset, err := contractInstance.AddAsset(opts, assetNameB32, common.HexToAddress(assetContractAddr))
	if err != nil {
		log.Fatalf("failed to add asset: %v\n", err)
	}
	fmt.Printf("%s\n", addAsset.Hash().String())
}

func getAssetContractAddress(opts *bind.CallOpts, client *ethclient.Client, whitelistAssetsContractAddress string, assetName string) {
	contractInstance, err := whitelistAssetsAbi.NewWhitelistAssets(common.HexToAddress(whitelistAssetsContractAddress), client)
	if err != nil {
		log.Fatalf("can't init WhitelistAssets instance: %v \n", err)
	}

	assetNameB32 := [32]byte{}
	copy(assetNameB32[:], []byte(assetName))

	assetContractAddress, err := contractInstance.CreditAddr(opts, assetNameB32)
	if err != nil {
		log.Fatalf("failed to add asset: %v\n", err)
	}
	fmt.Printf("%s\n", assetContractAddress.Hex())
}

func addWarpAsAdmin(opts *bind.TransactOpts, client *ethclient.Client, veloTokenContractAddress string, warpContractAddress string) {
	contractInstance, err := veloTokenAbi.NewVeloToken(common.HexToAddress(veloTokenContractAddress), client)
	if err != nil {
		log.Fatalf("can't init Velo Token instance: %v \n", err)
	}

	addWarpTx, err := contractInstance.AddWhitelistAdmin(opts, common.HexToAddress(warpContractAddress))
	if err != nil {
		log.Fatalf("failed to add warp as an admin: %v\n", err)
	}
	fmt.Printf("%s\n", addWarpTx.Hash().String())

	removeSelf, err := contractInstance.RenounceWhitelistAdmin(opts)
	if err != nil {
		log.Fatalf("failed to renouce self as an admin: %v\n", err)
	}
	fmt.Printf("%s\n", removeSelf.Hash().String())
}

func unlockAsset(opts *bind.TransactOpts, client *ethclient.Client, warpContractAddress string, assetName string, targetAddress string, amount *big.Int) {
	contractInstance, err := warpAbi.NewWarp(common.HexToAddress(warpContractAddress), client)
	if err != nil {
		log.Fatalf("can't init Velo Token instance: %v \n", err)
	}

	assetNameB32 := [32]byte{}
	copy(assetNameB32[:], []byte(assetName))

	unlockTx, err := contractInstance.Unlock(opts, common.HexToAddress(targetAddress), assetNameB32, amount)
	if err != nil {
		log.Fatalf("failed to add warp as an admin: %v\n", err)
	}
	fmt.Printf("%s\n", unlockTx.Hash().String())
}

func setPriceFeeder(opts *bind.TransactOpts, client *ethclient.Client, drsContractAddress string, currency string, pricefeederContractAddress string) {
	contractInstance, err := drsAbi.NewDigitalReserveSystem(common.HexToAddress(drsContractAddress), client)
	if err != nil {
		log.Fatalf("can't init DRS instance: %v \n", err)
	}

	currencyB32 := [32]byte{}
	copy(currencyB32[:], []byte(currency))

	setPriceFeederTx, err := contractInstance.SetPriceFeeder(opts, currencyB32, common.HexToAddress(pricefeederContractAddress))
	if err != nil {
		log.Fatalf("failed to set a price feeder: %v\n", err)
	}
	fmt.Printf("%s\n", setPriceFeederTx.Hash().String())
}

func updateTHBPrice(opts *bind.TransactOpts, client *ethclient.Client, priceFeederContractAddress string) {
	opts.Value = big.NewInt(100000000000000000)

	contractInstance, err := priceFeederTHBAbi.NewPriceFeeder(common.HexToAddress(priceFeederContractAddress), client)
	if err != nil {
		log.Fatalf("can't init price feeder: %v \n", err)
	}

	updatePriceTx, err := contractInstance.UpdatePrice(opts, big.NewInt(0))
	if err != nil {
		log.Fatalf("failed to update the price: %v\n", err)
	}
	fmt.Printf("%s\n", updatePriceTx.Hash().String())
}

func updateUSDPrice(opts *bind.TransactOpts, client *ethclient.Client, priceFeederContractAddress string) {
	opts.Value = big.NewInt(10000000000000000)

	contractInstance, err := priceFeederUSDAbi.NewPriceFeeder(common.HexToAddress(priceFeederContractAddress), client)
	if err != nil {
		log.Fatalf("can't init price feeder: %v \n", err)
	}

	updatePriceTx, err := contractInstance.UpdatePrice(opts, big.NewInt(0))
	if err != nil {
		log.Fatalf("failed to update the price: %v\n", err)
	}
	fmt.Printf("%s\n", updatePriceTx.Hash().String())
}

func getPrice(opts *bind.CallOpts, client *ethclient.Client, drsAddress string, currency string) {
	contractInstance, err := drsAbi.NewDigitalReserveSystem(common.HexToAddress(drsAddress), client)
	if err != nil {
		log.Fatalf("can't init drs contract: %v\n", err)
	}

	currencyB32 := [32]byte{}
	copy(currencyB32[:], []byte(currency))

	price, err := contractInstance.GetPrice(opts, currencyB32)
	if err != nil {
		log.Fatalf("failed to get price: %v\n", err)
	}
	fmt.Printf("VELO%v: %d\n", currency, price)
}

func approve(opts *bind.TransactOpts, client *ethclient.Client, veloTokenAddr string, spender string, amount *big.Int) {
	contractInstance, err := veloTokenAbi.NewVeloToken(common.HexToAddress(veloTokenAddr), client)
	if err != nil {
		log.Fatalf("can't init velo token contract: %v\n", err)
	}

	approveTx, err := contractInstance.Approve(opts, common.HexToAddress(spender), amount)
	if err != nil {
		log.Fatalf("failed to approve spender: %v\n", err)
	}
	fmt.Printf("approve tx: %v\n", approveTx.Hash().String())
}

func getAllowance(opts *bind.CallOpts, client *ethclient.Client, veloTokenAddr string, owner string, spender string) {
	contractInstance, err := veloTokenAbi.NewVeloToken(common.HexToAddress(veloTokenAddr), client)
	if err != nil {
		log.Fatalf("can't init velo token contract: %v\n", err)
	}

	allowances, err := contractInstance.Allowance(opts, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		log.Fatalf("failed to get allowances: %v\n", err)
	}
	fmt.Printf("allowances: %v\n", allowances)
}

func updateTrustedPartner(opts *bind.TransactOpts, client *ethclient.Client, drsAddress string, trustedPartnerAddr string) {
	contractInstance, err := drsAbi.NewDigitalReserveSystem(common.HexToAddress(drsAddress), client)
	if err != nil {
		log.Fatalf("can't init drs contract: %v\n", err)
	}

	updateTPsTx, err := contractInstance.SetTrustedPartner(opts, common.HexToAddress(trustedPartnerAddr))
	if err != nil {
		log.Fatalf("failed to update trusted partner: %v\n", err)
	}
	fmt.Printf("update TPs tx: %v\n", updateTPsTx.Hash().String())
}

func mint(
	opts *bind.TransactOpts,
	client *ethclient.Client,
	drsAddr string,
	peggedValue *big.Int,
	peggedCurrency string,
	assetName string,
	amount *big.Int,
	stellarAddr string,
	hashSecret string,
) {
	contractInstance, err := drsAbi.NewDigitalReserveSystem(common.HexToAddress(drsAddr), client)
	if err != nil {
		log.Fatalf("can't init drs contract: %v\n", err)
	}

	peggedCurrencyB32 := [32]byte{}
	copy(peggedCurrencyB32[:], []byte(peggedCurrency))

	assetNameB32 := [32]byte{}
	copy(assetNameB32[:], []byte(assetName))

	hashSecretB32 := [32]byte{}
	copy(hashSecretB32[:], []byte(hashSecret))

	mintTx, err := contractInstance.Mint(opts, peggedValue, peggedCurrencyB32, assetNameB32, amount, stellarAddr, hashSecretB32)
	if err != nil {
		log.Fatalf("failed to approve spender: %v\n", err)
	}
	fmt.Printf("mint tx: %v", mintTx.Hash().String())
}

func getCollateralOf(opts *bind.CallOpts, client *ethclient.Client, drsAddr string, assetName string, issuer string) {
	contractInstance, err := drsAbi.NewDigitalReserveSystem(common.HexToAddress(drsAddr), client)
	if err != nil {
		log.Fatalf("can't init drs contract: %v\n", err)
	}

	assetNameB32 := [32]byte{}
	copy(assetNameB32[:], []byte(assetName))

	collateral, err := contractInstance.CollateralOf(opts, assetNameB32, common.HexToAddress(issuer))
	if err != nil {
		log.Fatalf("failed to get collateral: %v\n", err)
	}
	fmt.Printf("collarter: (%v, %v)", assetName, collateral)
}

func checkRequiredCollateralWithFee(
	opts *bind.CallOpts,
	client *ethclient.Client,
	drsAddr string,
	peggedValue *big.Int,
	peggedCurrency string,
	amount *big.Int,
) {
	contractInstance, err := drsAbi.NewDigitalReserveSystem(common.HexToAddress(drsAddr), client)
	if err != nil {
		log.Fatalf("can't init drs contract: %v\n", err)
	}

	peggedCurrencyB32 := [32]byte{}
	copy(peggedCurrencyB32[:], []byte(peggedCurrency))

	collateral, fee, err := contractInstance.CheckRequiredCollateralWithFee(
		opts, peggedValue, peggedCurrencyB32, amount)
	if err != nil {
		log.Fatalf("failed to get collateral: %v\n", err)
	}
	fmt.Printf("(requiredCollateral, fee): (%v, %v)", collateral, fee)
}
