package command

import (
	"fmt"
	"os"
	"path/filepath"
)

func readFilesInDir(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	var filePaths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}
