package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/pascaldekloe/jwt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "postgres"
)

func main() {

	configuration, error := loadConfig("./configuration.yaml")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if error != nil {
		log.Fatal(error)
	}
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
