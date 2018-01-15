package main

import (
	. "models/session_token"
	"time"

	. "gopkg.in/check.v1"
)

var model = new(SessionToken)

func (s *ClientSuite) TestSessionTokenCreation(check *C) {
	new(SessionToken).Initialize().RemoveAll(nil)

	tokenCollection := new(SessionToken).Initialize()
	sessionToken := SessionToken{Token: "some-cr3epY-t0keN", ExpirestAt: time.After(time.Minute * 15)}

	insertError := tokenCollection.Insert(sessionToken)
	check.Assert(insertError, IsNil)
}
