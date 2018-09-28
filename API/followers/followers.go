package followers

import (
  "fmt"
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

func getUser(GHUser string, numFollowers int) (User, error) {

  context := context.Background()
  tokenService := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "TOKEN"},
  )
  tokenClient := oauth2.NewClient(context, tokenService)

  client := github.NewClient(tokenClient)

  var followerNameList []string

  followerObjs, _, err := client.Users.ListFollowers(context, GHUser, nil) //TODO: replace nil with correct option to request only the first numFollowers followers
  if err != nil {
    log.Fatal(err)
  }
  if err != nil {
    fmt.Printf("Problem in getting repository information %v\n", err)
    os.Exit(1)
  }

  if len(followerObjs) < numFollowers {
    numFollowers = len(followerObjs)
  }

  for _, followerObj := range followerObjs {
      followerNameList = append(followerNameList, *followerObj.Login)
  }
  followerNameList = followerNameList[:numFollowers] //NOTE: TEMP hack. SEE: TODO

  GHUserObj := User{Name: GHUser, Followers: followerNameList}

  return GHUserObj, err
}


func GenUserObjList(rootUser string, numLvls, numFollowers int) ([]User, error) {

  var userObjList []User
  var newUserObj User

  userObj, err := getUser(rootUser, numFollowers)
  if err != nil {
    log.Fatal(err)
  }
  userObjList = append(userObjList, userObj)
  for i := 1; i <= numLvls; i++ {
    for _, follower := range userObj.Followers{
      newUserObj, err = getUser(follower, numFollowers)
      userObjList = append(userObjList, newUserObj)
      if err != nil {
        log.Fatal(err)
      }
    }
    userObj = newUserObj
  }
  return userObjList, err
}

// DEBUG
// func main() {
//   userList, err := GenUserObjList("HauptJ", 3, 5)
//   if err != nil {
//     log.Fatal(err)
//   }
//
//   for _, userObj := range userList {
//     fmt.Printf("%+v\n", userObj)
//   }
// }
