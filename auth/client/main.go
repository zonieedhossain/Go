package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("mysupersecretpcdhase")

func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenarateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}
func GenarateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "hello"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something wrong : %s", err.Error())
		return "", nil
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
func main() {
	fmt.Println("My simple client")
	handleRequests()
}
