package entity

type CreateAccountInput struct {
	Passphrase          string
	SetAsDefaultAccount bool
}

type CreateAccountOutput struct {
	PublicAddress string
	PrivateKey    string
	IsDefault     bool
}
