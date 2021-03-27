package internal

import (
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

	return uA.Format("2006-01-02T15:04:05")

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
