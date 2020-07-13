package main

import (
	fmt "fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	// "github.com/gorilla/mux"
)

var mySigningKey = []byte("mysupersecretphrase")

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "sample server page")

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Println("not authorized")
		}

	})
}

func handlerequests() {
	// r := mux.NewRouter()
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("sampe server page")
	handlerequests()

}
