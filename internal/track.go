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
type TrackOptions struct {
	Interval    time.Duration
	MinStars    int
	AccessToken string
}

// Track tracks public GitHub repositories, continuously updating according to the given interval.
//
// The given interval must be greater than zero.
func Track(trackOptions *TrackOptions) error {

	headers := []string{"Owner", "Name", "Updated at (UTC)", "Star count"}
	pPrinter := NewPrettyPrinter(headers, " ", " |")

	for i := 0; ; <-time.Tick(trackOptions.Interval) {

		con := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: trackOptions.AccessToken},
		)
		tc := oauth2.NewClient(con, ts)
		client := github.NewClient(tc)

		// client := github.NewClient(nil)
		// con := context.Background()
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
