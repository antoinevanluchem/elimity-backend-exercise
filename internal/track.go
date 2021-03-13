package internal

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/go-github/v33/github"
)

// TrackOptions specifies the optional parameters to the Track function.
type TrackOptions struct {
	Interval  time.Duration
	MinStars  int
	TokenFile string
}

// Track tracks public GitHub repositories, continuously updating according to the given interval.
//
// The given interval must be greater than zero.
func Track(trackOptions *TrackOptions) error {
	fmt.Println("Dit is Track")
	for ; ; <-time.Tick(trackOptions.Interval) {
		client := github.NewClient(nil)
		con := context.Background()
		listOptions := github.ListOptions{PerPage: 3}
		searchOptions := &github.SearchOptions{ListOptions: listOptions, Sort: "updated"}

		query := fmt.Sprintf("is:public stars:>=%d", trackOptions.MinStars)
		fmt.Println("Dit is de query ", query)

		result, _, err := client.Search.Repositories(con, query, searchOptions)
		if err != nil {
			return err
		}

		headers := []string{"Owner", "Name", "Updated at (UTC)", "Star count"}
		prettyPrinter := NewPrettyPrinter(headers)
		prettyPrinter.PrintHeaders()

		for _, repository := range result.Repositories {
			print(*repository.Name)
			repoName := *getRepoName(repository)
			updatedAt := *getUpdatedAt(repository)
			stars := *getStars(repository)
			owner := *getOwner(repository)
			ownerName := *getOwnerName(&owner)

			row := map[string]string{"Owner": ownerName, "Name": repoName, "Updated at (UTC)": updatedAt.String(), "Star count": strconv.Itoa(stars)}
			prettyPrinter.AddRow(row)
			prettyPrinter.PrintRow(row)
		}

	}
}
