package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Ping struct {
	Message string
}

func PingHandler(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if err := json.NewEncoder(rw).Encode(Ping{"welcome"}); err != nil {
		panic(err)
	}
}
