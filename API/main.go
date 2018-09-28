package main

import (
  //"fmt"
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"

  "./followers"
)

func followersEndPt(writer http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  userList, err := followers.GenUserObjList(params["user"], 3, 5)
  if err != nil {
    respondWithError(writer, http.StatusBadRequest, "Well, something went wrong...")
    return
  } else {
    respondWithJson(writer, http.StatusOK, userList)
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

  if err := http.ListenAndServe(":8880", router); err != nil {
    log.Fatal(err)
  }
}
