package internal

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

// TrackOptions specifies the optional parameters to the Track function.
// Interval = repository update interval
// MinStars = minimal star count of repositories
// AccesToken = acces token for authenticated requests
type TrackOptions struct {
	Interval    time.Duration
	MinStars    int
	AccessToken AccessToken
}

// Track tracks public GitHub repositories, optional parameters must be given via a TrackOptions struct.
func Track(trackOptions *TrackOptions) error {

	headers := []string{"Owner", "Name", "Updated at (UTC)", "Star count"}
	pPrinter := NewPrettyPrinter(headers, " ", " |")

	i := 0

	for ; ; <-time.Tick(trackOptions.Interval) {

		client, con := getNewClient(trackOptions.AccessToken)

		listOptions := github.ListOptions{PerPage: 3}
		searchOptions := &github.SearchOptions{ListOptions: listOptions, Sort: "updated"}

		query := fmt.Sprintf("is:public stars:>=%d", trackOptions.MinStars)

		result, _, err := client.Search.Repositories(con, query, searchOptions)
		if err != nil {
			return err
		}

		for _, repository := range result.Repositories {
			repoName := *getRepoName(repository)
			updatedAt := *getUpdatedAt(repository)
			stars := *getStars(repository)
			owner := *getOwner(repository)
			ownerName := *getOwnerName(&owner)

			row := map[string]string{"Owner": ownerName, "Name": repoName, "Updated at (UTC)": updatedAt.String(), "Star count": strconv.Itoa(stars)}
			pPrinter.AddRow(row)
		}

		if i == 0 {
			pPrinter.Print()
		} else {
			pPrinter.PrintLastNRows(len(result.Repositories))
		}

		i++

	}
}

// Helper function to get a new github client and context
// If a accessToken is given the github client sends authenticated requests, if no accessToken is given the github client sends anonymous requests
func getNewClient(accessToken AccessToken) (*github.Client, context.Context) {

	con := context.Background()

	if accessToken.Token == "" {
		return github.NewClient(nil), con

	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken.Token},
	)
	tc := oauth2.NewClient(con, ts)

	return github.NewClient(tc), con

}
