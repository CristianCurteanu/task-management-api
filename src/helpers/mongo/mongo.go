package mongo

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func CreateInitialSession(tableName string, index *mgo.Index) *mgo.Collection {
	mgo.SetDebug(true)
	session, err := mgo.Dial(os.Getenv("DB_HOST"))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	SetIndex(session, index, tableName)

	return session.DB(os.Getenv("DB_NAME")).C(tableName)
}

func SetIndex(s *mgo.Session, index *mgo.Index, name string) {
	session := s.Copy()
	defer session.Close()
	c := session.DB(os.Getenv("DB_NAME")).C(name)

	err := c.EnsureIndex(*index)
	if err != nil {
		log.Fatal(err)
	}
}
