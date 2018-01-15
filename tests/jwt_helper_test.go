package main

import (
	"helpers/jwt"
	"testing"

	jwtGo "github.com/dgrijalva/jwt-go"
)

func TestJwtEncoder(test *testing.T) {
	type ClientTokenClaims struct {
		Uuid string `json:"uuid"`
		jwtGo.StandardClaims
	}

	claims := ClientTokenClaims{
		"test_$tR1nG",
		jwtGo.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	token, err := jwt.Encoder(claims)
	if err != nil {
		test.Fatalf("Token is unable to create")
	}

	if token == "" {
		test.Fatalf("Something went wrong with token encoding")
	}

	tokenStruct, err := jwt.Decoder(token, &ClientTokenClaims{})
	claimsDecoded, decoded := tokenStruct.Claims.(*ClientTokenClaims)

	if decoded {
		if claimsDecoded.Uuid != claims.Uuid {
			test.Fatalf("The token is not decoded properly")
		}
	} else {
		test.Fatalf("Error with decoding the JWT")
	}
}
