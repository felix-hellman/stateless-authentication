package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
)

func main() {
	configuration, error := loadConfig("./configuration.yaml")
	if error != nil {
		log.Fatal(error)
	}
	println("Secret" + configuration.Application.Secret)
	http.HandleFunc("/anonymous", func(w http.ResponseWriter, r *http.Request) {
		var claims jwt.Claims
		claims.Subject = "anonymous"
		claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
		claims.Set = map[string]interface{}{"email_verified": false}
		token, _ := claims.HMACSign("HS256", []byte(configuration.Application.Secret))
		tokenString := string(token)
		fmt.Fprintf(w, tokenString)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
