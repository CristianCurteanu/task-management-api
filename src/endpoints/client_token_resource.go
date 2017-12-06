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

var clientParams ClientParams

func ClientTokenResource(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&clientParams); err != nil {
		panic(err.Error())
	}
	uuid := UuidGenerator()
	key := JwtEncoder(uuid, clientParams.Email)
	client := Client{Email: clientParams.Email, Uuid: string(uuid), Key: key}

	conn := client.Connect()
	if conn.NewRecord(client) {
		conn.Create(&client)
		json.NewEncoder(rw).Encode(Response{"OK"})
	} else {
		json.NewEncoder(rw).Encode(Response{"Already exists"})
	}

}

func UuidGenerator() []byte {
	uuid, error := exec.Command("uuidgen").Output()
	if error != nil {
		panic(error.Error())
	}
	return uuid
}

func JwtEncoder(uuid []byte, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: 15000,
	})
	tokenString, err := token.SignedString(uuid)
	if err != nil {
		panic(err.Error())
	}
	return tokenString
}
