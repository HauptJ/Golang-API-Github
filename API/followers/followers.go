package followers

import (
  "log"
  "os"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
  "context"
)

type User struct {
  Name string
  Followers []string
}

func getUserObjGHAPI(GHUser string, numFollowers int) (User, error) {

  token := os.Getenv("TOKEN")
  context := context.Background()
  tokenService := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: token},
  )
  tokenClient := oauth2.NewClient(context, tokenService)

  client := github.NewClient(tokenClient)

  var followerNameList []string

  followerObjs, _, err := client.Users.ListFollowers(context, GHUser, nil) //TODO: Improve this by passing option to request only numFollowers followers
  if err != nil {
    log.Printf("Problem getting follower information %v\n", err)
    os.Exit(1)
  }

  if len(followerObjs) < numFollowers {
    numFollowers = len(followerObjs)
  }

  for _, followerObj := range followerObjs {
      followerNameList = append(followerNameList, *followerObj.Login)
  }
  followerNameList = followerNameList[:numFollowers] //NOTE SEE: TODO

  GHUserObj := User{Name: GHUser, Followers: followerNameList}

  return GHUserObj, err
}


func GenUserObjList(rootUser string, numLvls, numFollowers int) ([]User, error) {

  var userObjList []User
  var newUserObj User

  userObj, err := getUserObjGHAPI(rootUser, numFollowers)
  if err != nil {
    log.Fatal(err)
  }
  userObjList = append(userObjList, userObj)
  for i := 1; i <= numLvls; i++ {
    for _, follower := range userObj.Followers{
      newUserObj, err = getUserObjGHAPI(follower, numFollowers)
      userObjList = append(userObjList, newUserObj)
      if err != nil {
        log.Fatal(err)
      }
    }
    userObj = newUserObj
  }
  return userObjList, err
}
