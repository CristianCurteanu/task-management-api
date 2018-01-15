package endpoints

import (
	"encoding/json"
	auth "helpers/credentials"
	. "models/session_token"
	"net/http"
	"os"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type SessionParams struct {
	Email string `bson:"email" json:"email"`
	Token string `bson:"token" json:"token"`
}

type PasswordClaims struct {
	Password string `json:"uuid"`
	jwtGo.StandardClaims
}

var sessionParams SessionParams

// POST /session
// Return token of an existing session
func CreateNewSessionResource(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sessionParams); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Response{"Error, check with administrator"})
		return
	}

	credentials := auth.Credentials{Email: sessionParams.Email, Password: ""}
	password, err := credentials.DecodePasswordToken(sessionParams.Token)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Response{"Error, check with administrator"})
		return
	}

	user, authErr := credentials.Authenticate()
	if authErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Response{"Error, check with administrator"})
		return
	}

	coll := new(SessionToken).Initialize()
	loc, _ := time.LoadLocation("UTC")
	token := Hashing{Key: []byte(os.Getenv("JWT_KEY")), Value: []byte{sessionParams.Email}}

	encrypted, err := token.Encrypt()
	if err != nil {
		t.Fatalf(err.Error())
	}

	sessionToken := SessionToken{Token: encrypted, ExpirestAt: time.Now().In(loc).Add(15 * time.Minute)}
	insertionError := sessionTokenTable.Insert(sessionToken)
	if insertionError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Response{"Error, check with administrator"})
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(TokenResponse{encrypted})
}

// DELETE /session
func DestroySessionResource(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func Authentication(email, password string) string {

}
