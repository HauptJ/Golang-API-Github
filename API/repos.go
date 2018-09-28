package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"context"
	"fmt"
	"os"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "36f30921dd1a94c1fb2022df19972d9b82682e70"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	repos, _, err := client.Repositories.List(context, "HauptJ", nil)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

  //var packages []Package

  for _, repo := range repos {
    // packages = append(packages, Package{
    // 	FullName: *repo.FullName,
    // 	Description: *repo.Description,
    // 	ForksCount: *repo.ForksCount,
    // 	StarsCount: *repo.StargazersCount,
    // })
    fmt.Println(*repo.FullName)
    }
    //fmt.Println(pack)


  // for _, pack := range packages {
  //   fmt.Printf("%+v\n", pack)
  // }
}
