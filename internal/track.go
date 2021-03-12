package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v33/github"
)

type TrackOptions struct {
	interval  time.Duration
	minStars  int
	tokenFile string
}

// Track tracks public GitHub repositories, continuously updating according to the given interval.
//
// The given interval must be greater than zero.
func Track(trackOptions *TrackOptions) error {
	fmt.Println("Dit is Track")
	for ; ; <-time.Tick(trackOptions.interval) {
		client := github.NewClient(nil)
		con := context.Background()
		listOptions := github.ListOptions{PerPage: 3}
		searchOptions := &github.SearchOptions{ListOptions: listOptions, Sort: "updated"}

		query := fmt.Sprintf("is:public stars:>=%d", trackOptions.minStars)
		fmt.Println("Dit is de query ", query)

		result, _, err := client.Search.Repositories(con, query, searchOptions)
		if err != nil {
			return err
		}
		for _, repository := range result.Repositories {
			repoName := *repository.Name
			fmt.Println("repo Name ok")
			updatedAt := *repository.UpdatedAt
			fmt.Println("updated at ok")
			owner := *repository.Owner
			fmt.Printf("owner ok %v", owner)
			// ownerName := *owner.Name
			// fmt.Println("owner Name ok")

			fmt.Println("Let's go:")
			fmt.Println(repoName)
			// fmt.Println(ownerName)
			fmt.Println(updatedAt)
			fmt.Println("Dit is for loop in Track")
		}
	}
}
