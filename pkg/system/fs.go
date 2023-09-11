package system

import (
	"os"
	"path/filepath"
)

func dataDirPath(path string) string {
	basePath := os.Getenv("DATA_DIR")
	if basePath == "" {
		basePath = "/tmp/lilypad/data"
	}
	return filepath.Join(basePath, path)
}

func DataDir(path string) (string, error) {
	dirPath := dataDirPath(path)
	err := os.MkdirAll(dirPath, 0755) // 0755 is the file permission
	if err != nil {
		return "", err
	}
	return dirPath, nil
}
