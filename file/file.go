package file

import (
	"errors"
	"os"
)

func WriteFile(fileName string, content []byte) error {
	if fileName == "" {
		return errors.New("file name should not be empty")
	}
	return os.WriteFile(fileName, content, 0644)
}

func ReadFile(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.New("file name should not be empty")
	}
	byteContent, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(byteContent), nil
}
