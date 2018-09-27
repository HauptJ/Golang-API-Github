package main

import (
  "fmt"
  //"strings"
  "encoding/json"
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

  var locFollowers []User

  client := &http.Client{}

  // map to store returned JSON
  var res map[string]interface{}

  // generate GET request URL
  followerQuerryBegin := GITHUB_USER_API_ENDPOINT + "/" + githubID + "/followers"

  req, err := http.NewRequest("GET", followerQuerryBegin, nil)
  if err != nil {
    panic(err)
  }

  req.Header.Add("Accept", "application/json")

  q := req.URL.Query()
  q.Add("page", GITHUB_API_PAGES)
  q.Add("per_page", strconv.Itoa(int(numFollowers)))
  req.URL.RawQuery = q.Encode()
  fmt.Println(req.URL.String())

  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  //defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(&res)
  json.NewDecoder(resp.Body).Decode(&locFollowers)
  resp_body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(resp.Status)
  fmt.Println(resp_body)
  fmt.Println(resp.Body)
  //json.Unmarshal([]byte(resp.Body), &locFollowers)
  fmt.Println(res)
  fmt.Println(locFollowers)
  fmt.Printf("Followers : %+v", locFollowers)


}


func main() {

  var followers []User

  githubID := "HauptJ"
  var numLevels  uint8 = 1
  var numFollowers uint16 = 10
  GitHubFollowerAPI(githubID, numLevels, numFollowers, &followers)
  fmt.Println(&followers)
}
