package entity

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

type MintByCreditInput struct {
	Passphrase   string
	AssetCode    string
	CreditAmount string
}

type MintByCreditOutput struct {
	TxHash string
}
