package entity

type Account struct {
	PublicAddress       string `json:"publicAddress"`
	EncryptedPrivateKey []byte `json:"encryptedPrivateKey"`
	Nonce               []byte `json:"nonce"`
	IsDefault           bool   `json:"-"`
}

type CreateAccountInput struct {
	Passphrase          string
	SetAsDefaultAccount bool
}

type CreateAccountOutput struct {
	PublicAddress string
	PrivateKey    string
	IsDefault     bool
}
