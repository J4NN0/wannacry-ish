package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"strings"
)

// DecryptFile decrypts the content of the given file and writes the result to the output file.
func DecryptFile(inputFile string, key []byte) error {
	if !strings.HasSuffix(inputFile, encryptedFileExt) {
		return nil
	}

	ciphertext, err := os.ReadFile(inputFile)
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

	// The nonce size should be the same as aesGCM.NonceSize().
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}

	// Extract the nonce and the actual ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	err = os.WriteFile(inputFile, plaintext, 0644)
	if err != nil {
		return err
	}

	err = removeEncryptedFileExt(inputFile)
	if err != nil {
		return err
	}

	return nil
}

// removeCryExtension renames a file by removing the .cry extension if it exists.
func removeEncryptedFileExt(filePath string) error {
	newFilePath := strings.TrimSuffix(filePath, encryptedFileExt)

	err := os.Rename(filePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}
