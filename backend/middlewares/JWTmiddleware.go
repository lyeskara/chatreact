package middlewares

import (
	"backend/models"
	"context"
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

		var Token_string string
		if r.URL.Path == "/ws" {
			Token_string = r.URL.Query().Get("token")
			if Token_string == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid or expired token")
				return
			}

			token, err := jwt.Parse(Token_string, func(token *jwt.Token) (interface{}, error) {
				// validate the signing algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// provide the key to validate the signature
				return []byte(models.Secret), nil
			})
			if err != nil {
				log.Println(err)
			}
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid or expired token")
				return
			}

			var user models.User
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				user.Username = claims["username"].(string)
			}

			ctx := context.WithValue(r.Context(), models.UserName, user.Username)

			// Token is valid, proceed to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			Token_string = r.Header.Get("Authorization") // remove "Bearer " prefix
			var JWT string
			for i := 7; i < len(Token_string); i++ {
				char := string(Token_string[i])
				JWT += char
			}

			if JWT == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid or expired token")
				return
			}

			token, err := jwt.Parse(JWT, func(token *jwt.Token) (interface{}, error) {
				// validate the signing algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// provide the key to validate the signature
				return []byte(models.Secret), nil
			})
			if err != nil {
				log.Println(err)
			}
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid or expired token")
				return
			}

			var user models.User
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				user.Username = claims["username"].(string)
			}

			ctx := context.WithValue(r.Context(), models.UserName, user.Username)

			// Token is valid, proceed to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		}

	})
}
