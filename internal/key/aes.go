package key

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
)

const aesEncryptedKeyPath = "encrypted_aes_key.bin"

// GenerateAESKey generates a new AES key.
func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// EncryptAESKey encrypts an AES key using an RSA public key.
func EncryptAESKey(aesKey []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	encryptedKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, aesKey)
	if err != nil {
		return nil, err
	}
	return encryptedKey, nil
}

// DecryptAESKey decrypts an AES key using the RSA private key.
func DecryptAESKey(encryptedKey []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	aesKey, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedKey)
	if err != nil {
		return nil, err
	}
	return aesKey, nil
}

// SaveEncryptedAESKey saves the encrypted AES key to a file.
func SaveEncryptedAESKey(encryptedKey []byte) error {
	return os.WriteFile(aesEncryptedKeyPath, encryptedKey, 0644)
}

// LoadEncryptedAESKey loads the encrypted AES key from a binary file.
func LoadEncryptedAESKey() ([]byte, error) {
	return os.ReadFile(aesEncryptedKeyPath)
}
