package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v33/github"
)

// Track tracks public GitHub repositories, continuously updating according to the given interval.
//
// The given interval must be greater than zero.
func Track(interval time.Duration, minStars int) error {
	fmt.Println("Dit is Track")
	for ; ; <-time.Tick(interval) {
		client := github.NewClient(nil)
		con := context.Background()
		listOptions := github.ListOptions{PerPage: 3}
		searchOptions := &github.SearchOptions{ListOptions: listOptions, Sort: "updated"}

		query := fmt.Sprintf("is:public stars:>=%d", minStars)
		fmt.Println("Dit is de query ", query)

		result, _, err := client.Search.Repositories(con, query, searchOptions)
		if err != nil {
			return err
		}
		for _, repository := range result.Repositories {
			fmt.Println(*repository.Name)
			fmt.Println(*repository.Owner.Name)
			fmt.Println(*repository.UpdatedAt)
			fmt.Println("Dit is for loop in Track")
		}
	}
}
