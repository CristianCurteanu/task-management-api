package endpoints

import (
	"fmt"
	. "models/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UserResource(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := User{Name: "First", Email: "some@email.com"}
	fmt.Println(user)
	fmt.Fprintf(rw, "application/json", "{ping: \"welcome\"}")
}
