package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
  "log"
)

func Home(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to the home page!")
}

func TweetsIndex(w http.ResponseWriter, r *http.Request) {
  var result Tweets
  session := DbConnect()
  tweets := session.DB("goapi").C("tweets")
  err := tweets.Find(nil).Limit(100).Iter().All(&result)
  out, err := json.Marshal(result)
  log.Printf("moo")
  if err != nil {
    fmt.Fprintf(w, "error")
  }
  fmt.Fprintf(w, string(out))
}

func TweetsShow(w http.ResponseWriter, r *http.Request) {
  var result Tweet
  session := DbConnect()
  tweets := session.DB("goapi").C("tweets")
  vars := mux.Vars(r)
  id := bson.ObjectId(vars["id"])
  err := tweets.Find(bson.M{"_id": id}).One(&result)
  if err != nil {
    fmt.Fprintf(w, "error")
  }
  out, err := json.Marshal(result)
  if err != nil {
    fmt.Fprintf(w, "error")
  }
  fmt.Fprintf(w, string(out))
}
