/*
FILE: followers.go
DESC: /followers API endpoint functionality
LAST MODIFIED: 30-SEPT-18 by Joshua Haupt
*/

package followers

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type User struct {
	Name      string
	Followers []string
}

/*
DESC: Returns a list of 5 followers for the specified GithUb user
IN: GHUser: the specified user, numFollowers: the number of followers to return
OUT: GHUserObj: Object of type User which contains a Name and a list of followers - SUCCESS; err - FAILURE
*/
func getUserObjGHAPI(GHUser string, numFollowers uint8) (User, error) {

	token := os.Getenv("TOKEN")
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	var followerNameList []string

	opt := &github.ListOptions{Page: 1, PerPage: int(numFollowers)}
	followerObjs, _, err := client.Users.ListFollowers(context, GHUser, opt)
	if err != nil {
		log.Printf("Problem getting follower information %v\n", err)
	}

	if len(followerObjs) < int(numFollowers) {
		numFollowers = uint8(len(followerObjs))
	}

	for _, followerObj := range followerObjs {
		followerNameList = append(followerNameList, *followerObj.Login)
	}

	GHUserObj := User{Name: GHUser, Followers: followerNameList}

	return GHUserObj, err
}

/*
DESC: Returns a list of 5 followers for the specified GithUb user as well as the followers of the followers numLevls levels deep.
IN: rootUser: the initial user, numLvls: the number of levels deep to run, numFollowers: the number of followers to list for each user
OUT: userObjList: a list of objects of type User - SUCCESS; err - FAILURE
*/
func GenUserObjList(rootUser string, numLvls, numFollowers uint8) ([]User, error) {

	var userObjList []User
	var newUserObj User

	userObj, err := getUserObjGHAPI(rootUser, numFollowers)
	if err != nil {
		log.Printf("Problem getting user follower information %v\n", err)
	}
	userObjList = append(userObjList, userObj)
	for i := 1; i <= int(numLvls); i++ {
		for _, follower := range userObj.Followers {
			if IsDuplicateUser(follower, userObjList) == false {
				newUserObj, err = getUserObjGHAPI(follower, numFollowers)
				userObjList = append(userObjList, newUserObj)
				if err != nil {
					log.Printf("Problem getting user follower information %v\n", err)
				}
			}
		}
		userObj = newUserObj
	}
	return userObjList, err
}

/*
DESC: Checks if a user's follower has already been visited
IN: follower: the follower's name: userObjList: the current list of users
OUT: bool: true if the follower has already been visited, false if the user has never been visited
*/
func IsDuplicateUser(follower string, userObjList []User) bool {
	for _, user := range userObjList {
		if user.Name == follower {
			return true
		}
	}
	return false
}
