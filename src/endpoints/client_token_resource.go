package endpoints

import (
	"encoding/json"
	. "models/client"
	"net/http"
	"os/exec"

	jwt "github.com/dgrijalva/jwt-go"
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
	key := JwtEncoder(uuid, clientParams.Email)

	client := Client{Email: clientParams.Email, Uuid: string(uuid), Key: key}
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

func UuidGenerator() []byte {
	uuid, error := exec.Command("uuidgen").Output()
	if error != nil {
		panic(error.Error())
	}
	return uuid
}

func JwtEncoder(uuid []byte, val string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    val,
		ExpiresAt: 15000,
	})
	tokenString, err := token.SignedString(uuid)
	if err != nil {
		panic(err.Error())
	}
	return tokenString
}
