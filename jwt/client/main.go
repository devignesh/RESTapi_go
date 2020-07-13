package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	// jwt "github.com/dgrijava/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func GenerateJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	cliams := token.Claims.(jwt.MapClaims)

	cliams["authorized"] = true
	cliams["user"] = "vignesh"
	cliams["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("simple client")

	tokenString, err := GenerateJwt()

	if err != nil {
		fmt.Println("erroe while genereaating token")
	}
	fmt.Println(tokenString)
}
