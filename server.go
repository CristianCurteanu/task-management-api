package main

import (
	"endpoints"
	"log"
	"net/http"
	_ "strings"

	"github.com/julienschmidt/httprouter"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	router := httprouter.New()
	router.GET("/", endpoints.UserResource)
	router.GET("/ping", endpoints.PingHandler)
	router.POST("/client/token", endpoints.ClientTokenResource)
	err := http.ListenAndServe(":3035", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
