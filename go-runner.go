package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/pascaldekloe/jwt"
        "time"
)

func main() {

        http.HandleFunc("/anonymous", func(w http.ResponseWriter, r *http.Request) {
                var claims jwt.Claims
                claims.Subject = "anonymous"
                claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
                claims.Set = map[string]interface{}{"email_verified": false}
                token, _ := claims.HMACSign("HS256", []byte("secret"))
                token_string := string(token)
                fmt.Fprintf(w, token_string)
        })
        log.Fatal(http.ListenAndServe(":8080", nil))
}

