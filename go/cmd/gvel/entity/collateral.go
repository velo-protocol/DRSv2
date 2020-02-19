package entity

// CollateralHealthCheck
type CollateralHealthCheckInput struct {
	Passphrase string
}

type CollateralHealthCheckOutput struct {
	CollateralAssetAddress string
	CollateralAssetCode    string
	RequiredAmount         string
	PresentAmount          string
}
