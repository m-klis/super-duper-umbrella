package middlewares

import (
	"gochicoba/helpers"
	"gochicoba/models"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			// w.Write([]byte(`empty authorization`))
			helpers.ErrorResponse(w, r, http.StatusUnauthorized, "failed", "empty authorization")
			return
		}

		splitToken := strings.Split(token, " ")
		if splitToken[1] == "" {
			// w.Write([]byte(`empty token`))
			helpers.ErrorResponse(w, r, http.StatusUnauthorized, "failed", "empty token")
			return
		}

		claim := &models.Claims{}

		jwtToken, err := jwt.ParseWithClaims(splitToken[1], claim, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if jwtToken.Raw != splitToken[1] {
			// w.Write([]byte(`invalid token is different`))
			helpers.ErrorResponse(w, r, http.StatusUnauthorized, "failed", "invalid token")
			return
		}
		if err != nil {
			// w.Write([]byte(err.Error()))
			helpers.ErrorResponse(w, r, http.StatusUnauthorized, "failed", err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
