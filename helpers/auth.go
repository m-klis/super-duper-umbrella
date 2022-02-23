package helpers

import (
	"errors"
	"gochicoba/models"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Authentication(r *http.Request) error {
	token := r.Header.Get("Authorization")
	if token == "" {
		return errors.New("empty authorization")
	}

	splitToken := strings.Split(token, " ")
	if splitToken[1] == "" {
		return errors.New("empty token")
	}

	claim := &models.Claims{}

	jwtToken, err := jwt.ParseWithClaims(splitToken[1], claim, func(t *jwt.Token) (interface{}, error) { return []byte(os.Getenv("SECRET_KEY")), nil })
	if jwtToken.Raw != splitToken[1] {
		return errors.New("invalid token is different")
	}
	if err != nil {
		return err
	}

	// log.Println(jwtToken)
	return nil
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNjQ1NTQ1NDY3fQ.TP0JoMG3YLEhavgiLA9OUe5Fyd49CGqrZlq-vAbfMX8
