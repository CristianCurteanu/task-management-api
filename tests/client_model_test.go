package main

import (
	. "models/client"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

func (s *ClientSuite) TestClientCreation(check *C) {
	cleanupDatabase()
	insertionError := newClient()
	check.Assert(insertionError, IsNil)
}

func (s *ClientSuite) TestClientInitializeReturnCollection(check *C) {
	clientTable := new(Client).Initialize()
	check.Assert(reflect.TypeOf(clientTable), Equals, reflect.TypeOf(new(mgo.Collection)))
}

func (s *ClientSuite) TestClientFindAll(check *C) {
	cleanupDatabase()
	newClient()
	clientTable := new(Client).Initialize()
	results := []Client{}
	resultErrors := clientTable.Find(bson.M{}).All(&results)
	check.Assert(reflect.TypeOf(resultErrors), Not(Equals), reflect.TypeOf(new(error)))
	check.Assert(len(results), Equals, 1)
	check.Assert(results[0].Email, Equals, "some_fake_email@gmail.com")
	check.Assert(results[0].Uuid, Equals, "someUUID")
	check.Assert(results[0].Key, Equals, "$oM3Fr3@Kin6K3Y")
}

func newClient() error {
	client := Client{Email: "some_fake_email@gmail.com", Uuid: "someUUID", Key: "$oM3Fr3@Kin6K3Y"}
	clientTable := new(Client).Initialize()
	return clientTable.Insert(client)
}

func cleanupDatabase() {
	new(Client).Initialize().RemoveAll(nil)
}
