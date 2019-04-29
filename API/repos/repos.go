/*
FILE: repos.go
DESC: /repos API endpoint functionality
LAST MODIFIED: 30-SEPT-18 by Joshua Haupt
*/

package repos

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type RepoName struct {
	Owner string
	Name  string
}

type User struct {
	Name  string
	Repos []RepoName
}

type Repo struct {
	RepoName
	Stargazers []string
}

type UserRepo struct {
	Users []User
	Repos []Repo
}

/*
DESC: Calls the GitHub API and creates and returns an object of type User which contains a list of repositories for the specified user.
IN: GHUser: the specified username to list repositories for, numRepos: the number of repositories to list for the specified user
OUT: An object of type User which contains the list of repositories for a specified user - SUCCESSl; err - FAILURE
*/
func getUserObjGHAPI(GHUser string, numRepos uint8) (User, error) {

	token := os.Getenv("TOKEN")
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	var repoList []RepoName

	opt := &github.RepositoryListOptions{
		Visibility:  "public",
		ListOptions: github.ListOptions{Page: 1, PerPage: int(numRepos)},
	}
	repoObjs, _, err := client.Repositories.List(context, GHUser, opt)
	if err != nil {
		log.Printf("Problem getting user repo information %v\n", err)
	}

	if len(repoObjs) < int(numRepos) {
		numRepos = uint8(len(repoObjs))
	}

	for _, repoObj := range repoObjs {
		repoNameObj := RepoName{Owner: *repoObj.Owner.Login, Name: *repoObj.Name}
		repoList = append(repoList, repoNameObj)
	}

	GHUserObj := User{Name: GHUser, Repos: repoList}

	return GHUserObj, err
}

/*
DESC: Calls the GitHub API and creates and returns an object of type Repo which contains a list of stargazers for a specified repository.
IN: GHRepo: of type RepoName - the specified GitHub repository, numStargazers: the number of stargazers to list for the specified repo
OUT: An object of type Repo which contains the list of stargazers for a specified repository - SUCCESS; err - FAILURE
*/
func getRepoObjGHAPI(GHRepo RepoName, numStargazers uint8) (Repo, error) {

	token := os.Getenv("TOKEN")
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	var stargazerList []string

	opt := &github.ListOptions{Page: 1, PerPage: int(numStargazers)}
	stargazerUserObjs, _, err := client.Activity.ListStargazers(context, GHRepo.Owner, GHRepo.Name, opt)
	if err != nil {
		log.Printf("Problem getting stargazer information %v\n", err)
	}

	if len(stargazerUserObjs) < int(numStargazers) {
		numStargazers = uint8(len(stargazerUserObjs))
	}

	for _, stargazerUserObj := range stargazerUserObjs {
		stargazerList = append(stargazerList, *stargazerUserObj.User.Login)
	}

	repoObj := Repo{RepoName: GHRepo, Stargazers: stargazerList}

	return repoObj, err
}

/*
DESC: Returns a list of 5 repositories for a specified GitHub user and a list of 5 stargazers for each repository. The lists go numLvls levels deep.
IN: rootUser: the initial username to list repositories for, numLvls: the number of levels deep to run, numRepos: the number of repositories to list for the specified user, numStargazers: the number of stargazers to list for the specified repo
OUT: An object of type UserRepo which contains the list of User objects and the list of Repo objects
*/
func GenRepoStargazerLists(rootUser string, numLvls, numRepos, numStargazers uint8) (UserRepo, error) {

	var userObjList []User
	var newUserObj User
	var repoObjList []Repo
	var repoObj Repo
	var newRepoObj Repo

	userObj, err := getUserObjGHAPI(rootUser, numRepos)
	if err != nil {
		log.Printf("Problem getting user repo information %v\n", err)
	}
	userObjList = append(userObjList, userObj)
	for _, repo := range userObj.Repos {
		repoObj, err = getRepoObjGHAPI(repo, numStargazers)
		if err != nil {
			log.Printf("Problem getting user repo information %v\n", err)
		}
		repoObjList = append(repoObjList, repoObj)
	}
	for i := 1; i <= int(numLvls); i++ {
		for _, stargazer := range repoObj.Stargazers {
			newUserObj, err = getUserObjGHAPI(stargazer, numRepos)
			if err != nil {
				log.Printf("Problem getting user information %v\n", err)
			}
			userObjList = append(userObjList, newUserObj)
			for _, repo := range newUserObj.Repos {
				newRepoObj, err = getRepoObjGHAPI(repo, numStargazers)
				if err != nil {
					log.Printf("Problem getting user repo information %v\n", err)
				}
				repoObjList = append(repoObjList, newRepoObj)
			}
			repoObj = newRepoObj
		}
		userObj = newUserObj
	}

	outputObj := UserRepo{Users: userObjList, Repos: repoObjList}
	return outputObj, err
}
