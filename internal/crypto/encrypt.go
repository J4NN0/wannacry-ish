package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const encryptedFileExt = ".cry"

// EncryptFile encrypts the content of the given file and writes the result to the output file.
func EncryptFile(inputFile string, key []byte) error {
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Create a nonce. The nonce size should be the same as aesGCM.NonceSize().
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Encrypt the data
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	err = os.WriteFile(inputFile, ciphertext, 0644)
	if err != nil {
		return err
	}

	err = addEncryptedFileExt(inputFile)
	if err != nil {
		return err
	}

	return nil
}

// changeFileExtension renames a file by replacing its extension with a new one.
func addEncryptedFileExt(filePath string) error {
	dir := filepath.Dir(filePath)
	baseName := filepath.Base(filePath)

	newFileName := fmt.Sprintf("%s%s", baseName, encryptedFileExt)
	newFilePath := filepath.Join(dir, newFileName)

	err := os.Rename(filePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}
