package main

import (
	. "config/router"
	"log"
	"net/http"
	_ "strings"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	err := http.ListenAndServe(":3035", NewRoutes())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
