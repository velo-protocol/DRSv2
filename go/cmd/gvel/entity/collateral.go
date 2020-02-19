package entity

// CollateralHealthCheck
type CollateralHealthCheckInput struct {
	Passphrase string
}

type CollateralHealthCheckOutput struct {
	CollateralAssetAddress   string
	CollateralAssetCode      string
	RequiredCollateralAmount string
	CollateralPoolAmount     string
}
