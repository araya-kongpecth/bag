package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/araya-kongpecth/mux-miniproject/handlers"
)

var jwtKey = []byte("secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	fmt.Println("Authenticate function run!!!!")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &handlers.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims,

			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("Authenticate pass!!!")
		// If the user is authorized, call the next handler
		next.ServeHTTP(w, r)
	})
}
