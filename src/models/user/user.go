package models

import (
	"helpers/encryption"
	"helpers/mongo"

	mgo "gopkg.in/mgo.v2"
)

const USER_TABLE = "users"

type User struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

var index = mgo.Index{
	Key:        []string{"email"},
	Unique:     true,
	DropDups:   true,
	Background: true,
	Sparse:     true,
}

func (u *User) Initialize() *mgo.Collection {
	return mongo.CreateInitialSession(USER_TABLE, &index)
}

func (u *User) Authorize(password string) bool {
	decrypted, err := encryption.Decrypt(u.Password)
	if err != nil {
		return false
	} else {
		return decrypted == password
	}
}
