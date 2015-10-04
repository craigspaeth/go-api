package main

import (
  "gopkg.in/mgo.v2"
  "log"
  "os"
)

func DbConnect() *mgo.Session {
  session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
  if err != nil {
    panic(err)
  }
  col := session.DB("goapi").C("tweets")
  err = col.Insert(&Tweet{Body: "Hello World"})
  if err != nil {
    log.Fatal(err)
  }
  return session
}
