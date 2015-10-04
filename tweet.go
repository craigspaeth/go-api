package main

import (
  "time"
  "gopkg.in/mgo.v2/bson"
)

type Tweet struct {
	Id        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"created_at"`
}

type Tweets []Tweet
