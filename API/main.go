package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v18/github"
  "golang.org/x/oauth2"
)

type Test struct {
  Name string
  //ID int64
}

// Fetch all the public organizations' membership of a user.
//
func FetchOrganizations(username string) ([]*github.Organization, error) {
  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "153ada3fa801f79fa88e0932133809711c8ee728"},
  )
  tc := oauth2.NewClient(ctx, ts)
  // get go-github client
  client := github.NewClient(tc)
	orgs, _, err := client.Organizations.List(ctx, username, nil)


  for _, org := range orgs {
    test := &Test{
      Name: org.GetLogin(),
      //ID: *org.ID,
    }
    fmt.Printf("%+v\n", test.Name)
  }

  // for i, organization := range orgs {
	// 	fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
  //   fmt.Printf("%v. %v\n", i+1, organization)
  //   fmt.Println(organization)
	// }


	return orgs, err
}

func main() {
	var username string

	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

  fmt.Println(username)

	organizations, err := FetchOrganizations(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
    //fmt.Printf("%v. %v\n", i+1, organization.GetName())
     //fmt.Printf("%v. %v\n", i+1, organization.ID)
    // fmt.Println(organization)
	}
}
