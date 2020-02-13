package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"github.com/pkg/errors"

	"github.com/Evrynetlabs/evrynet-node/common/hexutil"
	"github.com/Evrynetlabs/evrynet-node/crypto"
	"io"
)

func GenerateKeypair() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyString := hexutil.Encode(crypto.FromECDSA(privateKey))[2:]

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, privateKeyString, nil
}

func Encrypt(rawMessage []byte, passphrase string) ([]byte, []byte, error) {
	hasher := sha256.New()
	hasher.Write([]byte(passphrase))

	c, err := aes.NewCipher(hasher.Sum(nil))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build cipher from passphrase")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build gcm")
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to rand nonce")
	}

	encryptedSeed := gcm.Seal(nonce, nonce, rawMessage, nil)

	return encryptedSeed, nonce, nil
}

func Decrypt(cipherText []byte, passphrase string) ([]byte, error) {
	hasher := sha256.New()
	hasher.Write([]byte(passphrase))

	c, err := aes.NewCipher(hasher.Sum(nil))
	if err != nil {
		return nil, errors.Wrap(err, "failed to build cipher from passphrase")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build GCM")
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("nonce size is larger than cipher text")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decipher and authenticate")
	}

	return plaintext, nil
}

func PrivateKeyToKeyPair(privateKey string) (*string, *string, error) {
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, nil, errors.New("invalid private key format")
	}

	publicKeyECDSA := privKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &address, &privateKey, nil
}
