/*
FILE: main.go
DESC: DRIVER - Initializes router to serve /followers and /repos API endpoints
LAST MODIFIED: 30-SEPT-18 by Joshua Haupt
*/

package main

import (
	"./followers"
	"./repos"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/*
CONSTANT DECLARATIONS
*/
const NUMLVLS = 3
const NUMFOLLOWERS = 5
const NUMREPOS = 5
const NUMSTARGAZERS = 5

/*
DESC: Returns a list of 5 followers for the specified GithUb user as well as the followers of the followers 3 levels deep.
*/
func followersEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userList, err := followers.GenUserObjList(params["user"], NUMLVLS, NUMFOLLOWERS)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Error with followers endpoint")
		return
	} else {
		respondWithJson(writer, http.StatusOK, userList)
	}
}

/*
DESC: Returns a list of 5 repositories for a specified GitHub user and a list of 5 stargazers for each repository. The lists are nested 3 levels deep.
*/
func reposEndPt(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userRepoList, err := repos.GenRepoStargazerLists(params["user"], NUMLVLS, NUMREPOS, NUMSTARGAZERS)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Error with repos endpoint")
		return
	} else {
		respondWithJson(writer, http.StatusOK, userRepoList)
	}
}

func respondWithError(writer http.ResponseWriter, code int, msg string) {
	respondWithJson(writer, code, map[string]string{"ERROR": msg})
}

func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/followers/{user}", followersEndPt).Methods("GET")
	router.HandleFunc("/repos/{user}", reposEndPt).Methods("GET")

	if err := http.ListenAndServe(":8880", router); err != nil {
		log.Fatal(err)
	}
}
