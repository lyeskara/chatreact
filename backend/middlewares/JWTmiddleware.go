package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JWTmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" || r.URL.Path == "/signup" {
			next.ServeHTTP(w, r)
			return
		}

		Token_string := r.Header.Get("Authorization") // remove "Bearer " prefix
		var JWT string

		for i := 7; i < len(Token_string); i++ {
			char := string(Token_string[i])
			JWT += char
		}
		
		token, err := jwt.Parse(JWT, func(token *jwt.Token) (interface{}, error) {
			// validate the signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// provide the key to validate the signature
			return []byte("KDks9v2f9EK1B1RZ"), nil
		})
		if err != nil {
			log.Fatal(err)
		}
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid or expired token")
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
