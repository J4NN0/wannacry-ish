package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const rsaBits = 2048
const rsaPrivateKeyPath = "rsa_private_key.pem"

// GenerateRSAKeyPair generates an RSA private and public key pair.
func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// SavePrivateKey saves the RSA private key to a PEM file.
func SavePrivateKey(privateKey *rsa.PrivateKey) error {
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	return os.WriteFile(rsaPrivateKeyPath, privateKeyPEM, 0600)
}

// LoadPrivateKey loads an RSA private key from a PEM file.
func LoadPrivateKey() (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(rsaPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// ExtractPublicKey extracts the RSA public key from a private key.
func ExtractPublicKey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	return &privateKey.PublicKey
}
