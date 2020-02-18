package entity

type MintCreditFromCollateralInput struct {
	AssetCode        string
	CollateralAmount string
	Passphrase       string
}

type MintCreditFromCollateralOutput struct {
	AssetCode  string
	MintAmount string
	TxHash     string
}
