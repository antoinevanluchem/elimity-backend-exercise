package internal

import (
	"errors"
	"os"
)

// A struct to save the accessToken
type AccessToken struct {
	Token string
}

// Function to read the access token at a specified path
// If no path is provided (empty string), an empty access token (empty string) is returned
func ReadTokenFile(path string) (AccessToken, error) {

	if path == "" {
		return AccessToken{}, nil
	}

	accessToken, err := os.ReadFile(path)
	if err != nil {
		return AccessToken{}, errors.New("something went wrong when reading the file")
	}

	return AccessToken{Token: string(accessToken)}, nil

}
