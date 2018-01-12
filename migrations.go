package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type People struct {
	FirstName string
	LastName  string
	Hobbies   []Hobby
}

type Hobby struct {
	Name string
}

func main() {
	mgo.SetDebug(true)
	session, err := mgo.Dial("localhost")
	defer session.Close()

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	conn := session.DB("test").C("people")

	person := People{FirstName: "Label", LastName: "Springs", Hobbies: []Hobby{Hobby{Name: "Fishing"}, Hobby{Name: "Reading"}}}
	errors := conn.Insert(person)

	if errors != nil {
		log.Fatal("==> Insertion Failed Error:", err)
		return
	}
	result := []People{}
	retrievalErr := conn.Find(bson.M{}).All(&result)
	if retrievalErr != nil {
		log.Fatal("Failed find person:", retrievalErr)
		return
	}
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("test").C("people")

	index := mgo.Index{
		Key:        []string{"firstname"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
