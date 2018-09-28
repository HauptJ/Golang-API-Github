package main

import (
  "fmt"
  "log"
  //"strings"
  // "encoding/json"
  "net/http"
   "strconv"
  "io/ioutil"
  //"bytes"
  . "./models"
)

/*
  Constant Declarations
*/
const GITHUB_USER_API_ENDPOINT = "https://api.github.com/users"
const GITHUB_API_PAGES = "1"

func GitHubFollowerAPI(githubID string, numLevels uint8, numFollowers uint16, followers *[]User) /*followers []User*/ {

  // generate GET request URL
  followerQuerry := GITHUB_USER_API_ENDPOINT + "/" + githubID + "/followers?page=1&per_page=" + strconv.Itoa(int(numFollowers))

  fmt.Println(followerQuerry)

  res, err := http.Get(followerQuerry)

  if err != nil {
    panic(err)
  }

  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    panic(err)
  }

  if res.StatusCode != 200 {
    log.Fatal("Unexpected status code", res.StatusCode)
  }

  fmt.Printf("Body: %s\n", body)

}


func main() {

  var followers []User

  githubID := "HauptJ"
  var numLevels  uint8 = 1
  var numFollowers uint16 = 10
  GitHubFollowerAPI(githubID, numLevels, numFollowers, &followers)
  fmt.Println(&followers)
}
