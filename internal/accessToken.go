package internal

import (
	"errors"
	"os"
)

// Function to read the access token at a specified path
// If no path is provided (empty string), an empty access token is returned
func ReadTokenFile(path string) (string, error) {

	// if path == "" {
	// 	return AccessToken{}, nil
	// }

	accessToken, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("something went wrong when reading the file")
	}

	return string(accessToken), nil

}
