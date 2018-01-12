package client

import (
	"helpers/mongo"

	mgo "gopkg.in/mgo.v2"
)

const CLIENT_TABLE = "client"

type Client struct {
	Email string
	Uuid  string
	Key   string
}

var index = mgo.Index{
	Key:        []string{"email"},
	Unique:     true,
	DropDups:   true,
	Background: true,
	Sparse:     true,
}

func (c *Client) Initialize() *mgo.Collection {
	return mongo.CreateInitialSession(CLIENT_TABLE, &index)
}
