package endpoints

import (
	"encoding/json"
	. "models/client"
	"net/http"
	"os/exec"

	jwt "helpers/jwt"

	jwtGo "github.com/dgrijalva/jwt-go"

	"github.com/julienschmidt/httprouter"
)

type ClientParams struct {
	Email string `bson:"email" json:"email"`
}

type Response struct {
	Message string
}

type TokenResponse struct {
	Token string
}

var clientParams ClientParams

func ClientTokenResource(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&clientParams); err != nil {
		panic(err.Error())
	}

	uuid := UuidGenerator()
	type ClientTokenClaims struct {
		Uuid string `json:"uuid"`
		jwtGo.StandardClaims
	}

	claims := ClientTokenClaims{
		uuid,
		jwtGo.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	key, _ := jwt.Encoder(claims)

	client := Client{Email: clientParams.Email, Uuid: uuid, Key: key}
	clientTable := new(Client).Initialize()
	insertionError := clientTable.Insert(client)

	if insertionError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Response{"Error, check with administrator"})
	} else {
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(TokenResponse{key})
	}
}

func UuidGenerator() string {
	uuid, error := exec.Command("uuidgen").Output()
	if error != nil {
		panic(error.Error())
	}
	return string(uuid)
}
