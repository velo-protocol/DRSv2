package entity

type MintCreditByCollateralInput struct {
	AssetCode        string
	CollateralAmount string
	Passphrase       string
}

type MintCreditByCollateralOutput struct {
	AssetCode  string
	MintAmount string
	TxHash     string
}
