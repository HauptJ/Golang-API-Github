package repos

import (
  "log"
  "os"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
  "context"
)

type RepoName struct {
  Owner string
  Name string
}

type User struct {
  Name string
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

func getUserObjGHAPI(GHUser string, numRepos int) (User, error) {
  context := context.Background()
  tokenService := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "TOKEN"},
  )
  tokenClient := oauth2.NewClient(context, tokenService)

  client := github.NewClient(tokenClient)

  var repoList []RepoName

  repoObjs, _, err := client.Repositories.List(context, GHUser, nil)
  if err != nil {
    log.Printf("Problem getting user repo information %v\n", err)
    os.Exit(1)
  }

  if len(repoObjs) < numRepos {
    numRepos = len(repoObjs)
  }

  for _, repoObj := range repoObjs {
    repoNameObj := RepoName{Owner: *repoObj.Owner.Login, Name: *repoObj.Name} //TODO: Improve this by passing option to request only numRepos repos
    repoList = append(repoList, repoNameObj)
  }
  repoList = repoList[:numRepos] //NOTE SEE: TODO

  GHUserObj := User{Name: GHUser, Repos: repoList}

  return GHUserObj, err
}

func getRepoObjGHAPI(repo RepoName, numStargazers int) (Repo, error) {
  context := context.Background()
  tokenService := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "TOKEN"},
  )
  tokenClient := oauth2.NewClient(context, tokenService)

  client := github.NewClient(tokenClient)

  var stargazerList []string

  stargazerUserObjs, _, err := client.Activity.ListStargazers(context, repo.Owner, repo.Name, nil)
  if err != nil {
    log.Printf("Problem getting stargazer information %v\n", err)
  }

  if len(stargazerUserObjs) < numStargazers {
    numStargazers = len(stargazerUserObjs)
  }

  for _, stargazerUserObj := range stargazerUserObjs {
    stargazerList = append(stargazerList, *stargazerUserObj.User.Login) //TODO: Improve this by passing option to request only numStargazers stargazers
  }
  stargazerList  = stargazerList[:numStargazers] //NOTE SEE: TODO

  repoObj := Repo{RepoName: repo, Stargazers: stargazerList}

  return repoObj, err
}

func GenRepoStargazerLists(rootUser string, numLvls, numRepos, numStargazers int) (UserRepo, error) {

  var userObjList []User
  var newUserObj User
  var repoObjList []Repo
  var repoObj Repo
  var newRepoObj Repo

  userObj, err := getUserObjGHAPI(rootUser, numRepos)
  if err != nil {
    log.Fatal(err)
  }
  userObjList = append(userObjList, userObj)
  for _, repo := range userObj.Repos {
    repoObj, err = getRepoObjGHAPI(repo, numStargazers)
    if err != nil {
      log.Fatal(err)
    }
    repoObjList = append(repoObjList, repoObj)
  }
  for i := 1; i <= numLvls; i++ {
    for _, stargazer := range repoObj.Stargazers {
      newUserObj, err = getUserObjGHAPI(stargazer, numRepos)
      if err != nil {
        log.Fatal(err)
      }
      userObjList = append(userObjList, newUserObj)
      for _, repo := range newUserObj.Repos {
        newRepoObj, err = getRepoObjGHAPI(repo, numStargazers)
        if err != nil {
          log.Fatal(err)
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
