package internal

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

// Struct to define a Gothub client
type Client struct {
	GithubClient *github.Client
	Context      context.Context
}

func GetNewClient(path string) (Client, error) {

	if path == "" {
		return getUnauthenticatedClient(), nil
	}

	accessToken, err := readTokenFile(path)
	if err != nil {
		return Client{}, err
	}

	if accessToken == "" {
		return getUnauthenticatedClient(), nil
	}

	return getAuthenticatedClient(accessToken), nil

}

// Helper function to read the access token at a specified path
func readTokenFile(path string) (string, error) {

	accessToken, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("something went wrong when reading the file")
	}

	return string(accessToken), nil

}

func getAuthenticatedClient(accessToken string) Client {

	con := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(con, ts)

	return Client{GithubClient: github.NewClient(tc), Context: con}

}

func getUnauthenticatedClient() Client {
	return Client{GithubClient: github.NewClient(nil), Context: context.Background()}
}
