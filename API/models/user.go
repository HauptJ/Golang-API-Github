package models

import (
  //"strings"
)

type User struct {
  Login string `json:"login"`
  ID string `json:"id"`
  Level uint
  Followers []User
  Repos []Repo
}


// type GithubUser struct {
//   login string
//   id int
//   node_id string
//   avatar_url string
//   gravatar_id string
//   url string
//   html_url string
//   followers_url string
//   following_url string
//   gists_url string
//   starred_url string
//   subscriptions_url string
//   organizations_url string
//   repos_url string
//   events_url string
//   Received_events_url string `json:"received_events_url"`
//   Type string `json:"type"`
//   Site_admin bool `json:"site_admin"`
// }
