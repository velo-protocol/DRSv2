package entity

// SetupCredit
type SetupCreditInput struct {
	Passphrase          string
	AssetCode           string
	PeggedValue         string
	PeggedCurrency      string
	CollateralAssetCode string
}
type SetupCreditOutput struct {
	TxHash              string
	CreditOwnerAddress  string
	AssetCode           string
	AssetAddress        string
	PeggedValue         string
	PeggedCurrency      string
	CollateralAssetCode string
}

// MintCreditByCollateral
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

// MintCreditByCredit
type MintCreditByCreditInput struct {
	Passphrase   string
	AssetCode    string
	CreditAmount string
}
type MintCreditByCreditOutput struct {
	TxHash              string
	AssetCode           string
	MintAmount          string
	AssetAddress        string
	CollateralAssetCode string
	CollateralAmount    string
}
