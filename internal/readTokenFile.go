package internal

import (
	"errors"
	"os"
)

func ReadTokenFile(path string) (string, error) {

	if path == "" {
		return "", nil
	}

	accessToken, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("something went wrong when reading the file")
	}

	return string(accessToken), nil

}
