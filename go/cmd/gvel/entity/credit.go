package entity

type SetupCreditInput struct {
	Passphrase          string
	AssetCode           string
	PeggedValue         string
	PeggedCurrency      string
	CollateralAssetCode string
}

type SetupCreditOutput struct {
	AssetCode           string
	AssetAddress        string
	PeggedValue         string
	PeggedCurrency      string
	CollateralAssetCode string
}
