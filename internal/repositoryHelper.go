package internal

import (
	"strconv"

	"github.com/google/go-github/v33/github"
)

func getRepoName(repo *github.Repository) string {

	if v := repo.Name; v != nil {
		return *v
	}

	return ""
}

func getUpdatedAt(repo *github.Repository) string {

	var uA github.Timestamp

	if v := repo.UpdatedAt; v != nil {
		uA = *v
	} else {
		uA = github.Timestamp{}
	}

	year, month, day := uA.Date()
	date := strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day)

	hour, min, sec := uA.Clock()
	time := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	updatedAt := date + "T" + time

	return updatedAt

}

func getOwner(repo *github.Repository) github.User {
	if v := repo.Owner; v != nil {
		return *v
	}

	return github.User{}
}

func getStars(repo *github.Repository) int {

	if v := repo.StargazersCount; v != nil {
		return *v
	}

	return -1
}

func getOwnerName(owner *github.User) string {

	if v := owner.Company; v != nil {
		return *v
	} else if v := owner.Name; v != nil {
		return *v
	} else if v := owner.Login; v != nil {
		return *v
	}

	return ""
}
