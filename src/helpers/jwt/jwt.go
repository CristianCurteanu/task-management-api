package jwt

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func Encoder(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Decoder(token string, claimType jwt.Claims) (*jwt.Token, error) {
	tk, err := jwt.ParseWithClaims(token, claimType, func(tk *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return tk, nil
	}
	return nil, err
}
