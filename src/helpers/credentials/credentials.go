package credentials

import (
	"errors"
	"helpers/jwt"
	. "models/user"

	jwtGo "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type Credentials struct {
	Email    string
	Password string
}

type PasswordClaims struct {
	Password string `json:"uuid"`
	jwtGo.StandardClaims
}

func (c *Credentials) DecodePasswordToken(token string) (string, error) {
	tokenStruct, err := jwt.Decoder(token, &PasswordClaims{})
	if err != nil {
		return "", err
	}

	passwordStruct, decoded := tokenStruct.Claims.(*PasswordClaims)
	if decoded {
		c.Password = passwordStruct.Password
		return passwordStruct.Password, nil
	}
	return "", errors.New("Something went wrong during decoding")
}

func (c *Credentials) Authenticate() (User, error) {
	coll := new(User).Initialize()
	user := User{}
	err := coll.Find(bson.M{"email": c.Email}).Select(bson.M{}).One(&user)
	if err != nil {
		return User{}, err
	} else if user.Authorize(c.Password) {
		return user, nil
	}
	return User{}, errors.New("Unauthorized")
}

func (c *Credentials) Authorize() (user User) {
	return
}
