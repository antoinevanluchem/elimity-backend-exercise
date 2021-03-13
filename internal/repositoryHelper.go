package internal

import (
	"github.com/google/go-github/v33/github"
)

func getRepoName(repo *github.Repository) (repoName *string) {

	if v := repo.Name; v != nil {
		repoName = v
	} else {
		rN := ""
		repoName = &rN
	}

	return
}

func getUpdatedAt(repo *github.Repository) (updatedAt *github.Timestamp) {

	if v := repo.UpdatedAt; v != nil {
		updatedAt = v
	} else {
		uA := github.Timestamp{}
		updatedAt = &uA
	}

	return

}

func getOwner(repo *github.Repository) (owner *github.User) {

	if v := repo.Owner; v != nil {
		owner = v
	} else {
		o := github.User{}
		owner = &o
	}

	return
}

func getStars(repo *github.Repository) (stars *int) {

	if v := repo.StargazersCount; v != nil {
		stars = v
	} else {
		s := 0
		stars = &s
	}

	return
}

func getOwnerName(owner *github.User) (ownerName *string) {

	if v := owner.Name; v != nil {
		ownerName = v
	} else if v := owner.Company; v != nil {
		ownerName = v
	} else if v := owner.Login; v != nil {
		ownerName = v
	} else {
		oN := ""
		ownerName = &oN
	}

	return
}
