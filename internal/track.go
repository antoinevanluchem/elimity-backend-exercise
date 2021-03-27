package internal

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/google/go-github/v33/github"
)

// TrackOptions specifies the optional parameters to the Track function.
// Interval = repository update interval
// MinStars = minimal star count of repositories
// AccesToken = acces token for authenticated requests
// Client = struct of type Client that contains the github client and context
type TrackOptions struct {
	Interval    time.Duration
	MinStars    int
	AccessToken string
	Client      Client
}

// Track tracks public GitHub repositories, optional parameters must be given via a TrackOptions struct.
func Track(trackOptions TrackOptions) error {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	headers := "Owner\tName\tUpdated at (UTC)\tStar count"
	fmt.Fprintln(w, headers)

	// headers := []string{"Owner", "Name", "Updated at (UTC)", "Star count"}
	// pPrinter := NewPrettyPrinter(headers, " ", " |")

	client := trackOptions.Client.GithubClient
	con := trackOptions.Client.Context

	// i := 0

	for ; ; <-time.Tick(trackOptions.Interval) {

		listOptions := github.ListOptions{PerPage: 3}
		searchOptions := &github.SearchOptions{ListOptions: listOptions, Sort: "updated"}

		query := fmt.Sprintf("is:public stars:>=%d", trackOptions.MinStars)

		result, _, err := client.Search.Repositories(con, query, searchOptions)
		if err != nil {
			return err
		}

		for _, repository := range result.Repositories {
			repoName := getRepoName(repository)
			updatedAt := getUpdatedAt(repository)
			stars := getStars(repository)
			owner := getOwner(repository)
			ownerName := getOwnerName(&owner)

			row := fmt.Sprintf("%s\t"+"%s\t"+"%s\t"+"%d", ownerName, repoName, updatedAt, stars)
			fmt.Fprintln(w, row)
			// row := map[string]string{"Owner": ownerName, "Name": repoName, "Updated at (UTC)": updatedAt, "Star count": strconv.Itoa(stars)}

			// pPrinter, err = pPrinter.AddRow(row)
			// if err != nil {
			// 	return err
			// }

		}

		w.Flush()

		// if i == 0 {
		// 	pPrinter.Print()
		// } else {
		// 	err := pPrinter.PrintLastNRows(len(result.Repositories))
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		// i++

	}
}
