package main

import (
  "github.com/gorilla/mux"
  "github.com/codegangsta/negroni"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/api/v1/tweets", TweetsIndex)
  router.HandleFunc("/api/v1/tweets/{id}", TweetsShow)
  n := negroni.Classic()
  n.UseHandler(router)
  n.Run(":3000")
}
