package entity

type Account struct {
	PublicAddress       string `json:"publicAddress"`
	EncryptedPrivateKey []byte `json:"encryptedPrivateKey"`
	Nonce               []byte `json:"nonce"`
	IsDefault           bool   `json:"-"`
}
