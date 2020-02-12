package entity

type ImportAccountInput struct {
	PrivateKey          string
	Passphrase          string
	SetAsDefaultAccount bool
}

type ImportAccountOutput struct {
	ImportedAccountAddress string
}
