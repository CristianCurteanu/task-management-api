package router

import (
	"endpoints"

	"github.com/julienschmidt/httprouter"
)

func NewRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", endpoints.UserResource)
	router.GET("/ping", endpoints.PingHandler)
	router.POST("/client/token", endpoints.ClientTokenResource)
	return router
}
