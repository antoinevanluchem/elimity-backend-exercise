package internal

import (
	"context"
	"fmt"
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
		for _, repository := range result.Repositories {
			repoName := *getRepoName(repository)
			updatedAt := *getUpdatedAt(repository)
			stars := *getStars(repository)
			owner := *getOwner(repository)
			ownerName := *getOwnerName(&owner)

			//TODO: pretty print

			fmt.Println("-----------")
			fmt.Println(repoName)
			fmt.Println(ownerName)
			fmt.Println(updatedAt)
			fmt.Println(stars)
			fmt.Println("-----------")
		}
	}
}
