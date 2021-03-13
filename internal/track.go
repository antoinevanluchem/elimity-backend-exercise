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
			repoName := *repository.Name
			updatedAt := *repository.UpdatedAt
			stars := *repository.StargazersCount
			owner := *repository.Owner
			ownerName := getOwnerName(&owner)

			// stars := hoe df krijg je stars?
			//TODO: pretty print

			fmt.Println("Let's go:")
			fmt.Println(repoName)
			fmt.Println(ownerName)
			fmt.Println(updatedAt)
			fmt.Println(stars)
		}
	}
}

func getOwnerName(owner *github.User) (ownerName string) {

	if v := owner.Name; v != nil {
		ownerName = *v
	} else if v := owner.Company; v != nil {
		ownerName = *v
	} else if v := owner.Login; v != nil {
		ownerName = *v
	} else {
		ownerName = ""
	}

	return
}
