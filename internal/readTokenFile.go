package internal

import (
	"errors"
	"fmt"
	"os"
)

func ReadTokenFile(path string) (string, error) {

	fmt.Printf("Path is %s \n", path)

	if path == "" {
		return "", nil
	}

	accessToken, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("something went wrong when reading the file")
	}

	return string(accessToken), nil

}
