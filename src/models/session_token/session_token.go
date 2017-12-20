package models

import (
	"helpers/mongo"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const SESSION_TOKEN_TABLE = "session_tokens"

type SessionToken struct {
	Token      string
	ExpirestAt time.Time
}

var index = mgo.Index{
	Key:        []string{"token", "expiresat"},
	Unique:     true,
	DropDups:   true,
	Background: true,
	Sparse:     true,
}

func (c *SessionToken) Initialize() *mgo.Collection {
	return mongo.CreateInitialSession(SESSION_TOKEN_TABLE, &index)
}
