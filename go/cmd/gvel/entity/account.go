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

type SetDefaultAccountInput struct {
	Account string
}

type SetDefaultAccountOutput struct {
	Account string
}

type ImportAccountInput struct {
	PrivateKey          string
	Passphrase          string
	SetAsDefaultAccount bool
}

type ImportAccountOutput struct {
	ImportedAccountAddress string
}

type ExportAccountInput struct {
	PublicAddress string
	Passphrase    string
}

type ExportAccountOutput struct {
	PublicAddress string
	PrivateKey    string
}
