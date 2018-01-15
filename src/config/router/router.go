package router

import (
	"endpoints"
	"helpers/encryption"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func NewRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", endpoints.UserResource)
	router.GET("/ping", endpoints.PingHandler)

	// Client apps
	router.POST("/client/token", endpoints.ClientTokenResource)

	// Sessions
	router.POST("/session", endpoints.CreateNewSessionResource)
	router.DELETE("/session", endpoints.DestroySessionResource)

	// Users

	// Boards

	// Columns

	return router
}

func TokenAuth(handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := []byte(r.Header.Get("Authorization"))

		decr := encryption.Hashing{Key: []byte(os.Getenv("JWT_KEY")), Value: token}
		decrypted, err := decr.Decrypt()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
